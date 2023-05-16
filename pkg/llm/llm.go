package llm

import (
	"context"
	"os"

	"github.com/otiai10/openaigo"
)

const defaultOpenaiChatModel = "gpt-3.5-turbo"

type LLM interface {
	Chat(message string) (string, error)
}

type llm struct {
	client *openaigo.Client
}

func New(apiKey string) LLM {
	if apiKey == "" {
		apiKey = os.Getenv("OPENAI_API_KEY")
	}
	client := openaigo.NewClient(apiKey)
	return &llm{client: client}
}

func (llm *llm) Chat(message string) (string, error) {
	request := openaigo.ChatCompletionRequestBody{
		Model: defaultOpenaiChatModel,
		Messages: []openaigo.ChatMessage{
			{Role: "user", Content: message},
		},
	}
	ctx := context.Background()
	response, err := llm.client.Chat(ctx, request)
	if err != nil {
		return "", nil
	}
	return response.Choices[0].Message.Content, nil
}

type MockLLM struct {
	chatResponses map[string]string
}

func (llm *MockLLM) Chat(message string) (string, error) {
	return llm.chatResponses[message], nil
}
