package lagrange

import (
	"github.com/Iceinu-Project/Iceinu/adapters"
	"github.com/Iceinu-Project/Iceinu/cache"
	"github.com/Iceinu-Project/Iceinu/ice"
	"github.com/Iceinu-Project/Iceinu/log"
	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/client/auth"
	"os"
	"os/signal"
	"strconv"
	"syscall"
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

// Cache 消息缓存
var Cache *cache.IceCacheManager

// AdapterLagrangeGo LagrangeGo适配器
type AdapterLagrangeGo struct{}

// Init 初始化适配器
func (a *AdapterLagrangeGo) Init() error {
	// 读取配置文件
	AdapterConfigInit()
	// 初始化消息缓存
	log.Debug("正在初始化LagrangeGo的消息缓存...")
	Cache = cache.NewIceCacheManager(AdapterLagrangeConf.CacheSize, AdapterLagrangeConf.CacheExpire)
	// 日志输出
	appInfo := auth.AppList["linux"]["3.2.10-25765"]
	deviceInfo := auth.NewDeviceInfo(AdapterLagrangeConf.Lagrange.Account)
	qqClientInstance := client.NewClient(uint32(AdapterLagrangeConf.Lagrange.Account), appInfo, AdapterLagrangeConf.Lagrange.SignServer)
	qqClientInstance.SetLogger(GetProtocolLogger())
	qqClientInstance.UseDevice(deviceInfo)

	// 尝试读取签名文件
	data, err := os.ReadFile("signature.bin")
	if err != nil {
		log.Warn("读取签名文件时发生错误:", err)
	} else {
		sig, err := auth.UnmarshalSigInfo(data, true)
		if err != nil {
			log.Warn("加载签名文件时发生错误:", err)
		} else {
			qqClientInstance.UseSig(sig)
		}
	}
	// 保存Client实例
	Client = qqClientInstance
	return nil
}

// SubscribeEvents 订阅事件
func (a *AdapterLagrangeGo) SubscribeEvents() error {
	BindEvents()
	return nil
}

// Start 启动适配器
func (a *AdapterLagrangeGo) Start() error {
	// 在函数结束时释放Client并尝试保存签名
	defer Client.Release()
	defer SaveSignature()
	// 事件订阅
	err := a.SubscribeEvents()
	if err != nil {
		return err
	}
	SetAllSubscribes()
	// 登录
	err = Login()
	if err != nil {
		return err
	}
	// 推送适配器连接事件
	ice.MakeAdapterConnectEvent(InfosLagrangeAdapter.Name, InfosLagrangeAdapter.Model, strconv.Itoa(int(Client.Uin)), Client.NickName())

	// 主协程关闭通道
	mc := make(chan os.Signal, 2)
	signal.Notify(mc, os.Interrupt, syscall.SIGTERM)
	for {
		switch <-mc {
		case os.Interrupt, syscall.SIGTERM:
			return nil
		}
	}
}

// GetAdapterInfo 获取适配器信息
func (a *AdapterLagrangeGo) GetAdapterInfo() *adapters.AdapterInfo {
	return &InfosLagrangeAdapter
}

// GetUserTree 获取用户树
func (a *AdapterLagrangeGo) GetUserTree() *adapters.UserTree {
	// 获取好友列表
	friends := Client.GetCachedAllFriendsInfo()
	var friendList []string
	for friend, _ := range friends {
		friendList = append(friendList, strconv.Itoa(int(friend)))
	}
	// 获取群列表
	groups := Client.GetCachedAllGroupsInfo()
	var groupList []string
	for group, _ := range groups {
		groupList = append(groupList, strconv.Itoa(int(group)))
	}
	// 复制群列表为频道列表
	channelList := make([]string, len(groupList))
	copy(channelList, groupList)
	// 返回用户树
	return &adapters.UserTree{
		SelfId:   strconv.Itoa(int(Client.Uin)),
		Platform: "NTQQ",
		NodeId:   ice.GetSelfNodeId(),
		Users:    friendList,
		Groups:   groupList,
		Channels: channelList,
	}
}

// Refresh 刷新适配器数据
func (a *AdapterLagrangeGo) Refresh() error {
	// 刷新LagrangeGo的缓存
	err := Client.RefreshFriendCache()
	if err != nil {
		return err
	}
	err = Client.RefreshAllRkeyInfoCache()
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
	return nil
}

func (a *AdapterLagrangeGo) ClearCache() {
	Cache.Clear()
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

func GetAdapter() adapters.IceinuAdapter {
	return &AdapterLagrangeGo{}
}
