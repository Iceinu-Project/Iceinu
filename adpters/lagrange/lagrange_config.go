package lagrange

import (
	"github.com/Iceinu-Project/Iceinu/config"
	"github.com/Iceinu-Project/Iceinu/log"
)

var AdapterLagrangeConf *AdapterConfig

// AdapterConfig 适配器配置
type AdapterConfig struct {
	CacheSize   int       `toml:"message_cache_size"`   // 消息缓存大小
	CacheExpire int       `toml:"message_cache_expire"` // 消息缓存过期时间
	Lagrange    LgrConfig `toml:"lagrange"`
}

type LgrConfig struct {
	SignServer      string `toml:"sign_server"`
	MusicSignServer string `toml:"music_sign_server"`
	Account         int    `toml:"account"`
	Password        string `toml:"password"`
}

// AdapterConfigInit 初始化适配器配置
func AdapterConfigInit() {
	AdapterLagrangeConf = &AdapterConfig{
		CacheSize:   200,
		CacheExpire: 600,
		Lagrange: LgrConfig{
			SignServer:      "https://sign.lagrangecore.org/api/sign/25765",
			MusicSignServer: "",
			Account:         0,
			Password:        "",
		},
	}
	err := config.ProcessConfig(AdapterLagrangeConf, "lagrange_config.toml")
	if err != nil {
		log.Errorf("初始化LagrangeGo适配器配置文件失败: %v", err)
		return
	}
}
