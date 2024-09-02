package openai

import (
	"context"
	"github.com/sashabaranov/go-openai"
)

type DocumentUploadRequest struct {
	FileUrl string
}

type DocumentUploadReply struct {
	ID string
}

func UploadDocument(ctx context.Context, request *DocumentUploadRequest) (*DocumentUploadReply, error) {
	client.CreateVectorStore(ctx, openai.FileBytesRequest{
		Name:    "",
		Bytes:   nil,
		Purpose: openai.PurposeAssistants,
	})
}
