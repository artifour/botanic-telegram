package types

type PassportFile struct {
	BaseFile
	FileDate int `json:"file_date"`
}
