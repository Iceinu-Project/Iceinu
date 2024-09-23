package satori

import "time"

// EventSatori 基于Satori标准设计的事件接收结构体，出于方便使用进行了部分修改和拓展
type EventSatori struct {
	Id        uint64       `json:"id,omitempty"`        // 事件ID
	Type      string       `json:"type,omitempty"`      // 事件类型
	Platform  string       `json:"platform,omitempty"`  // 接收者平台名称
	SelfId    string       `json:"selfId,omitempty"`    // 接收者平台账号
	Timestamp int64        `json:"timestamp,omitempty"` // 事件推送的时间戳
	Argv      *Argv        `json:"argv,omitempty"`      // 交互指令
	Button    *Button      `json:"button,omitempty"`    // 交互按钮
	Channel   *Channel     `json:"channel,omitempty"`   // 事件所属的频道
	Group     *Group       `json:"group,omitempty"`     // 事件所属的群组
	Login     *Login       `json:"login,omitempty"`     // 事件的登录信息
	Member    *GroupMember `json:"member,omitempty"`    // 事件的目标成员
	Message   *Message     `json:"message,omitempty"`   // 事件的消息
	Operator  *User        `json:"operator,omitempty"`  // 事件的操作者
	Role      *GroupRole   `json:"role,omitempty"`      // 事件的目标角色
	User      *User        `json:"user,omitempty"`      // 事件的目标用户
}

// Channel 频道结构体
type Channel struct {
	Id       string `json:"id,omitempty"`       // 频道ID
	Type     uint8  `json:"type,omitempty"`     // 频道类型, 0: 文本频道, 1: 私聊频道，2：分类频道，3：语音频道
	Name     string `json:"name,omitempty"`     // 频道名称
	ParentId string `json:"parentId,omitempty"` // 父频道ID
}

// Group 群组结构体
type Group struct {
	Id          string `json:"id,omitempty"`          // 群组ID
	Name        string `json:"name,omitempty"`        // 群组名称
	Avatar      string `json:"avatar,omitempty"`      // 群组头像
	Maxcount    uint32 `json:"maxcount,omitempty"`    // 群组最大人数
	MemberCount uint32 `json:"memberCount,omitempty"` // 群组成员数量
}

// GroupMember 群组成员结构体
type GroupMember struct {
	User     *User     `json:"user,omitempty"`
	Nickname string    `json:"nickname,omitempty"`
	Avatar   string    `json:"avatar,omitempty"`
	JoinedAt time.Time `json:"joinedAt"`
}

// GroupRole 群组角色结构体
type GroupRole struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Argv 交互指令结构体
type Argv struct {
	Name      string        `json:"name,omitempty"`
	Arguments []interface{} `json:"arguments,omitempty"`
	Options   interface{}   `json:"options,omitempty"`
}

// Button 交互按钮结构体
type Button struct {
	Id string `json:"id,omitempty"`
}

// Login 登录信息结构体
type Login struct {
	User      *User    `json:"user,omitempty"`      // 用户对象
	SelfId    string   `json:"selfId,omitempty"`    // 平台账号
	Platform  string   `json:"platform,omitempty"`  // 平台名称
	Status    uint8    `json:"status,omitempty"`    // 登录状态，0为离线，1为在线，2为连接中，3为断开连接，4为重新连接
	Features  []string `json:"features,omitempty"`  // 平台特性列表
	ProxyUrls []string `json:"proxyUrls,omitempty"` // 代理路由列表
}

// Message 消息结构体
type Message struct {
	Id              string           `json:"id,omitempty"`
	Content         string           `json:"content,omitempty"`
	Channel         *Channel         `json:"channel,omitempty"`
	Group           *Group           `json:"group,omitempty"`
	Member          *GroupMember     `json:"member,omitempty"`
	User            *User            `json:"user,omitempty"`
	CreatedAt       time.Time        `json:"createdAt"`
	UpdatedAt       time.Time        `json:"updatedAt"`
	MessageElements *[]ElementSatori `json:"messageElements,omitempty"`
}

// PagedList 分页列表结构体
type PagedList struct {
	Data interface{} `json:"data,omitempty"`
	Next string      `json:"next,omitempty"`
}

// User 用户结构体
type User struct {
	Id       string `json:"id,omitempty"`       // 用户ID
	Name     string `json:"name,omitempty"`     // 用户名称
	Nickname string `json:"nickname,omitempty"` // 用户昵称
	Avatar   string `json:"avatar,omitempty"`   // 用户头像
	IsBot    bool   `json:"isBot,omitempty"`    // 是否是机器人
}
