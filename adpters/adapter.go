package adpters

// IceinuAdapter Iceinu的适配器接口，编写适配器时需要实现这个接口
type IceinuAdapter interface {
	Init() error                  // 初始化适配器
	SubscribeEvents() error       // 订阅事件，用于操作适配器客户端
	Start() error                 // 启动适配器
	GetAdapterInfo() *AdapterInfo // 获取适配器元信息
	GetUserTree() *UserTree       // 获取完整用户树
}

// AdapterInfo 适配器元信息
type AdapterInfo struct {
	Name      string   // 适配器名称
	Version   string   // 适配器版本
	Model     string   // 适配器模型
	Platform  []string // 适配器平台
	Author    []string // 适配器作者
	License   []string // 适配器许可证
	Repo      string   // 适配器仓库地址
	Introduce string   // 适配器简介
}

// UserTree 用户树结构
//
// Iceinu由于本身设计了分布式/集群式的架构，所以需要保证各个节点不会重复处理数据，这需要维护一个用户树结构
//
// 简而言之，每个适配器和客户端连接时都会将客户端的频道/群组/好友信息处理成用户树结构，这个用户树结构会被上传到主节点
//
// 主节点会将所有适配器的用户树结构根据优先级合并成一个完整的用户树结构，广播给每个子节点，从而限制每个节点可以处理的用户范围
//
// 当节点创建新连接/失去可用性/更新用户数据时会重新向主节点发送用户树结构，主节点会根据用户树结构更新节点的用户范围
type UserTree struct {
	SelfId   string   // 适配器自身ID
	Platform string   // 平台
	Users    []string // 用户列表
	Groups   []string // 群组列表
	Channels []string // 频道列表
}
