package telegram

type videoNote struct {
	file
	Length   int       `json:"length"`
	Duration int       `json:"duration"`
	Thumb    photoSize `json:"thumb"`
	FileSize int       `json:"file_size"`
}
