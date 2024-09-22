package satori

import "time"

// EventSatori 基于Satori标准设计的事件接收结构体，出于方便使用进行了部分修改和拓展
type EventSatori struct {
	EventId   uint64       // 事件ID
	EventType string       // 事件类型
	Platform  string       // 接收者平台名称
	SelfId    string       // 接收者平台账号
	Timestamp int64        // 事件推送的时间戳
	Argv      *Argv        // 交互指令
	Button    *Button      // 交互按钮
	Channel   *Channel     // 事件所属的频道
	Group     *Group       // 事件所属的群组
	Login     *Login       // 事件的登录信息
	Member    *GroupMember // 事件的目标成员
	Message   *Message     // 事件的消息
	Operator  *User        // 事件的操作者
	Role      *GroupRole   // 事件的目标角色
	User      *User        // 事件的目标用户
}

// Channel 频道结构体
type Channel struct {
	Id       string // 频道ID
	Type     uint8  // 频道类型, 0: 文本频道, 1: 私聊频道，2：分类频道，3：语音频道
	Name     string // 频道名称
	ParentId string // 父频道ID
}

// Group 群组结构体
type Group struct {
	Id          string // 群组ID
	Name        string // 群组名称
	Avatar      string // 群组头像
	Maxcount    uint32 // 群组最大人数
	MemberCount uint32 // 群组成员数量
}

// GroupMember 群组成员结构体
type GroupMember struct {
	User     *User
	Nickname string
	Avatar   string
	JoinedAt time.Time
}

// GroupRole 群组角色结构体
type GroupRole struct {
	Id   string
	Name string
}

// Argv 交互指令结构体
type Argv struct {
	Name      string
	Arguments []interface{}
	Options   interface{}
}

// Button 交互按钮结构体
type Button struct {
	Id string
}

// Login 登录信息结构体
type Login struct {
	User      *User    // 用户对象
	SelfId    string   // 平台账号
	Platform  string   // 平台名称
	Status    uint8    // 登录状态，0为离线，1为在线，2为连接中，3为断开连接，4为重新连接
	Features  []string // 平台特性列表
	ProxyUrls []string // 代理路由列表
}

// Message 消息结构体
type Message struct {
	Id              string
	Content         string
	Channel         *Channel
	Group           *Group
	Member          *GroupMember
	User            *User
	CreatedAt       time.Time
	UpdatedAt       time.Time
	MessageElements *[]ElementSatori
}

// PagedList 分页列表结构体
type PagedList struct {
	Data interface{}
	Next string
}

// User 用户结构体
type User struct {
	Id       string // 用户ID
	Name     string // 用户名称
	Nickname string // 用户昵称
	Avatar   string // 用户头像
	IsBot    bool   // 是否是机器人
}
