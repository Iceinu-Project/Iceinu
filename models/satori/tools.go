package satori

import (
	"strconv"
	"strings"
)

// ParseSatori 从Satori标准XHTML消息中解析出Satori标准事件元素
func ParseSatori(content string) *[]ElementSatori {
	return nil
}

// ElementsToSatori 将Satori标准事件元素转换为Satori标准XHTML消息字符串
func ElementsToSatori(elements []ElementSatori) string {
	var sb strings.Builder

	for _, element := range elements {
		switch e := element.(type) {
		case *TextElement:
			sb.WriteString(escapeXML(e.Text))
		case *AtElement:
			sb.WriteString(`<at id="` + escapeXML(e.Id) + `" name="` + escapeXML(e.Name) + `" role="` + escapeXML(e.Role) + `" type="` + escapeXML(e.Type) + `"/>`)
		case *SharpElement:
			sb.WriteString(`<sharp id="` + escapeXML(e.Id) + `" name="` + escapeXML(e.Name) + `"/>`)
		case *AElement:
			sb.WriteString(`<a href="` + escapeXML(e.Href) + `"/>`)
		case *ImgElement:
			sb.WriteString(`<img src="` + escapeXML(e.Src) + `" title="` + escapeXML(e.Title) + `" cache="` + boolToString(e.Cache) + `" timeout="` + uint32ToString(e.Timeout) + `" width="` + uint32ToString(e.Width) + `" height="` + uint32ToString(e.Height) + `"/>`)
		case *AudioElement:
			sb.WriteString(`<audio src="` + escapeXML(e.Src) + `" title="` + escapeXML(e.Title) + `" cache="` + boolToString(e.Cache) + `" timeout="` + uint32ToString(e.Timeout) + `" duration="` + uint32ToString(e.Duration) + `" poster="` + escapeXML(e.Poster) + `"/>`)
		case *VideoElement:
			sb.WriteString(`<video src="` + escapeXML(e.Src) + `" title="` + escapeXML(e.Title) + `" cache="` + boolToString(e.Cache) + `" timeout="` + uint32ToString(e.Timeout) + `" width="` + uint32ToString(e.Width) + `" height="` + uint32ToString(e.Height) + `" duration="` + uint32ToString(e.Duration) + `" poster="` + escapeXML(e.Poster) + `"/>`)
		case *FileElement:
			sb.WriteString(`<file src="` + escapeXML(e.Src) + `" title="` + escapeXML(e.Title) + `" cache="` + boolToString(e.Cache) + `" timeout="` + uint32ToString(e.Timeout) + `" poster="` + escapeXML(e.Poster) + `"/>`)
		case *StrongElement:
			sb.WriteString(`<strong>`)
			sb.WriteString(ElementsToSatori(*e.Elements))
			sb.WriteString(`</strong>`)
		case *EmElement:
			sb.WriteString(`<em>`)
			sb.WriteString(ElementsToSatori(*e.Elements))
			sb.WriteString(`</em>`)
		case *InsElement:
			sb.WriteString(`<ins>`)
			sb.WriteString(ElementsToSatori(*e.Elements))
			sb.WriteString(`</ins>`)
		case *DelElement:
			sb.WriteString(`<del>`)
			sb.WriteString(ElementsToSatori(*e.Elements))
			sb.WriteString(`</del>`)
		case *SpoilerElement:
			sb.WriteString(`<spoiler>`)
			sb.WriteString(ElementsToSatori(*e.Elements))
			sb.WriteString(`</spoiler>`)
		case *CodeElement:
			sb.WriteString(`<code>`)
			sb.WriteString(ElementsToSatori(*e.Elements))
			sb.WriteString(`</code>`)
		case *SupElement:
			sb.WriteString(`<sup>`)
			sb.WriteString(ElementsToSatori(*e.Elements))
			sb.WriteString(`</sup>`)
		case *SubElement:
			sb.WriteString(`<sub>`)
			sb.WriteString(ElementsToSatori(*e.Elements))
			sb.WriteString(`</sub>`)
		case *BrElement:
			sb.WriteString(`<br/>`)
		case *HrElement:
			sb.WriteString(`<hr/>`)
		case *PElement:
			sb.WriteString(`<p>`)
			sb.WriteString(ElementsToSatori(*e.Elements))
			sb.WriteString(`</p>`)
		case *MessageElement:
			sb.WriteString(`<message id="` + escapeXML(e.Id) + `" forward="` + boolToString(e.Forward) + `">`)
			sb.WriteString(ElementsToSatori(*e.Elements))
			sb.WriteString(`</message>`)
		case *QuoteElement:
			sb.WriteString(`<quote id="` + escapeXML(e.Id) + `" name="` + escapeXML(e.Name) + `" groupId="` + escapeXML(e.GroupId) + `" channelId="` + escapeXML(e.ChannelId) + `" timestamp="` + int64ToString(e.Timestamp) + `">`)
			sb.WriteString(ElementsToSatori(*e.Elements))
			sb.WriteString(`</quote>`)
		case *AuthorElement:
			sb.WriteString(`<author id="` + escapeXML(e.Id) + `" name="` + escapeXML(e.Name) + `" avatar="` + escapeXML(e.Avatar) + `"/>`)
		case *ButtonElement:
			sb.WriteString(`<button id="` + escapeXML(e.Id) + `" type="` + escapeXML(e.Type) + `" href="` + escapeXML(e.Href) + `" text="` + escapeXML(e.Text) + `" theme="` + escapeXML(e.Theme) + `"/>`)
		case *FaceElement:
			sb.WriteString(`<face id="` + uint16ToString(e.Id) + `" isLargeFace="` + boolToString(e.IsLargeFace) + `"/>`)
		case *NodeElement:
			sb.WriteString(`<node groupId="` + int64ToString(e.GroupId) + `" senderId="` + int64ToString(e.SenderId) + `" senderName="` + escapeXML(e.SenderName) + `" time="` + int32ToString(e.Time) + `">`)
			sb.WriteString(ElementsToSatori(*e.Message))
			sb.WriteString(`</node>`)
		case *UnsupportedElement:
			sb.WriteString(`<unsupported type="` + escapeXML(e.Type) + `"/>`)
		}
	}

	return sb.String()
}

// escapeXML 转义特殊字符
func escapeXML(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(s, "&", "&amp;"), "\"", "&quot;"), "<", "&lt;"), ">", "&gt;")
}

// boolToString 将布尔值转换为字符串
func boolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

// uint32ToString 将 uint32 转换为字符串
func uint32ToString(u uint32) string {
	return strconv.FormatUint(uint64(u), 10)
}

// int64ToString 将 int64 转换为字符串
func int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

// int32ToString 将 int32 转换为字符串
func int32ToString(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

// uint16ToString 将 uint16 转换为字符串
func uint16ToString(u uint16) string {
	return strconv.FormatUint(uint64(u), 10)
}
