package lagrange

import "github.com/Iceinu-Project/iceinu/config"

type AdapterLagrangeConfig struct {
	Account      int    `toml:"account"`
	Password     string `toml:"password"`
	SignUrl      string `toml:"sign_url"`
	MusicSignUrl string `toml:"music_sign_url"`
}

var manager = config.GetManager()

func RegisterConfig() {
	manager.InitConfig("lagrange.toml", &AdapterLagrangeConfig{
		Account:      0,
		Password:     "",
		SignUrl:      "https://sign.lagrangecore.org/api/sign/25765",
		MusicSignUrl: "",
	})
}
