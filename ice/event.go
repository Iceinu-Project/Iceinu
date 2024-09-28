package ice

import "github.com/Iceinu-Project/Iceinu/adapters"

// IceinuEvent Iceinu全局事件结构体
//
// 默认的全局事件总线结构体，在直接使用Iceinu框架进行二次开发时使用这个事件结构
//
// 引用Iceinu框架作为第三方库时如需自定义自己的事件结构体，需要自行实现事件总线
//
// 事件种类：
// 0：WebSocket心跳事件
//
// 1：节点请求建立WebSocket连接事件
//
// 2：节点成功连接事件
//
// 3：节点断开连接事件
//
// 4：节点推送数据事件（主节点到子节点）
//
// 5：节点用户推送事件（子节点到主节点）
//
// 6：节点请求数据事件（子节点到主节点）
//
// 7：节点请求数据事件（主节点到子节点）
//
// 8：适配器连接事件
//
// 9：适配器断开连接事件
//
// 10：消息接收事件
//
// 11：消息发送事件
//
// 12：节点失活事件（子节点到主节点）
//
// 13：节点重新激活事件（子节点到主节点）
//
// 14：数据回调事件，用于节点间的数据传输
type IceinuEvent struct {
	Type      uint8       `json:"type"`
	From      string      `json:"from"`      // 消息事件来源节点ID
	Target    string      `json:"target"`    // 消息事件目标节点ID
	Timestamp int64       `json:"timestamp"` // 事件推送的时间戳
	Summary   string      `json:"summary"`   // 事件摘要，一般是承载的事件的事件类型，用于在事件总线层快速识别事件类型
	Event     interface{} `json:"event"`     // 事件内容，用于承载不同类型的消息，使用时需要进行断言
}

// WebSocketHeartbeatEvent 0：WebSocket心跳事件结构体
type WebSocketHeartbeatEvent struct {
	OK bool `json:"ok"`
}

// NodeConnectRequestEvent 1：节点请求建立WebSocket连接事件结构体
type NodeConnectRequestEvent struct {
	Mode           string `json:"mode"`            // 组网模式
	AdapterModel   string `json:"adapter_model"`   // 适配器模型
	PluginVerifier string `json:"plugin_verifier"` // 插件校验值
}

// NodeConnectedEvent 2：节点成功连接事件结构体
type NodeConnectedEvent struct {
	OK bool `json:"ok"`
}

// NodeDisconnectedEvent 3：节点断开连接事件结构体
type NodeDisconnectedEvent struct {
	OK bool `json:"ok"`
}

// NodePushDataEvent 4：节点推送数据事件（主节点到子节点）结构体
type NodePushDataEvent struct {
	Data interface{} `json:"data"` // 推送的数据
}

// NodeUserPushEvent 5：节点用户推送事件（子节点到主节点）结构体
type NodeUserPushEvent struct {
	UserTree adapters.UserTree `json:"user_tree"` // 用户树
}

// NodeRequestDataEvent 6：节点请求数据事件（子节点到主节点）结构体
type NodeRequestDataEvent struct {
	RequestSerial int32  `json:"request_serial"` // 请求序列号，用于标识请求防止接受错误
	DataType      string `json:"data_type"`      // 请求的数据类型
	Key           string `json:"key"`            // 请求的数据键
	Query         string `json:"query"`          // 请求的查询内容
}

// RequestNodeDataEvent 7：节点请求数据事件（主节点到子节点）结构体
type RequestNodeDataEvent struct {
	RequestSerial int32  `json:"request_serial"` // 请求序列号，用于标识请求防止接受错误
	DataType      string `json:"data_type"`      // 请求的数据类型
	Key           string `json:"key"`            // 请求的数据键
	Query         string `json:"query"`          // 请求的查询内容
}

// AdapterConnectEvent 8：适配器连接事件结构体
type AdapterConnectEvent struct {
	AdapterType  string `json:"adapter_type"`  // 适配器类型
	AdapterModel string `json:"adapter_model"` // 适配器模型
	UserId       string `json:"user_id"`       // 用户ID
	UserName     string `json:"user_name"`     // 用户名称
}

// AdapterDisconnectEvent 9：适配器断开连接事件结构体
type AdapterDisconnectEvent struct {
	OK bool `json:"ok"`
}

// NodeDeactiveEvent 12：节点失活事件（子节点到主节点）结构体
type NodeDeactiveEvent struct {
	OK bool `json:"ok"`
}

// NodeReactiveEvent 13：节点重新激活事件（子节点到主节点）结构体
type NodeReactiveEvent struct {
	OK bool `json:"ok"`
}

// DataCallbackEvent 14：数据回调事件，用于节点间的数据传输结构体
type DataCallbackEvent struct {
	OK            bool        `json:"ok"`             // 回调是否成功
	DataType      string      `json:"data_type"`      // 数据类型
	Data          interface{} `json:"data"`           // 回调的数据
	RequestSerial int32       `json:"request_serial"` // 请求序列号，用于标识请求防止接受错误
}
