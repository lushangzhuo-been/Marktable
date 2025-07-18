package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

func InitLog() *logrus.Logger {
	//实例化
	logger := logrus.New()
	//设置输出
	logger.SetOutput(os.Stdout)
	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{})
	return logger
}
