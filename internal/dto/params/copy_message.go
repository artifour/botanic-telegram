package params

import "github.com/artifour/botanic-telegram/internal/dto/types"

type CopyMessage struct {
	ChatId                   string                `json:"chat_id"`
	FromChatId               string                `json:"from_chat_id"`
	MessageId                int                   `json:"message_id"`
	Caption                  string                `json:"caption,omitempty"`
	ParseMode                bool                  `json:"parse_mode,omitempty"`
	CaptionEntities          []types.MessageEntity `json:"caption_entities,omitempty"`
	DisableNotification      bool                  `json:"disable_notification,omitempty"`
	ProtectContent           bool                  `json:"protect_content,omitempty"`
	ReplyToMessageId         int                   `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                  `json:"allow_sending_without_reply,omitempty"`
	// TODO ReplyMarkup
}
