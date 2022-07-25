package types

type PassportFile struct {
	File
	FileSize int `json:"file_size"`
	FileDate int `json:"file_date"`
}
