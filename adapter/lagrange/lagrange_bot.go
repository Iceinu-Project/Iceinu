package lagrange

import (
	"github.com/Iceinu-Project/iceinu/elements"
	"github.com/Iceinu-Project/iceinu/resource"
	"github.com/LagrangeDev/LagrangeGo/message"
	"strconv"
	"time"
)

type BotLagrange struct {
}

func (b BotLagrange) GetChannel(channelId string) *resource.Channel {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) GetChannelList(groupId string) []*resource.Channel {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) GetChannelListByToken(next string) *resource.PagedList {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) CreateChannel(groupId string, data *resource.Channel) (*resource.Channel, error) {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) UpdateChannel(groupId string, data *resource.Channel) error {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) DeleteChannel(channelId string) error {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) MuteChannel(channelId string, duration uint32) error {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) CreateUserChannel(userId string, groupId string) (*resource.Channel, error) {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) GetGroup(groupId string) *resource.Group {
	groupUin, _ := strconv.Atoi(groupId)
	group := LgrClient.GetCachedGroupInfo(uint32(groupUin))
	return &resource.Group{
		Id:          strconv.Itoa(int(group.GroupUin)),
		Name:        group.GroupName,
		Avatar:      group.Avatar,
		Maxcount:    group.MaxMember,
		MemberCount: group.MemberCount,
	}
}

func (b BotLagrange) GetGroupList() []*resource.Group {
	info, _ := LgrClient.GetAllGroupsInfo()
	groups := make([]*resource.Group, 0)
	for _, group := range info {
		groups = append(groups, &resource.Group{
			Id:          strconv.Itoa(int(group.GroupUin)),
			Name:        group.GroupName,
			Avatar:      group.Avatar,
			Maxcount:    group.MaxMember,
			MemberCount: group.MemberCount,
		})
	}
	return groups
}

func (b BotLagrange) GetGroupListByToken(next string) *resource.PagedList {
	info, _ := LgrClient.GetAllGroupsInfo()
	groups := make([]*resource.Group, 0)
	for _, group := range info {
		groups = append(groups, &resource.Group{
			Id:          strconv.Itoa(int(group.GroupUin)),
			Name:        group.GroupName,
			Avatar:      group.Avatar,
			Maxcount:    group.MaxMember,
			MemberCount: group.MemberCount,
		})
	}
	next = ""
	return &resource.PagedList{
		Data: groups,
		Next: "",
	}
}

func (b BotLagrange) ApproveGroupInvite(messageId string, approve bool, comment string) error {
	messageId = ""
	approve = false
	comment = ""
	return nil
}

func (b BotLagrange) GetGroupMember(groupId string, userId string) *resource.GroupMember {
	groupIdInt, _ := strconv.Atoi(groupId)
	userIdInt, _ := strconv.Atoi(userId)
	members, err := LgrClient.GetGroupMembersData(uint32(groupIdInt))
	if err != nil {
		return nil
	}
	if member, ok := members[uint32(userIdInt)]; ok {
		return &resource.GroupMember{
			User: &resource.User{
				Id:       userId,
				Name:     member.MemberName,
				Nickname: member.DisplayName(),
				Avatar:   member.Avatar,
				IsBot:    false,
			},
			Nickname: member.DisplayName(),
			Avatar:   member.Avatar,
			JoinedAt: time.Unix(int64(member.JoinTime), 0),
		}
	}
	// 否则返回空
	return nil
}

func (b BotLagrange) GetGroupMemberList(groupId string) []*resource.GroupMember {
	groupIdInt, _ := strconv.Atoi(groupId)
	members, err := LgrClient.GetGroupMembersData(uint32(groupIdInt))
	if err != nil {
		return nil
	}
	groupMembers := make([]*resource.GroupMember, 0)
	for _, member := range members {
		groupMembers = append(groupMembers, &resource.GroupMember{
			User: &resource.User{
				Id:       strconv.Itoa(int(member.Uin)),
				Name:     member.MemberName,
				Nickname: member.DisplayName(),
				Avatar:   member.Avatar,
				IsBot:    false,
			},
			Nickname: member.DisplayName(),
			Avatar:   member.Avatar,
			JoinedAt: time.Unix(int64(member.JoinTime), 0),
		})
	}
	return groupMembers
}

func (b BotLagrange) GetGroupMemberListByToken(groupId string, _ string) *resource.PagedList {
	groupIdInt, _ := strconv.Atoi(groupId)
	members, err := LgrClient.GetGroupMembersData(uint32(groupIdInt))
	if err != nil {
		return nil
	}
	groupMembers := make([]*resource.GroupMember, 0)
	for _, member := range members {
		groupMembers = append(groupMembers, &resource.GroupMember{
			User: &resource.User{
				Id:       strconv.Itoa(int(member.Uin)),
				Name:     member.MemberName,
				Nickname: member.DisplayName(),
				Avatar:   member.Avatar,
				IsBot:    false,
			},
			Nickname: member.DisplayName(),
			Avatar:   member.Avatar,
			JoinedAt: time.Unix(int64(member.JoinTime), 0),
		})
	}
	return &resource.PagedList{
		Data: groupMembers,
		Next: "",
	}
}

func (b BotLagrange) KickGroupMember(groupId string, userId string, permanent bool) error {
	groupIdInt, _ := strconv.Atoi(groupId)
	userIdInt, _ := strconv.Atoi(userId)
	err := LgrClient.GroupKickMember(uint32(groupIdInt), uint32(userIdInt), permanent)
	if err != nil {
		return err
	}
	return nil
}

func (b BotLagrange) MuteGroupMember(groupId string, userId string, duration uint32) error {
	groupIdInt, _ := strconv.Atoi(groupId)
	userIdInt, _ := strconv.Atoi(userId)
	err := LgrClient.GroupMuteMember(uint32(groupIdInt), uint32(userIdInt), duration)
	if err != nil {
		return err
	}
	return nil
}

func (b BotLagrange) ApproveGroupRequest(messageId string, approve bool, comment string) error {
	// 晚点实现
	panic("implement me")
}

// SetGroupMemberRole 设置群成员角色，在NTQQ中实际上只能是设置管理员
func (b BotLagrange) SetGroupMemberRole(groupId string, userId string, _ string) error {
	groupIdInt, _ := strconv.Atoi(groupId)
	userIdInt, _ := strconv.Atoi(userId)
	err := LgrClient.GroupSetAdmin(uint32(groupIdInt), uint32(userIdInt), true)
	if err != nil {
		return err
	}
	return nil
}

// UnsetGroupMemberRole 取消群成员角色，在NTQQ中实际上只能是取消管理员
func (b BotLagrange) UnsetGroupMemberRole(groupId string, userId string, _ string) error {
	groupIdInt, _ := strconv.Atoi(groupId)
	userIdInt, _ := strconv.Atoi(userId)
	err := LgrClient.GroupSetAdmin(uint32(groupIdInt), uint32(userIdInt), false)
	if err != nil {
		return err
	}
	return nil
}

func (b BotLagrange) GetGroupRoleList(_ string) []*resource.GroupRole {
	return []*resource.GroupRole{
		{
			Id:   "1",
			Name: "管理员",
		},
		{
			Id:   "2",
			Name: "群员",
		},
	}
}

func (b BotLagrange) GetGroupRoleListByToken(groupId string, next string) *resource.PagedList {
	return &resource.PagedList{
		Data: []*resource.GroupRole{
			{
				Id:   "1",
				Name: "管理员",
			},
			{
				Id:   "2",
				Name: "群员",
			},
		},
		Next: "",
	}
}

func (b BotLagrange) CreateGroupRole(_ string, _ *resource.GroupRole) (*resource.GroupRole, error) {
	return nil, nil
}

func (b BotLagrange) UpdateGroupRole(_ string, _ string, _ *resource.GroupRole) error {
	return nil
}

func (b BotLagrange) DeleteGroupRole(_ string, _ string) error {
	return nil
}

func (b BotLagrange) GetLoginInfo() *resource.Login {
	return nil
}

func (b BotLagrange) SendContent(groupId string, channelId string, content string) (*resource.Message, error) {
	groupIdInt, _ := strconv.Atoi(groupId)
	channelIdInt, _ := strconv.Atoi(channelId)
	var msg []message.IMessageElement
	if groupIdInt == 0 {
		msg = append(msg, message.NewText(content))
		res, err := LgrClient.SendPrivateMessage(uint32(channelIdInt), msg)
		if err != nil {
			return nil, err
		}
		return &resource.Message{
			Id:              strconv.Itoa(int(res.Id)),
			Content:         res.ToString(),
			Channel:         nil,
			Group:           nil,
			Member:          nil,
			User:            nil,
			CreatedAt:       time.Unix(int64(res.Time), 0),
			UpdatedAt:       time.Time{},
			MessageElements: ConvertIceElement(res.Elements),
		}, nil
	} else {
		msg = append(msg, message.NewText(content))
		res, err := LgrClient.SendGroupMessage(uint32(groupIdInt), msg)
		if err != nil {
			return nil, err
		}
		return &resource.Message{
			Id:              strconv.Itoa(int(res.Id)),
			Content:         res.ToString(),
			Channel:         nil,
			Group:           nil,
			Member:          nil,
			User:            nil,
			CreatedAt:       time.Unix(int64(res.Time), 0),
			UpdatedAt:       time.Time{},
			MessageElements: ConvertIceElement(res.Elements),
		}, nil
	}
}

func (b BotLagrange) GetMessage(channelId string, messageId string) (*resource.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) RecallMessage(channelId string, messageId string) error {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) UpdateMessage(channelId string, messageId string, content string) error {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) GetMessageList(channelId string, limit uint32, order bool) []*resource.Message {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) GetMessageListByRange(channelId string, messageId string, start uint32, count uint32) []*resource.Message {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) CreateReaction(channelId string, messageId string, emoji string) error {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) DeleteReaction(channelId string, messageId string, emoji string, userId string) error {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) ClearReaction(channelId string, messageId string, emoji string) error {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) GetReactionList(channelId string, messageId string, emoji string) []resource.User {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) GetReactionListByToken(channelId string, messageId string, emoji string, next string) *resource.PagedList {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) GetUser(userId string) *resource.User {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) GetFriendList() []*resource.User {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) ApproveFriendRequest(messageId string, approve string, comment string) error {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) Send(groupId string, channelId string, elements []elements.IceinuMessageElement) (*resource.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) SendSatori(groupId string, channelId string, satori string) (*resource.Message, error) {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) SendPoke(groupId string, channelId string, userId string) error {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) GetSelfUserId() string {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) GetSelfUserName() string {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) GetSelfAvatarUrl() string {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) GetSeldUserNickname() string {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) GetGroupAvatarUrl(groupId string) string {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) RefreshUserListCache() error {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) RefreshGroupListCache() error {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) RefreshGroupMemberCache(groupId string, userId string) error {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) RefreshGroupAllMembersCache(groupId string) error {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) RefreshChannelListCache(groupId string) error {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) RenameGroup(groupId string, name string) error {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) RenameGroupMember(groupId string, userId string, nickname string) error {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) RemarkGroup(groupId string, remark string) error {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) SetGroupGlobalMute(groupId string, mute bool) error {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) LeaveGroup(groupId string) error {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) SetGroupMemberTitle(groupId string, userId string, title string) error {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) UploadChannelFile(channelId string, filePath string) error {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) UploadGroupFile(groupId string, filePath string, targetFilePath string) error {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) GetGroupFileSystemInfo(groupId string) *resource.GroupFileSystemInfo {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) GetGroupFilesByFolder(groupId string, folderId string) interface{} {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) GetGroupRootFiles(groupId string) interface{} {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) MoveGroupFile(groupId string, fileId string, parentFolder string, targetFolderId string) error {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) DeleteGroupFile(groupId string, fileId string) error {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) CreateGroupFolder(groupId string, folderName string, parentFolder string) error {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) RenameGroupFolder(groupId string, folderId string, newFolderName string) error {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) DeleteGroupFolder(groupId string, folderId string) error {
	//TODO implement me
	panic("implement me")
}

func (b BotLagrange) GetOriginalClient() interface{} {
	//TODO implement me
	panic("implement me")
}

func GetBot() *BotLagrange {
	return &BotLagrange{}
}
