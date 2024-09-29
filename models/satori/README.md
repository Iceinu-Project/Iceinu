# Satori Model

参考：https://satori.js.org/zh-CN/introduction.html

氷犬默认的消息事件接收/发送模型，参考了跨平台的Satori标准，可以用于实现适配不同平台的消息适配器。

在使用Satori Model的情况下，可以直接构建Satori XHTML风格的消息进行发送，氷犬会自动将其解析成Satori元素切片并传递到事件系统中。

## 事件类型

### 消息事件
- `PrivateMessageEvent` 私聊消息事件，特指用户直接向机器人发送的消息，在不同平台上的表现可能略有区别，比如部分平台上的私聊实际上是特殊的频道。
- `ChannelMessageEvent` 频道消息事件，用户在频道中发送的消息，部分平台上的群组概念和频道是重合的。（如QQ、微信等）
- `TempMessageEvent` 临时消息事件，一种特别的私聊事件，部分平台允许非好友用户向机器人发送消息，这种消息会被视为临时消息。
- `NotifyEvent` 通知事件，部分平台具有此类事件，如戳一戳。
- `FriendRecallEvent` 撤回事件，用户撤回消息时触发。
- `ChannelRecallEvent` 撤回事件，频道中撤回消息时触发。

### 好友事件

- `RenameEvent` 好友昵称变更事件，好友修改昵称时触发。
- `FriendRequestEvent` 好友请求事件，用户请求添加好友时触发。

### 群组事件

- `GroupJoinEvent` 机器人加入群组事件，机器人加入群组时触发。
- `GroupLeaveEvent` 机器人退出群组事件，机器人退出群组时触发。
- `GroupInvitedEvent` 被邀请入群事件，机器人被邀请入群时触发。
- `GroupMemberJoinEvent` 群成员加入事件，群组有新成员加入时触发。
- `GroupMemberLeaveEvent` 群成员离开事件，群组有成员离开时触发。
- `GroupMuteEvent` 群组禁言事件，群组中有成员被禁言时触发。
- `GroupRoleChangeEvent` 群组角色变更事件，群组中有成员角色变更时触发。
- `GroupRenameEvent` 群组名称变更事件，群组名称变更时触发。
- `GroupTitleChangeEvent` 群组头衔变更事件，部分平台具有这个特性。
- `GroupDigestEvent` 群组头衔变更事件，部分平台具有这个特性。

## API

基于Satori的资源API并参考了其他Bot框架的一些实现，氷犬提供了一系列可供调用的事件API，用于使开发者可以快速与适配器进行交互。

调用API实际上是向事件总线发送事件，部分API需要通过临时增加一个事件监听器监听数据回调事件来实现API调用的异步回调。

### 适配器API

```go
GetImage(url string) string // 获取图片URL，用于解决一些平台的图片防盗链问题
Refresh() // 刷新适配器缓存，用于重新加载适配器配置
```

### Satori API

```go
// 频道相关
func GetChannel(channelId string) *Channel // 获取指定频道信息
func GetChannelList(groupId string) []*Channel // 获取群组中的频道列表
func CreateChannel(groupId string, data *Channel) *Channel // 在群组中创建频道
func UpdateChannel(channelId string, data *Channel) *Channel // 更新群组中指定频道的信息
func DeleteChannel(channelId string) // 删除指定频道
func MuteChannel(channelId string, duration int64) // 禁言指定群组频道，传入时长0时表示解除禁言
func CreateUserChannel(userId string, groupId string) *Channel // 创建用户私聊频道
// 群组相关
func GetGroup(groupId string) *Group // 获取指定群组信息
func GetGroupList() []*Group // 获取机器人的群组列表
func ApproveGroup(messageId string, approve bool, comment string) // 处理来自群组的邀请
// 群组成员相关
func GetGroupMember(groupId string, userId string) *GroupMember // 获取指定群组的指定成员信息
func GetGroupMemberList(groupId string) []*GroupMember // 获取指定群组的成员列表
func KickGroupMember(groupId string, userId string, permanent bool) // 从指定群组中踢出指定成员
func MuteGroupMember(groupId string, userId string, duration int64) // 禁言指定群组成员，传入时长0时表示解除禁言
func ApproveGroupMember(messageId string, approve bool, comment string) // 处理加群请求
// 群组角色相关
func SetGroupMemberRole(groupId string, userId string, roleId string) // 设置指定群组成员的角色
func UnsetGroupMemberRole(groupId string, userId string, roleId string) // 取消指定群组成员的角色
func GetGroupRoleList(groupId string) []*GroupRole // 获取指定群组的角色列表
func CreateGroupRole(groupId string, role *GroupRole) *GroupRole // 在指定群组中创建角色
func UpdateGroupRole(groupId string, roleId string, role *GroupRole) *GroupRole // 修改指定群组中的角色
func DeleteGroupRole(groupId string, roleId string) // 删除指定群组中的角色
// 登录信息相关
func GetLoginInfo() *Login // 获取登录信息
// 消息相关
func SendMessage(channelId string, message string) *Message // 发送消息
func RecallMessage(channelId string, messageId string) // 撤回消息
func EditMessage(channelId string, messageId string, message string) // 编辑消息
func GetMessage(channelId string, messageId string) *Message // 获取消息
func GetMessageList(channelId string, limit int64, offset int64) []*Message // 获取消息列表
// 表态相关
func CreateReaction(channelId string, messageId string, reaction string) // 创建表态
func DeleteReaction(channelId string, messageId string, reaction string, userId string) // 删除表态
func ClearReaction(channelId string, messageId string, reaction string) // 清除表态
func GetReactionList(channelId string, messageId string, reaction string) []*Reaction // 获取表态列表
// 用户相关
func GetUser(userId string) *User // 获取用户信息
func GetFriendList() []*User // 获取好友列表
func ApproveFriend(messageId string, approve bool, comment string) // 处理好友请求
```


### 扩展API

```go
func SendMessageElements(channelId string, elements ...Element) *Message // 直接发送消息元素
func Reply(messageId string, message string) *Message // 回复消息
func ReplyElements(messageId string, elements ...Element) *Message // 直接回复消息元素
```