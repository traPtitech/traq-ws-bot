package traq_ws_bot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

const (
	defaultOrigin       = "wss://q.trap.jp"
	botGatewayPath      = "/api/v3/bots/ws"
	authorizationScheme = "Bearer"
)

// Options Bot のオプション
type Options struct {
	// AccessToken （必須）BOTのアクセストークンを指定します。
	AccessToken string
	// Origin traQオリジンを指定します。
	// e.g. https://q.trap.jp
	Origin string
}

// Bot WebSocket BOT
type Bot struct {
	accessToken string
	origin      string
	handlers    map[string][]func(interface{})

	send chan<- *rawMessage
	c    *websocket.Conn
}

// NewBot Bot を作成します。
func NewBot(options *Options) (*Bot, error) {
	if options.AccessToken == "" {
		return nil, errors.New("access token is needed")
	}
	origin := options.Origin
	if origin == "" {
		origin = defaultOrigin
	}
	return &Bot{
		accessToken: options.AccessToken,
		origin:      origin,
		handlers:    make(map[string][]func(interface{})),
	}, nil
}

// Start WebSocketに接続し、イベントの送信と受信を始めます。
// 成功した場合、ブロックします。
func (b *Bot) Start() error {
	c, _, err := websocket.DefaultDialer.Dial(b.origin+botGatewayPath, http.Header{
		"Authorization": []string{authorizationScheme + " " + b.accessToken},
	})
	if err != nil {
		return fmt.Errorf("traq-ws-bot encountered an error while dialing %s: %w", b.origin+botGatewayPath, err)
	}

	done := make(chan struct{})
	send := make(chan *rawMessage)
	b.send = send
	b.c = c
	go b.readLoop(done)
	b.writeLoop(done, send)
	return nil
}

func (b *Bot) sendMessage(m *rawMessage) {
	b.send <- m
}

func (b *Bot) SendRTCState(channelID uuid.UUID, states ...[2]string) {
	if len(states) == 0 {
		b.sendMessage(&rawMessage{
			t:    websocket.TextMessage,
			data: []byte(fmt.Sprintf("rtcstate:%s:", channelID)),
		})
		return
	}

	elems := make([]string, 0, 2+len(states)*2)
	elems = append(elems, "rtcstate", channelID.String())
	for _, state := range states {
		elems = append(elems, state[0], state[1])
	}
	b.sendMessage(&rawMessage{
		t:    websocket.TextMessage,
		data: []byte(strings.Join(elems, ":")),
	})
}

// Close 接続を切断します。
func (b *Bot) Close() error {
	if b.c == nil {
		return nil
	}
	return b.c.Close()
}

type rawMessage struct {
	t    int
	data []byte
}

type eventMessage struct {
	Type string      `json:"type"`
	Body interface{} `json:"body"`
}

func (b *Bot) readLoop(done chan<- struct{}) {
	defer close(done)
	for {
		t, p, err := b.c.ReadMessage()
		if err != nil {
			return
		}

		switch t {
		case websocket.TextMessage:
			var m eventMessage
			if err := json.NewDecoder(bytes.NewReader(p)).Decode(&m); err != nil {
				b.sendMessage(&rawMessage{t: websocket.CloseMessage, data: websocket.FormatCloseMessage(websocket.CloseUnsupportedData, "unexpected json format")})
				log.Println("traq-ws-bot: unexpected json format, closing connection")
				return
			}
			go b.handle(m.Type, m.Body)
		case websocket.CloseMessage:
			return
		}
	}
}

func (b *Bot) writeLoop(done <-chan struct{}, send <-chan *rawMessage) {
	for {
		select {
		case <-done:
			return
		case m := <-send:
			err := b.c.WriteMessage(m.t, m.data)
			if err != nil {
				return
			}
		}
	}
}
