package config

import "github.com/sirupsen/logrus"

const layoutDateTime = "2006-01-02T15:04:05"

func Logrus() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: layoutDateTime,
	})
	
	return log
}
