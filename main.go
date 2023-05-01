package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"strings"
	"github.com/vinayprograms/chatgpt-cli/chatgpt"
)

func main() {
	fmt.Println("Welcome to ChatGPT CLI!")

	var apiKeyFile string
	var organization string

	flag.StringVar(&apiKeyFile, "k", "", "Path to API key file.")
	flag.StringVar(&organization, "o", "", "Organization name used when setting up the key")
	flag.Parse()
	

	if apiKeyFile == "" {
		fmt.Println("Error: API key file not provided. Use the '-k' flag to specify the file path.")
		return
	}
	if organization == "" {
		fmt.Println("Error: Organization name not provided. Use the '-o' flag to specify.")
		return
	}
	// Read the API key from the file.
	apiKey, err := readAPIKeyFromFile(apiKeyFile)
	if err != nil {
		fmt.Printf("Error reading API key from file: %v\n", err)
		return
	}

	client := chatgpt.NewClient(apiKey, organization)

	var chatHistory []chatgpt.Message

	chatHistory = append(chatHistory, 
		chatgpt.Message{
			Role: "user", 
			Content: `
Tell me a joke that nobody has heard about!
			`})

	req := &chatgpt.CompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages:  chatHistory,
		Temperature: 0.7,
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

func readAPIKeyFromFile(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(content)), nil
}
