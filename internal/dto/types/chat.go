package telegramic

type chat struct {
	Id int `json:"id"`
	Type string `json:"type"`
	Title string `json:"title"`
	Username string `json:"username"`
	FirstName string `json:"first_name"`
	LastName string    `json:"last_name"`
	Photo    chatPhoto `json:"photo"`
	Bio      string    `json:"bio"`
	HasPrivatForwards bool `json:"has_privat_forwards"`
	JoinToSendMessages bool `json:"join_to_send_messages"`
	JoinByRequest bool `json:"join_by_request"`
	Description string `json:"description"`
	InviteLink string `json:"invite_link"`
	PinnedMessage string `json:"pinned_message"`
}
