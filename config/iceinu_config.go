package config

type IceinuConfig struct {
	BotName  string         `toml:"bot_name"`
	LogLevel string         `toml:"log_level"`
	Database DatabaseConfig `toml:"database"`
}

type DatabaseConfig struct {
	URL string `toml:"url"`
}

func init() {
	// 注册Iceinu的配置文件
	manager.InitConfig("iceinu.toml", &IceinuConfig{
		BotName:  "Iceinu",
		LogLevel: "INFO",
		Database: DatabaseConfig{
			URL: "sqlite://iceinu.db",
		},
	})
}
