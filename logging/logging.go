package logging

import (
	"github.com/sirupsen/logrus"
	"time"
)

func SetLogger(log *logrus.Logger) {
	log = logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(
		&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: time.RFC3339Nano,
		},
	)
}
