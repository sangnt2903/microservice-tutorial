package openai

import "github.com/sashabaranov/go-openai"

var client *openai.Client

func init() {
	client = openai.NewClient("")
}
