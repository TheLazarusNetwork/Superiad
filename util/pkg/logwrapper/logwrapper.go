package logwrapper

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log *logrus.Entry

func Init() {
	hostname, err := os.Hostname()
	Log = logrus.New().WithFields(logrus.Fields{
		"hostname": hostname,
	})
	if err != nil {
		Log.Warnf("Error in getting hostname: %s", err)
	}
	Log.Logger.SetFormatter(&logrus.JSONFormatter{})

}
