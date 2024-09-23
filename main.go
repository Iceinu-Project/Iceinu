package main

import (
	"fmt"
	"github.com/Iceinu-Project/iceinu/ice"
	"github.com/Iceinu-Project/iceinu/log"
	"time"
)

// Iceinu的程序入口
// 可以参照文档来对其进行修改

func main() {
	// 定义日志格式
	formatter := &LogFormatter{}
	log.SetFormatter(formatter)
	// 输出日志
	log.Info("Hello, World!")
	// 创建事件总线
	bus := ice.NewEventBus()

	customPublishLogger := func(event *ice.IceinuEvent, next func(event *ice.IceinuEvent)) {
		log.Infof("Publish event: %v", event)
		next(event)
		log.Infof("Event published: %v", event)
	}

	// 定义订阅者中间件：错误恢复
	subscriberRecovery := func(event *ice.IceinuEvent, next func(event *ice.IceinuEvent)) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("处理事件时发生错误：%v\n", err)
			}
		}()
		next(event)
	}

	// 添加中间件
	bus.UseGlobalPublishMiddleware(customPublishLogger)
	bus.UseSubscribeMiddleware(subscriberRecovery)

	// 订阅事件，获取订阅 ID
	subID := bus.Subscribe(1, "NodeConnect", func(event *ice.IceinuEvent) {
		fmt.Println("处理事件内容：", event.Event)
		// 模拟错误
		// panic("模拟的错误")
	})

	// 发布事件
	event := &ice.IceinuEvent{
		Type:      1,
		From:      "node1",
		Target:    "node2",
		Timestamp: time.Now().Unix(),
		Summary:   "NodeConnect",
		Event:     "节点连接事件",
	}
	bus.Publish(event)

	// 等待异步处理完成
	time.Sleep(1 * time.Second)

	// 取消订阅
	bus.Unsubscribe(1, "NodeConnect", subID)
}
