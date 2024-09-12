package log

import (
	"fmt"
	"github.com/KyokuKong/gradients"
	"github.com/sirupsen/logrus"
	"os"
)

var logger *logrus.Logger

// 初始化日志Logger
func init() {
	// 创建一个新的logrus.Logger实例
	logger = logrus.New()

	// 设置日志级别为InfoLevel
	logger.SetLevel(logrus.DebugLevel)

	// 设置输出格式为JSON格式
	logger.SetFormatter(&IceinuFormatter{})

	// 将日志输出到标准输出
	logger.SetOutput(os.Stdout)
}

type IceinuFormatter struct{}

func (f *IceinuFormatter) Format(entry *logrus.Entry) ([]byte, error) {
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

// GetLogger 获取日志Logger实例
func GetLogger() *logrus.Logger {
	return logger
}

// SetLoggerFormatter 设置自定义日志输出
func SetLoggerFormatter(formatter logrus.Formatter) {
	logger.SetFormatter(formatter)
}
