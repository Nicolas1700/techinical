package repository

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
)

type ChatGptApi interface {
	ChatGptMessague(ctx context.Context, message string) (result string, err error)
}

var _ ChatGptApi = (*chatGptApi)(nil)

type chatGptApi struct {
	keyOpenIa string
}

func NewChatGptApi(keyOpenIa string) ChatGptApi {
	return &chatGptApi{
		keyOpenIa: keyOpenIa,
	}
}

func (api chatGptApi) ChatGptMessague(ctx context.Context, message string) (result string, err error) {
	client := openai.NewClient(api.keyOpenIa)
	// Definimos el cuerpo de la peticion
	resp, err := client.CreateChatCompletion(ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
		},
	)
	if err != nil {
		return
	}
	result = resp.Choices[0].Message.Content
	return

}
