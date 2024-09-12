package config

import (
	"encoding/json"
	"fmt"
	"github.com/Iceinu-Project/iceinu/log"
	"os"
)

var logger = log.GetLogger()

type Config struct {
	SignServerURL      string `json:"SignServerURL"`
	MusicSignServerURL string `json:"MusicSignServerURL"`
	IgnoreSelf         bool   `json:"IgnoreSelf"`
	Master             struct {
		ListenHost string `json:"ListenHost"`
		ListenPort int    `json:"ListenPort"`
	} `json:"Master"`
	Node struct {
		IsNode    bool   `json:"IsNode"`
		MasterURL string `json:"MasterURL"`
	} `json:"Node"`
}

var config *Config

// LoadConfig 从 JSON 文件加载配置
func LoadConfig() (*Config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	decoder := json.NewDecoder(file)
	var config Config
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// SaveConfig 将配置保存到 JSON 文件
func SaveConfig(config *Config) error {
	file, err := os.Create("config.json")
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(config)
	if err != nil {
		return err
	}

	return nil
}

// InitConfig 对配置进行初始化处理
func InitConfig() {
	logger.Debug("正在初始化配置文件读取...")
	// 检测是否存在config.json文件
	_, err := os.Stat("config.json")
	if err != nil {
		// 如果不存在，则创建一个新的配置文件
		config = &Config{
			SignServerURL:      "",
			MusicSignServerURL: "",
			IgnoreSelf:         false,
			Master: struct {
				ListenHost string `json:"ListenHost"`
				ListenPort int    `json:"ListenPort"`
			}{
				ListenHost: "127.0.0.1",
				ListenPort: 8080,
			},
			Node: struct {
				IsNode    bool   `json:"IsNode"`
				MasterURL string `json:"MasterURL"`
			}{
				IsNode:    false,
				MasterURL: "",
			},
		}
		err := SaveConfig(config)
		if err != nil {
			return
		}
		// 等待用户输入回车
		logger.Warn("检测到程序所在目录中没有config.json文件，已自动创建一个新的配置文件，请在编辑之后按回车继续。")
		_, err = fmt.Scanln()
		if err != nil {
			return
		}
	}

	config, err = LoadConfig()
	if err != nil {
		return
	}
	logger.Debug("配置文件读取成功")
}

// GetConfig 获取配置
func GetConfig() *Config {
	return config
}
