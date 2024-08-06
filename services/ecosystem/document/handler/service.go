package handler

import (
	"SAI/pkg/contexts"
	"SAI/pkg/log"
	"SAI/proto/generated/ecosystem/document/document_pb"
	"context"
)

type Service struct {
	document_pb.UnimplementedServiceServer
}

func (s *Service) Ping(ctx context.Context, request *document_pb.PingRequest) (*document_pb.PingReply, error) {
	contexts.Logger(ctx).Info(func() string {
		return "Service Ping"
	}, log.NewField("hi", 22))
	contexts.Logger(ctx).Info(func() string {
		return "Service Ping"
	}, log.NewField("hi", 22))
	contexts.Logger(ctx).Info(func() string {
		return "Service Ping"
	}, log.NewField("hi", 22))
	contexts.Logger(ctx).Info(func() string {
		return "Service Ping"
	}, log.NewField("hi", 22))
	return &document_pb.PingReply{}, nil
}

func New() *Service {
	return &Service{}
}
