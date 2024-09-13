package ice

import "github.com/LagrangeDev/LagrangeGo/client"

// Bot 对LagrangeGo的QQClient的封装，使其更加简单易用
type Bot struct {
	LgrangeClient *client.QQClient
}

// NewBot 创建一个新的Bot实例
func NewBot(lgClient *client.QQClient) *Bot {
	return &Bot{
		LgrangeClient: lgClient,
	}
}
