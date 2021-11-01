package traq_ws_bot

import (
	"encoding/json"
	"log"

	"github.com/traPtitech/traq-ws-bot/event"
	"github.com/traPtitech/traq-ws-bot/payload"
)

func (b *Bot) handle(event string, p interface{}) {
	for _, h := range b.handlers[event] {
		h(p)
	}
}

// OnEvent 任意のイベントについてハンドラを登録します。
func (b *Bot) OnEvent(event string, h func(i interface{})) {
	b.handlers[event] = append(b.handlers[event], h)
}

// OnError ERROR イベントハンドラを登録します。
func (b *Bot) OnError(h func(message string)) {
	b.OnEvent(event.Error, func(i interface{}) {
		message, ok := i.(string)
		if !ok {
			log.Println("[traq-ws-bot] Unexpected payload on ERROR")
			return
		}
		h(message)
	})
}

// OnPing PING イベントハンドラを登録します。
func (b *Bot) OnPing(h func(p *payload.Ping)) {
	b.OnEvent(event.Ping, func(i interface{}) {
		b, _ := json.Marshal(i)
		var p payload.Ping
		if err := json.Unmarshal(b, &p); err != nil {
			log.Println("[traq-ws-bot] Unexpected payload on PING")
			return
		}
		h(&p)
	})
}

// OnJoined JOINED イベントハンドラを登録します。
func (b *Bot) OnJoined(h func(p *payload.Joined)) {
	b.OnEvent(event.Joined, func(i interface{}) {
		b, _ := json.Marshal(i)
		var p payload.Joined
		if err := json.Unmarshal(b, &p); err != nil {
			log.Println("[traq-ws-bot] Unexpected payload on JOINED")
			return
		}
		h(&p)
	})
}

// OnLeft LEFT イベントハンドラを登録します。
func (b *Bot) OnLeft(h func(p *payload.Left)) {
	b.OnEvent(event.Left, func(i interface{}) {
		b, _ := json.Marshal(i)
		var p payload.Left
		if err := json.Unmarshal(b, &p); err != nil {
			log.Println("[traq-ws-bot] Unexpected payload on LEFT")
			return
		}
		h(&p)
	})
}

// OnMessageCreated MESSAGE_CREATED イベントハンドラを登録します。
func (b *Bot) OnMessageCreated(h func(p *payload.MessageCreated)) {
	b.OnEvent(event.MessageCreated, func(i interface{}) {
		b, _ := json.Marshal(i)
		var p payload.MessageCreated
		if err := json.Unmarshal(b, &p); err != nil {
			log.Println("[traq-ws-bot] Unexpected payload on MESSAGE_CREATED")
			return
		}
		h(&p)
	})
}

// OnMessageUpdated MESSAGE_UPDATED イベントハンドラを登録します。
func (b *Bot) OnMessageUpdated(h func(p *payload.MessageUpdated)) {
	b.OnEvent(event.MessageUpdated, func(i interface{}) {
		b, _ := json.Marshal(i)
		var p payload.MessageUpdated
		if err := json.Unmarshal(b, &p); err != nil {
			log.Println("[traq-ws-bot] Unexpected payload on MESSAGE_UPDATED")
			return
		}
		h(&p)
	})
}

// OnMessageDeleted MESSAGE_DELETED イベントハンドラを登録します。
func (b *Bot) OnMessageDeleted(h func(p *payload.MessageDeleted)) {
	b.OnEvent(event.MessageDeleted, func(i interface{}) {
		b, _ := json.Marshal(i)
		var p payload.MessageDeleted
		if err := json.Unmarshal(b, &p); err != nil {
			log.Println("[traq-ws-bot] Unexpected payload on MESSAGE_DELETED")
			return
		}
		h(&p)
	})
}

// OnBotMessageStampsUpdated BOT_MESSAGE_STAMPS_UPDATED イベントハンドラを登録します。
func (b *Bot) OnBotMessageStampsUpdated(h func(p *payload.BotMessageStampsUpdated)) {
	b.OnEvent(event.BotMessageStampsUpdated, func(i interface{}) {
		b, _ := json.Marshal(i)
		var p payload.BotMessageStampsUpdated
		if err := json.Unmarshal(b, &p); err != nil {
			log.Println("[traq-ws-bot] Unexpected payload on BOT_MESSAGE_STAMPS_UPDATED")
			return
		}
		h(&p)
	})
}

// OnDirectMessageCreated DIRECT_MESSAGE_CREATED イベントハンドラを登録します。
func (b *Bot) OnDirectMessageCreated(h func(p *payload.DirectMessageCreated)) {
	b.OnEvent(event.DirectMessageCreated, func(i interface{}) {
		b, _ := json.Marshal(i)
		var p payload.DirectMessageCreated
		if err := json.Unmarshal(b, &p); err != nil {
			log.Println("[traq-ws-bot] Unexpected payload on DIRECT_MESSAGE_CREATED")
			return
		}
		h(&p)
	})
}

// OnDirectMessageUpdated DIRECT_MESSAGE_UPDATED イベントハンドラを登録します。
func (b *Bot) OnDirectMessageUpdated(h func(p *payload.DirectMessageUpdated)) {
	b.OnEvent(event.DirectMessageUpdated, func(i interface{}) {
		b, _ := json.Marshal(i)
		var p payload.DirectMessageUpdated
		if err := json.Unmarshal(b, &p); err != nil {
			log.Println("[traq-ws-bot] Unexpected payload on DIRECT_MESSAGE_UPDATED")
			return
		}
		h(&p)
	})
}

// OnDirectMessageDeleted DIRECT_MESSAGE_DELETED イベントハンドラを登録します。
func (b *Bot) OnDirectMessageDeleted(h func(p *payload.DirectMessageDeleted)) {
	b.OnEvent(event.DirectMessageDeleted, func(i interface{}) {
		b, _ := json.Marshal(i)
		var p payload.DirectMessageDeleted
		if err := json.Unmarshal(b, &p); err != nil {
			log.Println("[traq-ws-bot] Unexpected payload on DIRECT_MESSAGE_DELETED")
			return
		}
		h(&p)
	})
}

// OnChannelCreated CHANNEL_CREATED イベントハンドラを登録します。
func (b *Bot) OnChannelCreated(h func(p *payload.ChannelCreated)) {
	b.OnEvent(event.ChannelCreated, func(i interface{}) {
		b, _ := json.Marshal(i)
		var p payload.ChannelCreated
		if err := json.Unmarshal(b, &p); err != nil {
			log.Println("[traq-ws-bot] Unexpected payload on CHANNEL_CREATED")
			return
		}
		h(&p)
	})
}

// OnChannelTopicChanged CHANNEL_TOPIC_CHANGED イベントハンドラを登録します。
func (b *Bot) OnChannelTopicChanged(h func(p *payload.ChannelTopicChanged)) {
	b.OnEvent(event.ChannelTopicChanged, func(i interface{}) {
		b, _ := json.Marshal(i)
		var p payload.ChannelTopicChanged
		if err := json.Unmarshal(b, &p); err != nil {
			log.Println("[traq-ws-bot] Unexpected payload on CHANNEL_TOPIC_CHANGED")
			return
		}
		h(&p)
	})
}

// OnUserCreated USER_CREATED イベントハンドラを登録します。
func (b *Bot) OnUserCreated(h func(p *payload.UserCreated)) {
	b.OnEvent(event.UserCreated, func(i interface{}) {
		b, _ := json.Marshal(i)
		var p payload.UserCreated
		if err := json.Unmarshal(b, &p); err != nil {
			log.Println("[traq-ws-bot] Unexpected payload on USER_CREATED")
			return
		}
		h(&p)
	})
}

// OnStampCreated STAMP_CREATED イベントハンドラを登録します。
func (b *Bot) OnStampCreated(h func(p *payload.StampCreated)) {
	b.OnEvent(event.StampCreated, func(i interface{}) {
		b, _ := json.Marshal(i)
		var p payload.StampCreated
		if err := json.Unmarshal(b, &p); err != nil {
			log.Println("[traq-ws-bot] Unexpected payload on STAMP_CREATED")
			return
		}
		h(&p)
	})
}

// OnTagAdded TAG_ADDED イベントハンドラを登録します。
func (b *Bot) OnTagAdded(h func(p *payload.TagAdded)) {
	b.OnEvent(event.TagAdded, func(i interface{}) {
		b, _ := json.Marshal(i)
		var p payload.TagAdded
		if err := json.Unmarshal(b, &p); err != nil {
			log.Println("[traq-ws-bot] Unexpected payload on TAG_ADDED")
			return
		}
		h(&p)
	})
}

// OnTagRemoved TAG_REMOVED イベントハンドラを登録します。
func (b *Bot) OnTagRemoved(h func(p *payload.TagRemoved)) {
	b.OnEvent(event.TagRemoved, func(i interface{}) {
		b, _ := json.Marshal(i)
		var p payload.TagRemoved
		if err := json.Unmarshal(b, &p); err != nil {
			log.Println("[traq-ws-bot] Unexpected payload on TAG_REMOVED")
			return
		}
		h(&p)
	})
}
