package ice

import (
	"time"

	"github.com/Iceinu-Project/iceinu/resource"
)

// PlatformEvent Iceinu的基础平台事件结构体，基本参考了Satori的设计，用于实现跨平台的统一事件格式
//
// 当然，也参考了Satori的资源系统
type PlatformEvent struct {
	EventId   uint64                // 事件ID
	EventType string                // 事件类型
	Platform  string                // 接收者平台名称
	SelfId    string                // 接收者平台账号
	Timestamp time.Time             // 事件推送的时间戳
	Argv      *resource.Argv        // 交互指令
	Button    *resource.Button      // 交互按钮
	Channel   *resource.Channel     // 事件所属的频道
	Group     *resource.Group       // 事件所属的群组
	Login     *resource.Login       // 事件的登录信息
	Member    *resource.GroupMember // 事件的目标成员
	Message   *resource.Message     // 事件的消息
	Operator  *resource.User        // 事件的操作者
	Role      *resource.GroupRole   // 事件的目标角色
	User      *resource.User        // 事件的目标用户
}
