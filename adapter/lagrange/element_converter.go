package lagrange

import (
	"github.com/Iceinu-Project/iceinu/elements"
	"github.com/LagrangeDev/LagrangeGo/message"
	"strconv"
	"strings"
	"time"
)

// ConvertIceElement 将LagrangeGo的元素转换为Iceinu的元素
func ConvertIceElement(e []message.IMessageElement) *[]elements.IceinuMessageElement {
	// 从LagrangeGo的元素转换为Iceinu的元素
	var IceinuElements []elements.IceinuMessageElement
	// 遍历传入的元素
	for _, ele := range e {
		switch ele.Type() {
		// 将元素依次对应转换并传入
		case message.Text:
			ele := ele.(*message.TextElement)
			// 检测文本中是否包含换行符
			if strings.Contains(ele.Content, "\n") {
				// 如果包含换行符，进行拆分和处理
				textParts := strings.Split(ele.Content, "\n")
				for i, part := range textParts {
					// 将每段文本添加到 IceinuElements
					IceinuElements = append(IceinuElements, &elements.TextElement{Text: part})
					// 如果不是最后一段文本，则插入 BrElement
					if i < len(textParts)-1 {
						IceinuElements = append(IceinuElements, &elements.BrElement{})
					}
				}
			} else {
				// 如果不包含换行符，直接添加文本元素
				IceinuElements = append(IceinuElements, &elements.TextElement{Text: ele.Content})
			}
		case message.At:
			ele := ele.(*message.AtElement)
			IceinuElements = append(IceinuElements, &elements.AtElement{
				Id:   strconv.Itoa(int(ele.TargetUin)),
				Name: ele.Display,
				Role: "",
				Type: strconv.Itoa(int(ele.Type())),
			})
		case message.Face:
			ele := ele.(*message.FaceElement)
			IceinuElements = append(IceinuElements, &elements.FaceElement{
				Id: ele.FaceID,
			})
		case message.Voice:
			ele := ele.(*message.VoiceElement)
			IceinuElements = append(IceinuElements, &elements.AudioElement{
				Src:      ele.Url,
				Title:    ele.Name,
				Duration: ele.Size,
				Poster:   "",
			})
		case message.Image:
			ele := ele.(*message.ImageElement)
			IceinuElements = append(IceinuElements, &elements.ImageElement{
				Src:    ele.Url,
				Width:  ele.Width,
				Height: ele.Height,
				Title:  ele.ImageId,
			})
		case message.File:
			ele := ele.(*message.FileElement)
			IceinuElements = append(IceinuElements, &elements.FileElement{
				Src:   ele.FileUrl,
				Title: ele.FileName,
			})
		case message.Reply:
			ele := ele.(*message.ReplyElement)
			IceinuElements = append(IceinuElements, &elements.QuoteElement{
				UserId:    strconv.Itoa(int(ele.SenderUin)),
				UserName:  ele.SenderUid,
				GroupId:   strconv.Itoa(int(ele.GroupUin)),
				Timestamp: time.Unix(int64(ele.Time), 0),
				Elements:  ConvertIceElement(ele.Elements),
			})
		case message.Forward:
			ele := ele.(*message.ForwardMessage)
			IceinuElements = append(IceinuElements, &elements.MessageElement{
				Forward:  true,
				Elements: UnzipNodes(ele.Nodes),
			})

		default:
			IceinuElements = append(IceinuElements, &elements.UnsupportedElement{Type: strconv.Itoa(int(ele.Type()))})
		}
	}
	return &IceinuElements
}

func UnzipNodes(n []*message.ForwardNode) *[]elements.IceinuMessageElement {
	var IceinuElements []elements.IceinuMessageElement
	for _, node := range n {
		IceinuElements = append(IceinuElements, &elements.NodeElement{
			GroupId:    node.GroupId,
			SenderId:   node.SenderId,
			SenderName: node.SenderName,
			Time:       node.Time,
			Message:    ConvertIceElement(node.Message),
		})
	}
	return &IceinuElements
}
