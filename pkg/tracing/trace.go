package tracing

import (
	"SAI/pkg/conf"
	"SAI/pkg/log/jaeger_wrapper"
	"SAI/pkg/log/log_level"
	"SAI/pkg/tracing/span"
)

var (
	traceEnable, _   = conf.GetBool("trace", "enable", true)
	traceLevel, _    = conf.GetString("trace", "level", log_level.InfoLevel)
	traceProvider, _ = conf.GetString("trace", "provider", "jaeger")
)

type ITracer interface {
	StartSpan(operationName string) span.Span
}

func StartSpan(operationName string) span.Span {
	return t.StartSpan(operationName)
}

type tracer struct {
	trace ITracer
}

func (t *tracer) StartSpan(operationName string) span.Span {
	if traceEnable {
		return t.trace.StartSpan(operationName)
	}
	return span.NoopSpan()
}

var t = newTracer()

func newTracer() ITracer {
	t := &tracer{}

	switch traceProvider {
	case "jaeger":
		t.trace = jaeger_wrapper.New()
	}

	return t
}
