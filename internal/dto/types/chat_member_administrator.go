package types

// ChatMemberAdministrator
// Represents a chat member that has some additional privileges.
// https://core.telegram.org/bots/api#chatmemberadministrator
type ChatMemberAdministrator struct {
	BaseChatMember
	CanBeEdited         bool   `json:"can_be_edited"`
	IsAnonymous         bool   `json:"is_anonymous"`
	CanManageChat       bool   `json:"can_manage_chat"`
	CanDeleteMessages   bool   `json:"can_delete_messages"`
	CanManageVideoChats bool   `json:"can_manage_video_chats"`
	CanRestrictMembers  bool   `json:"can_restrict_members"`
	CanPromoteMembers   bool   `json:"can_promote_members"`
	CanChangeInfo       bool   `json:"can_change_info"`
	CanInviteUsers      bool   `json:"can_invite_users"`
	CanPostMessages     bool   `json:"can_post_messages"`
	CanEditMessages     bool   `json:"can_edit_messages"`
	CanPinMessages      bool   `json:"can_pin_messages"`
	CustomTitle         string `json:"custom_title"`
}
