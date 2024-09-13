package lagrange

import (
	"github.com/Iceinu-Project/iceinu/ice"
	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/message"
)

// 向lagrangego的subscriber注册一系列处理函数，将其事件转发到iceinu的事件总线便于截取处理

func SetIceinuHandler() {
	Manager.RegisterPrivateMessageHandler(func(client *client.QQClient, event *message.PrivateMessage) {
		// 将私聊事件转换成Iceinu的通用消息结构
		ice.IceBus.Push("PrivateMessageEvent", ice.PrivateMessageEvent{
			Client: client,
			Message: &ice.Message{
				MsgId:         event.Id,
				MsgInternalId: event.InternalId,
				ClientSeq:     0,
				EventId:       event.Target,
				EventName:     "",
				UserId:        event.Sender.Uin,
				Sender:        event.Sender,
				Elements:      event.Elements,
			},
		})
	})

	Manager.RegisterGroupMessageHandler(func(client *client.QQClient, event *message.GroupMessage) {
		// 将群聊事件转换成Iceinu的通用消息结构
		ice.IceBus.Push("GroupMessageEvent", ice.GroupMessageEvent{
			Client: client,
			Message: &ice.Message{
				MsgId:         event.Id,
				MsgInternalId: event.InternalId,
				ClientSeq:     0,
				EventId:       event.GroupUin,
				EventName:     event.GroupName,
				UserId:        event.Sender.Uin,
				Sender:        event.Sender,
				Elements:      event.Elements,
			},
		})

	})

}
