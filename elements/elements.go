package elements

import (
	"encoding/base64"
	"fmt"
	"io"
	"strings"
	"time"
)

var DefaultThumb, _ = base64.StdEncoding.DecodeString("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")

// IceinuMessageElement Iceinu的通用消息元素接口，参考了Satori的标准消息元素设计
// https://satori.js.org/zh-CN/protocol/elements.html
// 基于对不同平台的支持，Iceinu对标准消息元素进行扩展和调整来方便一些平台特殊设计的实现
// 带*号的消息元素是Iceinu自定义的消息元素，方便实现一些针对平台设计的特殊功能
type IceinuMessageElement interface {
	GetType() string  // 获取消息元素类型
	ToSatori() string // 转换为Satori消息元素字符串
}

// TextElement 文本消息元素
type TextElement struct {
	Text string
}

func (t *TextElement) GetType() string {
	return "text"
}

func (t *TextElement) ToSatori() string {
	return t.Text
}

// AtElement At提及消息元素
type AtElement struct {
	Id   string // 目标用户ID
	Name string // 目标用户名称
	Role string // 目标用户角色
	Type string // At请求类型，0为全体成员，1为指定成员
}

func (a *AtElement) GetType() string {
	return "at"
}

func (a *AtElement) ToSatori() string {
	var attributes []string
	if a.Id != "" {
		attributes = append(attributes, fmt.Sprintf("id=\"%s\"", a.Id))
	}
	if a.Name != "" {
		attributes = append(attributes, fmt.Sprintf("name=\"%s\"", a.Name))
	}
	if a.Role != "" {
		attributes = append(attributes, fmt.Sprintf("role=\"%s\"", a.Role))
	}
	if a.Type != "" {
		attributes = append(attributes, fmt.Sprintf("type=\"%s\"", a.Type))
	}
	return fmt.Sprintf("<at %s/>", strings.Join(attributes, " "))
}

// SharpElement Sharp提及频道消息元素
type SharpElement struct {
	Id   string // 目标频道ID
	Name string // 目标频道名称
}

func (s *SharpElement) GetType() string {
	return "sharp"
}

func (s *SharpElement) ToSatori() string {
	var attributes []string
	if s.Id != "" {
		attributes = append(attributes, fmt.Sprintf("id=\"%s\"", s.Id))
	}
	if s.Name != "" {
		attributes = append(attributes, fmt.Sprintf("name=\"%s\"", s.Name))
	}
	return fmt.Sprintf("<sharp %s/>", strings.Join(attributes, " "))
}

// LinkElement A超链接消息元素
type LinkElement struct {
	Href string // 链接地址
}

func (a *LinkElement) GetType() string {
	return "link"
}

func (a *LinkElement) ToSatori() string {
	if a.Href != "" {
		return fmt.Sprintf("<a href=\"%s\"/>", a.Href)
	}
	return "<a/>"
}

// ImageElement 图片消息元素
type ImageElement struct {
	// 用于接收图片

	ImageId string // 图片ID
	Src     string // 图片源地址
	Title   string // 图片标题
	Width   uint32 // 图片宽度
	Height  uint32 // 图片高度

	EffectId int  // 图片特效ID
	IsFlash  bool // 是否是闪图

	// 用于发送图片
	Summary string        // 图片描述
	Path    string        // 图片路径或URL
	Stream  io.ReadSeeker // 图片流
}

func (i *ImageElement) GetType() string {
	return "image"
}

func (i *ImageElement) ToSatori() string {
	var attributes []string
	if i.ImageId != "" {
		attributes = append(attributes, fmt.Sprintf("id=\"%s\"", i.ImageId))
	}
	if i.Src != "" {
		attributes = append(attributes, fmt.Sprintf("src=\"%s\"", i.Src))
	}
	if i.Title != "" {
		attributes = append(attributes, fmt.Sprintf("title=\"%s\"", i.Title))
	}
	if i.Width != 0 {
		attributes = append(attributes, fmt.Sprintf("width=\"%d\"", i.Width))
	}
	if i.Height != 0 {
		attributes = append(attributes, fmt.Sprintf("height=\"%d\"", i.Height))
	}
	if i.EffectId != 0 {
		attributes = append(attributes, fmt.Sprintf("effect=\"%d\"", i.EffectId))
	}
	if i.IsFlash {
		attributes = append(attributes, fmt.Sprintf("flash=\"%t\"", i.IsFlash))
	}
	if i.Summary != "" {
		attributes = append(attributes, fmt.Sprintf("summary=\"%s\"", i.Summary))
	}
	if i.Path != "" {
		attributes = append(attributes, fmt.Sprintf("path=\"%s\"", i.Path))
	}
	return fmt.Sprintf("<img %s/>", strings.Join(attributes, " "))
}

// AudioElement 音频消息元素
type AudioElement struct {
	Src      string
	Title    string
	Duration uint32
	Poster   string
	Stream   io.ReadSeeker
	Summary  string
	Path     string
}

func (a *AudioElement) GetType() string {
	return "audio"
}

func (a *AudioElement) ToSatori() string {
	var attributes []string
	if a.Src != "" {
		attributes = append(attributes, fmt.Sprintf("src=\"%s\"", a.Src))
	}
	if a.Title != "" {
		attributes = append(attributes, fmt.Sprintf("title=\"%s\"", a.Title))
	}
	if a.Duration != 0 {
		attributes = append(attributes, fmt.Sprintf("duration=\"%d\"", a.Duration))
	}
	if a.Poster != "" {
		attributes = append(attributes, fmt.Sprintf("poster=\"%s\"", a.Poster))
	}
	if a.Summary != "" {
		attributes = append(attributes, fmt.Sprintf("summary=\"%s\"", a.Summary))
	}
	if a.Path != "" {
		attributes = append(attributes, fmt.Sprintf("path=\"%s\"", a.Path))
	}
	return fmt.Sprintf("<audio %s/>", strings.Join(attributes, " "))
}

// VideoElement 视频消息元素
type VideoElement struct {
	Src      string
	Title    string
	Width    uint32
	Height   uint32
	Duration uint32
	Poster   string
	Path     string
}

func (v *VideoElement) GetType() string {
	return "video"
}

func (v *VideoElement) ToSatori() string {
	var attributes []string
	if v.Src != "" {
		attributes = append(attributes, fmt.Sprintf("src=\"%s\"", v.Src))
	}
	if v.Title != "" {
		attributes = append(attributes, fmt.Sprintf("title=\"%s\"", v.Title))
	}
	if v.Width != 0 {
		attributes = append(attributes, fmt.Sprintf("width=\"%d\"", v.Width))
	}
	if v.Height != 0 {
		attributes = append(attributes, fmt.Sprintf("height=\"%d\"", v.Height))
	}
	if v.Duration != 0 {
		attributes = append(attributes, fmt.Sprintf("duration=\"%d\"", v.Duration))
	}
	if v.Poster != "" {
		attributes = append(attributes, fmt.Sprintf("poster=\"%s\"", v.Poster))
	}
	if v.Path != "" {
		attributes = append(attributes, fmt.Sprintf("path=\"%s\"", v.Path))
	}
	return fmt.Sprintf("<video %s/>", strings.Join(attributes, " "))
}

// ShortVideoElement 短视频消息元素
type ShortVideoElement struct {
	Title    string
	Src      string
	Duration uint32
	Summary  string
	Stream   io.ReadSeeker
	Path     string
}

func (s *ShortVideoElement) GetType() string {
	return "shortvideo"
}

func (s *ShortVideoElement) ToSatori() string {
	var attributes []string
	if s.Title != "" {
		attributes = append(attributes, fmt.Sprintf("title=\"%s\"", s.Title))
	}
	if s.Src != "" {
		attributes = append(attributes, fmt.Sprintf("src=\"%s\"", s.Src))
	}
	if s.Duration != 0 {
		attributes = append(attributes, fmt.Sprintf("duration=\"%d\"", s.Duration))
	}
	if s.Summary != "" {
		attributes = append(attributes, fmt.Sprintf("summary=\"%s\"", s.Summary))
	}
	if s.Path != "" {
		attributes = append(attributes, fmt.Sprintf("path=\"%s\"", s.Path))
	}
	return fmt.Sprintf("<shortvideo %s/>", strings.Join(attributes, " "))
}

// FileElement 文件消息元素
type FileElement struct {
	Src    string
	Title  string
	Poster string
	Size   uint64
	Stream io.ReadSeeker
	Path   string
}

func (f *FileElement) GetType() string {
	return "file"
}

func (f *FileElement) ToSatori() string {
	var attributes []string
	if f.Src != "" {
		attributes = append(attributes, fmt.Sprintf("src=\"%s\"", f.Src))
	}
	if f.Title != "" {
		attributes = append(attributes, fmt.Sprintf("title=\"%s\"", f.Title))
	}
	if f.Poster != "" {
		attributes = append(attributes, fmt.Sprintf("poster=\"%s\"", f.Poster))
	}
	if f.Size != 0 {
		attributes = append(attributes, fmt.Sprintf("size=\"%d\"", f.Size))
	}
	if f.Path != "" {
		attributes = append(attributes, fmt.Sprintf("path=\"%s\"", f.Path))
	}
	return fmt.Sprintf("<file %s/>", strings.Join(attributes, " "))
}

// StrongElement 粗体消息元素
type StrongElement struct {
	Text string
}

func (b *StrongElement) GetType() string {
	return "strong"
}

func (b *StrongElement) ToSatori() string {
	return fmt.Sprintf("<b>%s</b>", b.Text)
}

// EmElement 斜体消息元素
type EmElement struct {
	Text string
}

func (e *EmElement) GetType() string {
	return "em"
}

func (e *EmElement) ToSatori() string {
	return fmt.Sprintf("<em>%s</em>", e.Text)
}

// InsElement 下划线消息元素
type InsElement struct {
	Text string
}

func (i *InsElement) GetType() string {
	return "ins"
}

func (i *InsElement) ToSatori() string {
	return fmt.Sprintf("<ins>%s</ins>", i.Text)
}

// DelElement 删除线消息元素
type DelElement struct {
	Text string
}

func (d *DelElement) GetType() string {
	return "del"
}

func (d *DelElement) ToSatori() string {
	return fmt.Sprintf("<del>%s</del>", d.Text)
}

// SplElement 剧透消息元素
type SplElement struct {
	Text string
}

func (s *SplElement) GetType() string {
	return "spl"
}

func (s *SplElement) ToSatori() string {
	return fmt.Sprintf("<spl>%s</spl>", s.Text)
}

// CodeElement 代码消息元素
type CodeElement struct {
	Text string
}

func (c *CodeElement) GetType() string {
	return "code"
}

func (c *CodeElement) ToSatori() string {
	return fmt.Sprintf("<code>%s</code>", c.Text)
}

// SupElement 上标消息元素
type SupElement struct {
	Text string
}

func (s *SupElement) GetType() string {
	return "sup"
}

func (s *SupElement) ToSatori() string {
	return fmt.Sprintf("<sup>%s</sup>", s.Text)
}

// SubElement 下标消息元素
type SubElement struct {
	Text string
}

func (s *SubElement) GetType() string {
	return "sub"
}

func (s *SubElement) ToSatori() string {
	return fmt.Sprintf("<sub>%s</sub>", s.Text)
}

// BrElement 换行消息元素
type BrElement struct{}

func (b *BrElement) GetType() string {
	return "br"
}

func (b *BrElement) ToSatori() string {
	return "<br/>"
}

// HrElement *水平线消息元素
type HrElement struct{}

func (h *HrElement) GetType() string {
	return "hr"
}

func (h *HrElement) ToSatori() string {
	return "<hr/>"
}

// AuthorElement 作者消息元素
type AuthorElement struct {
	Id     string
	Name   string
	Avatar string
}

func (a *AuthorElement) GetType() string {
	return "author"
}

func (a *AuthorElement) ToSatori() string {
	var attributes []string
	if a.Id != "" {
		attributes = append(attributes, fmt.Sprintf("id=\"%s\"", a.Id))
	}
	if a.Name != "" {
		attributes = append(attributes, fmt.Sprintf("name=\"%s\"", a.Name))
	}
	if a.Avatar != "" {
		attributes = append(attributes, fmt.Sprintf("avatar=\"%s\"", a.Avatar))
	}
	return fmt.Sprintf("<author %s/>", strings.Join(attributes, " "))
}

// ButtonElement 按钮消息元素
type ButtonElement struct {
	Id   string
	Type string
	Href string
	Text string
}

func (b *ButtonElement) GetType() string {
	return "button"
}

func (b *ButtonElement) ToSatori() string {
	var attributes []string
	if b.Id != "" {
		attributes = append(attributes, fmt.Sprintf("id=\"%s\"", b.Id))
	}
	if b.Type != "" {
		attributes = append(attributes, fmt.Sprintf("type=\"%s\"", b.Type))
	}
	if b.Href != "" {
		attributes = append(attributes, fmt.Sprintf("href=\"%s\"", b.Href))
	}
	if b.Text != "" {
		attributes = append(attributes, b.Text)
	}
	return fmt.Sprintf("<button %s/>", strings.Join(attributes, " "))
}

// FaceElement *表情消息元素，可以传递用于一些平台内置的表情
type FaceElement struct {
	Id          uint16
	IsLargeFace bool
}

func (f *FaceElement) GetType() string {
	return "face"
}

func (f *FaceElement) ToSatori() string {
	var attributes []string
	if f.Id != 0 {
		attributes = append(attributes, fmt.Sprintf("id=\"%d\"", f.Id))
	}
	if f.IsLargeFace {
		attributes = append(attributes, "size=\"large\"")
	}
	return fmt.Sprintf("<face %s/>", strings.Join(attributes, " "))
}

// QuoteElement *引用消息元素，可以引用消息中的某一部分内容
//
// # Satori中虽然也标注了引用消息元素，但是没有给出具体设计，这里主要用于QQ的引用消息设计
//
// 一般只用于解析接受到的回复消息事件，主动进行回复消息需要通过对应适配器的bot API进行
type QuoteElement struct {
	UserId    string
	UserName  string
	GroupId   string
	Timestamp time.Time
	Elements  *[]IceinuMessageElement
}

func (q *QuoteElement) GetType() string {
	return "quote"
}

func (q *QuoteElement) ToSatori() string {
	var attributes []string
	if q.UserId != "" {
		attributes = append(attributes, fmt.Sprintf("id=\"%s\"", q.UserId))
	}
	if q.UserName != "" {
		attributes = append(attributes, fmt.Sprintf("name=\"%s\"", q.UserName))
	}
	if q.GroupId != "" {
		attributes = append(attributes, fmt.Sprintf("group=\"%s\"", q.GroupId))
	}
	if !q.Timestamp.IsZero() {
		attributes = append(attributes, fmt.Sprintf("time=\"%s\"", q.Timestamp.Format(time.RFC3339)))
	}

	// 构建开始标签，包含属性
	openingTag := fmt.Sprintf("<quote %s>", strings.Join(attributes, " "))

	// 将子元素转换为 Satori 字符串
	var elements []string
	if q.Elements != nil {
		for _, element := range *q.Elements {
			elements = append(elements, element.ToSatori())
		}
	}

	// 定义结束标签
	closingTag := "</quote>"

	// 组合所有部分
	return openingTag + strings.Join(elements, "") + closingTag
}

// UnsupportedElement 未支持的消息元素
type UnsupportedElement struct {
	Type string
}

func (u *UnsupportedElement) GetType() string {
	return "unsupported"
}

func (u *UnsupportedElement) ToSatori() string {
	return fmt.Sprintf("<!-- Unsupported element type: %s -->", u.Type)
}

// MessageElement 消息封装元素
type MessageElement struct {
	Id        string
	Forward   bool
	NtForward bool
	Elements  *[]IceinuMessageElement
}

func (m *MessageElement) GetType() string {
	return "message"
}

func (m *MessageElement) ToSatori() string {
	var attributes []string
	if m.Id != "" {
		attributes = append(attributes, fmt.Sprintf("id=\"%s\"", m.Id))
	}
	if m.Forward {
		attributes = append(attributes, "forward=\"true\"")
	}

	// 构建开始标签，包含属性
	openingTag := fmt.Sprintf("<message %s>", strings.Join(attributes, " "))

	// 将子元素转换为 Satori 字符串
	var elements []string
	if m.Elements != nil {
		for _, element := range *m.Elements {
			elements = append(elements, element.ToSatori())
		}
	}

	// 定义结束标签
	closingTag := "</message>"

	// 组合所有部分
	return openingTag + strings.Join(elements, "") + closingTag
}

// NodeElement *消息节点封装元素，用于QQ平台的节点合并转发消息
type NodeElement struct {
	GroupId    int64
	SenderId   int64
	SenderName string
	Time       int32
	Message    *[]IceinuMessageElement
}

func (n *NodeElement) GetType() string {
	return "node"
}

func (n *NodeElement) ToSatori() string {
	var attributes []string
	if n.GroupId != 0 {
		attributes = append(attributes, fmt.Sprintf("group=\"%d\"", n.GroupId))
	}
	if n.SenderId != 0 {
		attributes = append(attributes, fmt.Sprintf("sender_id=\"%d\"", n.SenderId))
	}
	if n.SenderName != "" {
		attributes = append(attributes, fmt.Sprintf("sender_name=\"%s\"", n.SenderName))
	}
	if n.Time != 0 {
		attributes = append(attributes, fmt.Sprintf("time=\"%d\"", n.Time))
	}

	// 构建开始标签，包含属性
	openingTag := fmt.Sprintf("<node %s>", strings.Join(attributes, " "))

	// 将消息内容转换为 Satori 字符串
	var messages []string
	if n.Message != nil {
		for _, element := range *n.Message {
			messages = append(messages, element.ToSatori())
		}
	}

	// 定义结束标签
	closingTag := "</node>"

	// 组合所有部分
	return openingTag + strings.Join(messages, "") + closingTag
}
