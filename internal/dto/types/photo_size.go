package types

type PhotoSize struct {
	File
	Width    int `json:"width"`
	Height   int `json:"height"`
	FileSize int `json:"file_size"`
}
