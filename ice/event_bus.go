package ice

import "sync"

// iceinu的事件总线实现，启动后生成全局共享的Bus实例
// 一般情况下事件的推送和订阅由适配器和框架完成，但是也可以自己直接调用Bus实例的方法来绕过已有的事件封装（不推荐）

func init() {
	// 创建事件总线
	Bus = CreateBus()
}

// Bus 全局事件总线实例
var Bus *EventBus

// Middleware 中间件结构体，可以通过实现中间件来对事件的发布过程进行一些处理
type Middleware func(eventType string, payload interface{})

// Event 在事件总线中传递的事件结构体
type Event struct {
	EventType string
	Payload   interface{}
}

// EventChan 事件总线中的事件通道
type EventChan chan Event

// EventBus 事件总线结构体
type EventBus struct {
	subscribers map[string][]EventChan
	lock        sync.RWMutex
	middlewares []Middleware
}

// CreateBus 创建事件总线，这个函数一般不需要单独调用，iceinu会在运行时自动创建事件总线
func CreateBus() *EventBus {
	newBus := &EventBus{
		subscribers: make(map[string][]EventChan),
	}
	return newBus
}

// Subscribe 订阅指定类型的事件
func (bus *EventBus) Subscribe(eventType string, ch EventChan) {
	bus.lock.Lock()
	defer bus.lock.Unlock()
	bus.subscribers[eventType] = append(bus.subscribers[eventType], ch)
}

// Unsubscribe 取消订阅指定类型的事件
func (bus *EventBus) Unsubscribe(eventType string, ch EventChan) {
	bus.lock.Lock()
	defer bus.lock.Unlock()
	subscribers := bus.subscribers[eventType]
	for i, subscriber := range subscribers {
		if subscriber == ch {
			bus.subscribers[eventType] = append(subscribers[:i], subscribers[i+1:]...)
			break
		}
	}
}

// Publish 推送事件到事件总线
func (bus *EventBus) Publish(eventType string, payload interface{}) {
	// 调用所有中间件
	for _, mw := range bus.middlewares {
		mw(eventType, payload)
	}

	// 持有锁
	bus.lock.RLock()
	subscribers, found := bus.subscribers[eventType]
	bus.lock.RUnlock()

	// 进入事件处理
	if found {
		event := Event{EventType: eventType, Payload: payload}
		// 把每个订阅的事件丢给协程处理避免阻塞
		for _, ch := range subscribers {
			go func(ch EventChan) {
				ch <- event
			}(ch)
		}
	}
}

// AddMiddleware 向事件总线中添加自定义的中间件
func (bus *EventBus) AddMiddleware(mw Middleware) {
	bus.middlewares = append(bus.middlewares, mw)
}
