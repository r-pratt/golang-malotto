/*
	@author: Robert Pratt <robert.pratt[at]homelab.farm>
	@package: logging
	@description: logging wrapper
	@created: 2020-12-05
*/
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
