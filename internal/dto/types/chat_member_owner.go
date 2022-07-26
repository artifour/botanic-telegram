package types

// ChatMemberOwner
// Represents a chat member that owns the chat and has all administrator privileges.
// https://core.telegram.org/bots/api#chatmemberowner
type ChatMemberOwner struct {
	BaseChatMember
	IsAnonymous bool   `json:"is_anonymous"`
	CustomTitle string `json:"custom_title"`
}
