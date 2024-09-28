package ice

import (
	"gorm.io/gorm"
	"time"
)

// IceinuNodeData 氷犬的节点元信息数据表，框架启动时自动初始化
type IceinuNodeData struct {
	gorm.Model
	NodeId       string // 节点ID，一般是UUID，每一个节点只会有一个UUID，除非数据被重置
	AdapterModel string // 适配器模型
}

// IceinuPluginList 氷犬的插件信息数据表，用于存储插件的启用状态，并在各个节点之间同步
type IceinuPluginList struct {
	PluginId       string    `gorm:"primaryKey"` // 插件ID，一般是名字
	IsEnabled      bool      // 是否启用
	UpdateTime     time.Time // 操作时间
	IsWhiteList    bool      // 是否是白名单模式，这个模式会反转禁用逻辑
	BanUsers       string    // 对哪些用户禁用该插件的触发器
	BanGroups      string    // 对哪些群禁用该插件的触发器
	BanChannels    string    // 对哪些频道禁用该插件的触发器
	BanNodes       string    // 对哪些节点禁用该插件
	IsDetectUpdate bool      // 是否检测更新
}
