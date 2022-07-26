package types

type Animation struct {
	BaseFile
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Duration int    `json:"duration"`
	Thumb    string `json:"thumb"`
	FileName string `json:"file_name"`
	MimeType string `json:"mime_type"`
}
