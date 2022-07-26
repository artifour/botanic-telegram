package types

type BaseChatMember struct {
	Status string `json:"status"`
	User   User   `json:"user"`
}
