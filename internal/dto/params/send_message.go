package params

import "github.com/artifour/telegramic/internal/dto/types"

type SendMessage struct {
	ChatId                   string                `json:"chat_id"`
	Text                     string                `json:"text"`
	ParseMode                string                `json:"parse_mode,omitempty"`
	Entities                 []types.MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview    bool                  `json:"disable_web_page_preview,omitempty"`
	DisableNotification      bool                  `json:"disable_notification,omitempty"`
	ProtectContent           bool                  `json:"protect_content,omitempty"`
	ReplyToMessageId         int                   `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                  `json:"allow_sending_without_reply,omitempty"`
	// TODO ReplyMarkup
}
