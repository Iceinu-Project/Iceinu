package ice

import (
	"math/rand"
	"sync"
	"time"
)

// 全局随机数生成器和互斥锁
var (
	rng  *rand.Rand
	once sync.Once
	mu   sync.Mutex
)

// initRNG 初始化随机数生成器
func initRNG() {
	source := rand.NewSource(time.Now().UnixNano())
	rng = rand.New(source)
}

// GenerateEventID 使用 math/rand 生成随机的 uint64 事件 ID
func GenerateEventID() uint64 {
	once.Do(initRNG)
	mu.Lock()
	defer mu.Unlock()
	return rng.Uint64()
}
