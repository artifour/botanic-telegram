package types

// InlineKeyboardMarkup
// This object represents an inline keyboard that appears right next to the message it belongs to.
// https://core.telegram.org/bots/api#inlinekeyboardmarkup
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}
