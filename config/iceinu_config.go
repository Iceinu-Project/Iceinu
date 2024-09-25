package config

import "github.com/Iceinu-Project/Iceinu/log"

var IceConf *IceinuConfig

type IceinuConfig struct {
	LogLevel string         `toml:"log_level"` // 日志级别
	Node     NodeConfig     `toml:"node"`      // 节点配置
	Database DatabaseConfig `toml:"database"`  // 数据库配置
	Cache    CacheConfig    `toml:"cache"`     // 缓存配置
}

type NodeConfig struct {
	IsEnableNode bool   `toml:"is_enable_node"` // 是否启用节点连接
	Mode         string `toml:"node_mode"`      // 运行模式，可选值为Dist（分布式），Static（静态集群），Dynamic（动态集群）
	IsMaster     bool   `toml:"is_master"`      // 是否为主节点
	MasterURL    string `toml:"master_url"`     // 主节点地址
}

type DatabaseConfig struct {
	IsEnableRemoteDatabase bool   `toml:"is_enable_remote_database"` // 是否启用远程数据库
	DatabaseType           string `toml:"database_type"`             // 数据库类型，可选值为MySQL，PostgreSQL
	DatabaseURL            string `toml:"database_url"`              // 数据库连接地址
}

type CacheConfig struct {
	MaxCacheSize uint32 `toml:"max_cache_size"` // 内置缓存最大容量
	CacheExpire  uint32 `toml:"cache_expire"`   // 内置缓存过期时间
}

// IceConfigInit 初始化内置配置文件
func IceConfigInit() {
	log.Debugf("正在初始化内置配置文件...")
	// 初始化内置配置文件
	IceConf = &IceinuConfig{
		LogLevel: "INFO",
		Node: NodeConfig{
			IsEnableNode: false,
			Mode:         "Dist",
			IsMaster:     true,
			MasterURL:    "",
		},
		Database: DatabaseConfig{
			IsEnableRemoteDatabase: false,
			DatabaseType:           "PostgreSQL",
			DatabaseURL:            "",
		},
		Cache: CacheConfig{
			MaxCacheSize: 100,
			CacheExpire:  600,
		},
	}
	err := ProcessConfig(IceConf, "ice_config.toml")
	if err != nil {
		log.Errorf("初始化内置配置文件失败: %v", err)
		return
	}
}
