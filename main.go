package main

import (
	"flag"
	"gtihub.com/Iceinu-Project/iceinu/adapter"
	"gtihub.com/Iceinu-Project/iceinu/adapter/lagrange"
	"gtihub.com/Iceinu-Project/iceinu/config"
	"gtihub.com/Iceinu-Project/iceinu/ice"
	"gtihub.com/Iceinu-Project/iceinu/logger"
)

func main() {
	logger.Infof("正在启动Iceinu，请稍等...")

	// 解析启动参数
	IsDebug := flag.Bool("debug", false, "是否启用调试模式")
	flag.Parse()

	// 设置调试日志等级
	if *IsDebug {
		logger.SetLevel("DEBUG")
	}

	// 初始化适配器，默认使用LagrangeGo适配器
	var a adapter.IceAdapter
	a = &lagrange.AdapterLagrange{}

	// 读取配置文件
	lagrange.RegisterConfig()

	// 处理所有配置文件
	cm := config.GetManager()
	err := cm.LoadConfigs()
	if err != nil {
		logger.Fatalf("读取配置文件时发生错误：%v", err)
	}

	// 设置日志等级
	iceConf := cm.Get("iceinu.toml").(*config.IceinuConfig)
	if !*IsDebug {
		logger.Debugf("调试模式已启用，跳过默认的日志等级设置")
		if iceConf.LogLevel != "" {
			logger.SetLevel(iceConf.LogLevel)
		}
	}

	// 自定义事件总线中间件
	ice.Bus.AddMiddleware(func(eventType string, payload interface{}) {
		logger.Debugf("事件推送：%s", eventType)
	})

	// 启动适配器初始化
	a.Init()
}
