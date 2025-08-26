package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func Init() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Log.Out = file
	} else {
		Log.Warn("Failed to log to file, using default stderr")
	}
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}
