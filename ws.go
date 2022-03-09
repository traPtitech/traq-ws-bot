package traqWSBot

import (
	"sync"

	"github.com/gorilla/websocket"
)

// rawMessage indicates a raw WebSocket message.
type rawMessage struct {
	t    int
	data []byte
}

// wsConn is a light-weight wrapper of *websocket.Conn.
type wsConn struct {
	conn   *websocket.Conn
	send   chan *rawMessage
	closed bool
	sync.RWMutex

	textMessageHandler func(p []byte)
}

func newWSConn(conn *websocket.Conn) *wsConn {
	return &wsConn{
		conn:   conn,
		send:   make(chan *rawMessage),
		closed: false,
	}
}

func (w *wsConn) OnTextMessage(h func(p []byte)) {
	w.textMessageHandler = h
}

func (w *wsConn) Start() {
	go w.writeLoop()
	w.readLoop()
}

func (w *wsConn) WriteMessage(m *rawMessage) {
	w.send <- m
}

func (w *wsConn) readLoop() {
	defer w.close()

	for {
		t, p, err := w.conn.ReadMessage()
		if err != nil {
			return
		}

		switch t {
		case websocket.TextMessage:
			w.textMessageHandler(p)
		case websocket.BinaryMessage:
			// Not supported, just ignore it
		}
	}
}

func (w *wsConn) writeLoop() {
	defer w.close()

	for {
		m, ok := <-w.send
		if !ok {
			return
		}

		if err := w.conn.WriteMessage(m.t, m.data); err != nil {
			return
		}

		if m.t == websocket.CloseMessage {
			return
		}
	}
}

func (w *wsConn) close() {
	w.Lock()
	defer w.Unlock()

	if w.closed {
		return
	}
	w.closed = true
	_ = w.conn.Close()
	close(w.send)
}
