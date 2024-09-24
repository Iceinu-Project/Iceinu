package main

import (
	"github.com/Iceinu-Project/iceinu/config"
	"github.com/Iceinu-Project/iceinu/log"
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
}
