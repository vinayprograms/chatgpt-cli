package chatgpt

import (
	"bytes"
	"errors"
	"io/ioutil"
	"encoding/json"
	"net/http"
)

const chatGPTAPIEndpoint = "https://api.openai.com/v1/chat/completions"

type Client struct {
	APIKey string
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type CompletionRequest struct {
	APIKey    string    `json:"-"`
	Messages  []Message `json:"messages"`
	MaxTokens int       `json:"max_tokens"`
}

func NewClient(apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
	}
}

func (c *Client) GenerateCompletion(req *CompletionRequest) (string, error) {
	client := &http.Client{}

	reqJSON, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	httpReq, err := http.NewRequest("POST", chatGPTAPIEndpoint, bytes.NewBuffer(reqJSON))
	if err != nil {
		return "", err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+c.APIKey)

	resp, err := client.Do(httpReq)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("request to ChatGPT API failed with status: " + resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result struct {
		Choices []struct {
			Text string `json:"text"`
		} `json:"choices"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	if len(result.Choices) > 0 {
		return result.Choices[0].Text, nil
	}

	return "", errors.New("no response from ChatGPT API")
}
