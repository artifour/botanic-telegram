package types

// PollOption
// This object contains information about one answer option in a poll.
// https://core.telegram.org/bots/api#polloption
type PollOption struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voterCount"`
}
