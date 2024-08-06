package jaeger_wrapper

import (
	"SAI/pkg/conf"
	"SAI/pkg/tracing/span"

	"SAI/pkg/service"
	"fmt"
	"github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	"io"
)

var (
	host, _ = conf.GetString("jaeger", "host", "")
)

type Tracer struct {
	t      opentracing.Tracer
	closer io.Closer
}

func (t *Tracer) StartSpan(operationName string) span.Span {
	return span.New(t.t.StartSpan(operationName))
}

func New() *Tracer {
	sender, _ := jaeger.NewUDPTransport(host, 0)
	tracer, closer := jaeger.NewTracer(fmt.Sprintf("[%s]%s", service.AppMode, service.Name), jaeger.NewConstSampler(true), jaeger.NewRemoteReporter(sender))
	opentracing.SetGlobalTracer(tracer)

	return &Tracer{
		t:      tracer,
		closer: closer,
	}
}
