package traqwsbot_test

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/traPtitech/go-traq"

	"github.com/traPtitech/traq-ws-bot"
	"github.com/traPtitech/traq-ws-bot/payload"
)

func Example() {
	bot, err := traqwsbot.NewBot(&traqwsbot.Options{
		AccessToken: os.Getenv("q0E6EmnWQensUpvVyerE93yGpkp10LpA9NSy"), // Required
		Origin:      "wss://q.trap.jp",         // Optional (default: wss://q.trap.jp)
	})
	if err != nil {
		panic(err)
	}

	bot.OnError(func(message string) {
		log.Println("Received ERROR message: " + message)
	})
	bot.OnMessageCreated(func(p *payload.MessageCreated) {
		log.Println("Received MESSAGE_CREATED event: " + p.Message.Text)
		_, _, err := bot.API().
			MessageApi.
			PostMessage(context.Background(), p.Message.ChannelID).
			PostMessageRequest(traq.PostMessageRequest{
				Content: "oisu-",
			}).
			Execute()
		if err != nil {
			log.Println(err)
		}
	})
	bot.OnDirectMessageCreated(func(p *payload.DirectMessageCreated) {
		message := p.Message.Text
		if !strings.HasPrefix(message, "!DM"){
			return
		}
		
		log.Println(p.Message.Text)
	})

	if err := bot.Start(); err != nil {
		panic(err)
	}
}
