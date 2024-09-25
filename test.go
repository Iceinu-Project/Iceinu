package main

import (
	"github.com/Iceinu-Project/Iceinu/ice"
	"time"
)

// 这不是Iceinu的自动单元测试入口，只是一个临时的测试函数

func Test() {
	// 发布事件测试
	// 发布事件
	ice.Bus.Publish(&ice.IceinuEvent{
		Type:      0,
		From:      "node0",
		Target:    "node0",
		Timestamp: time.Now().Unix(),
		Summary:   "WebsocketHeartBeat",
		Event:     ice.WebSocketHeartbeatEvent{OK: true},
	})
}
