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
	Port          string //大写表示其他文件可以访问
	JwtSignKey    string
	JwtEepireTime int64 //token维持时间：分钟
	Username      string
	Password      string
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
	//设置默认环境变量
	viper.SetDefault("JWT_SIGN_KEY", "ou_fan")
	viper.SetDefault("JWT_EXPIRE_TIME", 120)
	viper.SetDefault("PORT", ":8080")
	viper.SetDefault("USERNAME", "oufan")
	viper.SetDefault("PASSWORD", "dot")
	logLevel := viper.GetString("LOG_LEVEL") //获取程序配置
	Port = viper.GetString("PORT")
	JwtSignKey = viper.GetString("JWT_SIGN_KEY")
	JwtEepireTime = viper.GetInt64("JWT_EXPIRE_TIME")
	Username = viper.GetString("USERNAME")
	Password = viper.GetString("PASSWORD")
	//加载日志输出格式
	initLogConfig(logLevel)
}
