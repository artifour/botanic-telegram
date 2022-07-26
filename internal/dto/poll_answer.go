package telegram

type pollAnswer struct {
	PollId string `json:"poll_id"`
	// TODO user
	OptionIds []int `json:"option_ids"`
}
