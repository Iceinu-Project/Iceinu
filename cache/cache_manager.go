package cache

import (
	"bytes"
	"encoding/gob"
	"github.com/Iceinu-Project/Iceinu/log"
	"github.com/coocood/freecache"
)

// IceCacheManager Iceinu的缓存管理器实例，实现了一系列的缓存管理方法，便于直接使用
type IceCacheManager struct {
	MaxSize    uint32           // 缓存最大容量，单位为MB
	ExpireTime uint32           // 缓存过期时间，单位为秒
	Cache      *freecache.Cache // 缓存实例
}

// NewIceCacheManager 创建新的缓存管理器实例
func NewIceCacheManager(maxSize, expireTime uint32) *IceCacheManager {
	log.Debugf("缓存管理器初始化完成，最大容量：%dMB，过期时间：%d秒", maxSize, expireTime)
	return &IceCacheManager{
		MaxSize:    maxSize * 1024 * 1024,
		ExpireTime: expireTime,
		Cache:      freecache.NewCache(int(maxSize * 1024 * 1024)),
	}
}

// GetCache 直接获取缓存实例
func (icm *IceCacheManager) GetCache() *freecache.Cache {
	return icm.Cache
}

// Set 设置缓存数据
func (icm *IceCacheManager) Set(key string, value interface{}) error {
	// 创建一个字节缓冲区
	var buf bytes.Buffer

	// 创建一个新的编码器，并将值编码到缓冲区中
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(value)
	if err != nil {
		return err
	}

	// 将序列化后的数据存储到缓存中
	return icm.Cache.Set([]byte(key), buf.Bytes(), int(icm.ExpireTime))
}

// Get 获取缓存数据，value 必须是指针类型以便解码后进行填充
func (icm *IceCacheManager) Get(key string, value interface{}) error {
	// 从缓存中获取数据
	data, err := icm.Cache.Get([]byte(key))
	if err != nil {
		return err
	}

	// 创建一个字节缓冲区，并使用解码器解码数据到提供的 value 中
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(value)
}

// Del 删除缓存数据
func (icm *IceCacheManager) Del(key string) {
	icm.Cache.Del([]byte(key))
}

// Clear 清空缓存
func (icm *IceCacheManager) Clear() {
	icm.Cache.Clear()
}

// Update 更新缓存数据
func (icm *IceCacheManager) Update(key string, value interface{}) error {
	// 先删除旧数据
	icm.Del(key)

	// 再设置新数据
	return icm.Set(key, value)
}

// GetMaxSize 获取缓存最大容量
func (icm *IceCacheManager) GetMaxSize() uint32 {
	return icm.MaxSize
}

// GetExpireTime 获取缓存过期时间
func (icm *IceCacheManager) GetExpireTime() uint32 {
	return icm.ExpireTime
}

// SetWithExpire 设置缓存数据并同时指定过期时间
func (icm *IceCacheManager) SetWithExpire(key string, value interface{}, expire int) error {
	// 创建一个字节缓冲区
	var buf bytes.Buffer

	// 创建一个新的编码器，并将值编码到缓冲区中
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(value)
	if err != nil {
		return err
	}

	// 将序列化后的数据存储到缓存中
	return icm.Cache.Set([]byte(key), buf.Bytes(), expire)
}
