package adapter

import (
	"github.com/Iceinu-Project/iceinu/elements"
	"github.com/Iceinu-Project/iceinu/resource"
)

// Bot Iceinu的客户端接口API，用于实现iceinu对平台客户端的直接操作
type Bot interface {
	// 首先是基于Satori标准的API
	// 这部分可以参考Satori文档的资源部分，但是有一定的不同，为了方便兼容各种使用方式进行了一些扩展
	// 比如没有直接支持分页这个东西，方便使用
	// https://satori.js.org/zh-CN/resources/channel.html

	GetChannel(channelId string) *resource.Channel                                   // 获取频道信息
	GetChannelList(groupId string) []*resource.Channel                               // 获取指定群组中的频道列表
	GetChannelListByToken(next string) *resource.PagedList                           // 获取指定群组中的频道列表(分页)
	CreateChannel(groupId string, data *resource.Channel) (*resource.Channel, error) // 创建频道
	UpdateChannel(groupId string, data *resource.Channel) error                      // 更新频道
	DeleteChannel(channelId string) error                                            // 删除频道
	MuteChannel(channelId string, duration uint32) error                             // 禁言频道
	CreateUserChannel(userId string, groupId string) (*resource.Channel, error)      // 创建用户（私聊）频道

	GetGroup(groupId string) *resource.Group                                 // 获取群组信息
	GetGroupList() []*resource.Group                                         // 获取群组列表
	GetGroupListByToken(next string) *resource.PagedList                     // 获取群组列表(分页)
	ApproveGroupInvite(messageId string, approve bool, comment string) error // 处理群组邀请bot加入请求

	GetGroupMember(groupId string, userId string) *resource.GroupMember        // 获取指定群组成员信息
	GetGroupMemberList(groupId string) []*resource.GroupMember                 // 获取群组成员列表
	GetGroupMemberListByToken(groupId string, next string) *resource.PagedList // 获取群组成员列表(分页)
	KickGroupMember(groupId string, userId string, permanent bool) error       // 踢出群组成员
	MuteGroupMember(groupId string, userId string, duration uint32) error      // 禁言群组成员
	ApproveGroupRequest(messageId string, approve bool, comment string) error  // 处理群组加入请求

	SetGroupMemberRole(groupId string, userId string, roleId string) error                 // 设置群组成员角色权限
	UnsetGroupMemberRole(groupId string, userId string, roleId string) error               // 取消群组成员角色权限
	GetGroupRoleList(groupId string) []*resource.GroupRole                                 // 获取群组角色权限列表
	GetGroupRoleListByToken(groupId string, next string) *resource.PagedList               // 获取群组角色权限列表(分页)
	CreateGroupRole(groupId string, role *resource.GroupRole) (*resource.GroupRole, error) // 创建群组角色权限
	UpdateGroupRole(groupId string, roleId string, role *resource.GroupRole) error         // 更新群组角色权限
	DeleteGroupRole(groupId string, roleId string) error                                   // 删除群组角色权限

	GetLoginInfo() *resource.Login // 获取登录信息

	SendContent(channelId string, content string) (*resource.Message, error)                                  // 发送纯文本消息
	GetMessage(channelId string, messageId string) (*resource.Message, error)                                 // （从缓存中）获取指定消息
	RecallMessage(channelId string, messageId string) error                                                   // 撤回指定消息
	UpdateMessage(channelId string, messageId string, content string) error                                   // 编辑指定消息
	GetMessageList(channelId string, limit uint32, order bool) []*resource.Message                            // 获取一定数量的频道消息列表
	GetMessageListByRange(channelId string, messageId string, start uint32, count uint32) []*resource.Message // 获取频道消息列表(可指定范围)

	CreateReaction(channelId string, messageId string, emoji string) error                                    // 添加消息反应
	DeleteReaction(channelId string, messageId string, emoji string, userId string) error                     // 删除消息反应
	ClearReaction(channelId string, messageId string, emoji string) error                                     // 清除消息反应
	GetReactionList(channelId string, messageId string, emoji string) []resource.User                         // 获取消息反应列表
	GetReactionListByToken(channelId string, messageId string, emoji string, next string) *resource.PagedList // 获取消息反应列表(分页)

	GetUser(userId string) *resource.User                                        // 获取指定用户信息
	GetFriendList() []*resource.User                                             // 获取好友列表信息
	ApproveFriendRequest(messageId string, approve string, comment string) error // 处理好友请求

	// Iceinu的特有API
	// 其中一部分是对各个平台功能的扩展适配，还有一部分是其他功能的快捷方式

	Send(elements []elements.IceinuMessageElement) (*resource.Message, error) // 直接发送Iceinu通用元素
	SendSatori(satori string) (*resource.Message, error)                      // 发送Satori XHTML格式的消息字符串，自动解析成通用元素并发送
	SendPoke(userId string) error                                             // 发送戳一戳
	SendGroupPoke(groupId string) error                                       // 发送群组戳一戳
	SendChannelPoke(channelId string) error                                   // 发送频道戳一戳

	GetSelfUserId() string                   // 获取自己的用户ID
	GetSelfUserName() string                 // 获取自己的用户名
	GetSelfAvatarUrl() string                // 获取自己的头像URL
	GetSeldUserNickname() string             // 获取自己的昵称
	GetGroupAvatarUrl(groupId string) string // 获取指定群组的头像URL

	RefreshUserListCache() error                                 // 刷新用户列表
	RefreshGroupListCache() error                                // 刷新群组列表
	RefreshGroupMemberCache(groupId string, userId string) error // 刷新指定群组的指定成员的信息
	RefreshGroupAllMembersCache(groupId string) error            // 刷新指定群组所有成员的信息
	RefreshChannelListCache(groupId string) error                // 刷新指定群组的频道列表

	RenameGroup(groupId string, name string) error                          // 重命名群组
	RenameGroupMember(groupId string, userId string, nickname string) error // 重命名群组成员
	RemarkGroup(groupId string, remark string) error                        // 设置群组备注
	SetGroupGlobalMute(groupId string, mute bool) error                     // 设置群组全员禁言
	LeaveGroup(groupId string) error                                        // 退出群组
	SetGroupMemberTitle(groupId string, userId string, title string) error  // 给群组成员设置头衔

	// 这部分功能接口设计主要来自LagrangeGo，但是也可能在其他NTQQ平台上实现

	UploadChannelFile(channelId string, filePath string) error                                     // 向频道上传文件
	UploadGroupFile(groupId string, filePath string, targetFilePath string) error                  // 向群组上传文件
	GetGroupFileSystemInfo(groupId string) *resource.GroupFileSystemInfo                           // 获取群组文件系统信息(暂未确定)
	GetGroupFilesByFolder(groupId string, folderId string) interface{}                             // 获取群组文件夹内的文件列表(暂未确定)
	GetGroupRootFiles(groupId string) interface{}                                                  // 获取群组根目录文件列表(暂未确定)
	MoveGroupFile(groupId string, fileId string, parentFolder string, targetFolderId string) error // 移动群组文件
	DeleteGroupFile(groupId string, fileId string) error                                           // 删除群组文件
	CreateGroupFolder(groupId string, folderName string, parentFolder string) error                // 创建群组文件夹
	RenameGroupFolder(groupId string, folderId string, newFolderName string) error                 // 重命名群组文件夹
	DeleteGroupFolder(groupId string, folderId string) error                                       // 删除群组文件夹

	// GetOriginalClient 获取适配器的原始客户端对象，部分适配器可能不需要这个东西，只是方便直接传递原本的客户端实例
	GetOriginalClient() interface{} // 获取原始客户端对象
}
