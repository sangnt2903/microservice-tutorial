package span

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"

	logfield "SAI/pkg/log"
)

type span struct {
	logger logfield.Logger
	span   opentracing.Span
}

func (s *span) Info(event func() string, fields ...logfield.Field) {
	var keyValues []interface{}
	keyValues = append(keyValues, "event", event())
	for _, field := range fields {
		keyValues = append(keyValues, field.Key, field.Value)
	}

	s.span.LogKV(keyValues...)
	s.logger.Info(event, fields...)
}

func (s *span) Debug(event func() string, fields ...logfield.Field) {
	var keyValues []interface{}
	keyValues = append(keyValues, "event", event())
	for _, field := range fields {
		keyValues = append(keyValues, field.Key, field.Value)
	}

	s.span.LogKV(keyValues...)
	s.logger.Debug(event, fields...)
}

func (s *span) Error(event func() string, fields ...logfield.Field) {
	var keyValues []interface{}
	keyValues = append(keyValues, "event", event())
	for _, field := range fields {
		keyValues = append(keyValues, field.Key, field.Value)
	}

	s.span.SetTag("error", "true")
	s.span.LogKV(keyValues...)
	s.logger.Error(event, fields...)
}

func (s *span) Warn(event func() string, fields ...logfield.Field) {
	var keyValues []interface{}
	keyValues = append(keyValues, "event", event())
	for _, field := range fields {
		keyValues = append(keyValues, field.Key, field.Value)
	}

	s.span.LogKV(keyValues...)
	s.logger.Warn(event, fields...)
}

func (s *span) Panic(event func() string, fields ...logfield.Field) {
	var keyValues []interface{}
	keyValues = append(keyValues, "event", event())
	for _, field := range fields {
		keyValues = append(keyValues, field.Key, field.Value)
	}

	s.span.LogKV(keyValues...)
	s.logger.Panic(event, fields...)
}

func (s *span) Fatal(event func() string, fields ...logfield.Field) {
	var keyValues []interface{}
	keyValues = append(keyValues, "event", event())
	for _, field := range fields {
		keyValues = append(keyValues, field.Key, field.Value)
	}

	s.span.LogKV(keyValues...)
	s.logger.Fatal(event, fields...)
}

func (s *span) SetTag(key string, value string) {
	s.span.SetTag(key, value)
}

func (s *span) AddField(key string, value string) {
	s.span.LogFields(log.String(key, value))
}

func (s *span) Finish() {
	s.span.Finish()
}

type noopSpan struct{}

func (s *noopSpan) Info(event func() string, fields ...logfield.Field) {

}

func (s *noopSpan) Debug(event func() string, fields ...logfield.Field) {

}

func (s *noopSpan) Error(event func() string, fields ...logfield.Field) {

}

func (s *noopSpan) Warn(event func() string, fields ...logfield.Field) {

}

func (s *noopSpan) Panic(event func() string, fields ...logfield.Field) {

}

func (s *noopSpan) Fatal(event func() string, fields ...logfield.Field) {

}

func (s *noopSpan) SetTag(key string, value string) {}

func (s *noopSpan) AddField(key string, value string) {}

func (s *noopSpan) Finish() {}
