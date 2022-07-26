package types

// Video
// This object represents a video file.
// https://core.telegram.org/bots/api#video
type Video struct {
	BaseFile
	Width    int       `json:"width"`
	Height   int       `json:"height"`
	Duration int       `json:"duration"`
	Thumb    PhotoSize `json:"thumb"`
	FileName string    `json:"file_name"`
	MimeType string    `json:"mime_type"`
}
