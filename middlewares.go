package main

import (
	"github.com/Iceinu-Project/Iceinu/ice"
	"github.com/Iceinu-Project/Iceinu/log"
	"github.com/Iceinu-Project/Iceinu/models/satori"
)

// SetCustomMiddleWare 设置自定义中间件
func SetCustomMiddleWare() {

	// 这个函数中定义了一系列事件总线的中间件，你可以直接参考这些中间件以及氷犬的在线文档来编写中间件

	ice.UseGlobalPublishMiddleware(func(event *ice.IceinuEvent, next func(event *ice.IceinuEvent)) {
		// 默认的事件监听中间件示例，它会在任意事件发布时运行
		log.Debugf("接收到来自节点 %s 的 %s 事件", event.From, event.Summary)
		next(event)
	})

	ice.UseTypePublishMiddleware(8, func(event *ice.IceinuEvent, next func(event *ice.IceinuEvent)) {
		// 适配器连接事件监听中间件，它会在适配器连接事件发布时运行
		// 提取事件里的参数
		adapterConnectEvent := event.Event.(*ice.AdapterConnectEvent)
		log.Infof("来自节点 %s 的%s连接成功，机器人Id %s，机器人名称 %s", event.From, adapterConnectEvent.AdapterType, adapterConnectEvent.UserId, adapterConnectEvent.UserName)
		next(event)
	})

	ice.UseSummaryPublishMiddleware("PrivateMessageEvent", func(event *ice.IceinuEvent, next func(event *ice.IceinuEvent)) {
		// Satori私聊消息事件监听中间件，它会在私聊消息事件发布时运行
		privateMessageEvent := event.Event.(*satori.EventSatori)
		log.Infof("[%s][%s][私聊]%s@%s：%s", privateMessageEvent.Platform, privateMessageEvent.User.Name, privateMessageEvent.Operator.Id, privateMessageEvent.Operator.Nickname, satori.ElementsToSatori(*privateMessageEvent.Message.MessageElements))
		next(event)
	})

	ice.UseSummaryPublishMiddleware("ChannelMessageEvent", func(event *ice.IceinuEvent, next func(event *ice.IceinuEvent)) {
		// Satori群聊消息事件监听中间件，它会在群聊消息事件发布时运行
		groupMessageEvent := event.Event.(*satori.EventSatori)
		log.Infof("[%s][%s][群聊][%s]%s@%s：%s", groupMessageEvent.Platform, groupMessageEvent.User.Name, groupMessageEvent.Group.Name, groupMessageEvent.Operator.Id, groupMessageEvent.Operator.Nickname, satori.ElementsToSatori(*groupMessageEvent.Message.MessageElements))
		next(event)
	})

	ice.UseSummaryPublishMiddleware("FriendRecallEvent", func(event *ice.IceinuEvent, next func(event *ice.IceinuEvent)) {
		// Satori好友撤回消息事件监听中间件，它会在好友撤回消息事件发布时运行
		friendRecallEvent := event.Event.(*satori.EventSatori)
		log.Infof("[%s][%s][私聊]%s@%s 撤回了一条消息", friendRecallEvent.Platform, friendRecallEvent.User.Name, friendRecallEvent.Operator.Id, friendRecallEvent.Operator.Nickname)
		next(event)
	})

	ice.UseSummaryPublishMiddleware("ChannelRecallEvent", func(event *ice.IceinuEvent, next func(event *ice.IceinuEvent)) {
		// Satori群聊撤回消息事件监听中间件，它会在群聊撤回消息事件发布时运行
		groupRecallEvent := event.Event.(*satori.EventSatori)
		log.Infof("[%s][%s][群聊][%s]%s@%s 撤回了一条消息", groupRecallEvent.Platform, groupRecallEvent.User.Name, groupRecallEvent.Group.Name, groupRecallEvent.Operator.Id, groupRecallEvent.Operator.Nickname)
		next(event)
	})

	ice.UseSummaryPublishMiddleware("FriendRequestEvent", func(event *ice.IceinuEvent, next func(event *ice.IceinuEvent)) {
		// Satori好友请求事件监听中间件，它会在好友请求事件发布时运行
		friendRequestEvent := event.Event.(*satori.EventSatori)
		log.Infof("[%s][%s]收到来自%s的好友请求：%s", friendRequestEvent.Platform, friendRequestEvent.User.Name, friendRequestEvent.Operator.Id, friendRequestEvent.Meta.Comment)
		next(event)
	})
}
