package telegramic

type animation struct {
	FileId string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	Width int `json:"width"`
	Height int `json:"height"`
	Duration int `json:"duration"`
	Thumb string `json:"thumb"`
	FileName string `json:"file_name"`
	MimeType string `json:"mime_type"`
	FileSize int `json:"file_size"`
}
