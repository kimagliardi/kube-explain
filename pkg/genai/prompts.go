package genai

import (
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func resourceUsagePrompt(resourceDescription, model string) openai.ChatCompletionRequest {
	return openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo, // replace later
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "You are a skilled Kubernetes administrator. Your task is to provide concise summaries of Kubernetes resources based on their descriptions. Focus on extracting and highlighting critical information such as status, configuration details, resource allocations, events, and any warnings or anomalies. Your summary should distill the essential details into a brief overview, offering quick insights into the resource's current state, usage, and health.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: fmt.Sprintf("I need a summary of this Kubernetes resource. Please highlight the key aspects like status, resource usage, and any critical events or warnings that need attention. Here's the output of the `kubectl describe` command: %v", resourceDescription),
			},
		}}
}
