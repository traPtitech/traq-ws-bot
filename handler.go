package traqwsbot

import (
	"encoding/json"
	"log"

	"github.com/traPtitech/traq-ws-bot/event"
	"github.com/traPtitech/traq-ws-bot/payload"
)

func (b *Bot) handleMultiCast(event string, raw json.RawMessage) {
	for _, h := range b.handlers[event] {
		go h(raw)
	}
}

// OnEvent 指定したイベントに対してハンドラを登録します。
func (b *Bot) OnEvent(event string, h func(rawPayload json.RawMessage)) {
	b.handlers[event] = append(b.handlers[event], h)
}

// OnAnyEvent 任意のイベントに対してハンドラを登録します。
func (b *Bot) OnAnyEvent(h func(e string, rawPayload json.RawMessage)) {
	for _, e := range event.AllEvents {
		b.OnEvent(e, func(rawPayload json.RawMessage) {
			h(e, rawPayload)
		})
	}
}

// makeDecoder is a utility function to be used with (*Bot).OnEvent.
func makeDecoder[P any](next func(p *P)) func(raw json.RawMessage) {
	return func(raw json.RawMessage) {
		var p P
		if err := json.Unmarshal(raw, &p); err != nil {
			log.Printf("[traq-ws-bot] Unexpected payload while unmarshaling to %T: %s\n", p, err)
			return
		}
		next(&p)
	}
}

// OnError ERROR イベントハンドラを登録します。
func (b *Bot) OnError(h func(message string)) {
	b.OnEvent(event.Error, makeDecoder(func(p *string) { h(*p) }))
}
func (b *Bot) OnPing(h func(p *payload.Ping)) {
	b.OnEvent(event.Ping, makeDecoder(h))
}
func (b *Bot) OnJoined(h func(p *payload.Joined)) {
	b.OnEvent(event.Joined, makeDecoder(h))
}
func (b *Bot) OnLeft(h func(p *payload.Left)) {
	b.OnEvent(event.Left, makeDecoder(h))
}
func (b *Bot) OnMessageCreated(h func(p *payload.MessageCreated)) {
	b.OnEvent(event.MessageCreated, makeDecoder(h))
}
func (b *Bot) OnMessageUpdated(h func(p *payload.MessageUpdated)) {
	b.OnEvent(event.MessageUpdated, makeDecoder(h))
}
func (b *Bot) OnMessageDeleted(h func(p *payload.MessageDeleted)) {
	b.OnEvent(event.MessageDeleted, makeDecoder(h))
}
func (b *Bot) OnBotMessageStampsUpdated(h func(p *payload.BotMessageStampsUpdated)) {
	b.OnEvent(event.BotMessageStampsUpdated, makeDecoder(h))
}
func (b *Bot) OnDirectMessageCreated(h func(p *payload.DirectMessageCreated)) {
	b.OnEvent(event.DirectMessageCreated, makeDecoder(h))
}
func (b *Bot) OnDirectMessageUpdated(h func(p *payload.DirectMessageUpdated)) {
	b.OnEvent(event.DirectMessageUpdated, makeDecoder(h))
}
func (b *Bot) OnDirectMessageDeleted(h func(p *payload.DirectMessageDeleted)) {
	b.OnEvent(event.DirectMessageDeleted, makeDecoder(h))
}
func (b *Bot) OnChannelCreated(h func(p *payload.ChannelCreated)) {
	b.OnEvent(event.ChannelCreated, makeDecoder(h))
}
func (b *Bot) OnChannelTopicChanged(h func(p *payload.ChannelTopicChanged)) {
	b.OnEvent(event.ChannelTopicChanged, makeDecoder(h))
}
func (b *Bot) OnUserCreated(h func(p *payload.UserCreated)) {
	b.OnEvent(event.UserCreated, makeDecoder(h))
}
func (b *Bot) OnUserActivated(h func(p *payload.UserActivated)) {
	b.OnEvent(event.UserActivated, makeDecoder(h))
}
func (b *Bot) OnStampCreated(h func(p *payload.StampCreated)) {
	b.OnEvent(event.StampCreated, makeDecoder(h))
}
func (b *Bot) OnTagAdded(h func(p *payload.TagAdded)) {
	b.OnEvent(event.TagAdded, makeDecoder(h))
}
func (b *Bot) OnTagRemoved(h func(p *payload.TagRemoved)) {
	b.OnEvent(event.TagRemoved, makeDecoder(h))
}
