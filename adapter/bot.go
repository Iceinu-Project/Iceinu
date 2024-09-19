package adapter

import (
	"github.com/Iceinu-Project/iceinu/elements"
	"github.com/Iceinu-Project/iceinu/resource"
)

// Bot Iceinu的客户端接口API，用于实现iceinu对平台客户端的直接操作
type Bot interface {
	// 首先是基于Satori标准的API
	// 这部分可以参考Satori文档的资源部分，但是有一定的不同
	// 比如没有直接支持分页这个东西，方便使用
	// https://satori.js.org/zh-CN/resources/channel.html

	GetChannel(channelId string) *resource.Channel
	GetChannelList(groupId string) []*resource.Channel
	CreateChannel(groupId string, data *resource.Channel) (*resource.Channel, error)
	UpdateChannel(groupId string, data *resource.Channel) error
	DeleteChannel(channelId string) error
	MuteChannel(channelId string, duration uint32) error
	CreateUserChannel(userId string, groupId string) (*resource.Channel, error)

	GetGroup(groupId string) *resource.Group
	GetGroupList() []*resource.Group
	ApproveGroupInvite(messageId string, approve bool, comment string) error

	GetGroupMember(groupId string, userId string) *resource.GroupMember
	GetGroupMemberList(groupId string) []*resource.GroupMember
	KickGroupMember(groupId string, userId string, permanent bool) error
	MuteGroupMember(groupId string, userId string, duration uint32) error
	ApproveGroupRequest(messageId string, approve bool, comment string) error

	SetGroupMemberRole(groupId string, userId string, roleId string) error
	UnsetGroupMemberRole(groupId string, userId string, roleId string) error
	GetGroupRoleList(groupId string)
	CreateGroupRole(groupId string, role *resource.GroupRole) (*resource.GroupRole, error)
	UpdateGroupRole(groupId string, roleId string, role *resource.GroupRole) error
	DeleteGroupRole(groupId string, roleId string) error

	GetLoginInfo() *resource.Login

	SendContent(channelId string, content string) (*resource.Message, error)
	GetMessage(channelId string, messageId string) (*resource.Message, error)
	RecallMessage(channelId string, messageId string) error
	UpdateMessage(channelId string, messageId string, content string) error
	GetMessageList(channelId string, limit uint32, order bool) []*resource.Message

	CreateReaction(channelId string, messageId string, emoji string) error
	DeleteReaction(channelId string, messageId string, emoji string, userId string) error
	ClearReaction(channelId string, messageId string, emoji string)
	GetReactionList(channelId string, messageId string, emoji string) []resource.User

	GetUser(userId string) *resource.User
	GetFriendList() []*resource.User
	ApproveFriendRequest(messageId string, approve string, comment string) error

	// Iceinu的特有API
	// 其中一部分是对各个平台功能的扩展适配，还有一部分是其他功能的快捷方式

	Send(elements []elements.IceinuMessageElement) (*resource.Message, error) // 直接发送Iceinu通用元素
	SendSatori(satori string) (*resource.Message, error)                      // 发送Satori XHTML格式的消息字符串，自动解析成通用元素并发送
	SendPoke(userId string) error                                             // 发送戳一戳
	SendGroupPoke(groupId string) error                                       // 发送群组戳一戳
	SendChannelPoke(channelId string) error                                   // 发送频道戳一戳

	GetSelfUserId() string       // 获取自己的用户ID
	GetSelfUserName() string     // 获取自己的用户名
	GetSelfAvatarUrl() string    // 获取自己的头像URL
	GetSeldUserNickname() string // 获取自己的昵称

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
	GetGroupFileSystemInfo(groupId string) interface{}                                             // 获取群组文件系统信息(暂未确定)
	GetGroupFilesByFolder(groupId string, folderId string) interface{}                             // 获取群组文件夹内的文件列表(暂未确定)
	GetGroupRootFiles(groupId string) interface{}                                                  // 获取群组根目录文件列表(暂未确定)
	MoveGroupFile(groupId string, fileId string, parentFolder string, targetFolderId string) error // 移动群组文件
	DeleteGroupFile(groupId string, fileId string) error                                           // 删除群组文件
	CreateGroupFolder(groupId string, folderName string, parentFolder string) error                // 创建群组文件夹
	RenameGroupFolder(groupId string, folderId string, newFolderName string) error                 // 重命名群组文件夹
	DeleteGroupFolder(groupId string, folderId string) error                                       // 删除群组文件夹
}
