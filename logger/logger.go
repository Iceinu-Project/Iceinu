package logger

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
	logger.SetLevel(logrus.InfoLevel)

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

// Info 输出Info级别的日志
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Infof 格式化输出Info级别的日志
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Debug 输出Debug级别的日志
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Debugf 格式化输出Debug级别的日志
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// Warn 输出Warn级别的日志
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Warnf 格式化输出Warn级别的日志
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// Error 输出Error级别的日志
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Errorf 格式化输出Error级别的日志
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

// Fatal 输出Fatal级别的日志
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Fatalf 格式化输出Fatal级别的日志
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

// Panic 输出Panic级别的日志
func Panic(args ...interface{}) {
	logger.Panic(args...)
}

// Panicf 格式化输出Panic级别的日志
func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

func SetLevel(level string) {
	l, err := logrus.ParseLevel(level)
	if err != nil {
		logger.Warnf("设置日志等级失败：%v", err)
		return
	}
	logger.Infof("日志等级已设置为：%s", level)
	logger.SetLevel(l)
}
