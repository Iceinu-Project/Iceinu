package lagrange

import (
	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/client/event"
	"github.com/LagrangeDev/LagrangeGo/message"
)

// 定义各个消息类型的处理函数

type PrivateMessageHandler func(client *client.QQClient, message *message.PrivateMessage)
type GroupMessageHandler func(client *client.QQClient, message *message.GroupMessage)
type TempMessageHandler func(client *client.QQClient, message *message.TempMessage)

type SelfPrivateMessageHandler func(client *client.QQClient, message *message.PrivateMessage)
type SelfGroupMessageHandler func(client *client.QQClient, message *message.GroupMessage)
type SelfTempMessageHandler func(client *client.QQClient, message *message.TempMessage)

// 定义各个事件类型的处理函数

type GroupJoinEventHandler func(client *client.QQClient, event *event.GroupMemberIncrease)  // bot进群
type GroupLeaveEventHandler func(client *client.QQClient, event *event.GroupMemberDecrease) // bot退群

type GroupInvitedEventHandler func(client *client.QQClient, event *event.GroupInvite)                      // bot被邀请进群
type GroupMemberJoinRequestEventHandler func(client *client.QQClient, event *event.GroupMemberJoinRequest) // 加群邀请
type GroupMemberJoinEventHandler func(client *client.QQClient, event *event.GroupMemberIncrease)           // 成员加群
type GroupMemberLeaveEventHandler func(client *client.QQClient, event *event.GroupMemberDecrease)          // 成员退群
type GroupMuteEventHandler func(client *client.QQClient, event *event.GroupMute)                           // 禁言事件
type GroupRecallEventHandler func(client *client.QQClient, event *event.GroupRecall)
type GroupMemberPermissionChangedEventHandler func(client *client.QQClient, event *event.GroupMemberPermissionChanged)
type GroupNameUpdatedEventHandler func(client *client.QQClient, event *event.GroupNameUpdated)                   // 群更名
type MemberSpecialTitleUpdatedEventHandler func(client *client.QQClient, event *event.MemberSpecialTitleUpdated) // 群成员特殊头衔更新
type NewFriendRequestHandler func(client *client.QQClient, event *event.NewFriendRequest)                        // 新朋友请求
type FriendRecallEventHandler func(client *client.QQClient, event *event.FriendRecall)                           // 好友消息撤回
type RenameEventHandler func(client *client.QQClient, event *event.Rename)                                       // 好友昵称更改
type FriendNotifyEventHandler func(client *client.QQClient, event event.INotifyEvent)                            // 好友通知
type GroupNotifyEventHandler func(client *client.QQClient, event event.INotifyEvent)                             // 群通知

// SubscribeManager 订阅管理器实现，用于向LagrangeGo注册各类消息/事件处理函数
type SubscribeManager struct {
	privateMessageHandlers                    []PrivateMessageHandler
	groupMessageHandlers                      []GroupMessageHandler
	tempMessageHandlers                       []TempMessageHandler
	selfPrivateMessageHandlers                []SelfPrivateMessageHandler
	selfGroupMessageHandlers                  []SelfGroupMessageHandler
	selfTempMessageHandlers                   []SelfTempMessageHandler
	groupJoinEventHandlers                    []GroupJoinEventHandler
	groupLeaveEventHandlers                   []GroupLeaveEventHandler
	groupInviteEventHandlers                  []GroupInvitedEventHandler
	groupMemberJoinRequestEventHandlers       []GroupMemberJoinRequestEventHandler
	groupMemberJoinEventHandlers              []GroupMemberJoinEventHandler
	groupMemberLeaveEventHandlers             []GroupMemberLeaveEventHandler
	groupMuteEventHandlers                    []GroupMuteEventHandler
	groupRecallEventHandlers                  []GroupRecallEventHandler
	groupMemberPermissionChangedEventHandlers []GroupMemberPermissionChangedEventHandler
	groupNameUpdatedEventHandlers             []GroupNameUpdatedEventHandler
	memberSpecialTitleUpdatedEventHandlers    []MemberSpecialTitleUpdatedEventHandler
	newFriendRequestHandlers                  []NewFriendRequestHandler
	friendRecallEventHandlers                 []FriendRecallEventHandler
	renameEventHandlers                       []RenameEventHandler
	friendNotifyEventHandlers                 []FriendNotifyEventHandler
	groupNotifyEventHandlers                  []GroupNotifyEventHandler
}

// Manager 全局 SubscribeManager 实例
var Manager = &SubscribeManager{}

// RegisterPrivateMessageHandler 注册私聊消息处理函数
func (sm *SubscribeManager) RegisterPrivateMessageHandler(handler PrivateMessageHandler) {
	sm.privateMessageHandlers = append(sm.privateMessageHandlers, handler)
}

// RegisterGroupMessageHandler 注册群消息处理函数
func (sm *SubscribeManager) RegisterGroupMessageHandler(handler GroupMessageHandler) {
	sm.groupMessageHandlers = append(sm.groupMessageHandlers, handler)
}

// RegisterTempMessageHandler 注册临时消息处理函数
func (sm *SubscribeManager) RegisterTempMessageHandler(handler TempMessageHandler) {
	sm.tempMessageHandlers = append(sm.tempMessageHandlers, handler)
}

// RegisterSelfPrivateMessageHandler 注册自己的私聊消息处理函数
func (sm *SubscribeManager) RegisterSelfPrivateMessageHandler(handler SelfPrivateMessageHandler) {
	sm.selfPrivateMessageHandlers = append(sm.selfPrivateMessageHandlers, handler)
}

// RegisterSelfGroupMessageHandler 注册自己的群消息处理函数
func (sm *SubscribeManager) RegisterSelfGroupMessageHandler(handler SelfGroupMessageHandler) {
	sm.selfGroupMessageHandlers = append(sm.selfGroupMessageHandlers, handler)
}

// RegisterSelfTempMessageHandler 注册自己的临时消息处理函数
func (sm *SubscribeManager) RegisterSelfTempMessageHandler(handler SelfTempMessageHandler) {
	sm.selfTempMessageHandlers = append(sm.selfTempMessageHandlers, handler)
}

// RegisterGroupJoinEventHandler 注册群成员加入事件处理函数
func (sm *SubscribeManager) RegisterGroupJoinEventHandler(handler GroupJoinEventHandler) {
	sm.groupJoinEventHandlers = append(sm.groupJoinEventHandlers, handler)
}

// RegisterGroupLeaveEventHandler 注册群成员离开事件处理函数
func (sm *SubscribeManager) RegisterGroupLeaveEventHandler(handler GroupLeaveEventHandler) {
	sm.groupLeaveEventHandlers = append(sm.groupLeaveEventHandlers, handler)
}

// RegisterGroupInviteEventHandler 注册群邀请事件处理函数
func (sm *SubscribeManager) RegisterGroupInviteEventHandler(handler GroupInvitedEventHandler) {
	sm.groupInviteEventHandlers = append(sm.groupInviteEventHandlers, handler)

}

// RegisterGroupMemberJoinRequestEventHandler 注册群成员加群邀请事件处理函数
func (sm *SubscribeManager) RegisterGroupMemberJoinRequestEventHandler(handler GroupMemberJoinRequestEventHandler) {
	sm.groupMemberJoinRequestEventHandlers = append(sm.groupMemberJoinRequestEventHandlers, handler)
}

// RegisterGroupMemberJoinEventHandler 注册群成员加入事件处理函数
func (sm *SubscribeManager) RegisterGroupMemberJoinEventHandler(handler GroupMemberJoinEventHandler) {
	sm.groupMemberJoinEventHandlers = append(sm.groupMemberJoinEventHandlers, handler)
}

// RegisterGroupMemberLeaveEventHandler 注册群成员离开事件处理函数
func (sm *SubscribeManager) RegisterGroupMemberLeaveEventHandler(handler GroupMemberLeaveEventHandler) {
	sm.groupMemberLeaveEventHandlers = append(sm.groupMemberLeaveEventHandlers, handler)
}

// RegisterGroupMuteEventHandler 注册群成员禁言事件处理函数
func (sm *SubscribeManager) RegisterGroupMuteEventHandler(handler GroupMuteEventHandler) {
	sm.groupMuteEventHandlers = append(sm.groupMuteEventHandlers, handler)
}

// RegisterGroupRecallEventHandler 注册群消息撤回事件处理函数
func (sm *SubscribeManager) RegisterGroupRecallEventHandler(handler GroupRecallEventHandler) {
	sm.groupRecallEventHandlers = append(sm.groupRecallEventHandlers, handler)
}

// RegisterGroupMemberPermissionChangedEventHandler 注册群成员权限变更事件处理函数
func (sm *SubscribeManager) RegisterGroupMemberPermissionChangedEventHandler(handler GroupMemberPermissionChangedEventHandler) {
	sm.groupMemberPermissionChangedEventHandlers = append(sm.groupMemberPermissionChangedEventHandlers, handler)
}

// RegisterGroupNameUpdatedEventHandler 注册群名称变更事件处理函数
func (sm *SubscribeManager) RegisterGroupNameUpdatedEventHandler(handler GroupNameUpdatedEventHandler) {
	sm.groupNameUpdatedEventHandlers = append(sm.groupNameUpdatedEventHandlers, handler)
}

// RegisterMemberSpecialTitleUpdatedEventHandler 注册群成员特殊头衔变更事件处理函数
func (sm *SubscribeManager) RegisterMemberSpecialTitleUpdatedEventHandler(handler MemberSpecialTitleUpdatedEventHandler) {
	sm.memberSpecialTitleUpdatedEventHandlers = append(sm.memberSpecialTitleUpdatedEventHandlers, handler)
}

// RegisterNewFriendRequestHandler 注册新朋友请求处理函数
func (sm *SubscribeManager) RegisterNewFriendRequestHandler(handler NewFriendRequestHandler) {
	sm.newFriendRequestHandlers = append(sm.newFriendRequestHandlers, handler)
}

// RegisterFriendRecallEventHandler 注册好友消息撤回事件处理函数
func (sm *SubscribeManager) RegisterFriendRecallEventHandler(handler FriendRecallEventHandler) {
	sm.friendRecallEventHandlers = append(sm.friendRecallEventHandlers, handler)
}

// RegisterRenameEventHandler 注册好友昵称更改事件处理函数
func (sm *SubscribeManager) RegisterRenameEventHandler(handler RenameEventHandler) {
	sm.renameEventHandlers = append(sm.renameEventHandlers, handler)
}

// RegisterFriendNotifyEventHandler 注册好友通知事件处理函数
func (sm *SubscribeManager) RegisterFriendNotifyEventHandler(handler FriendNotifyEventHandler) {
	sm.friendNotifyEventHandlers = append(sm.friendNotifyEventHandlers, handler)
}

// RegisterGroupNotifyEventHandler 注册群通知事件处理函数
func (sm *SubscribeManager) RegisterGroupNotifyEventHandler(handler GroupNotifyEventHandler) {
	sm.groupNotifyEventHandlers = append(sm.groupNotifyEventHandlers, handler)
}

// SetAllSubscribes 设置所有订阅处理
func SetAllSubscribes() {
	Client.PrivateMessageEvent.Subscribe(func(client *client.QQClient, event *message.PrivateMessage) {
		for _, handler := range Manager.privateMessageHandlers {
			handler(client, event)
		}
	})

	Client.GroupMessageEvent.Subscribe(func(client *client.QQClient, event *message.GroupMessage) {
		for _, handler := range Manager.groupMessageHandlers {
			handler(client, event)
		}
	})

	Client.TempMessageEvent.Subscribe(func(client *client.QQClient, event *message.TempMessage) {
		for _, handler := range Manager.tempMessageHandlers {
			handler(client, event)
		}
	})

	Client.SelfPrivateMessageEvent.Subscribe(func(client *client.QQClient, event *message.PrivateMessage) {
		for _, handler := range Manager.selfPrivateMessageHandlers {
			handler(client, event)
		}
	})

	Client.SelfGroupMessageEvent.Subscribe(func(client *client.QQClient, event *message.GroupMessage) {
		for _, handler := range Manager.selfGroupMessageHandlers {
			handler(client, event)
		}
	})

	Client.SelfTempMessageEvent.Subscribe(func(client *client.QQClient, event *message.TempMessage) {
		for _, handler := range Manager.selfTempMessageHandlers {
			handler(client, event)
		}
	})

	Client.GroupJoinEvent.Subscribe(func(client *client.QQClient, event *event.GroupMemberIncrease) {
		for _, handler := range Manager.groupJoinEventHandlers {
			handler(client, event)
		}
	})

	Client.GroupLeaveEvent.Subscribe(func(client *client.QQClient, event *event.GroupMemberDecrease) {
		for _, handler := range Manager.groupLeaveEventHandlers {
			handler(client, event)
		}
	})

	Client.GroupInvitedEvent.Subscribe(func(client *client.QQClient, event *event.GroupInvite) {
		for _, handler := range Manager.groupInviteEventHandlers {
			handler(client, event)
		}
	})

	Client.GroupMemberJoinRequestEvent.Subscribe(func(client *client.QQClient, event *event.GroupMemberJoinRequest) {
		for _, handler := range Manager.groupMemberJoinRequestEventHandlers {
			handler(client, event)
		}
	})

	Client.GroupMemberJoinEvent.Subscribe(func(client *client.QQClient, event *event.GroupMemberIncrease) {
		for _, handler := range Manager.groupMemberJoinEventHandlers {
			handler(client, event)
		}
	})

	Client.GroupMemberLeaveEvent.Subscribe(func(client *client.QQClient, event *event.GroupMemberDecrease) {
		for _, handler := range Manager.groupMemberLeaveEventHandlers {
			handler(client, event)
		}
	})

	Client.GroupMuteEvent.Subscribe(func(client *client.QQClient, event *event.GroupMute) {
		for _, handler := range Manager.groupMuteEventHandlers {
			handler(client, event)
		}
	})

	Client.GroupRecallEvent.Subscribe(func(client *client.QQClient, event *event.GroupRecall) {
		for _, handler := range Manager.groupRecallEventHandlers {
			handler(client, event)
		}
	})

	Client.GroupMemberPermissionChangedEvent.Subscribe(func(client *client.QQClient, event *event.GroupMemberPermissionChanged) {
		for _, handler := range Manager.groupMemberPermissionChangedEventHandlers {
			handler(client, event)
		}
	})

	Client.GroupNameUpdatedEvent.Subscribe(func(client *client.QQClient, event *event.GroupNameUpdated) {
		for _, handler := range Manager.groupNameUpdatedEventHandlers {
			handler(client, event)
		}
	})

	Client.MemberSpecialTitleUpdatedEvent.Subscribe(func(client *client.QQClient, event *event.MemberSpecialTitleUpdated) {
		for _, handler := range Manager.memberSpecialTitleUpdatedEventHandlers {
			handler(client, event)
		}
	})

	Client.NewFriendRequestEvent.Subscribe(func(client *client.QQClient, event *event.NewFriendRequest) {
		for _, handler := range Manager.newFriendRequestHandlers {
			handler(client, event)
		}
	})

	Client.FriendRecallEvent.Subscribe(func(client *client.QQClient, event *event.FriendRecall) {
		for _, handler := range Manager.friendRecallEventHandlers {
			handler(client, event)
		}
	})

	Client.RenameEvent.Subscribe(func(client *client.QQClient, event *event.Rename) {
		for _, handler := range Manager.renameEventHandlers {
			handler(client, event)
		}
	})

	Client.FriendNotifyEvent.Subscribe(func(client *client.QQClient, event event.INotifyEvent) {
		for _, handler := range Manager.friendNotifyEventHandlers {
			handler(client, event)
		}
	})

	Client.GroupNotifyEvent.Subscribe(func(client *client.QQClient, event event.INotifyEvent) {
		for _, handler := range Manager.groupNotifyEventHandlers {
			handler(client, event)
		}
	})
}
