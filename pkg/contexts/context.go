package contexts

import (
	"SAI/pkg/tracing/span"
	"context"
)

type serverCtx struct{}
type traceCtx struct{}

type serverData struct {
}

func NewServerContext(ctx context.Context) context.Context {
	d := ctx.Value(serverCtx{})
	if d != nil {
		return ctx
	}

	return context.WithValue(ctx, serverCtx{}, &serverData{})
}

func SpanContext(ctx context.Context, s span.Span) context.Context {
	_, ok := ctx.Value(traceCtx{}).(span.Span)
	if ok {
		return ctx
	}
	return context.WithValue(ctx, traceCtx{}, s)
}

func Logger(ctx context.Context) span.Span {
	s, ok := ctx.Value(traceCtx{}).(span.Span)
	if ok {
		return s
	}

	return nil
}
