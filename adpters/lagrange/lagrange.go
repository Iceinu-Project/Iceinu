package lagrange

import (
	"github.com/Iceinu-Project/Iceinu/adpters"
	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/client/auth"
)

// InfosLagrangeAdapter LagrangeGo适配器元信息
var InfosLagrangeAdapter = adpters.AdapterInfo{
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
	//TODO implement me
	panic("implement me")
}

func (AdapterLagrangeGo) SubscribeEvents() error {
	//TODO implement me
	panic("implement me")
}

func (AdapterLagrangeGo) Start() error {
	//TODO implement me
	panic("implement me")
}

func (AdapterLagrangeGo) GetAdapterInfo() *adpters.AdapterInfo {
	return &InfosLagrangeAdapter
}

func (AdapterLagrangeGo) GetUserTree() *adpters.UserTree {
	//TODO implement me
	panic("implement me")
}
