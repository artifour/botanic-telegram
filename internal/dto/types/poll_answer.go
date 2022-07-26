package types

// PollAnswer
// This object represents an answer of a user in a non-anonymous poll.
// https://core.telegram.org/bots/api#pollanswer
type PollAnswer struct {
	PollId    string `json:"poll_id"`
	User      User   `json:"user"`
	OptionIds []int  `json:"option_ids"`
}
