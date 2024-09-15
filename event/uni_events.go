package event

import "time"

// Iceinu的通用事件相当一大部分是参考了LagrangeGo的事件设计
// https://github.com/LagrangeDev/LagrangeGo
// 更准确一点来讲，是对LagrangeGo的各个消息事件进行了统一和简化，方便进行开发
// 应该说Iceinu本身从一开始就更贴近于对LagrangeGo的封装和简化，对其他适配器的支持是用来以防万一的XD

// UniMessageEvent 通用消息事件，包含了群聊消息、私聊消息和临时会话消息等事件类型的封装
type UniMessageEvent struct {
	MsgId             uint32        // 消息Id
	MsgTime           time.Time     // 消息时间戳
	InternalId        uint32        // 内部Id
	ClientSeq         uint32        // 客户端序列号
	UserId            uint32        // 用户Id，指消息的发送方
	UserName          string        // 用户名称
	UserNickname      string        // 用户昵称
	AnonymousId       string        // 匿名名称
	AnonymousNickname string        // 匿名昵称
	CardName          string        // 卡片名称
	IsFriend          bool          // 是否为好友
	EnvironmentId     uint32        // 环境Id，指消息的接收方，在私聊中是对方的Id，在群聊中是群聊的Id
	EnvironmentName   string        // 环境名称
	Elements          []interface{} // 消息元素列表
	OriginalObject    interface{}   // 仅供LagrangeGo的初始对象存储
}

// ---好友相关事件---

// NewFriendRequestEvent 新的好友请求事件
type NewFriendRequestEvent struct {
	UserId       uint32 // 用户Id
	UserName     string // 用户名
	UserNickname string // 用户昵称
	RequestMsg   string // 请求信息
	Source       string // 请求来源
}

// FriendRecallEvent 好友撤回消息事件
type FriendRecallEvent struct {
	UserId    uint32    // 用户Id
	UserName  string    // 用户名称
	Sequence  uint64    // 序列号
	Timestamp time.Time // 时间戳
	Random    uint32    // 随机数（？）
}

// FriendRenameEvent 好友改名修改事件
type FriendRenameEvent struct {
	Type         uint32 // 类型，0为自身，1为好友
	UserId       uint32 // 用户Id
	UserName     string // 用户名称
	UserNickname string // 用户昵称
}

// FriendPokeEvent 好友戳一戳事件
type FriendPokeEvent struct {
	SenderId   uint32 // 发送者Id
	ReceiverId uint64 // 接收者Id
	Suffix     string // 后缀
	Action     string // 行为
}

// ---群聊相关事件---

// GroupEvent 群聊事件(暂时不明用法）
type GroupEvent struct {
	GroupId uint32 // 群聊Id
}

// GroupMemberPermissionChangedEvent 群成员权限变更事件
type GroupMemberPermissionChangedEvent struct {
	GroupEvent
	UserId   uint32 // 权限变动用户Id
	UserName string // 权限变动用户名
	IsAdmin  bool   // 是否成为群管理员
}

// GroupNameUpdatedEvent 群名修改事件
type GroupNameUpdatedEvent struct {
	GroupId  uint32 // 群聊Id
	NewName  string // 新群名
	UserId   uint32 // 操作者用户Id
	UserName string // 操作者用户名
}

// GroupMuteEvent 群禁言事件
type GroupMuteEvent struct {
	GroupEvent
	UserId     uint32 // 操作者用户Id
	UserName   string // 操作者用户名
	TargetId   uint32 // 目标用户Id
	TargetName string // 目标用户名
	Duration   uint32 // 持续时间
}

// GroupRecallEvent 群消息撤回事件
type GroupRecallEvent struct {
	GroupEvent
	UserId     uint32    // 操作者用户Id
	UserName   uint32    // 操作者用户名称
	AuthorId   uint32    // 原消息发送者Id
	AuthorName uint32    // 原消息发送者名称
	Sequence   uint64    // 序列号
	Timestamp  time.Time // 时间戳
	Random     uint32    // 随机数（？）
}

// GroupMemberJoinRequestEvent 群成员加群请求事件
type GroupMemberJoinRequestEvent struct {
	GroupEvent
	UserId       uint32 // 用户Id
	UserName     string // 用户名称
	UserNickname string // 用户昵称
	InvitorId    uint32 // 邀请者Id
	InvitorName  string // 邀请者名称
	Answer       string // 问题/答案
	RequestSeq   uint64 // 加群请求序列号
}

// GroupMemberIncreaseEvent 群成员增加事件
type GroupMemberIncreaseEvent struct {
	GroupEvent
	UserId      uint32 // 用户Id
	UserName    string // 用户名
	InvitorId   uint32 // 邀请者Id
	InvitorName string // 邀请者名称
	JoinType    uint32 // 加入方式
}

// GroupMemberDecreaseEvent 群成员减少事件
type GroupMemberDecreaseEvent struct {
	GroupEvent
	UserId       uint32 // 用户Id
	UserName     string // 用户名称
	OperatorId   uint32 // 操作者Id
	OperatorName string // 操作者名称
	ExitType     uint32 // 退出方式
}

// GroupDigestEvent 精华消息事件
type GroupDigestEvent struct {
	GroupId          uint32    // 群聊Id
	MsgId            uint32    // 消息Id
	InternalMsgId    uint32    // 消息内部Id
	OperationType    uint32    // 操作类型，1为设置精华消息，2为取消精华消息
	Timestamp        time.Time // 时间戳
	SenderId         uint32    // 原消息发送者Id
	SenderNickname   string    // 原消息发送者昵称
	OperatorId       uint32    // 操作者Id
	OperatorNickname string    // 操作者昵称
}

// GroupPokeEvent 群戳一戳事件
type GroupPokeEvent struct {
	GroupId    uint32 // 群聊Id
	SenderId   uint32 // 发送者Id
	ReceiverId uint32 // 接收者Id
	Suffix     string // 后缀
	Action     string // 行为
}

// MemberSpecialTitleUpdatedEvent 群成员头衔更新事件
type MemberSpecialTitleUpdatedEvent struct {
	GroupId uint32 // 群聊Id
	UserId  uint32 // 用户Id
	Title   string // 头衔
}

// GroupInviteEvent 群邀请事件
type GroupInviteEvent struct {
	GroupId         uint32 // 群聊Id
	InvitorId       uint32 // 邀请者Id
	InvitorName     string // 邀请者名称
	InvitorNickname string // 邀请者昵称
	RequestSeq      uint64 // 请求序列号
}
