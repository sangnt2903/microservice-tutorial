package handler

import (
	"SAI/proto/generated/ecosystem/document/document_pb"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io/ioutil"
	"net/http"
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

func TestOpenAI(t *testing.T) {
	url := "http://localhost:1337/v1/chat/completions"

	// Create the request body
	requestBody := map[string]interface{}{
		"model":    "gpt-3.5-turbo",
		"provider": "Pizzagpt",
		"messages": []map[string]string{
			{"role": "system", "content": "Bạn là một reviewer nổi tiếng với 300k follows trên instagram, bạn sẽ viết review cho những địa điểm theo yêu cầu bên dưới. Ngôn từ bạn dùng teencode và bức phá"},
			{"role": "user", "content": "Viết 3 reviews, mỗi review khoảng trên 300 từ về nhà ở phú mỹ hưng sử dụng icon và hashtag ở định dạng json only (khong can markdown format) để tôi có thể parse sử dụng xen kẽ ngôi thứ 3 và thứ 1 để review. Không cần xuống dòng chỉ cần sử dụng dấu chấm. Ngôn ngữ năng động và hài hước"},
		},
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the headers
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Authorization", "Bearer "+apiKey)

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var response openai.ChatCompletionResponse
	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Print the response
	fmt.Println("Response:", string(response.Choices[0].Message.Content))
}
