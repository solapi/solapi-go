package groups

import "github.com/solapi/solapi-go/v2/messages"

// Request/Response types for group APIs

type CreateGroupRequest struct {
	Name string `json:"name,omitempty"`
}

type CreateGroupResponse struct {
	GroupID   string             `json:"groupId"`
	GroupInfo messages.GroupInfo `json:"groupInfo,omitempty"`
}

type AddGroupMessagesRequest struct {
	Messages        []messages.Message `json:"messages"`
	AllowDuplicates *bool              `json:"allowDuplicates,omitempty"`
}

type GroupActionResponse struct {
	GroupInfo         messages.GroupInfo       `json:"groupInfo"`
	FailedMessageList []messages.FailedMessage `json:"failedMessageList"`
}

type ListMessagesQuery struct {
	StartKey string
	Limit    int
}

// Scheduling
type ScheduleRequest struct {
	ScheduledDate string `json:"scheduledDate"`
}

// Remove messages from group
type RemoveGroupMessagesRequest struct {
	MessageIDs []string `json:"messageIds"`
}

// List groups
type ListGroupsQuery struct {
	StartKey string
	Limit    int
}

type ListGroupsResponse struct {
	GroupList map[string]messages.DetailGroupMessageResponse `json:"groupList"`
	Limit     int                                            `json:"limit"`
	StartKey  *string                                        `json:"startKey"`
	NextKey   string                                         `json:"nextKey"`
}

// Single group response
type GroupResponse struct {
	GroupID   string             `json:"groupId"`
	GroupInfo messages.GroupInfo `json:"groupInfo,omitempty"`
}
