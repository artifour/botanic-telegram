package types

// VideoNote
// This object represents a video message (available in Telegram apps as of v.4.0).
// https://core.telegram.org/bots/api#videonote
type VideoNote struct {
	BaseFile
	Length   int       `json:"length"`
	Duration int       `json:"duration"`
	Thumb    PhotoSize `json:"thumb"`
	FileSize int       `json:"file_size"`
}
