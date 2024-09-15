package lagrange

import (
	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/client/auth"
	"github.com/sirupsen/logrus"
	"gtihub.com/Iceinu-Project/iceinu/adapter"
	"gtihub.com/Iceinu-Project/iceinu/event"
	"gtihub.com/Iceinu-Project/iceinu/ice"
	"gtihub.com/Iceinu-Project/iceinu/logger"
	"os"
	"time"
)

type AdapterLagrange struct {
}

type Bot struct {
	*client.QQClient
}

var LgrClient *Bot

func (lgr *AdapterLagrange) GetMeta() *adapter.IceAdapterMeta {
	return &adapter.IceAdapterMeta{
		AdapterName: "Lagrange Adapter",
		Version:     "Beta0.0.1",
		Platform:    "NTQQ",
		Author: []string{
			"Kyoku",
		},
		Introduce: "基于Lagrange的NTQQ适配器，内置了LagrangeGo，无需再连接额外的协议端。",
	}
}

func (lgr *AdapterLagrange) Init() {
	logger.Infof("正在初始化Lagrange适配器，适配器当前版本: %s", lgr.GetMeta().Version)

	// 发送一个适配器初始化事件
	ice.Bus.Publish("AdapterInitEvent", event.AdapterInitEvent{
		Timestamp:   time.Time{},
		AdapterMeta: lgr.GetMeta(),
	})

	// 创建LagrangeGo的客户端实例
	plogger := logger.GetProtocolLogger()
	appInfo := auth.AppList["linux"]["3.2.10-25765"]
	deviceInfo := auth.NewDeviceInfo(3291183200)
	qqClientInstance := client.NewClient(3291183200, appInfo, "https://sign.lagrangecore.org/api/sign/25765")
	qqClientInstance.SetLogger(plogger)
	qqClientInstance.UseDevice(deviceInfo)

	data, err := os.ReadFile("signature.bin")
	if err != nil {
		logrus.Warnln("读取签名文件时发生错误:", err)
	} else {
		sig, err := auth.UnmarshalSigInfo(data, true)
		if err != nil {
			logrus.Warnln("加载签名文件时发生错误:", err)
		} else {
			qqClientInstance.UseSig(sig)
		}
	}
	LgrClient = &Bot{QQClient: qqClientInstance}

	defer LgrClient.Release()
	defer SaveSignature()

	// 登录
	err = Login()
	if err != nil {
		return
	}

	// 设置事件订阅器，将LagrangeGo的事件转换并发送到iceinu的事件总线上
	SetAllSubscribes()
}

// Login 登录
func Login() error {
	// 声明 err 变量并进行错误处理
	err := LgrClient.Login("", "qrcode.png")
	if err != nil {
		logrus.Errorln("登录时发生错误:", err)
		return err
	}
	// 推送登录事件
	ice.Bus.Publish("AdapterLoginEvent", event.AdapterInitEvent{
		Timestamp: time.Time{},
	})
	return nil
}

// SaveSignature 保存sign信息
func SaveSignature() {
	data, err := LgrClient.Sig().Marshal()
	if err != nil {
		logrus.Errorln("生成签名文件时发生错误err:", err)
		return
	}
	err = os.WriteFile("signature.bin", data, 0644)
	if err != nil {
		logrus.Errorln("写入签名文件时发生错误 err:", err)
		return
	}
	logrus.Infoln("签名已被写入签名文件")
}
