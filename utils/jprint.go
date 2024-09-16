package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func JPrint(input interface{}) {
	// 创建 JSON 编码器并禁用 HTML 转义
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetEscapeHTML(false) // 禁用自动转义

	// 美化输出
	encoder.SetIndent("", "  ")

	// 尝试将输入转换为 JSON 字符串并输出到 stdout
	err := encoder.Encode(input)
	if err != nil {
		fmt.Println("无法将输入转换为 JSON:", err)
		return
	}
}
