package groups

import (
	"encoding/json"

	"github.com/solapi/solapi-go/v2/messages"
)

// Request/Response types for group APIs

type CreateGroupRequest struct {
	Name string `json:"name,omitempty"`
}

// CreateGroupResponse shares GroupResponse's shape and decoding.
type CreateGroupResponse = GroupResponse

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

// UnmarshalJSON accepts both {groupId, groupInfo{...}} and the flat group document
// returned by GET /messages/v4/groups/{groupId}, filling GroupInfo from either.
func (g *GroupResponse) UnmarshalJSON(data []byte) error {
	type wrapped struct {
		GroupID   string              `json:"groupId"`
		GroupInfo *messages.GroupInfo `json:"groupInfo"`
	}
	var w wrapped
	if err := json.Unmarshal(data, &w); err != nil {
		return err
	}
	g.GroupID = w.GroupID
	if w.GroupInfo != nil {
		g.GroupInfo = *w.GroupInfo
		return nil
	}
	var flat messages.GroupInfo
	if err := json.Unmarshal(data, &flat); err != nil {
		return err
	}
	g.GroupInfo = flat
	if g.GroupID == "" {
		g.GroupID = flat.GroupID
	}
	return nil
}
