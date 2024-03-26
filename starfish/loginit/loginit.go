package loginit

import (
	"io"
	"starfish/starfish/globalflag"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLog() {

	logrus.SetLevel(logrus.Level(viper.GetInt("loglevel")))
	if globalflag.IsDaemon {

		logrus.Infoln("Enable daemon mode")
		globalflag.IsDaemon = true
		logFile := &lumberjack.Logger{
			Filename:   viper.GetString("logfilepath") + "/log.txt", // 日志文件名
			MaxSize:    10,                                          // 最大文件大小（MB）
			MaxBackups: 3,                                           // 最大备份文件数
			MaxAge:     7,                                           // 保留的最大天数
		}
		logrus.SetOutput(io.MultiWriter(logrus.StandardLogger().Out, logFile))
		// 开启日志输出的文件行号
		logrus.SetReportCaller(true)
	} else {
		logrus.SetOutput(logrus.StandardLogger().Out)
		logrus.Infoln("Enable console mode")
	}
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02T15:04:05Z07:00",
	})

}
