package resource

type Channel struct {
	Id       string // 频道ID
	Type     uint8  // 频道类型, 0: 文本频道, 1: 私聊频道，2：分类频道，3：语音频道
	Name     string // 频道名称
	ParentId string // 父频道ID
}
