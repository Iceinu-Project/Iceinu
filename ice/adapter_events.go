package ice

import (
	"gtihub.com/Iceinu-Project/iceinu/adapter"
	"time"
)

// AdapterInitEvent 适配器初始化事件
type AdapterInitEvent struct {
	Timestamp   time.Time
	AdapterMeta *adapter.IceAdapterMeta
}

// AdapterConnectEvent 适配器连接事件
type AdapterConnectEvent struct {
	Timestamp time.Time
}

// AdapterDisconnectEvent 适配器断开连接事件
type AdapterDisconnectEvent struct {
	Timestamp time.Time
}

// AdapterLoginEvent 适配器登录事件
type AdapterLoginEvent struct {
	Timestamp time.Time
}
