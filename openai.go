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
			Model:    openai.GPT3Dot5Turbo16K,
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
			Content: PROMPT_SM_DOCUMENTATION,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: string(json),
		},
	})
	return content, err
}

func (oai *OpenAI) CreateCodeUpdateRequest(command CodeUpdateCommand) (string, error) {
	json, err := json.Marshal(command.Codebase)
	if err != nil {
		return "", err
	}

	content, err := oai.completionRequest([]openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: PROMPT_SM_CODE_UPDATE,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "TASK NAME:" + command.Name,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "TASK DESCRIPTION:" + command.Description,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "TASK ACCEPTANCE CRITERIA:" + command.AcceptanceCriteria,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: "CODEBASE",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: string(json),
		},
	})

	return content, err
}
