package log

type Logger interface {
	Info(event func() string, fields ...Field)
	Debug(event func() string, fields ...Field)
	Error(event func() string, fields ...Field)
	Warn(event func() string, fields ...Field)
	Panic(event func() string, fields ...Field)
	Fatal(event func() string, fields ...Field)
}
