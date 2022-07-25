package types

type Audio struct {
	File
	Duration  int       `json:"duration"`
	Performer string    `json:"performer"`
	Title     string    `json:"title"`
	FileName  string    `json:"file_name"`
	MimeType  string    `json:"mime_type"`
	FileSize  string    `json:"file_size"`
	Thumb     PhotoSize `json:"thumb"`
}
