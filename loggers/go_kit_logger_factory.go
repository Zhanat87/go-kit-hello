package loggers

import (
	"os"

	"github.com/go-kit/kit/log"
)

type GoKitLoggerFactory struct{}

func (s *GoKitLoggerFactory) CreateLogger() log.Logger {
	logger := log.NewLogfmtLogger(os.Stderr)
	logger = &serializedLogger{Logger: logger}
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	return logger
}
