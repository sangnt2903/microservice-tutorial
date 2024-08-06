package handler

import (
	"SAI/proto/generated/ecosystem/document/document_pb"
	"context"
	"time"
)

func (s *Service) CreateDocument(ctx context.Context, request *document_pb.DocumentCreateRequest) (*document_pb.DocumentCreateReply, error) {
	time.Sleep(time.Second * 1)
	return &document_pb.DocumentCreateReply{}, nil
}
