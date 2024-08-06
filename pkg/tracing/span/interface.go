package span

import (
	"SAI/pkg/log"
	"github.com/opentracing/opentracing-go"
)

type Span interface {
	SetTag(key string, value string)
	AddField(key string, value string)
	Finish()

	Info(event func() string, fields ...log.Field)
	Debug(event func() string, fields ...log.Field)
	Error(event func() string, fields ...log.Field)
	Warn(event func() string, fields ...log.Field)
	Panic(event func() string, fields ...log.Field)
	Fatal(event func() string, fields ...log.Field)
}

func NoopSpan() Span {
	return &noopSpan{}
}

func New(s opentracing.Span) Span {
	return &span{
		logger: log.GetLogger(),
		span:   s,
	}
}
