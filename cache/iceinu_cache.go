package cache

import (
	"github.com/Iceinu-Project/iceinu/config"
	"github.com/Iceinu-Project/iceinu/log"
)

var IceCache *IceCacheManager

// 初始化缓存管理器
func init() {
	log.Debugf("正在初始化内置缓存管理器...")
	IceCache = NewIceCacheManager(config.IceConf.Cache.MaxCacheSize, config.IceConf.Cache.CacheExpire)
}
