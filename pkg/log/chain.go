package log

import (
	"SAI/pkg/conf"
)

type chainLogger struct {
	loggers []Logger
}

func (chain *chainLogger) Info(event func() string, fields ...Field) {
	for _, logger := range chain.loggers {
		logger.Info(event, fields...)
	}
}

func (chain *chainLogger) Debug(event func() string, fields ...Field) {
	for _, logger := range chain.loggers {
		logger.Debug(event, fields...)
	}
}

func (chain *chainLogger) Error(event func() string, fields ...Field) {
	for _, logger := range chain.loggers {
		logger.Error(event, fields...)
	}
}

func (chain *chainLogger) Warn(event func() string, fields ...Field) {
	for _, logger := range chain.loggers {
		logger.Warn(event, fields...)
	}
}

func (chain *chainLogger) Panic(event func() string, fields ...Field) {
	for _, logger := range chain.loggers {
		logger.Panic(event, fields...)
	}
}

func (chain *chainLogger) Fatal(event func() string, fields ...Field) {
	for _, logger := range chain.loggers {
		logger.Fatal(event, fields...)
	}
}

var (
	loggers, _ = conf.GetStringSlice("log", "loggers")
)

var defaultLogger Logger

func init() {
	defaultLogger = _new()
}

func _new() Logger {
	c := &chainLogger{}
	for _, logger := range loggers {
		switch logger {
		case "console":
			c.Add(newConsoleLogger())
		}
	}

	return c
}

func GetLogger() Logger {
	return defaultLogger
}

func (chain *chainLogger) Add(logger Logger) {
	chain.loggers = append(chain.loggers, logger)
}
