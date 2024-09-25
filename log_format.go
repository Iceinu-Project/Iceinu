package main

import (
	"fmt"
	"github.com/Iceinu-Project/IceGradient"
	"github.com/sirupsen/logrus"
)

// LogFormatter 可以通过修改这个结构体的Format方法来设置你想要的日志格式
type LogFormatter struct{}

func (f *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 根据日志级别设置不同的颜色
	var textColor string
	var levelColor string
	var levelText string
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = gradient.DarkGreen
		textColor = gradient.DarkGreen
		levelText = "DEBUG"
	case logrus.InfoLevel:
		levelColor = gradient.DarkCyan
		textColor = gradient.Reset
		levelText = "_INFO"
	case logrus.WarnLevel:
		levelColor = gradient.Orange
		textColor = gradient.DarkOrange
		levelText = "_WARN"
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = gradient.Red
		textColor = gradient.Red
		levelText = "ERROR"
	default:
		levelColor = gradient.White
		textColor = gradient.White
		levelText = "UNKNOWN"
	}

	// 构建日志格式,可以按需修改
	logMsg := fmt.Sprintf(
		"%s• %s %s[%s%s%s] %s%s\n",
		gradient.Gray,
		entry.Time.Format("2006-01-02 15:04:05"),
		textColor,
		levelColor,
		levelText,
		textColor,
		entry.Message,
		gradient.Reset,
	)

	return []byte(logMsg), nil
}
