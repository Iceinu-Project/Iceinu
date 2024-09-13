package lagrange

import (
	"github.com/Iceinu-Project/iceinu/log"
	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/client/auth"
	"github.com/sirupsen/logrus"
	"os"
)

type Bot struct {
	*client.QQClient
}

var LgrClient *Bot

func Init() {
	logger := log.GetProtocolLogger()
	appInfo := auth.AppList["linux"]["3.2.10-25765"]
	deviceInfo := auth.NewDeviceInfo(3291183200)
	qqClientInstance := client.NewClient(3291183200, appInfo, "https://sign.lagrangecore.org/api/sign/25765")
	qqClientInstance.SetLogger(logger)
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
}

// Login 登录
func Login() error {
	// 声明 err 变量并进行错误处理
	err := LgrClient.Login("", "qrcode.png")
	if err != nil {
		logrus.Errorln("登录时发生错误:", err)
		return err
	}
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
