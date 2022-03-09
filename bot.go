package traq_ws_bot

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

const (
	defaultOrigin       = "wss://q.trap.jp"
	botGatewayPath      = "/api/v3/bots/ws"
	authorizationScheme = "Bearer"
	firstRetryWait      = 3 * time.Second
	maxRetryWait        = 10 * time.Minute
)

// Options Bot のオプション
type Options struct {
	// AccessToken BOTのアクセストークン (required)
	AccessToken string
	// Origin traQオリジン (default: wss://q.trap.jp)
	// e.g. wss://q.trap.jp, ws://localhost:3000
	Origin string
	// DisableAutoReconnect 接続が終了した、もしくは失敗した場合の自動再接続を無効化する (default: false)
	DisableAutoReconnect bool
}

// Bot WebSocket BOT
type Bot struct {
	op            *Options
	nextRetryWait time.Duration
	handlers      map[string][]func(json.RawMessage)

	conn *wsConn
}

// NewBot Bot を作成します。
func NewBot(options *Options) (*Bot, error) {
	if options.AccessToken == "" {
		return nil, errors.New("access token is required")
	}
	op := *options
	if op.Origin == "" {
		op.Origin = defaultOrigin
	}
	return &Bot{
		op:            &op,
		nextRetryWait: firstRetryWait,
		handlers:      make(map[string][]func(json.RawMessage)),
	}, nil
}

// Start WebSocketに接続し、イベントの送信と受信を始めます。
// 成功した場合、ブロックします。
func (b *Bot) Start() error {
	for {
		err := b.connect()
		if b.op.DisableAutoReconnect {
			return err
		}

		if err == nil {
			// once connected, but disconnected for some reason
			b.nextRetryWait = firstRetryWait
			log.Printf("[traq-ws-bot] Disconnected from WebSocket, retrying in %v ...\n", b.nextRetryWait)
			time.Sleep(b.nextRetryWait)
		} else {
			log.Printf("[traq-ws-bot] Encountered an error while dialing %s: %s\n", b.op.Origin+botGatewayPath, err)
			log.Printf("[traq-ws-bot] Failed to connect to WebSocket, retrying in %v ...\n", b.nextRetryWait)
			time.Sleep(b.nextRetryWait)
			// exponential backoff
			b.nextRetryWait = min(b.nextRetryWait*2, maxRetryWait)
		}
	}
}

func (b *Bot) connect() error {
	c, _, err := websocket.DefaultDialer.Dial(b.op.Origin+botGatewayPath, http.Header{
		"Authorization": []string{authorizationScheme + " " + b.op.AccessToken},
	})
	if err != nil {
		return err
	}

	log.Println("[traq-ws-bot] Connected! Now receiving events...")
	b.conn = newWSConn(c)
	b.conn.OnTextMessage(b.handleRawTextMessage)
	b.conn.Start()
	return nil
}

func (b *Bot) handleRawTextMessage(p []byte) {
	var m struct {
		Type string          `json:"type"`
		Body json.RawMessage `json:"body"`
	}
	if err := json.NewDecoder(bytes.NewReader(p)).Decode(&m); err != nil {
		b.conn.WriteMessage(&rawMessage{t: websocket.CloseMessage, data: websocket.FormatCloseMessage(websocket.CloseUnsupportedData, "unexpected json format")})
		log.Println("[traq-ws-bot] Unexpected json format, closing connection")
		return
	}
	b.handleMultiCast(m.Type, m.Body)
}

func (b *Bot) SendRTCState(channelID uuid.UUID, states ...[2]string) {
	elems := make([]string, 0, 2+len(states)*2)
	elems = append(elems, "rtcstate", channelID.String())
	for _, state := range states {
		elems = append(elems, state[0], state[1])
	}
	b.conn.WriteMessage(&rawMessage{
		t:    websocket.TextMessage,
		data: []byte(strings.Join(elems, ":") + ":"),
	})
}
