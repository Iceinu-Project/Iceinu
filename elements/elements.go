package elements

import (
	"fmt"
	"strings"
	"time"
)

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
	Id   string
	Name string
	Role string
	Type string
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
	Id   string
	Name string
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
	Href string
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
	Src     string
	Title   string
	Width   uint32
	Height  uint32
	Cache   bool
	Timeout string
}

func (i *ImageElement) GetType() string {
	return "image"
}

func (i *ImageElement) ToSatori() string {
	var attributes []string
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
	if i.Cache {
		attributes = append(attributes, fmt.Sprintf("cache=\"%t\"", i.Cache))
	}
	if i.Timeout != "" {
		attributes = append(attributes, fmt.Sprintf("timeout=\"%s\"", i.Timeout))
	}
	return fmt.Sprintf("<img %s/>", strings.Join(attributes, " "))
}

// AudioElement 音频消息元素
type AudioElement struct {
	Src      string
	Title    string
	Duration uint32
	Poster   string
	Cache    bool
	Timeout  string
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
	if a.Cache {
		attributes = append(attributes, fmt.Sprintf("cache=\"%t\"", a.Cache))
	}
	if a.Timeout != "" {
		attributes = append(attributes, fmt.Sprintf("timeout=\"%s\"", a.Timeout))
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
	Cache    bool
	Timeout  string
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
	if v.Cache {
		attributes = append(attributes, fmt.Sprintf("cache=\"%t\"", v.Cache))
	}
	if v.Timeout != "" {
		attributes = append(attributes, fmt.Sprintf("timeout=\"%s\"", v.Timeout))
	}
	return fmt.Sprintf("<video %s/>", strings.Join(attributes, " "))
}

// FileElement 文件消息元素
type FileElement struct {
	Src     string
	Title   string
	Poster  string
	Cache   bool
	Timeout string
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
	if f.Cache {
		attributes = append(attributes, fmt.Sprintf("cache=\"%t\"", f.Cache))
	}
	if f.Timeout != "" {
		attributes = append(attributes, fmt.Sprintf("timeout=\"%s\"", f.Timeout))
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
	if q.Elements != nil {
		var elements []string
		for _, element := range *q.Elements {
			elements = append(elements, element.ToSatori())
		}
		attributes = append(attributes, strings.Join(elements, ""))
	}
	return fmt.Sprintf("<quote %s/>", strings.Join(attributes, " "))
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
