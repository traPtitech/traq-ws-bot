package event

// Error ERRORイベント WebSocketのみ
const Error = "ERROR"

const (
	// Ping Pingイベント
	Ping = "PING"
	// Joined チャンネル参加イベント
	Joined = "JOINED"
	// Left チャンネル退出イベント
	Left = "LEFT"
	// MessageCreated メッセージ作成イベント or メンションメッセージ作成イベント
	MessageCreated = "MESSAGE_CREATED"
	// MessageUpdated メッセージ編集イベント
	MessageUpdated = "MESSAGE_UPDATED"
	// MessageDeleted メッセージ削除イベント
	MessageDeleted = "MESSAGE_DELETED"
	// BotMessageStampsUpdated BOTメッセージスタンプ更新イベント
	BotMessageStampsUpdated = "BOT_MESSAGE_STAMPS_UPDATED"
	// DirectMessageCreated ダイレクトメッセージ作成イベント
	DirectMessageCreated = "DIRECT_MESSAGE_CREATED"
	// DirectMessageUpdated ダイレクトメッセージ編集イベント
	DirectMessageUpdated = "DIRECT_MESSAGE_UPDATED"
	// DirectMessageDeleted ダイレクトメッセージ削除イベント
	DirectMessageDeleted = "DIRECT_MESSAGE_DELETED"
	// ChannelCreated チャンネル作成イベント
	ChannelCreated = "CHANNEL_CREATED"
	// ChannelTopicChanged チャンネルトピック変更イベント
	ChannelTopicChanged = "CHANNEL_TOPIC_CHANGED"
	// UserCreated ユーザー作成イベント
	UserCreated = "USER_CREATED"
	// StampCreated スタンプ作成イベント
	StampCreated = "STAMP_CREATED"
	// TagAdded タグ追加イベント
	TagAdded = "TAG_ADDED"
	// TagRemoved タグ削除イベント
	TagRemoved = "TAG_REMOVED"
)
