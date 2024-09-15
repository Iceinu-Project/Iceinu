package elements

import (
	"github.com/LagrangeDev/LagrangeGo/client/packets/pb/message"
	"github.com/LagrangeDev/LagrangeGo/client/packets/pb/service/oidb"
	"io"
	"time"
)

// 消息元素的设计基本就是复刻的LagrangeGo
// 这样转换过程非常方便XD（指LagrangeGo适配器）
// 参考：https://github.com/LagrangeDev/LagrangeGo/blob/master/message/elements.go

type ElementType int

type MessageElement interface {
	Type() ElementType
}

const (
	Text     ElementType = iota // 文本
	Image                       // 图片
	Face                        // 表情
	At                          // 艾特
	Reply                       // 回复
	Service                     // 服务
	Forward                     // 转发
	File                        // 文件
	Voice                       // 语音
	Video                       // 视频
	LightApp                    // 轻应用
)

type AtType int

type ForwardNode struct {
	GroupId    int64
	SenderId   int64
	SenderName string
	Time       int32
	Message    []MessageElement
}

type TextElement struct {
	Content string
}

type AtElement struct {
	TargetId   uint32
	TargetName string
	Display    string
	SubType    AtType
}

type FaceElement struct {
	FaceId      uint16
	IsLargeFace bool
}

type ReplyElement struct {
	ReplySeq   uint32
	SenderId   uint32
	SenderName string
	GroupId    uint32
	Timestamp  time.Time
	Elements   []MessageElement
}

type ImageElement struct {
	ImageId   string
	FileId    int64
	ImageType int32
	Size      uint32
	Width     uint32
	Height    uint32
	Url       string

	// EffectID show pic effect id.
	EffectID int32
	Flash    bool

	// send
	Summary     string
	Ext         string
	Md5         []byte
	Sha1        []byte
	MsgInfo     *oidb.MsgInfo
	Stream      io.ReadSeeker
	CompatFace  *message.CustomFace     // GroupImage
	CompatImage *message.NotOnlineImage // FriendImage
}

type FileElement struct {
	FileSize uint64
	FileName string
	FileMd5  []byte
	FileUrl  string
	FileId   string // group
	FileUUID string // private
	FileHash string

	// send
	FileStream io.ReadSeeker
	FileSha1   []byte
}

type ShortVideoElement struct {
	Name     string
	Uuid     []byte
	Size     uint32
	Url      string
	Duration uint32

	// send
	Thumb   []byte
	Summary string
	Md5     []byte
	Sha1    []byte
	Stream  io.ReadSeeker
	MsgInfo *oidb.MsgInfo
	Compat  *message.VideoFile
}

type VideoThumb struct {
	Stream io.ReadSeeker
	Size   uint32
	Md5    []byte
	Sha1   []byte
	Width  uint32
	Height uint32
}

type LightAppElement struct {
	AppName string
	Content string
}

type ForwardMessage struct {
	ResID string
	Nodes []*ForwardNode
}
