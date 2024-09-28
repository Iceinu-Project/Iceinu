package main

import (
	"github.com/Iceinu-Project/IceGradient"
	"github.com/Iceinu-Project/Iceinu/adapters"
	"github.com/Iceinu-Project/Iceinu/adapters/lagrange"
	"github.com/Iceinu-Project/Iceinu/config"
	"github.com/Iceinu-Project/Iceinu/ice"
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
	// 输出欢迎日志
	log.Infof("欢迎使用🧊" + gradient.Bold +
		gradient.GradientText("氷犬 Iceinu Bot", "#00d2ff", "#3a7bd5") + gradient.DarkGray + " | " +
		gradient.RGBToANSI(255, 255, 255) +
		gradient.GradientBackgroundText(" 通用的模块化 Go 聊天机器人框架 ", "#00d2ff", "#3a7bd5") +
		gradient.Reset)
	log.Infof("当前版本: " + gradient.Cyan + "Re:β 0.2.4")
	// 初始化数据库连接
	ice.InitLocalDatabase()
	// 设置日志级别
	log.SetLevel(config.IceConf.LogLevel)
	log.Debugf("调试模式已启用")
	// 设置自定义中间件
	SetCustomMiddleWare()
	// 如果启用集群/分布式模式，则尝试和主节点建立连接
	// 输出节点Id和集群模式
	log.Info("当前节点模式: " + gradient.Cyan + config.IceConf.Node.Mode)
	log.Info("Node Id: " + gradient.Cyan + ice.GetSelfNodeId())
	log.Info("MasterNode Id: " + gradient.Cyan + ice.GetMasterNodeId())
	// 启动适配器
	var adapter adapters.IceinuAdapter
	adapter = lagrange.GetAdapter()
	err := adapter.Init()
	if err != nil {
		return
	}
	adapter.Start()
}
