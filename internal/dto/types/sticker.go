package types

// Sticker
// This object represents a sticker.
// https://core.telegram.org/bots/api#sticker
type Sticker struct {
	BaseFile
	Width            int          `json:"width"`
	Height           int          `json:"height"`
	IsAnimated       bool         `json:"is_animated"`
	IsVideo          bool         `json:"is_video"`
	Thumb            PhotoSize    `json:"thumb"`
	Emoji            string       `json:"emoji"`
	SetName          string       `json:"set_name"`
	PremiumAnimation File         `json:"premium_animation"`
	MaskPosition     MaskPosition `json:"mask_position"`
}
