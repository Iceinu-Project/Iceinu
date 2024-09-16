package resource

import "time"

type Message struct {
	Id        string
	Content   string
	Channel   *Channel
	Guild     *Guild
	Member    *GuildMember
	User      *User
	CreatedAt time.Time
	UpdatedAt time.Time
}
