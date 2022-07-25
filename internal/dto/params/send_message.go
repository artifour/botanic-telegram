package telegramic

type sendMessage struct {
	ChatId string `json:"chat_id"`
	Text string `json:"text"`
	ParseMode string `json:"parse_mode,omitempty"`
	Entities []messageEntity `json:"entities,omitempty"`
}
