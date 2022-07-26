package types

// Document
// This object represents a general file (as opposed to photos, voice messages and audio files).
// https://core.telegram.org/bots/api#document
type Document struct {
	BaseFile
	Thumb    PhotoSize `json:"thumb"`
	FileName string    `json:"file_name"`
	MimeType string    `json:"mime_type"`
}
