package resource

type Login struct {
	User      *User    // 用户对象
	SelfId    string   // 平台账号
	Platform  string   // 平台名称
	Status    uint8    // 登录状态，0为离线，1为在线，2为连接中，3为断开连接，4为重新连接
	Features  []string // 平台特性列表
	ProxyUrls []string // 代理路由列表
}
