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

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type OpenAIClient struct {
	APIKey     string
	BaseURL    string
	HTTPClient HTTPClient
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
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

func (c *OpenAIClient) ChatCompletion(model string, systemPrompt string, prompt string) (*string, error) {
	if strings.TrimSpace(c.APIKey) == "" {
		return nil, fmt.Errorf("api key is required")
	}

	if strings.TrimSpace(prompt) == "" {
		return nil, fmt.Errorf("prompt is required")
	}

	if model == "" {
		model = defaultOpenAIModel
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

	requestBody := &ChatCompletionRequest{
		Model:       model,
		Messages:    messages,
		Temperature: 0.7,
		MaxTokens:   300,
	}

	encodedBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.chatCompletionsURL(), bytes.NewBuffer(encodedBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))

	httpClient := c.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, string(body))
	}

	chatCompletionResp := &ChatCompletionResponse{}
	if err := json.Unmarshal(body, chatCompletionResp); err != nil {
		return nil, err
	}

	if len(chatCompletionResp.Choices) == 0 {
		return nil, fmt.Errorf("no choices returned")
	}

	content := chatCompletionResp.Choices[0].Message.Content
	return &content, nil
}

func (c *OpenAIClient) chatCompletionsURL() string {
	baseURL := strings.TrimRight(c.BaseURL, "/")
	if baseURL == "" {
		baseURL = defaultOpenAIBaseURL
	}

	return fmt.Sprintf("%s/v1/chat/completions", baseURL)
}
