package starter

import (
	"github.com/sirupsen/logrus"
	"simple-skeleton/boot"
)

type LogStarter struct {
	boot.BaseStarter
}

func (l *LogStarter) Init(ctx boot.StaterContext) {
	// 定义日志格式
	formatter := &logrus.TextFormatter{}
	logrus.SetFormatter(formatter)

	// 日志级别
	level := ctx.Conf().App.Debug

	if level == true {
		logrus.SetLevel(logrus.DebugLevel)
	}

	formatter.ForceColors = true
	formatter.DisableColors = false
	formatter.FullTimestamp = true
	formatter.TimestampFormat = "2006-01-02 15:04:05.000"

	logrus.Info("日志配置成功")
}
