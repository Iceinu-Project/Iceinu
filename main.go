package main

import (
	"flag"
	"github.com/Iceinu-Project/iceinu/config"
	"github.com/Iceinu-Project/iceinu/handler"
	"github.com/Iceinu-Project/iceinu/lagrange"
	"github.com/Iceinu-Project/iceinu/log"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 定义命令行参数
	isNodeEnabled := flag.Bool("node", false, "启用子节点模式")
	isDebug := flag.Bool("debug", false, "输出调试模式日志")
	flag.Parse()

	// 初始化日志
	logger := log.GetLogger()

	// 检测并输出调试参数
	if *isDebug {
		logger.SetLevel(logrus.DebugLevel)
		logger.Debug("*Debug模式已启用")
		if *isNodeEnabled {
			logger.Debug("*子节点模式已启用")
		}
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}
	// 读取配置文件
	config.InitConfig()

	conf := config.GetConfig()

	logger.Info("🧊Iceinu 正在启动...")
	logger.Info("当前版本: v0.0.1")
	logger.Info("当前配置: ", conf)

	lagrange.Init()
	lagrange.Login()

	handler.BindHandler()

	lagrange.SetIceinuHandler()
	lagrange.SetAllSubscribes()

	defer lagrange.LgrClient.Release()
	defer lagrange.SaveSignature()

	// 主协程关闭通道
	mc := make(chan os.Signal, 2)
	signal.Notify(mc, os.Interrupt, syscall.SIGTERM)
	for {
		switch <-mc {
		case os.Interrupt, syscall.SIGTERM:
			return
		}
	}
}
