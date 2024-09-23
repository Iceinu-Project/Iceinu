package ice

// IceinuEvent Iceinu全局事件结构体
//
// 默认的全局事件总线结构体，在直接使用Iceinu框架进行二次开发时使用这个事件结构
//
// 引用Iceinu框架作为第三方库时如需自定义自己的事件结构体，需要自行实现事件总线
//
// 事件种类：
// 0：WebSocket心跳事件
//
// 1：节点连接事件
//
// 2：节点断开事件
//
// 3：节点更新推送事件
//
// 4：节点更新请求事件
//
// 5：消息接受事件（从子节点到主节点）
//
// 6：消息处理事件（从主节点到子节点）
type IceinuEvent struct {
	Type      uint8       `json:"type"`
	From      string      `json:"from"`      // 消息事件来源节点ID
	Target    string      `json:"target"`    // 消息事件目标节点ID
	Timestamp int64       `json:"timestamp"` // 事件推送的时间戳
	Summary   string      `json:"summary"`   // 事件摘要，一般是承载的事件的事件类型，用于在事件总线层快速识别事件类型
	Event     interface{} `json:"event"`     // 事件内容，用于承载不同类型的消息，使用时需要进行断言
}

type WebSocketHeartbeatEvent struct {
	OK bool `json:"ok"`
}

type NodeConnectEvent struct {
}

type NodeDisconnectEvent struct {
}

type NodeUpdatePushEvent struct {
}

type NodeUpdateRequestEvent struct {
}
