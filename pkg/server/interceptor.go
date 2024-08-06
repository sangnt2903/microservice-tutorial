package server

import (
	"SAI/pkg/contexts"
	"SAI/pkg/helper/json_helper"
	"SAI/pkg/log"
	"SAI/pkg/tracing"

	"context"
	"google.golang.org/grpc"
)

func requestInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		reqCtx := contexts.NewServerContext(ctx)

		span := tracing.StartSpan(info.FullMethod)
		defer span.Finish()
		spanCtx := contexts.SpanContext(reqCtx, span)

		contexts.Logger(spanCtx).Info(func() string {
			return "Call to handler " + info.FullMethod
		}, log.NewField("request", json_helper.SafeString(req)))

		resp, err := handler(spanCtx, req)
		if err != nil {
			span.SetTag("error", "true")
			contexts.Logger(spanCtx).Error(func() string {
				return "call to " + info.FullMethod + " failed"
			}, log.Error(err))
			return nil, err
		}

		contexts.Logger(spanCtx).Info(func() string {
			return "Call to handler " + info.FullMethod + " success"
		}, log.NewField("response", json_helper.SafeString(resp)))

		return resp, nil
	}
}
