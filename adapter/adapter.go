package adapter

// 继承IceAdapter接口即可实现适配器，具体的适配器实现模式参考文档

// IceAdapterMeta 适配器元信息
type IceAdapterMeta struct {
	AdapterName string   // 适配器名称
	Version     string   // 适配器版本
	Platform    string   // 适配器平台
	Author      []string // 适配器作者
	Introduce   string   // 适配器介绍
}

// IceAdapter Iceinu的适配器接口，用于实现适配器实例
type IceAdapter interface {
	GetMeta() *IceAdapterMeta // 获取适配器的元信息
	Init()                    // 适配器初始化
}
