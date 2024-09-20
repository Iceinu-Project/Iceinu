package lagrange

import (
	"fmt"
	"strconv"
	"time"

	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/message"

	"github.com/Iceinu-Project/iceinu/ice"
	"github.com/Iceinu-Project/iceinu/logger"
	"github.com/Iceinu-Project/iceinu/resource"
	"github.com/Iceinu-Project/iceinu/utils"
)

func SetAllHandler() {
	Manager.RegisterPrivateMessageHandler(func(client *client.QQClient, event *message.PrivateMessage) {
		self, _ := client.FetchUserInfoUin(client.Uin)
		e := ice.PlatformEvent{
			EventId:   uint64(event.Id),
			EventType: "PrivateMessageEvent",
			Platform:  "QQNT",
			SelfId:    strconv.Itoa(int(client.Uin)),
			Timestamp: time.Unix(int64(event.Time), 0),
			Group: &resource.Group{
				Id:     "",
				Name:   "",
				Avatar: "",
			},
			Channel: &resource.Channel{
				Id:       strconv.Itoa(int(event.Sender.Uin)),
				Type:     1,
				Name:     event.Sender.Uid,
				ParentId: "",
			},
			Message: &resource.Message{
				Id:              strconv.Itoa(int(event.InternalId)),
				Content:         event.ToString(),
				MessageElements: ConvertIceElement(event.Elements),
			},
			Operator: &resource.User{
				Id:       strconv.Itoa(int(event.Sender.Uin)),
				Name:     event.Sender.Uid,
				Nickname: event.Sender.Nickname,
				Avatar:   self.Avatar,
				IsBot:    false,
			},
			User: &resource.User{
				Id:       strconv.Itoa(int(event.Target)),
				Name:     client.GetUid(client.Uin),
				Nickname: client.NickName(),
				Avatar:   self.Avatar,
				IsBot:    false,
			},
		}
		logger.Infof("[私聊][%s]%s：%s", e.Operator.Id, e.Operator.Nickname, utils.SatorizeIceElements(e.Message.MessageElements))
		ice.Bus.Publish("PrivateMessageEvent", &e)
	})
	Manager.RegisterGroupMessageHandler(func(client *client.QQClient, event *message.GroupMessage) {
		groupinfo := client.GetCachedGroupInfo(event.GroupUin)
		fmt.Println(groupinfo)
		self, _ := client.FetchUserInfoUin(client.Uin)
		e := ice.PlatformEvent{
			EventId:   uint64(event.Id),
			EventType: "GroupMessageEvent",
			Platform:  "QQNT",
			SelfId:    strconv.Itoa(int(client.Uin)),
			Timestamp: time.Unix(int64(event.Time), 0),
			Channel: &resource.Channel{
				Id:       strconv.Itoa(int(event.GroupUin)),
				Type:     0,
				Name:     event.GroupName,
				ParentId: "",
			},
			Group: &resource.Group{
				Id:     strconv.Itoa(int(event.GroupUin)),
				Name:   event.GroupName,
				Avatar: groupinfo.Avatar,
			},
			Message: &resource.Message{
				Id:              strconv.Itoa(int(event.InternalId)),
				Content:         event.ToString(),
				MessageElements: ConvertIceElement(event.Elements),
			},
			Operator: &resource.User{
				Id:       strconv.Itoa(int(event.Sender.Uin)),
				Name:     event.Sender.Uid,
				Nickname: event.Sender.Nickname,
				Avatar:   self.Avatar,
				IsBot:    false,
			},
		}
		logger.Infof("[群聊][来自群%s][%s]%s：%s", e.Group.Id, e.Operator.Id, e.Operator.Nickname, utils.SatorizeIceElements(e.Message.MessageElements))
		ice.Bus.Publish("GroupMessageEvent", &e)
	})
}
