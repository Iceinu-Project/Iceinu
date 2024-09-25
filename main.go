package main

import (
	"github.com/Iceinu-Project/IceGradient"
	"github.com/Iceinu-Project/Iceinu/config"
	"github.com/Iceinu-Project/Iceinu/log"
)

// Iceinu的程序入口
// 可以参照文档来对其进行修改

func main() {
	// 定义日志格式
	formatter := &LogFormatter{}
	log.SetFormatter(formatter)
	log.Infof("正在启动Iceinu......")

	// 初始化内置配置文件读取
	config.IceConfigInit()
	// 设置日志级别
	log.SetLevel(config.IceConf.LogLevel)
	log.Debugf("调试模式已启用")
	// 输出欢迎日志
	log.Infof("欢迎使用🧊" + gradient.Bold +
		gradient.GradientText("氷犬 Iceinu Bot", "#00d2ff", "#3a7bd5") + gradient.DarkGray + " | " +
		gradient.RGBToANSI(255, 255, 255) +
		gradient.GradientBackgroundText(" 通用的模块化 Go 聊天机器人框架 ", "#00d2ff", "#3a7bd5") +
		gradient.Reset)
	log.Infof("当前版本: " + gradient.Cyan + "β0.1.3")
}
