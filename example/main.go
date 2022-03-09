package main

import (
	"log"
	"os"

	"github.com/traPtitech/traq-ws-bot"
	"github.com/traPtitech/traq-ws-bot/payload"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
}

func getEnvOrPanic(name string) string {
	s := os.Getenv(name)
	if s == "" {
		panic(name + " is required")
	}
	return s
}

func main() {
	bot, err := traqWSBot.NewBot(&traqWSBot.Options{
		AccessToken: getEnvOrPanic("ACCESS_TOKEN"),
		Origin:      getEnvOrPanic("TRAQ_ORIGIN"),
	})
	if err != nil {
		panic(err)
	}

	bot.OnError(func(message string) {
		log.Println("Received ERROR message: " + message)
	})
	bot.OnMessageCreated(func(p *payload.MessageCreated) {
		log.Println("Received MESSAGE_CREATED event: " + p.Message.Text)
	})

	if err := bot.Start(); err != nil {
		panic(err)
	}
}
