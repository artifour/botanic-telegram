package types

type Audio struct {
	BaseFile
	Duration  int       `json:"duration"`
	Performer string    `json:"performer"`
	Title     string    `json:"title"`
	FileName  string    `json:"file_name"`
	MimeType  string    `json:"mime_type"`
	Thumb     PhotoSize `json:"thumb"`
}
