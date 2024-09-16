package resource

import "time"

type Guild struct {
	Id     string // 群组ID
	Name   string // 群组名称
	Avatar string // 群组头像
}

type GuildMember struct {
	User     *User
	Nickname string
	Avatar   string
	JoinedAt time.Time
}

type GuildRole struct {
	Id   string
	Name string
}
