package lagrange

import (
	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/message"
	"gtihub.com/Iceinu-Project/iceinu/logger"
	"gtihub.com/Iceinu-Project/iceinu/resource"
	"gtihub.com/Iceinu-Project/iceinu/utils"
	"strconv"
	"time"

	"gtihub.com/Iceinu-Project/iceinu/ice"
)

func SetAllHandler() {
	Manager.RegisterPrivateMessageHandler(func(client *client.QQClient, event *message.PrivateMessage) {
		e := ice.PlatformEvent{
			EventId:   uint64(event.Id),
			EventType: "PrivateMessageEvent",
			Platform:  "QQNT",
			SelfId:    strconv.Itoa(int(client.Uin)),
			Timestamp: time.Unix(int64(event.Time), 0),
			Message: &resource.Message{
				Id:              strconv.Itoa(int(event.InternalId)),
				Content:         event.ToString(),
				MessageElements: ConvertIceElement(event.Elements),
			},
			Operator: &resource.User{
				Id:       strconv.Itoa(int(event.Sender.Uin)),
				Name:     event.Sender.Uid,
				Nickname: event.Sender.Nickname,
				Avatar:   "",
				IsBot:    false,
			},
			User: &resource.User{
				Id:       strconv.Itoa(int(event.Target)),
				Name:     client.GetUid(client.Uin),
				Nickname: client.NickName(),
				Avatar:   "",
				IsBot:    false,
			},
		}
		logger.Infof("[私聊][%s]%s：%s", e.Operator.Id, e.Operator.Nickname, utils.SatorizeIceElements(e.Message.MessageElements))
		ice.Bus.Publish("PrivateMessageEvent", &e)
	})
	Manager.RegisterGroupMessageHandler(func(client *client.QQClient, event *message.GroupMessage) {
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
			Guild: &resource.Guild{
				Id:     strconv.Itoa(int(event.GroupUin)),
				Name:   event.GroupName,
				Avatar: "",
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
				Avatar:   "",
				IsBot:    false,
			},
		}
		logger.Infof("[群聊][来自群%s][%s]%s：%s", e.Guild.Id, e.Operator.Id, e.Operator.Nickname, utils.SatorizeIceElements(e.Message.MessageElements))
		ice.Bus.Publish("GroupMessageEvent", &e)
	})
}
