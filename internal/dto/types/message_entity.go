package types

const (
	MessageEntityTypeMention       = "mention"
	MessageEntityTypeHashtag       = "hashtag"
	MessageEntityTypeCashtag       = "cashtag"
	MessageEntityTypeBotCommand    = "bot_command"
	MessageEntityTypeUrl           = "url"
	MessageEntityTypeEmail         = "email"
	MessageEntityTypePhoneNumber   = "phone_number"
	MessageEntityTypeBold          = "bold"
	MessageEntityTypeItalic        = "italic"
	MessageEntityTypeUnderline     = "underline"
	MessageEntityTypeStrikethrough = "strikethrough"
	MessageEntityTypeSpoiler       = "spoiler"
	MessageEntityTypeCode          = "code"
	MessageEntityTypePre           = "pre"
	MessageEntityTypeTextLink      = "text_link"
	MessageEntityTypeTextMention   = "text_mention"
)

// MessageEntity
// This object represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.
// https://core.telegram.org/bots/api#messageentity
type MessageEntity struct {
	Type     string `json:"type"`
	Offset   int    `json:"offset"`
	Length   int    `json:"length"`
	Url      string `json:"url"`
	User     User   `json:"user"`
	Language string `json:"language"`
}
