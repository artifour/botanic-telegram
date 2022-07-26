package types

type File struct {
	BaseFile
	FilePath string `json:"file_path"`
}
