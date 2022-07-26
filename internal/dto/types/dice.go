package types

// Dice
// This object represents an animated emoji that displays a random value.
// https://core.telegram.org/bots/api#dice
type Dice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}
