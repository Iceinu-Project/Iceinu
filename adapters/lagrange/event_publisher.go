package lagrange

import (
	"github.com/Iceinu-Project/Iceinu/ice"
	"github.com/Iceinu-Project/Iceinu/models/satori"
	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/message"
	"time"
)

func EventsBinder() {
	// 绑定私聊消息事件
	Client.PrivateMessageEvent.Subscribe(func(client *client.QQClient, event *message.PrivateMessage) {
		ice.Publish(&ice.IceinuEvent{
			Type:      5,
			From:      ice.GetSelfNodeId(),
			Target:    ice.GetMasterNodeId(),
			Timestamp: time.Now().Unix(),
			Summary:   "",
			Event:     &satori.EventSatori{},
		})
	})
}
