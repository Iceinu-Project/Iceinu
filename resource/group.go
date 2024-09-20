package resource

import "time"

type Group struct {
	Id          string // 群组ID
	Name        string // 群组名称
	Avatar      string // 群组头像
	Maxcount    uint32 // 群组最大人数
	MemberCount uint32 // 群组成员数量
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
