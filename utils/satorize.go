package utils

import (
	"github.com/Iceinu-Project/iceinu/elements"
	"strings"
)

// SatorizeIceElements 将Iceinu的元素切片解压成Satori字符串
func SatorizeIceElements(e *[]elements.IceinuMessageElement) string {
	var satoriStrings []string
	if e != nil {
		for _, element := range *e {
			satoriStrings = append(satoriStrings, element.ToSatori())
		}
	}
	return strings.Join(satoriStrings, "")
}
