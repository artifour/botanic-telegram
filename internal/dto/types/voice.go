package types

// Voice
// This object represents a voice note.
// https://core.telegram.org/bots/api#voice
type Voice struct {
	BaseFile
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type"`
}
