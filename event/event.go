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
	// UserActivated ユーザー凍結解除ベント
	UserActivated = "USER_ACTIVATED"
	// StampCreated スタンプ作成イベント
	StampCreated = "STAMP_CREATED"
	// TagAdded タグ追加イベント
	TagAdded = "TAG_ADDED"
	// TagRemoved タグ削除イベント
	TagRemoved = "TAG_REMOVED"
	// UserGroupCreated グループ作成イベント
	UserGroupCreated = "USER_GROUP_CREATED"
	// UserGroupUpdated グループ更新イベント
	UserGroupUpdated = "USER_GROUP_UPDATED"
	// UserGroupDeleted グループ削除イベント
	UserGroupDeleted = "USER_GROUP_DELETED"
	// UserGroupMemberAdded グループメンバー追加イベント
	UserGroupMemberAdded = "USER_GROUP_MEMBER_ADDED"
	// UserGroupMemberUpdated グループメンバー更新イベント
	UserGroupMemberUpdated = "USER_GROUP_MEMBER_UPDATED"
	// UserGroupMemberRemoved グループメンバー削除イベント
	UserGroupMemberRemoved = "USER_GROUP_MEMBER_REMOVED"
	// UserGroupAdminAdded グループ管理者追加イベント
	UserGroupAdminAdded = "USER_GROUP_ADMIN_ADDED"
	// UserGroupAdminRemoved グループ管理者削除イベント
	UserGroupAdminRemoved = "USER_GROUP_ADMIN_REMOVED"
)

var AllEvents = []string{
	Error,
	Ping,
	Joined,
	Left,
	MessageCreated,
	MessageUpdated,
	MessageDeleted,
	BotMessageStampsUpdated,
	DirectMessageCreated,
	DirectMessageUpdated,
	DirectMessageDeleted,
	ChannelCreated,
	ChannelTopicChanged,
	UserCreated,
	UserActivated,
	StampCreated,
	TagAdded,
	TagRemoved,
	UserGroupCreated,
	UserGroupUpdated,
	UserGroupDeleted,
	UserGroupMemberAdded,
	UserGroupMemberUpdated,
	UserGroupMemberRemoved,
	UserGroupAdminAdded,
	UserGroupAdminRemoved,
}
