package genai

import (
	"context"
	"log"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

type Client struct {
	client *openai.Client
}

func (c *Client) Summarize(resourceDescription string) (openai.ChatCompletionResponse, error) {
	{
		prompt := resourceUsagePrompt(resourceDescription, openai.GPT3Dot5Turbo)
		resp, err := c.client.CreateChatCompletion(
			context.Background(),
			prompt,
		)
		return resp, err
	}
}

func NewClient() *Client {
	openai_key := getAPIkey()
	client := openai.NewClient(openai_key)
	return &Client{client}
}

func getAPIkey() string {
	openai_key := os.Getenv("OPEN_AI_API_KEY")
	if openai_key == "" {
		log.Fatal("OPEN_AI_API_KEY is not set")
		return ""
	}
	return openai_key
}
