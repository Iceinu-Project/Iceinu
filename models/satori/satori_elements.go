package satori

// 参考：https://satori.js.org/zh-CN/protocol/elements.html
// 在Satori的标准元素基础上扩展了一部分元素，方便直接使用

// ElementSatori Satori标准事件元素接口
type ElementSatori interface {
	GetType() string
}

// TextElement 文本消息元素
type TextElement struct {
	Text string
}

func (t *TextElement) GetType() string {
	return "text"
}

// AtElement 提及用户消息元素
type AtElement struct {
	Id   string
	Name string
	Role string
	Type string
}

func (a *AtElement) GetType() string {
	return "at"
}

// SharpElement 提及频道消息元素
type SharpElement struct {
	Id   string
	Name string
}

func (s *SharpElement) GetType() string {
	return "sharp"
}

// AElement 超链接消息元素
type AElement struct {
	Href string
}

func (a *AElement) GetType() string {
	return "a"
}

// ImgElement 图片消息元素
type ImgElement struct {
	Src     string
	Title   string
	Cache   bool
	Timeout uint32
	Width   uint32
	Height  uint32
}

func (i *ImgElement) GetType() string {
	return "img"
}

// AudioElement 音频消息元素
type AudioElement struct {
	Src      string
	Title    string
	Cache    bool
	Timeout  uint32
	Duration uint32
	Poster   string
}

func (a *AudioElement) GetType() string {
	return "audio"
}

// VideoElement 视频消息元素
type VideoElement struct {
	Src      string
	Title    string
	Cache    bool
	Timeout  uint32
	Width    uint32
	Height   uint32
	Duration uint32
	Poster   string
}

func (v *VideoElement) GetType() string {
	return "video"
}

// FileElement 文件消息元素
type FileElement struct {
	Src     string
	Title   string
	Cache   bool
	Timeout uint32
	Poster  string
}

func (f *FileElement) GetType() string {
	return "file"
}

// StrongElement 加粗消息元素
type StrongElement struct {
	Elements *[]ElementSatori
}

func (s *StrongElement) GetType() string {
	return "strong"
}

// EmElement 斜体消息元素
type EmElement struct {
	Elements *[]ElementSatori
}

func (e *EmElement) GetType() string {
	return "em"
}

// InsElement 下划线消息元素
type InsElement struct {
	Elements *[]ElementSatori
}

func (i *InsElement) GetType() string {
	return "ins"
}

// DelElement 删除线消息元素
type DelElement struct {
	Elements *[]ElementSatori
}

func (d *DelElement) GetType() string {
	return "del"
}

// SpoilerElement 剧透消息元素
type SpoilerElement struct {
	Elements *[]ElementSatori
}

func (s *SpoilerElement) GetType() string {
	return "spoiler"
}

// CodeElement 代码消息元素
type CodeElement struct {
	Elements *[]ElementSatori
}

func (c *CodeElement) GetType() string {
	return "code"
}

// SupElement 上标消息元素
type SupElement struct {
	Elements *[]ElementSatori
}

func (s *SupElement) GetType() string {
	return "sup"
}

// SubElement 下标消息元素
type SubElement struct {
	Elements *[]ElementSatori
}

func (s *SubElement) GetType() string {
	return "sub"
}

// BrElement 换行消息元素
type BrElement struct {
}

func (b *BrElement) GetType() string {
	return "br"
}

// HrElement 分割线消息元素
type HrElement struct {
}

func (h *HrElement) GetType() string {
	return "hr"
}

// PElement 段落消息元素
type PElement struct {
	Elements *[]ElementSatori
}

func (p *PElement) GetType() string {
	return "p"
}

// MessageElement 消息元素
type MessageElement struct {
	Id       string
	Forward  bool
	Elements *[]ElementSatori
}

func (m *MessageElement) GetType() string {
	return "message"
}

// QuoteElement 引用消息元素
type QuoteElement struct {
	Id        string
	Name      string
	GroupId   string
	ChannelId string
	Timestamp int64
	Elements  *[]ElementSatori
}

func (q *QuoteElement) GetType() string {
	return "quote"
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

// ButtonElement 按钮消息元素
type ButtonElement struct {
	Id    string
	Type  string
	Href  string
	Text  string
	Theme string
}

func (b *ButtonElement) GetType() string {
	return "button"
}

// FaceElement 表情消息元素
type FaceElement struct {
	Id          uint16
	IsLargeFace bool
}

func (f *FaceElement) GetType() string {
	return "face"
}

// NodeElement 节点消息元素
type NodeElement struct {
	GroupId    int64
	SenderId   int64
	SenderName string
	Time       int32
	Message    *[]ElementSatori
}

func (n *NodeElement) GetType() string {
	return "node"
}

// UnsupportedElement 未支持的消息元素
type UnsupportedElement struct {
	Type string
}

func (u *UnsupportedElement) GetType() string {
	return "unsupported"
}
