package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

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
		log.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, err
}

func (oai *OpenAI) CreateDocumentationRequest(dir_info Directory) (string, error) {
	json, err := json.Marshal(dir_info)
	if err != nil {
		log.Println("Error marshalling directory info: ", err)
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
		log.Println("Error marshalling codebase: ", err)
		return "", err
	}

	promt := fmt.Sprintf(PROMPT_SM_CODE_UPDATE, command.Name, command.TaskType, command.Description, command.AcceptanceCriteria, string(json))

	content, err := oai.completionRequest([]openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: promt,
		},
	})

	return content, err
}

func (oai *OpenAI) CreateFileSummaryRequest(content string) (string, error) {
	content, err := oai.completionRequest([]openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: PROMPT_FILE_SUMMARY,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: content,
		},
	})

	return content, err
}

func (oai *OpenAI) CreateProjectSummaryRequest(content string) (string, error) {
	content, err := oai.completionRequest([]openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: PROMPT_PROJECT_SUMMARY,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: content,
		},
	})

	return content, err
}
