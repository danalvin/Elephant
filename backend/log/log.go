package log

import (
	"elephant/config"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log = logrus.New()

func init() {

	conf := config.GetConfig()

	log.Formatter = &logrus.TextFormatter{
		TimestampFormat: "2020-01-02 23:12:01",
		FullTimestamp:   true,
	}

	log.Out = os.Stdout 
	log.Formatter = &logrus.JSONFormatter{}

	log.SetOutput(&lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/elephant.log", conf.GetString("app.path_to_log_file")),
		MaxSize:    200, // mbs,
		MaxBackups: 2,
		MaxAge:     28, // days
	})

	log.SetLevel(logrus.InfoLevel)
}

// GetLogger - return an instance of our custom logger
func GetLogger() *logrus.Logger {
	return log
}
