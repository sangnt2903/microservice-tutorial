package log

import (
	"SAI/pkg/conf"
	"SAI/pkg/log/log_level"
	"fmt"
)

var (
	level, _  = conf.GetString("console", "level", log_level.InfoLevel)
	enable, _ = conf.GetBool("console", "enable", true)
)

type consoleLogger struct{}

func (l *consoleLogger) Info(event func() string, fields ...Field) {
	if enable && level >= log_level.InfoLevel {
		l.write(log_level.InfoLevel, getData(log_level.InfoLevel, event, fields...))
	}
}

func (l *consoleLogger) Debug(event func() string, fields ...Field) {
	if enable && level >= log_level.DebugLevel {
		l.write(log_level.DebugLevel, getData(log_level.DebugLevel, event, fields...))
	}
}

func (l *consoleLogger) Error(event func() string, fields ...Field) {
	if enable && level >= log_level.ErrorLevel {
		l.write(log_level.ErrorLevel, getData(log_level.ErrorLevel, event, fields...))
	}
}

func (l *consoleLogger) Warn(event func() string, fields ...Field) {
	if enable && level >= log_level.WarnLevel {
		l.write(log_level.WarnLevel, getData(log_level.WarnLevel, event, fields...))
	}
}

func (l *consoleLogger) Panic(event func() string, fields ...Field) {
	if enable && level >= log_level.PanicLevel {
		l.write(log_level.PanicLevel, getData(log_level.PanicLevel, event, fields...))
	}
}

func (l *consoleLogger) Fatal(event func() string, fields ...Field) {
	if enable && level >= log_level.FatalLevel {
		l.write(log_level.FatalLevel, getData(log_level.FatalLevel, event, fields...))
	}
}

func (l *consoleLogger) write(logLevel string, log string) {
	should := log_level.ToLevel(logLevel) >= log_level.ToLevel(level)
	if should {
		fmt.Println(log)
	}
}

func newConsoleLogger() *consoleLogger {
	return &consoleLogger{}
}
