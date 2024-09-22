package satori

// ElementSatori Satori标准事件元素接口
type ElementSatori interface {
	GetType() string
}

type TextElement struct {
	Text string
}

func (t *TextElement) GetType() string {
	return "text"
}

type AtElement struct {
	Id   string
	Name string
	Role string
	Type string
}

func (a *AtElement) GetType() string {
	return "at"
}

type SharpElement struct {
	Id   string
	Name string
}

func (s *SharpElement) GetType() string {
	return "sharp"
}

type AElement struct {
	Href string
}

func (a *AElement) GetType() string {
	return "a"
}

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

type StrongElement struct {
	Elements *[]ElementSatori
}

func (s *StrongElement) GetType() string {
	return "strong"
}

type EmElement struct {
	Elements *[]ElementSatori
}

func (e *EmElement) GetType() string {
	return "em"
}

type InsElement struct {
	Elements *[]ElementSatori
}

func (i *InsElement) GetType() string {
	return "ins"
}

type DelElement struct {
	Elements *[]ElementSatori
}

func (d *DelElement) GetType() string {
	return "del"
}

type SpoilerElement struct {
	Elements *[]ElementSatori
}

func (s *SpoilerElement) GetType() string {
	return "spoiler"
}

type CodeElement struct {
	Elements *[]ElementSatori
}

func (c *CodeElement) GetType() string {
	return "code"
}

type SupElement struct {
	Elements *[]ElementSatori
}

func (s *SupElement) GetType() string {
	return "sup"
}

type SubElement struct {
	Elements *[]ElementSatori
}

func (s *SubElement) GetType() string {
	return "sub"
}

type BrElement struct {
}

func (b *BrElement) GetType() string {
	return "br"
}

type HrElement struct {
}

func (h *HrElement) GetType() string {
	return "hr"
}

type PElement struct {
	Elements *[]ElementSatori
}

func (p *PElement) GetType() string {
	return "p"
}

type MessageElement struct {
	Id       string
	Forward  bool
	Elements *[]ElementSatori
}

func (m *MessageElement) GetType() string {
	return "message"
}

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

type AuthorElement struct {
	Id     string
	Name   string
	Avatar string
}

func (a *AuthorElement) GetType() string {
	return "author"
}

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
