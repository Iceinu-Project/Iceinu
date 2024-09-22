package main

import (
	"fmt"
	"github.com/Iceinu-Project/iceinu/log"
	"github.com/KyokuKong/gradients"
	"github.com/sirupsen/logrus"
)

// Iceinu的程序入口
// 可以参照文档来对这其进行修改

func main() {
	// 定义日志格式
	formatter := &LogFormatter{}
	log.SetFormatter(formatter)
	// 输出日志
	log.Info("Hello, World!")
}

// LogFormatter 可以通过修改这个结构体的Format方法来设置你想要的日志格式
//
// 其本身实现了Logrus的Formatter接口
type LogFormatter struct{}

func (f *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 根据日志级别设置不同的颜色
	var textColor string
	var levelColor string
	var levelText string
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gradients.DarkGreen
		textColor = gradients.DarkGreen
		levelText = "DEBUG"
	case logrus.InfoLevel:
		levelColor = gradients.DarkCyan
		textColor = gradients.White
		levelText = "_INFO"
	case logrus.WarnLevel:
		levelColor = gradients.Orange
		textColor = gradients.DarkYellow
		levelText = "_WARN"
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = gradients.Red
		textColor = gradients.Red
		levelText = "ERROR"
	default:
		levelColor = gradients.White
		textColor = gradients.White
		levelText = "UNKNOWN"
	}

	// 构建日志格式,可以按需修改
	logMsg := fmt.Sprintf(
		"%s• %s %s[%s%s%s] %s%s\n",
		gradients.Gray,
		entry.Time.Format("2006-01-02 15:04:05"),
		textColor,
		levelColor,
		levelText,
		textColor,
		entry.Message,
		gradients.ResetColor,
	)

	return []byte(logMsg), nil
}
