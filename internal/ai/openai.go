package openai

import (
	"context"
	"fmt"
	"os"
	openai "github.com/sashabaranov/go-openai"
)

func Openai(text string) string {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4oMini20240718,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: text,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return ""
	}

	fmt.Println(resp.Choices[0].Message.Content)
	return resp.Choices[0].Message.Content
}
