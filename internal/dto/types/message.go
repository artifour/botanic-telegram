package telegramic

type message struct {
	MessageId int `json:"message_id"`
	From user `json:"from"`
	SenderChat chat `json:"sender_chat"`
	Date int `json:"date"`
	Chat chat `json:"chat"`
	ForwardFrom user `json:"forward_from"`
	ForwardFromChat chat `json:"forward_from_chat"`
	ForwardFromMessageId int `json:"forward_from_message_id"`
	ForwardSignature string `json:"forward_signature"`
	ForwardSenderName string `json:"forward_sender_name"`
	ForwardDate int `json:"forward_date"`
	IsAutomaticForward bool `json:"is_automatic_forward"`
	ReplyToMessage *message `json:"reply_to_message"`
	ViaBot user `json:"via_bot"`
	EditDate int `json:"edit_date"`
	HasProtectedContent bool `json:"has_protected_content"`
	MediaGroupId int `json:"media_group_id"`
	AuthorSignature string `json:"author_signature"`
	Text string              `json:"text"`
	Entities []messageEntity `json:"entities"`
	Animation animation      `json:"animation"`

}
