package types

type CallbackGame struct {
	UserId             int  `json:"user_id"`
	Score              int  `json:"score"`
	Force              bool `json:"force"`
	DisableEditMessage bool `json:"disable_edit_message"`
	ChatId             int  `json:"chat_id"`
	MessageId          int  `json:"message_id"`
	InlineMessageId    int  `json:"inline_message_id"`
}
