package handler

import (
	"fmt"
	"github.com/Iceinu-Project/iceinu/ice"
)

func BindHandler() {
	ice.IceBus.Bind("PrivateMessageEvent", func(args ...interface{}) {
		// 处理私聊消息
		e := args[0].(ice.PrivateMessageEvent)
		fmt.Println("PrivateMessageEvent", e.Client.Uin, e.Message.UserId)
	})
	ice.IceBus.Bind("GroupMessageEvent", func(args ...interface{}) {
		// 处理群聊消息
		e := args[0].(ice.GroupMessageEvent)
		fmt.Println("GroupMessageEvent", e.Client.Uin, e.Message.UserId)
	})
}
