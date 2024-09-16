package lagrange

import (
	"fmt"
	"github.com/LagrangeDev/LagrangeGo/message"
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
