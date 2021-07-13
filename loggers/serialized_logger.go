package loggers

import (
	"sync"

	"github.com/go-kit/kit/log"
)

type serializedLogger struct {
	mtx sync.Mutex
	log.Logger
}

func (l *serializedLogger) Log(keyvals ...interface{}) error {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	return l.Logger.Log(keyvals...)
}
