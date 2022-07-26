package types

// MaskPosition
// This object describes the position on faces where a mask should be placed by default.
// https://core.telegram.org/bots/api#maskposition
type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float32 `json:"x_shift"`
	YShift float32 `json:"y_shift"`
	Scale  float32 `json:"scale"`
}
