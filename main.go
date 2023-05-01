package main

import (
	"fmt"
	"os"
	"github.com/vinayprograms/chatgpt-cli/chatgpt"
)

func main() {
	fmt.Println("Welcome to ChatGPT CLI!")

	// TODO: Parse CLI arguments and handle other requirements.

	// TODO: Replace this with the actual API key from the file specified by the "-k" flag.
	apiKey := "your-api-key"
	client := chatgpt.NewClient(apiKey)

	var chatHistory []chatgpt.Message

	chatHistory = append(chatHistory, chatgpt.Message{Role: "user", Content: "tell me a joke"})

	req := &chatgpt.CompletionRequest{
		Messages:  chatHistory,
		MaxTokens: 1000, // Limit the response length, adjust this as needed.
	}

	response, err := client.GenerateCompletion(req)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Add the AI response to the chat history.
	chatHistory = append(chatHistory, chatgpt.Message{Role: "assistant", Content: response})

	fmt.Printf("AI: %s\n", response)
}
