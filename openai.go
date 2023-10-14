package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

type OpenAI struct {
	jwt    string
	client *openai.Client
}

func (oai *OpenAI) Init(jwt string) {
	oai.jwt = jwt
	oai.client = openai.NewClient(jwt)
}

func (oai *OpenAI) completionRequest(messages []openai.ChatCompletionMessage) (string, error) {
	resp, err := oai.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: messages,
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, err
}

func (oai *OpenAI) CreateDocumentationRequest(dir_info Directory) (string, error) {
	json, err := json.Marshal(dir_info)
	if err != nil {
		return "", err
	}

	content, err := oai.completionRequest([]openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: DOCUMENTATION_SYSTEM_MESSAGE,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: string(json),
		},
	})
	return content, err
}