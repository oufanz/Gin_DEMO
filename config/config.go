package config

import (
	"Gin_DEMO/utils/logs"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	TimeFormat string = "2006-01-02 15:04:05"
)

var (
	Port string //大写表示其他文件可以访问
)

func initLogConfig(logLevel string) {
	//配置日志输出级别
	if logLevel == "debug" {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}
	//加入文件名行号
	logrus.SetReportCaller(true)
	//日志格式改为json
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: TimeFormat})
}
func init() {
	// fmt.Println("加载程序配置")
	logs.Debug(nil, "开始加载程序配置")
	//环境变量加载我们的程序配置
	viper.SetDefault("LOG_LEVEL", "debug")
	viper.AutomaticEnv()
	//配置程序启动端口号
	viper.SetDefault("PORT", ":8080")
	logLevel := viper.GetString("LOG_LEVEL") //获取程序配置
	Port = viper.GetString("PORT")
	//加载日志输出格式
	initLogConfig(logLevel)
}
