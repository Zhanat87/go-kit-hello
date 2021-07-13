package loggers

import (
	"os"

	"github.com/sirupsen/logrus"
)

type LogrusLoggerFactory struct{}

func (s *LogrusLoggerFactory) CreateLogger() *logrus.Logger {
	return &logrus.Logger{
		Out:   os.Stdout,
		Level: logrus.Level(5),
		Formatter: &logrus.TextFormatter{
			FullTimestamp: true,
		},
	}
}
