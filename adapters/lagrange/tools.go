package lagrange

import (
	"fmt"
	"github.com/Iceinu-Project/Iceinu/log"
	"github.com/Iceinu-Project/Iceinu/models/satori"
	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/client/entity"
	"github.com/LagrangeDev/LagrangeGo/message"
	"strconv"
	"strings"
)

// ToSatoriElements 将LagrangeGo的消息元素切片转换为Satori的消息元素切片
func ToSatoriElements(elements []message.IMessageElement) *[]satori.ElementSatori {
	// 创建存储Satori消息切片的变量
	var result []satori.ElementSatori
	// 遍历传入的元素
	for _, ele := range elements {
		// 通过消息元素的Type方法来确定消息类型
		switch ele.Type() {
		// 将元素依次对应转换并传入
		case message.Text:
			ele := ele.(*message.TextElement)
			// 检测文本中是否包含换行符
			if strings.Contains(ele.Content, "\n") {
				// 如果包含换行符，进行拆分和处理
				textParts := strings.Split(ele.Content, "\n")
				for i, part := range textParts {
					// 将每段文本添加到 result
					result = append(result, &satori.TextElement{Text: part})
					// 如果不是最后一段文本，则插入 BrElement
					if i < len(textParts)-1 {
						result = append(result, &satori.BrElement{})
					}
				}
			} else {
				// 如果不包含换行符，直接添加文本元素
				result = append(result, &satori.TextElement{Text: ele.Content})
			}
		case message.At:
			ele := ele.(*message.AtElement)
			result = append(result, &satori.AtElement{
				Id:   strconv.Itoa(int(ele.TargetUin)),
				Name: ele.Display,
				Role: "",
				Type: strconv.Itoa(int(ele.Type())),
			})
		case message.Face:
			ele := ele.(*message.FaceElement)
			result = append(result, &satori.FaceElement{
				Id: ele.FaceID,
			})
		case message.Voice:
			ele := ele.(*message.VoiceElement)
			result = append(result, &satori.AudioElement{
				Src:      ele.Url,
				Title:    ele.Name,
				Duration: ele.Size,
				Poster:   "",
				Cache:    false,
				Timeout:  0,
			})
		case message.Image:
			ele := ele.(*message.ImageElement)
			result = append(result, &satori.ImgElement{
				Src:     ele.Url,
				Width:   ele.Width,
				Height:  ele.Height,
				Title:   ele.ImageId,
				Cache:   false,
				Timeout: 0,
			})
		case message.File:
			ele := ele.(*message.FileElement)
			result = append(result, &satori.FileElement{
				Src:     ele.FileUrl,
				Title:   ele.FileName,
				Poster:  "",
				Cache:   false,
				Timeout: 0,
			})
		case message.Reply:
			ele := ele.(*message.ReplyElement)
			result = append(result, &satori.QuoteElement{
				Id:        strconv.Itoa(int(ele.SenderUin)),
				Name:      ele.SenderUid,
				ChannelId: strconv.Itoa(int(ele.GroupUin)),
				GroupId:   strconv.Itoa(int(ele.GroupUin)),
				Timestamp: int64(ele.Time),
				Elements:  ToSatoriElements(ele.Elements),
			})
		case message.Forward:
			ele := ele.(*message.ForwardMessage)
			result = append(result, &satori.MessageElement{
				Forward:  true,
				Elements: UnzipNodes(ele.Nodes),
			})

		default:
			result = append(result, &satori.UnsupportedElement{Type: strconv.Itoa(int(ele.Type()))})

		}
	}
	return &result
}

func UnzipNodes(n []*message.ForwardNode) *[]satori.ElementSatori {
	var result []satori.ElementSatori
	for _, node := range n {
		result = append(result, &satori.NodeElement{
			GroupId:    node.GroupId,
			SenderId:   node.SenderId,
			SenderName: node.SenderName,
			Time:       node.Time,
			Message:    ToSatoriElements(node.Message),
		})
	}
	return &result
}

// GetGroupMembersDataInCache 从缓存中获取群成员映射，如果缓存中没有则拉取并存入缓存
func GetGroupMembersDataInCache(client *client.QQClient, groupId uint32) map[uint32]*entity.GroupMember {
	// 尝试在内置缓存中查找群成员映射
	var groupMemberMap map[uint32]*entity.GroupMember
	err := Cache.Get(fmt.Sprintf("group_member_data_%d", groupId), &groupMemberMap)
	if err != nil {
		// 如果缓存中没有群成员映射，则创建一个
		groupMembersData, err := client.GetGroupMembersData(groupId)
		if err != nil {
			log.Warn("无法获取群成员列表数据:", err)
		}
		err = Cache.Set(fmt.Sprintf("group_member_data_%d", groupId), groupMembersData)
		if err != nil {
			log.Warn("无法缓存群成员映射数据:", err)
		}
		groupMemberMap = groupMembersData
		log.Debugf("%v群成员缓存数据更新完成，共%d个成员", groupId, len(groupMemberMap))
	}
	return groupMemberMap
}

// GetFriendsDataInCache 从缓存中获取好友映射，如果缓存中没有则拉取并存入缓存
func GetFriendsDataInCache(client *client.QQClient) map[uint32]*entity.Friend {
	// 尝试在内置缓存中查找好友映射
	var friendMap map[uint32]*entity.Friend
	err := Cache.Get("friend_data", &friendMap)
	if err != nil {
		// 如果缓存中没有好友映射，则创建一个
		friendData, err := client.GetFriendsData()
		if err != nil {
			log.Warn("无法获取好友列表数据:", err)
		}
		err = Cache.Set("friend_data", friendData)
		if err != nil {
			log.Warn("无法缓存好友映射数据:", err)
		}
		friendMap = friendData
		log.Debugf("好友缓存数据更新完成，共%d个好友", len(friendMap))
	}
	return friendMap
}

// GetSelfInfoInCache 从缓存中获取自身信息，如果缓存中没有则拉取并存入缓存
func GetSelfInfoInCache(client *client.QQClient) *entity.Friend {
	// 尝试在内置缓存中查找自身信息
	friendData := GetFriendsDataInCache(client)
	selfInfo, ok := friendData[client.Uin]
	if !ok {
		log.Warn("无法获取自身信息")
		return nil
	}
	return selfInfo
}
