package log_level

const (
	DebugLevel = "debug"
	InfoLevel  = "info"
	WarnLevel  = "warn"
	ErrorLevel = "error"
	PanicLevel = "panic"
	FatalLevel = "fatal"
)

type Level int

func ToLevel(level string) Level {
	switch level {
	case DebugLevel:
		return debugLevel
	case InfoLevel:
		return infoLevel
	case WarnLevel:
		return warnLevel
	case ErrorLevel:
		return errorLevel
	case PanicLevel:
		return panicLevel
	case FatalLevel:
		return fatalLevel
	default:
		return infoLevel
	}
}

const (
	debugLevel Level = iota
	infoLevel
	warnLevel
	errorLevel
	panicLevel
	fatalLevel
)
