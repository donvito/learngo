package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const (
	defaultOpenAIBaseURL = "https://api.openai.com"
	defaultOpenAIModel   = "gpt-4o-mini"

	roleSystem = "system"
	roleUser   = "user"
)

type OpenAIClient struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}

type OpenAIMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatCompletionRequest struct {
	Model       string          `json:"model"`
	Messages    []OpenAIMessage `json:"messages"`
	Temperature float64         `json:"temperature"`
	MaxTokens   int             `json:"max_tokens"`
}

type ChatCompletionResponse struct {
	Choices []struct {
		Message OpenAIMessage `json:"message"`
	} `json:"choices"`
}

func (c *OpenAIClient) ChatCompletion(model string, systemPrompt string, prompt string) (*string, error) {
	if strings.TrimSpace(prompt) == "" {
		return nil, fmt.Errorf("prompt is required")
	}

	selectedModel := model
	if selectedModel == "" {
		selectedModel = defaultOpenAIModel
	}

	messages := make([]OpenAIMessage, 0, 2)
	if systemPrompt != "" {
		messages = append(messages, OpenAIMessage{
			Role:    roleSystem,
			Content: systemPrompt,
		})
	}

	messages = append(messages, OpenAIMessage{
		Role:    roleUser,
		Content: prompt,
	})

	chatCompletionRequest := &ChatCompletionRequest{
		Model:       selectedModel,
		Messages:    messages,
		Temperature: 0,
		MaxTokens:   1024,
	}

	requestBody, err := json.Marshal(chatCompletionRequest)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.chatCompletionURL(), bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))

	client := c.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}

	chatCompletionResponse := &ChatCompletionResponse{}
	if err := json.Unmarshal(body, chatCompletionResponse); err != nil {
		return nil, err
	}

	if len(chatCompletionResponse.Choices) == 0 {
		return nil, fmt.Errorf("no choices")
	}

	content := chatCompletionResponse.Choices[0].Message.Content
	return &content, nil
}

func (c *OpenAIClient) chatCompletionURL() string {
	baseURL := c.BaseURL
	if baseURL == "" {
		baseURL = defaultOpenAIBaseURL
	}

	return fmt.Sprintf("%s/v1/chat/completions", strings.TrimRight(baseURL, "/"))
}
