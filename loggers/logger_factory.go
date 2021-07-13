package loggers

type LoggerFactory interface {
	CreateLogger() Logger
}
