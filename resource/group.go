package resource

import "time"

type Group struct {
	Id     string // 群组ID
	Name   string // 群组名称
	Avatar string // 群组头像
}

type GroupMember struct {
	User     *User
	Nickname string
	Avatar   string
	JoinedAt time.Time
}

type GroupRole struct {
	Id   string
	Name string
}
