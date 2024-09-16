package lagrange

import (
	"encoding/xml"
	"fmt"
	"github.com/LagrangeDev/LagrangeGo/message"
	"strconv"
	"strings"
)

func IElementsToSatoriMsg(elements []message.IMessageElement) string {
	var satoriMsg strings.Builder

	// 遍历 IMessageElement 数组
	for _, element := range elements {
		switch element.Type() {
		case message.Text:
			textElem := element.(*message.TextElement)
			satoriMsg.WriteString(textElem.Content)

		case message.At:
			atElem := element.(*message.AtElement)
			if atElem.TargetUin == 0 {
				satoriMsg.WriteString("<at type=\"all\"/>")
			} else {
				satoriMsg.WriteString(fmt.Sprintf("<at id=\"%d\" name=\"%s\"/>", atElem.TargetUin, atElem.Display))
			}

		case message.Voice:
			voiceElem := element.(*message.VoiceElement)
			satoriMsg.WriteString(fmt.Sprintf("<audio src=\"%s\"/>", voiceElem.Url))

		case message.Image:
			imageElem := element.(*message.ImageElement)
			satoriMsg.WriteString(fmt.Sprintf("<img src=\"%s\" width=\"%d\" height=\"%d\"/>", imageElem.Url, imageElem.Width, imageElem.Height))

		case message.File:
			fileElem := element.(*message.FileElement)
			satoriMsg.WriteString(fmt.Sprintf("<file src=\"%s\" title=\"%s\"/>", fileElem.FileUrl, fileElem.FileName))

		case message.Reply:
			replyElem := element.(*message.ReplyElement)
			satoriMsg.WriteString(fmt.Sprintf("<quote id=\"%d\"/>", replyElem.ReplySeq))

		default:
			// 未知类型元素的处理
			satoriMsg.WriteString(fmt.Sprintf("<!-- Unsupported element type: %d -->", element.Type()))
		}
	}

	return satoriMsg.String()
}

// SatoriMsg 定义 Satori XML 消息结构体
type SatoriMsg struct {
	XMLName  xml.Name     `xml:"message"`
	Elements []xmlElement `xml:",any"` // 任意消息元素
}

// 定义任意消息元素
type xmlElement struct {
	XMLName xml.Name   `xml:""`
	Attr    []xml.Attr `xml:",any,attr"` // 任意属性
	Content string     `xml:",chardata"` // 文本内容
}

// SatoriXMLToIMessageElements 将 Satori XML 消息解析为 []IMessageElement
func SatoriXMLToIMessageElements(xmlData string) ([]message.IMessageElement, error) {
	var satoriMsg SatoriMsg
	err := xml.Unmarshal([]byte(xmlData), &satoriMsg)
	if err != nil {
		return nil, fmt.Errorf("解析XML失败: %v", err)
	}

	var elements []message.IMessageElement

	// 遍历解析的XML元素
	for _, el := range satoriMsg.Elements {
		switch el.XMLName.Local {
		case "at":
			elements = append(elements, parseAtElement(el))
		case "audio":
			elements = append(elements, parseAudioElement(el))
		case "img":
			elements = append(elements, parseImageElement(el))
		case "file":
			elements = append(elements, parseFileElement(el))
		case "quote":
			elements = append(elements, parseReplyElement(el))
		default:
			// 处理文本内容
			if el.Content != "" {
				elements = append(elements, message.NewText(el.Content))
			}
		}
	}

	return elements, nil
}

// 解析 <at> 元素
func parseAtElement(el xmlElement) *message.AtElement {
	var targetUin uint32
	var display string
	for _, attr := range el.Attr {
		switch attr.Name.Local {
		case "id":
			id, _ := strconv.ParseUint(attr.Value, 10, 32)
			targetUin = uint32(id)
		case "name":
			display = attr.Value
		}
	}
	return message.NewAt(targetUin, display)
}

// 解析 <audio> 元素
func parseAudioElement(el xmlElement) *message.VoiceElement {
	var url string
	for _, attr := range el.Attr {
		if attr.Name.Local == "src" {
			url = attr.Value
		}
	}
	return &message.VoiceElement{Url: url}
}

// 解析 <img> 元素
func parseImageElement(el xmlElement) *message.ImageElement {
	var url string
	var width, height uint32
	for _, attr := range el.Attr {
		switch attr.Name.Local {
		case "src":
			url = attr.Value
		case "width":
			w, _ := strconv.Atoi(attr.Value)
			width = uint32(w)
		case "height":
			h, _ := strconv.Atoi(attr.Value)
			height = uint32(h)
		}
	}
	return &message.ImageElement{
		Url:    url,
		Width:  width,
		Height: height,
	}
}

// 解析 <file> 元素
func parseFileElement(el xmlElement) *message.FileElement {
	var fileUrl, fileName string
	for _, attr := range el.Attr {
		switch attr.Name.Local {
		case "src":
			fileUrl = attr.Value
		case "title":
			fileName = attr.Value
		}
	}
	return &message.FileElement{
		FileUrl:  fileUrl,
		FileName: fileName,
	}
}

// 解析 <quote> 元素（用于回复）
func parseReplyElement(el xmlElement) *message.ReplyElement {
	var replySeq uint32
	for _, attr := range el.Attr {
		if attr.Name.Local == "id" {
			id, _ := strconv.ParseUint(attr.Value, 10, 32)
			replySeq = uint32(id)
		}
	}
	return &message.ReplyElement{ReplySeq: replySeq}
}
