package params

type ForwardMessage struct {
	ChatId              string `json:"chat_id"`
	FromChatId          string `json:"from_chat_id"`
	DisableNotification bool   `json:"disable_notification,omitempty"`
	ProtectContent      bool   `json:"protect_content,omitempty"`
	MessageId           int    `json:"message_id"`
}
