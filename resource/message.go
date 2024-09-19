package resource

import (
	"github.com/Iceinu-Project/iceinu/elements"
	"time"
)

type Message struct {
	Id              string
	Content         string
	Channel         *Channel
	Group           *Group
	Member          *GroupMember
	User            *User
	CreatedAt       time.Time
	UpdatedAt       time.Time
	MessageElements *[]elements.IceinuMessageElement
}
