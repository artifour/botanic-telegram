package telegram

const (
	pollTypeRegular = "regular"
	pollTypeQuiz    = "quiz"
)

type poll struct {
	Id                    string       `json:"id"`
	Question              string       `json:"question"`
	Options               []pollOption `json:"options"`
	TotalVoterCount       int          `json:"total_voter_count"`
	IsClosed              bool         `json:"is_closed"`
	IsAnonymous           bool         `json:"is_anonymous"`
	Type                  string       `json:"type"`
	AllowsMultipleAnswers bool         `json:"allows_multiple_answers"`
	CorrectOptionId       int          `json:"correct_option_id"`
	Explanation           string       `json:"explanation"`
	// TODO ExplanationEntities
	OpenPeriod int `json:"open_period"`
	CloseDate  int `json:"close_date"`
}
