package ice

import "time"

// MakeWebsocketHeartbeatEvent 创建一个WebSocket心跳事件
func MakeWebsocketHeartbeatEvent() *IceinuEvent {
	e := &IceinuEvent{
		Type:      0,
		From:      GetSelfNodeId(),
		Target:    GetMasterNodeId(),
		Timestamp: time.Now().Unix(),
		Summary:   "WebSocketHeartbeatEvent",
		Event: &WebSocketHeartbeatEvent{
			OK: true,
		},
	}
	Publish(e)
	return e
}

// MakeNodeConnectRequestEvent 创建一个节点请求建立WebSocket连接事件
func MakeNodeConnectRequestEvent(mode string, adapterModel string, pluginVerifier string) *IceinuEvent {
	e := &IceinuEvent{
		Type:      1,
		From:      GetSelfNodeId(),
		Target:    GetMasterNodeId(),
		Timestamp: time.Now().Unix(),
		Summary:   "NodeConnectRequestEvent",
		Event: &NodeConnectRequestEvent{
			Mode:           mode,
			AdapterModel:   adapterModel,
			PluginVerifier: pluginVerifier,
		},
	}
	Publish(e)
	return e
}

// MakeNodeConnectedEvent 创建一个节点成功连接事件
func MakeNodeConnectedEvent() *IceinuEvent {
	e := &IceinuEvent{
		Type:      2,
		From:      GetSelfNodeId(),
		Target:    GetMasterNodeId(),
		Timestamp: time.Now().Unix(),
		Summary:   "NodeConnectedEvent",
		Event: &NodeConnectedEvent{
			OK: true,
		},
	}
	Publish(e)
	return e
}

// MakeNodeDisconnectedEvent 创建一个节点断开连接事件
func MakeNodeDisconnectedEvent() *IceinuEvent {
	e := &IceinuEvent{
		Type:      3,
		From:      GetSelfNodeId(),
		Target:    GetMasterNodeId(),
		Timestamp: time.Now().Unix(),
		Summary:   "NodeDisconnectedEvent",
		Event: &NodeDisconnectedEvent{
			OK: true,
		},
	}
	Publish(e)
	return e
}

// MakeNodePushDataEvent 创建一个节点推送数据事件
func MakeNodePushDataEvent(data interface{}) *IceinuEvent {
	e := &IceinuEvent{
		Type:      4,
		From:      GetSelfNodeId(),
		Target:    GetMasterNodeId(),
		Timestamp: time.Now().Unix(),
		Summary:   "NodePushDataEvent",
		Event:     data,
	}
	Publish(e)
	return e
}

// MakeNodeUserPushEvent 创建一个节点用户推送事件
func MakeNodeUserPushEvent(data interface{}) *IceinuEvent {
	e := &IceinuEvent{
		Type:      5,
		From:      GetSelfNodeId(),
		Target:    GetMasterNodeId(),
		Timestamp: time.Now().Unix(),
		Summary:   "NodeUserPushEvent",
		Event:     data,
	}
	Publish(e)
	return e
}

// MakeNodeRequestDataEvent 创建一个节点请求数据事件
func MakeNodeRequestDataEvent(dataType string, key string, query string) *IceinuEvent {
	e := &IceinuEvent{
		Type:      6,
		From:      GetSelfNodeId(),
		Target:    GetMasterNodeId(),
		Timestamp: time.Now().Unix(),
		Summary:   "NodeRequestDataEvent",
		Event: &NodeRequestDataEvent{
			DataType: dataType,
			Key:      key,
			Query:    query,
		},
	}
	Publish(e)
	return e
}

// MakeRequestNodeDataEvent 创建一个请求节点数据事件
func MakeRequestNodeDataEvent(dataType string, key string, query string) *IceinuEvent {
	e := &IceinuEvent{
		Type:      7,
		From:      GetSelfNodeId(),
		Target:    GetMasterNodeId(),
		Timestamp: time.Now().Unix(),
		Summary:   "RequestNodeDataEvent",
		Event: &RequestNodeDataEvent{
			DataType: dataType,
			Key:      key,
			Query:    query,
		},
	}
	Publish(e)
	return e
}

// MakeNodeDeactiveEvent 创建一个节点失活事件
func MakeNodeDeactiveEvent() *IceinuEvent {
	e := &IceinuEvent{
		Type:      12,
		From:      GetSelfNodeId(),
		Target:    GetMasterNodeId(),
		Timestamp: time.Now().Unix(),
		Summary:   "NodeActiveEvent",
		Event: &NodeDeactiveEvent{
			OK: true,
		},
	}
	Publish(e)
	return e
}

// MakeNodeReactiveEvent 创建一个节点重新激活事件
func MakeNodeReactiveEvent() *IceinuEvent {
	e := &IceinuEvent{
		Type:      13,
		From:      GetSelfNodeId(),
		Target:    GetMasterNodeId(),
		Timestamp: time.Now().Unix(),
		Summary:   "NodeActiveEvent",
		Event: &NodeReactiveEvent{
			OK: true,
		},
	}
	Publish(e)
	return e
}

// MakeAdapterConnectEvent 创建一个适配器连接事件
func MakeAdapterConnectEvent(adapterType string, model string, userId string, userName string) *IceinuEvent {
	e := &IceinuEvent{
		Type:      8,
		From:      GetSelfNodeId(),
		Target:    GetMasterNodeId(),
		Timestamp: time.Now().Unix(),
		Summary:   "AdapterConnectEvent",
		Event: &AdapterConnectEvent{
			AdapterType:  adapterType,
			AdapterModel: model,
			UserId:       userId,
			UserName:     userName,
		},
	}
	Publish(e)
	return e
}

// MakeAdapterDisconnectEvent 创建一个适配器断开连接事件
func MakeAdapterDisconnectEvent() *IceinuEvent {
	e := &IceinuEvent{
		Type:      9,
		From:      GetSelfNodeId(),
		Target:    GetMasterNodeId(),
		Timestamp: time.Now().Unix(),
		Summary:   "AdapterDisconnectEvent",
		Event: &AdapterDisconnectEvent{
			OK: true,
		},
	}
	Publish(e)
	return e
}
