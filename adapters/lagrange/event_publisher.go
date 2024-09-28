package lagrange

import (
	"github.com/Iceinu-Project/Iceinu/ice"
	"github.com/Iceinu-Project/Iceinu/models/satori"
	"github.com/LagrangeDev/LagrangeGo/client"
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
				Id:        uint64(event.Id),
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
			Summary:   "GroupMessageEvent",
			Event: &satori.EventSatori{
				Id:        uint64(event.Id),
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
			},
		})
	})
}
