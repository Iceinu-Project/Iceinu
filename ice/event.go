package ice

import (
	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/message"
)

// Message 消息事件消息，用于同时处理私聊和群聊信息
type Message struct {
	MsgId         uint32
	MsgInternalId uint32
	ClientSeq     uint32
	EventId       uint32
	EventName     string
	UserId        uint32
	Sender        *message.Sender
	Elements      []message.IMessageElement
}

type PrivateMessageEvent struct {
	Client  *client.QQClient
	Message *Message
}

type GroupMessageEvent struct {
	Client  *client.QQClient
	Message *Message
}
