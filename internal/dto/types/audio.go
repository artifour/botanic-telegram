package types

// Audio
// This object represents an audio file to be treated as music by the Telegram clients.
// https://core.telegram.org/bots/api#audio
type Audio struct {
	BaseFile
	Duration  int       `json:"duration"`
	Performer string    `json:"performer"`
	Title     string    `json:"title"`
	FileName  string    `json:"file_name"`
	MimeType  string    `json:"mime_type"`
	Thumb     PhotoSize `json:"thumb"`
}
