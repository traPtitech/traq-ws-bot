package payload

// Ping PINGイベントペイロード
type Ping struct {
	Base
}

// Joined JOINEDイベントペイロード
type Joined struct {
	Base
	// Channel 参加したチャンネル
	Channel Channel `json:"channel"`
}

// Left LEFTイベントペイロード
type Left struct {
	Base
	// Channel 退出したチャンネル
	Channel Channel `json:"channel"`
}

// MessageCreated MESSAGE_CREATEDイベントペイロード
type MessageCreated struct {
	Base
	// Message 投稿されたメッセージ
	Message Message `json:"message"`
}

// MessageUpdated MESSAGE_UPDATEDイベントペイロード
type MessageUpdated struct {
	Base
	// Message 更新されたメッセージ
	Message Message `json:"message"`
}

// MessageDeleted MESSAGE_DELETEDイベントペイロード
type MessageDeleted struct {
	Base
	// Message 削除されたメッセージ
	Message struct {
		// ID メッセージUUID
		ID string `json:"id"`
		// ChannelID 投稿先チャンネルUUID
		ChannelID string `json:"channelId"`
	} `json:"message"`
}

// BotMessageStampsUpdated BOT_MESSAGE_STAMPS_UPDATEDイベントペイロード
type BotMessageStampsUpdated struct {
	Base
	// MessageID スタンプの更新があったメッセージUUID
	MessageID string `json:"messageId"`
	// Stamps メッセージに現在ついている全てのスタンプ
	Stamps []MessageStamp `json:"stamps"`
}

// DirectMessageCreated DIRECT_MESSAGE_CREATEDイベントペイロード
type DirectMessageCreated struct {
	Base
	// Message 投稿されたメッセージ
	Message Message `json:"message"`
}

// DirectMessageUpdated DIRECT_MESSAGE_UPDATEDイベントペイロード
type DirectMessageUpdated struct {
	Base
	// Message 更新されたメッセージ
	Message Message `json:"message"`
}

// DirectMessageDeleted DIRECT_MESSAGE_DELETEDイベントペイロード
type DirectMessageDeleted struct {
	Base
	// Message 削除されたメッセージ
	Message struct {
		// ID メッセージUUID
		ID string `json:"id"`
		// UserID DMの宛先ユーザーUUID
		UserID string `json:"userId"`
		// ChannelID 投稿先チャンネルUUID
		ChannelID string `json:"channelId"`
	} `json:"message"`
}

// ChannelCreated CHANNEL_CREATEDイベントペイロード
type ChannelCreated struct {
	Base
	// Channel 作成されたチャンネル
	Channel Channel `json:"channel"`
}

// ChannelTopicChanged CHANNEL_TOPIC_CHANGEDイベントペイロード
type ChannelTopicChanged struct {
	Base
	// Channel 変更されたチャンネル
	Channel Channel `json:"channel"`
	// Topic 変更後のトピック
	Topic string `json:"topic"`
	// Updater トピック更新者
	Updater User `json:"updater"`
}

// UserCreated USER_CREATEDイベントペイロード
type UserCreated struct {
	Base
	// User 作成されたユーザー
	User User `json:"user"`
}

// UserActivated USER_ACTIVATEDイベントペイロード
type UserActivated struct {
	Base
	// User 凍結解除されたユーザー
	User User `json:"user"`
}

// StampCreated STAMP_CREATEDイベントペイロード
type StampCreated struct {
	Base
	// ID スタンプUUID
	ID string `json:"id"`
	// Name スタンプ名
	Name string `json:"name"`
	// FileID スタンプ画像ファイルUUID
	FileID string `json:"fileId"`
	// Creator スタンプを作成したユーザー
	Creator User `json:"creator"`
}

// TagAdded TAG_ADDEDイベントペイロード
type TagAdded struct {
	Base
	// TagID タグUUID
	TagID string `json:"tagId"`
	// Tag タグ名
	Tag string `json:"tag"`
}

// TagRemoved TAG_REMOVEDイベントペイロード
type TagRemoved struct {
	Base
	// TagID タグUUID
	TagID string `json:"tagId"`
	// Tag タグ名
	Tag string `json:"tag"`
}

// UserGroupCreated USER_GROUP_CREATEDイベントペイロード
type UserGroupCreated struct {
	Base
	// Group 作成されたグループ
	Group UserGroup `json:"group"`
}

// UserGroupUpdated USER_GROUP_UPDATEDイベントペイロード
type UserGroupUpdated struct {
	Base
	// GroupID 更新されたグループUUID
	GroupID string `json:"groupId"`
}

// UserGroupDeleted USER_GROUP_DELETEDイベントペイロード
type UserGroupDeleted struct {
	Base
	// GroupID 削除されたグループUUID
	GroupID string `json:"groupId"`
}

// UserGroupMemberAdded USER_GROUP_MEMBER_ADDEDイベントペイロード
type UserGroupMemberAdded struct {
	Base
	// GroupMember 追加されたグループメンバー情報
	GroupMember `json:"groupMember"`
}

// UserGroupMemberUpdated USER_GROUP_MEMBER_UPDATEDイベントペイロード
type UserGroupMemberUpdated struct {
	Base
	// GroupMember 更新されたグループメンバー情報
	GroupMember `json:"groupMember"`
}

// UserGroupMemberRemoved USER_GROUP_MEMBER_REMOVEDイベントペイロード
type UserGroupMemberRemoved struct {
	Base
	// GroupMember 削除されたグループメンバー情報
	GroupMember `json:"groupMember"`
}

// UserGroupAdminAdded USER_GROUP_ADMIN_ADDEDイベントペイロード
type UserGroupAdminAdded struct {
	Base
	// GroupMember 追加されたグループ管理者情報
	GroupMember `json:"groupMember"`
}

// UserGroupAdminRemoved USER_GROUP_ADMIN_REMOVEDイベントペイロード
type UserGroupAdminRemoved struct {
	Base
	// GroupMember 削除されたグループ管理者情報
	GroupMember `json:"groupMember"`
}
