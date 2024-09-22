package models

type IceinuEvent struct {
	// 事件种类，用于确定其中承载的消息类型
	// 0：WebSocket心跳事件
	// 1：节点连接事件
	// 2：节点断开事件
	// 3：节点更新推送事件
	// 4：节点更新请求事件
	// 5：消息接受事件（从子节点到主节点）
	// 6：消息处理事件（从主节点到子节点）
	Type      uint8       `json:"type"`
	From      string      `json:"from"`      // 消息事件来源节点ID
	Target    string      `json:"target"`    // 消息事件目标节点ID
	Timestamp int64       `json:"timestamp"` // 事件推送的时间戳
	Event     interface{} `json:"event"`     // 事件内容，用于承载不同类型的消息，由事件总线进行断言
}
