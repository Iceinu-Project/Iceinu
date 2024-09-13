package ice

import "sync"

var IceBus = NewEventBus()

// IceEventBus Iceinu的事件总线
type IceEventBus struct {
	handlers map[string][]interface{}
	lock     sync.RWMutex
}

// NewEventBus 创建一个新的事件总线
func NewEventBus() *IceEventBus {
	return &IceEventBus{
		handlers: make(map[string][]interface{}),
	}
}

// Bind 绑定一个事件处理函数
func (bus *IceEventBus) Bind(eventType string, handler interface{}) {
	bus.lock.Lock()
	defer bus.lock.Unlock()
	bus.handlers[eventType] = append(bus.handlers[eventType], handler)
}

// Push 推送一个事件
func (bus *IceEventBus) Push(eventType string, args ...interface{}) {
	bus.lock.RLock()
	defer bus.lock.RUnlock()
	if handlers, found := bus.handlers[eventType]; found {
		for _, handler := range handlers {
			go handler.(func(...interface{}))(args...)
		}
	}
}
