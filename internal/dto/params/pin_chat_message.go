package params

type PinChatMessage struct {
	ChatId              string `json:"chat_id"`
	MessageId           int    `json:"message_id"`
	DisableNotification bool   `json:"disable_notification,omitempty"`
}
