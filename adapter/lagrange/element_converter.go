package lagrange

import (
	"github.com/LagrangeDev/LagrangeGo/message"
	"gtihub.com/Iceinu-Project/iceinu/elements"
	"strconv"
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
			IceinuElements = append(IceinuElements, &elements.TextElement{Text: ele.Content})
		case message.At:
			ele := ele.(*message.AtElement)
			IceinuElements = append(IceinuElements, &elements.AtElement{
				Id:   strconv.Itoa(int(ele.TargetUin)),
				Name: ele.Display,
				Role: "",
				Type: "",
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
				Cache:    false,
				Timeout:  "",
			})
		case message.Image:
			ele := ele.(*message.ImageElement)
			IceinuElements = append(IceinuElements, &elements.ImageElement{
				Src:     ele.Url,
				Width:   ele.Width,
				Height:  ele.Height,
				Title:   ele.ImageId,
				Cache:   false,
				Timeout: "",
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
				UserId:    strconv.Itoa(int(ele.GroupUin)),
				UserName:  ele.SenderUid,
				GroupId:   strconv.Itoa(int(ele.GroupUin)),
				Timestamp: time.Unix(int64(ele.Time), 0),
				Elements:  ConvertIceElement(ele.Elements),
			})

		default:
			IceinuElements = append(IceinuElements, &elements.UnsupportedElement{Type: strconv.Itoa(int(ele.Type()))})
		}
	}
	return &IceinuElements
}
