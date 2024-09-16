package resource

import (
	"gtihub.com/Iceinu-Project/iceinu/elements"
	"time"
)

type Message struct {
	Id              string
	Content         string
	Channel         *Channel
	Guild           *Guild
	Member          *GuildMember
	User            *User
	CreatedAt       time.Time
	UpdatedAt       time.Time
	MessageElements *[]elements.IceinuMessageElement
}
