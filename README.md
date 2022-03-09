# traQ WebSocket BOT ライブラリ

traQのWebSocket BOTを作るためのライブラリです。

WebSocketへの接続と、WebSocket内でのデータの送信/受信を補助します。 APIへの接続は補助しないため、[sapphi-red/go-traq](https://github.com/sapphi-red/go-traq)
などと合わせて使用してください。

## Sample

```go
package main

import (
	"encoding/json"
	"log"

	"github.com/traPtitech/traq-ws-bot"
	"github.com/traPtitech/traq-ws-bot/payload"
)

func main() {
	// Create a bot instance
	b, err := traqWSBot.NewBot(&traqWSBot.Options{
		AccessToken: "access-token", // required
		Origin:      "wss://q.trap.jp",
	})
	if err != nil {
		panic(err)
	}

	// Set event handlers
	b.OnMessageCreated(func(p *payload.MessageCreated) {
		log.Println("Message created", p)
	})
	b.OnError(func(message string) {
		log.Println("Command error", message)
	})
	b.OnEvent("ANY_EVENT", func(rawPayload json.RawMessage) {
		log.Println("You can receive any events in case it is not implemented")
	})

	go func() {
		// Join Qall session
		b.SendRTCState("channel-id", [2]string{"qall.micmuted", "session-id"})
		// Leave Qall session
		b.SendRTCState("channel-id")
	}()

	// Connect to WS and start receiving events
	if err := b.Start(); err != nil {
		panic(err)
	}
}
```
