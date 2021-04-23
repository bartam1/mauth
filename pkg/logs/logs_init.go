package logs

import (
	"strconv"

	"github.com/bartam1/mauth/pkg/config"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func Init() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "time",
			logrus.FieldKeyLevel: "severity",
			logrus.FieldKeyMsg:   "message",
		},
	})

	if isLocalEnv, _ := strconv.ParseBool(config.Global.LOCAL_ENV); isLocalEnv {
		logrus.SetFormatter(&prefixed.TextFormatter{
			ForceFormatting: true,
		})
	}
	level, _ := strconv.Atoi(config.Global.LOG_LEVEL)
	logrus.SetLevel(logrus.Level(level))
}
