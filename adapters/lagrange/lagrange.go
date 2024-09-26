package lagrange

import (
	"github.com/Iceinu-Project/Iceinu/adapters"
	"github.com/Iceinu-Project/Iceinu/log"
	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/client/auth"
	"github.com/sirupsen/logrus"
	"os"
)

// InfosLagrangeAdapter LagrangeGo适配器元信息
var InfosLagrangeAdapter = adapters.AdapterInfo{
	Name:      "LagrangeGo适配器",
	Version:   "1.0.0",
	Model:     "Satori",
	Platform:  []string{"NTQQ"},
	Author:    []string{"Kyoku"},
	License:   []string{"MIT License"},
	Repo:      "https://github.com/Iceinu-Project/Iceinu",
	Introduce: "内置LagrangeGo的NTQQ适配器，无需外置协议端程序",
}

// Client LagrangeGo客户端实例
var Client *client.QQClient

// AdapterLagrangeGo LagrangeGo适配器
type AdapterLagrangeGo struct{}

func (AdapterLagrangeGo) Init() error {
	// 读取配置文件
	AdapterConfigInit()
	// 日志输出
	appInfo := auth.AppList["linux"]["3.2.10-25765"]
	deviceInfo := auth.NewDeviceInfo(AdapterLagrangeConf.Lagrange.Account)
	qqClientInstance := client.NewClient(uint32(AdapterLagrangeConf.Lagrange.Account), appInfo, AdapterLagrangeConf.Lagrange.SignServer)
	qqClientInstance.SetLogger(GetProtocolLogger())
	qqClientInstance.UseDevice(deviceInfo)

	// 尝试读取签名文件
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
	// 保存Client实例
	Client = qqClientInstance
	return nil
}

func (AdapterLagrangeGo) SubscribeEvents() error {
	EventsBinder()
	return nil
}

func (AdapterLagrangeGo) Start() error {
	// 在函数结束时释放Client并尝试保存签名
	defer Client.Release()
	defer SaveSignature()
	// 登录
	err := Login()
	if err != nil {
		return err
	}
	// 尝试刷新client的所有缓存
	err = Client.RefreshFriendCache()
	if err != nil {
		return err
	}
	err = Client.RefreshAllGroupsInfo()
	if err != nil {
		return err
	}
	err = Client.RefreshAllGroupMembersCache()
	if err != nil {
		return err
	}
	err = Client.RefreshAllRkeyInfoCache()
	if err != nil {
		return err
	}
	return nil
}

func (AdapterLagrangeGo) GetAdapterInfo() *adapters.AdapterInfo {
	return &InfosLagrangeAdapter
}

func (AdapterLagrangeGo) GetUserTree() *adapters.UserTree {
	//TODO implement me
	panic("implement me")
}

// Login 登录
func Login() error {
	// 声明 err 变量并进行错误处理
	err := Client.Login("", "qrcode.png")
	if err != nil {
		log.Error("登录时发生错误:", err)
		return err
	}
	return nil
}

// SaveSignature 保存sign信息
func SaveSignature() {
	data, err := Client.Sig().Marshal()
	if err != nil {
		log.Error("生成签名文件时发生错误err:", err)
		return
	}
	err = os.WriteFile("signature.bin", data, 0644)
	if err != nil {
		log.Error("写入签名文件时发生错误 err:", err)
		return
	}
	log.Info("签名已被写入签名文件")
}
