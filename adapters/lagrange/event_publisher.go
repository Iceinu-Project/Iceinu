package lagrange

import (
	"github.com/Iceinu-Project/Iceinu/ice"
	"github.com/Iceinu-Project/Iceinu/models/satori"
	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/client/event"
	"github.com/LagrangeDev/LagrangeGo/message"
	"strconv"
	"time"
)

func BindEvents() {
	// 注册私聊事件
	Manager.RegisterPrivateMessageHandler(func(client *client.QQClient, event *message.PrivateMessage) {
		// 尝试从适配器的缓存中获取自身信息
		selfInfo := GetSelfInfoInCache(client)
		// 尝试从适配器的缓存中获取好友信息
		friendInfo := GetFriendsDataInCache(client)
		ice.Publish(&ice.IceinuEvent{
			Type:      10,
			From:      ice.GetSelfNodeId(),
			Target:    ice.GetMasterNodeId(),
			Timestamp: time.Now().Unix(),
			Summary:   "PrivateMessageEvent",
			Event: &satori.EventSatori{
				Id:        ice.GenerateEventID(),
				Type:      "PrivateMessageEvent",
				Platform:  "QQNT",
				SelfId:    strconv.Itoa(int(client.Uin)),
				Timestamp: int64(event.Time),
				Argv:      nil,
				Button:    nil,
				Channel: &satori.Channel{
					Id:       strconv.Itoa(int(event.Sender.Uin)),
					Type:     1,
					Name:     event.Sender.Uid,
					ParentId: "",
				},
				Group: &satori.Group{
					Id:          "",
					Name:        "",
					Avatar:      "",
					Maxcount:    0,
					MemberCount: 0,
				},
				Login:  nil,
				Member: nil,
				Message: &satori.Message{
					Id:              strconv.Itoa(int(event.InternalId)),
					Content:         event.ToString(),
					Channel:         nil,
					Group:           nil,
					Member:          nil,
					User:            nil,
					CreatedAt:       int64(event.Time),
					UpdatedAt:       int64(event.Time),
					MessageElements: ToSatoriElements(event.Elements),
				},
				Operator: &satori.User{
					Id:       strconv.Itoa(int(event.Sender.Uin)),
					Name:     event.Sender.CardName,
					Nickname: event.Sender.Nickname,
					Avatar:   friendInfo[event.Sender.Uin].Avatar,
					IsBot:    false,
				},
				Role: nil,
				User: &satori.User{
					Id:       strconv.Itoa(int(event.Self)),
					Name:     client.NickName(),
					Nickname: client.NickName(),
					Avatar:   selfInfo.Avatar,
					IsBot:    false,
				},
				Meta: &satori.Meta{
					MessageId:  event.Id,
					InternalId: event.InternalId,
					ClientSeq:  event.ClientSeq,
				},
			},
		})
		// 将LagrangeGo的消息存入消息缓存
		err := Cache.Set(strconv.Itoa(int(event.InternalId)), event)
		if err != nil {
			return
		}
	})

	// 注册群聊事件
	Manager.RegisterGroupMessageHandler(func(client *client.QQClient, event *message.GroupMessage) {
		// 尝试从LagrangeGo的内置缓存中获取群信息
		groupInfo := client.GetCachedGroupInfo(event.GroupUin)
		// 尝试从适配器的缓存中获取自身信息
		selfInfo := GetSelfInfoInCache(client)
		// 尝试从适配器的缓存中获取群成员映射
		groupMemberData := GetGroupMembersDataInCache(client, event.GroupUin)
		ice.Publish(&ice.IceinuEvent{
			Type:      10,
			From:      ice.GetSelfNodeId(),
			Target:    ice.GetMasterNodeId(),
			Timestamp: time.Now().Unix(),
			Summary:   "ChannelMessageEvent",
			Event: &satori.EventSatori{
				Id:        ice.GenerateEventID(),
				Type:      "GroupMessageEvent",
				Platform:  "QQNT",
				SelfId:    strconv.Itoa(int(client.Uin)),
				Timestamp: int64(event.Time),
				Argv:      nil,
				Button:    nil,
				Channel: &satori.Channel{
					Id:       strconv.Itoa(int(event.GroupUin)),
					Type:     0,
					Name:     event.GroupName,
					ParentId: "",
				},
				Group: &satori.Group{
					Id:          strconv.Itoa(int(event.GroupUin)),
					Name:        event.GroupName,
					Avatar:      groupInfo.Avatar,
					Maxcount:    groupInfo.MaxMember,
					MemberCount: groupInfo.MemberCount,
				},
				Login:  nil,
				Member: nil,
				Message: &satori.Message{
					Id:              strconv.Itoa(int(event.InternalId)),
					Content:         event.ToString(),
					Channel:         nil,
					Group:           nil,
					Member:          nil,
					User:            nil,
					CreatedAt:       int64(event.Time),
					UpdatedAt:       int64(event.Time),
					MessageElements: ToSatoriElements(event.Elements),
				},
				Operator: &satori.User{
					Id:       strconv.Itoa(int(event.Sender.Uin)),
					Name:     event.Sender.CardName,
					Nickname: event.Sender.Nickname,
					Avatar:   groupMemberData[event.Sender.Uin].Avatar,
					IsBot:    false,
				},
				Role: nil,
				User: &satori.User{
					Id:       strconv.Itoa(int(client.Uin)),
					Name:     client.NickName(),
					Nickname: client.NickName(),
					Avatar:   selfInfo.Avatar,
					IsBot:    false,
				},
				Meta: &satori.Meta{
					MessageId:  event.Id,
					InternalId: event.InternalId,
				},
			},
		})
		// 将LagrangeGo的消息存入消息缓存
		err := Cache.Set(strconv.Itoa(int(event.InternalId)), event)
		if err != nil {
			return
		}
	})

	// 注册临时消息事件
	Manager.RegisterTempMessageHandler(func(client *client.QQClient, event *message.TempMessage) {
		// 尝试从LagrangeGo的内置缓存中获取群信息
		groupInfo := client.GetCachedGroupInfo(event.GroupUin)
		// 尝试从适配器的缓存中获取自身信息
		selfInfo := GetSelfInfoInCache(client)
		// 尝试从适配器的缓存中获取群成员映射
		groupMemberData := GetGroupMembersDataInCache(client, event.GroupUin)
		ice.Publish(&ice.IceinuEvent{
			Type:      10,
			From:      ice.GetSelfNodeId(),
			Target:    ice.GetMasterNodeId(),
			Timestamp: time.Now().Unix(),
			Summary:   "TempMessageEvent",
			Event: &satori.EventSatori{
				Id:        ice.GenerateEventID(),
				Type:      "TempMessageEvent",
				Platform:  "NTQQ",
				SelfId:    strconv.Itoa(int(client.Uin)),
				Timestamp: time.Now().Unix(),
				Argv:      nil,
				Button:    nil,
				Channel: &satori.Channel{
					Id:       strconv.Itoa(int(event.Sender.Uin)),
					Type:     1,
					Name:     event.Sender.Uid,
					ParentId: "",
				},
				Group: &satori.Group{
					Id:          strconv.Itoa(int(event.GroupUin)),
					Name:        event.GroupName,
					Avatar:      groupInfo.Avatar,
					Maxcount:    groupInfo.MaxMember,
					MemberCount: groupInfo.MemberCount,
				},
				Login:  nil,
				Member: nil,
				Message: &satori.Message{
					Id:              strconv.Itoa(int(event.Id)),
					Content:         event.ToString(),
					Channel:         nil,
					Group:           nil,
					Member:          nil,
					User:            nil,
					CreatedAt:       0,
					UpdatedAt:       0,
					MessageElements: ToSatoriElements(event.Elements),
				},
				Operator: &satori.User{
					Id:       strconv.Itoa(int(event.Sender.Uin)),
					Name:     event.Sender.CardName,
					Nickname: event.Sender.Nickname,
					Avatar:   groupMemberData[event.Sender.Uin].Avatar,
					IsBot:    false,
				},
				Role: nil,
				User: &satori.User{
					Id:       strconv.Itoa(int(client.Uin)),
					Name:     client.NickName(),
					Nickname: client.NickName(),
					Avatar:   selfInfo.Avatar,
					IsBot:    false,
				},
				Meta: &satori.Meta{
					MessageId: event.Id,
				},
			},
		})
		// 将LagrangeGo的消息存入消息缓存
		err := Cache.Set(strconv.Itoa(int(event.Id)), event)
		if err != nil {
			return
		}
	})

	// 注册好友消息撤回事件
	Manager.RegisterFriendRecallEventHandler(func(client *client.QQClient, event *event.FriendRecall) {
		// 尝试从适配器的缓存中获取自身信息
		selfInfo := GetSelfInfoInCache(client)
		// 尝试从适配器的缓存中获取好友信息
		friendInfo := GetFriendsDataInCache(client)
		ice.Publish(&ice.IceinuEvent{
			Type:      10,
			From:      ice.GetSelfNodeId(),
			Target:    ice.GetMasterNodeId(),
			Timestamp: time.Now().Unix(),
			Summary:   "FriendRecallEvent",
			Event: &satori.EventSatori{
				Id:        ice.GenerateEventID(),
				Type:      "FriendRecallEvent",
				Platform:  "QQNT",
				SelfId:    strconv.Itoa(int(client.Uin)),
				Timestamp: int64(event.Time),
				Argv:      nil,
				Button:    nil,
				Channel: &satori.Channel{
					Id:       strconv.Itoa(int(event.FromUin)),
					Type:     1,
					Name:     event.FromUid,
					ParentId: "",
				},
				Group: &satori.Group{
					Id:          "",
					Name:        "",
					Avatar:      "",
					Maxcount:    0,
					MemberCount: 0,
				},
				Login:   nil,
				Member:  nil,
				Message: nil,
				Operator: &satori.User{
					Id:       strconv.Itoa(int(event.FromUin)),
					Name:     friendInfo[event.FromUin].Nickname,
					Nickname: friendInfo[event.FromUin].Nickname,
					Avatar:   friendInfo[event.FromUin].Avatar,
					IsBot:    false,
				},
				Role: nil,
				User: &satori.User{
					Id:       strconv.Itoa(int(client.Uin)),
					Name:     client.NickName(),
					Nickname: client.NickName(),
					Avatar:   selfInfo.Avatar,
					IsBot:    false,
				},
				Meta: &satori.Meta{
					Sequence: event.Sequence,
					Random:   event.Random,
				},
			},
		})
	})

	// 注册群聊消息撤回事件
	Manager.RegisterGroupRecallEventHandler(func(client *client.QQClient, event *event.GroupRecall) {
		// 尝试从LagrangeGo的内置缓存中获取群信息
		groupInfo := client.GetCachedGroupInfo(event.GroupUin)
		// 尝试从适配器的缓存中获取群成员映射
		groupMemberData := GetGroupMembersDataInCache(client, event.GroupUin)
		ice.Publish(&ice.IceinuEvent{
			Type:      10,
			From:      ice.GetSelfNodeId(),
			Target:    ice.GetMasterNodeId(),
			Timestamp: time.Now().Unix(),
			Summary:   "ChannelRecallEvent",
			Event: &satori.EventSatori{
				Id:        ice.GenerateEventID(),
				Type:      "ChannelRecallEvent",
				Platform:  "QQNT",
				SelfId:    strconv.Itoa(int(client.Uin)),
				Timestamp: int64(event.Time),
				Argv:      nil,
				Button:    nil,
				Channel: &satori.Channel{
					Id:       strconv.Itoa(int(event.GroupUin)),
					Type:     0,
					Name:     groupInfo.GroupName,
					ParentId: "",
				},
				Group: &satori.Group{
					Id:          strconv.Itoa(int(event.GroupUin)),
					Name:        groupInfo.GroupName,
					Avatar:      groupInfo.Avatar,
					Maxcount:    groupInfo.MaxMember,
					MemberCount: groupInfo.MemberCount,
				},
				Login: nil,
				Member: &satori.GroupMember{
					User: &satori.User{
						Id:       strconv.Itoa(int(event.AuthorUin)),
						Name:     groupMemberData[event.AuthorUin].MemberCard,
						Nickname: "",
						Avatar:   "",
						IsBot:    false,
					},
					Nickname: groupMemberData[event.AuthorUin].MemberName,
					Avatar:   groupMemberData[event.AuthorUin].Avatar,
					JoinedAt: int64(groupMemberData[event.AuthorUin].JoinTime),
				},
				Message: nil,
				Operator: &satori.User{
					Id:       strconv.Itoa(int(event.OperatorUin)),
					Name:     groupMemberData[event.OperatorUin].MemberCard,
					Nickname: groupMemberData[event.OperatorUin].MemberName,
					Avatar:   groupMemberData[event.OperatorUin].Avatar,
					IsBot:    false,
				},
				Role: nil,
				User: &satori.User{
					Id:       strconv.Itoa(int(event.AuthorUin)),
					Name:     groupMemberData[event.AuthorUin].MemberCard,
					Nickname: groupMemberData[event.AuthorUin].MemberName,
					Avatar:   groupMemberData[event.AuthorUin].Avatar,
					IsBot:    false,
				},
				Meta: &satori.Meta{
					Sequence: event.Sequence,
					Random:   event.Random,
				},
			},
		})
	})

	// 注册好友改名事件
	Manager.RegisterRenameEventHandler(func(client *client.QQClient, event *event.Rename) {
		// 尝试从适配器的缓存中获取自身信息
		selfInfo := GetSelfInfoInCache(client)
		// 尝试从适配器的缓存中获取好友信息
		friendInfo := GetFriendsDataInCache(client)
		ice.Publish(&ice.IceinuEvent{
			Type:      10,
			From:      ice.GetSelfNodeId(),
			Target:    ice.GetMasterNodeId(),
			Timestamp: time.Now().Unix(),
			Summary:   "RenameEvent",
			Event: &satori.EventSatori{
				Id:        ice.GenerateEventID(),
				Type:      "RenameEvent",
				Platform:  "QQNT",
				SelfId:    strconv.Itoa(int(client.Uin)),
				Timestamp: time.Now().Unix(),
				Argv:      nil,
				Button:    nil,
				Channel:   nil,
				Group:     nil,
				Login:     nil,
				Member:    nil,
				Message:   nil,
				Operator: &satori.User{
					Id:       strconv.Itoa(int(event.Uin)),
					Name:     friendInfo[event.Uin].Nickname,
					Nickname: event.Nickname,
					Avatar:   friendInfo[event.Uin].Avatar,
					IsBot:    false,
				},
				Role: nil,
				User: &satori.User{
					Id:       strconv.Itoa(int(client.Uin)),
					Name:     client.NickName(),
					Nickname: client.NickName(),
					Avatar:   selfInfo.Avatar,
					IsBot:    false,
				},
				Meta: nil,
			},
		})
	})

	// 注册好友请求事件
	Manager.RegisterNewFriendRequestHandler(func(client *client.QQClient, event *event.NewFriendRequest) {
		// 尝试从适配器的缓存中获取自身信息
		selfInfo := GetSelfInfoInCache(client)
		// 拉取对方用户信息
		info, err := client.FetchUserInfoUin(event.SourceUin)
		if err != nil {
			return
		}
		ice.Publish(&ice.IceinuEvent{
			Type:      10,
			From:      ice.GetSelfNodeId(),
			Target:    ice.GetMasterNodeId(),
			Timestamp: time.Now().Unix(),
			Summary:   "NewFriendRequestEvent",
			Event: &satori.EventSatori{
				Id:        ice.GenerateEventID(),
				Type:      "NewFriendRequestEvent",
				Platform:  "QQNT",
				SelfId:    strconv.Itoa(int(client.Uin)),
				Timestamp: time.Now().Unix(),
				Argv:      nil,
				Button:    nil,
				Channel:   nil,
				Group:     nil,
				Login:     nil,
				Member:    nil,
				Message:   nil,
				Operator: &satori.User{
					Id:       strconv.Itoa(int(event.SourceUin)),
					Name:     event.Source,
					Nickname: event.SourceNick,
					Avatar:   info.Avatar,
					IsBot:    false,
				},
				Role: nil,
				User: &satori.User{
					Id:       strconv.Itoa(int(client.Uin)),
					Name:     client.NickName(),
					Nickname: client.NickName(),
					Avatar:   selfInfo.Avatar,
					IsBot:    false,
				},
				Meta: &satori.Meta{
					Comment: event.Msg,
				},
			},
		})
	})

}
