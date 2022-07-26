package types

type PhotoSize struct {
	BaseFile
	Width  int `json:"width"`
	Height int `json:"height"`
}
