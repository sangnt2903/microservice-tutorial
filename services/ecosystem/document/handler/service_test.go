package handler

import (
	"SAI/proto/generated/ecosystem/document/document_pb"
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

func TestService_Ping(t *testing.T) {
	clientConn, err := grpc.NewClient("localhost:18000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.Nil(t, err)
	assert.NotNil(t, clientConn)
	service := document_pb.NewServiceClient(clientConn)
	resp, err := service.Ping(context.Background(), &document_pb.PingRequest{})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}
