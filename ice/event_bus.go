package ice

import (
	"github.com/google/uuid"
	"sync"
)

var Bus *EventBus

// 初始化事件总线
func init() {
	Bus = NewEventBus()
}

// EventHandler 事件处理函数类型
type EventHandler func(event *IceinuEvent)

// PublishMiddleware 发布中间件类型
type PublishMiddleware func(event *IceinuEvent, next func(event *IceinuEvent))

// SubscribeMiddleware 订阅中间件类型
type SubscribeMiddleware func(event *IceinuEvent, next func(event *IceinuEvent))

// 订阅结构体，包含订阅 ID 和处理函数
type subscription struct {
	id      string
	handler EventHandler
}

// EventBus 事件总线结构
type EventBus struct {
	subscribers          map[uint8]map[string][]subscription // 按类型和摘要存储订阅者
	globalPublishMWs     []PublishMiddleware                 // 全局发布中间件
	typePublishMWs       map[uint8][]PublishMiddleware       // 指定类型发布中间件
	summaryPublishMWs    map[string][]PublishMiddleware      // 指定摘要发布中间件
	subscribeMiddlewares []SubscribeMiddleware               // 订阅者接收事件中间件
	lock                 sync.RWMutex
}

// NewEventBus 创建新的事件总线
func NewEventBus() *EventBus {
	return &EventBus{
		subscribers:       make(map[uint8]map[string][]subscription),
		typePublishMWs:    make(map[uint8][]PublishMiddleware),
		summaryPublishMWs: make(map[string][]PublishMiddleware),
	}
}

// 生成唯一的订阅 ID
func generateSubscriberID() string {
	return uuid.New().String()
}

// Subscribe 订阅事件，返回订阅 ID
func (bus *EventBus) Subscribe(eventType uint8, summary string, handler EventHandler) string {
	bus.lock.Lock()
	defer bus.lock.Unlock()

	if bus.subscribers[eventType] == nil {
		bus.subscribers[eventType] = make(map[string][]subscription)
	}

	subID := generateSubscriberID()
	sub := subscription{
		id:      subID,
		handler: handler,
	}

	bus.subscribers[eventType][summary] = append(bus.subscribers[eventType][summary], sub)

	return subID
}

// Unsubscribe 取消订阅，使用订阅 ID
func (bus *EventBus) Unsubscribe(eventType uint8, summary string, subID string) {
	bus.lock.Lock()
	defer bus.lock.Unlock()

	subs := bus.subscribers[eventType][summary]
	for i, sub := range subs {
		if sub.id == subID {
			bus.subscribers[eventType][summary] = append(subs[:i], subs[i+1:]...)
			break
		}
	}
}

// GetSubscribers 获取订阅者列表，返回订阅 ID 和处理函数
func (bus *EventBus) GetSubscribers(eventType uint8, summary string) []subscription {
	bus.lock.RLock()
	defer bus.lock.RUnlock()

	return bus.subscribers[eventType][summary]
}

// UseGlobalPublishMiddleware 添加全局发布中间件
func (bus *EventBus) UseGlobalPublishMiddleware(middleware PublishMiddleware) {
	bus.lock.Lock()
	defer bus.lock.Unlock()

	bus.globalPublishMWs = append(bus.globalPublishMWs, middleware)
}

// UseTypePublishMiddleware 添加指定类型发布中间件
func (bus *EventBus) UseTypePublishMiddleware(eventType uint8, middleware PublishMiddleware) {
	bus.lock.Lock()
	defer bus.lock.Unlock()

	bus.typePublishMWs[eventType] = append(bus.typePublishMWs[eventType], middleware)
}

// UseSummaryPublishMiddleware 添加指定摘要发布中间件
func (bus *EventBus) UseSummaryPublishMiddleware(summary string, middleware PublishMiddleware) {
	bus.lock.Lock()
	defer bus.lock.Unlock()

	bus.summaryPublishMWs[summary] = append(bus.summaryPublishMWs[summary], middleware)
}

// UseSubscribeMiddleware 添加订阅者接收事件中间件
func (bus *EventBus) UseSubscribeMiddleware(middleware SubscribeMiddleware) {
	bus.lock.Lock()
	defer bus.lock.Unlock()

	bus.subscribeMiddlewares = append(bus.subscribeMiddlewares, middleware)
}

// Publish 发布事件
func (bus *EventBus) Publish(event *IceinuEvent) {
	// 最终的发布函数
	finalPublish := func(event *IceinuEvent) {
		bus.lock.RLock()
		handlers := bus.collectHandlers(event)
		bus.lock.RUnlock()

		for _, handler := range handlers {
			wrappedHandler := bus.wrapSubscribeMiddlewares(handler)
			go wrappedHandler(event)
		}
	}

	// 包装发布中间件
	wrappedPublish := bus.wrapPublishMiddlewares(event, finalPublish)

	// 异步执行发布函数
	go wrappedPublish(event)
}

// 收集订阅者的处理函数
func (bus *EventBus) collectHandlers(event *IceinuEvent) []EventHandler {
	var handlers []EventHandler

	// 收集按类型订阅的处理函数
	if summaries, ok := bus.subscribers[event.Type]; ok {
		// 收集按摘要订阅的处理函数
		if subs, ok := summaries[event.Summary]; ok {
			for _, sub := range subs {
				handlers = append(handlers, sub.handler)
			}
		}
	}

	return handlers
}

// 包装发布中间件
func (bus *EventBus) wrapPublishMiddlewares(event *IceinuEvent, finalPublish func(event *IceinuEvent)) func(event *IceinuEvent) {
	bus.lock.RLock()
	var middlewares []PublishMiddleware

	// 添加全局发布中间件
	middlewares = append(middlewares, bus.globalPublishMWs...)

	// 添加指定类型的发布中间件
	if mw, ok := bus.typePublishMWs[event.Type]; ok {
		middlewares = append(middlewares, mw...)
	}

	// 添加指定摘要的发布中间件
	if mw, ok := bus.summaryPublishMWs[event.Summary]; ok {
		middlewares = append(middlewares, mw...)
	}
	bus.lock.RUnlock()

	// 按顺序应用中间件
	for i := len(middlewares) - 1; i >= 0; i-- {
		next := finalPublish
		mw := middlewares[i]
		finalPublish = func(event *IceinuEvent) {
			mw(event, next)
		}
	}
	return finalPublish
}

// 包装订阅者中间件
func (bus *EventBus) wrapSubscribeMiddlewares(handler EventHandler) EventHandler {
	bus.lock.RLock()
	middlewares := bus.subscribeMiddlewares
	bus.lock.RUnlock()

	// 按顺序应用中间件
	for i := len(middlewares) - 1; i >= 0; i-- {
		next := handler
		mw := middlewares[i]
		handler = func(event *IceinuEvent) {
			mw(event, next)
		}
	}
	return handler
}

func UseGlobalPublishMiddleware(middleware PublishMiddleware) {
	Bus.UseGlobalPublishMiddleware(middleware)
}

func UseTypePublishMiddleware(eventType uint8, middleware PublishMiddleware) {
	Bus.UseTypePublishMiddleware(eventType, middleware)
}

func UseSummaryPublishMiddleware(summary string, middleware PublishMiddleware) {
	Bus.UseSummaryPublishMiddleware(summary, middleware)
}

func UseSubscribeMiddleware(middleware SubscribeMiddleware) {
	Bus.UseSubscribeMiddleware(middleware)
}
