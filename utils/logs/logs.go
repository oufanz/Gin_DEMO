package logs

import "github.com/sirupsen/logrus"

//debug类型日志
func Debug(fields map[string]interface{}, msg string) {
	logrus.WithFields(fields).Debug(msg)

}
func Info(fields map[string]interface{}, msg string) {
	logrus.WithFields(fields).Info(msg)

}
