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
	apiBaseURL         = "https://api.openai.com"
	defaultOpenAIModel = "gpt-4o-mini"
)

type OpenAIClient struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}

type ResponsesRequest struct {
	Model           string `json:"model"`
	Instructions    string `json:"instructions,omitempty"`
	Input           string `json:"input"`
	MaxOutputTokens int    `json:"max_output_tokens,omitempty"`
}

type ResponsesResponse struct {
	ID     string `json:"id"`
	Object string `json:"object"`
	Model  string `json:"model"`
	Output []struct {
		ID      string `json:"id"`
		Type    string `json:"type"`
		Status  string `json:"status"`
		Content []struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"content"`
	} `json:"output"`
}

type OpenAIErrorResponse struct {
	Error struct {
		Message string `json:"message"`
		Type    string `json:"type"`
		Code    string `json:"code"`
	} `json:"error"`
}

func (o *OpenAIClient) GenerateText(model string, instructions string, prompt string) (string, error) {
	if strings.TrimSpace(o.APIKey) == "" {
		return "", fmt.Errorf("api key is required")
	}

	if strings.TrimSpace(prompt) == "" {
		return "", fmt.Errorf("prompt is required")
	}

	llm := model
	if llm == "" {
		llm = defaultOpenAIModel
	}

	openAIRequest := &ResponsesRequest{
		Model:           llm,
		Instructions:    instructions,
		Input:           prompt,
		MaxOutputTokens: 300,
	}

	openAIRequestJSON, err := json.Marshal(openAIRequest)
	if err != nil {
		return "", err
	}

	baseURL := o.BaseURL
	if baseURL == "" {
		baseURL = apiBaseURL
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/v1/responses", strings.TrimRight(baseURL, "/")), bytes.NewBuffer(openAIRequestJSON))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", o.APIKey))

	client := o.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d, reason: %s, message: %s", resp.StatusCode, resp.Status, parseOpenAIError(body))
	}

	openAIResponse := &ResponsesResponse{}
	err = json.Unmarshal(body, openAIResponse)
	if err != nil {
		return "", err
	}

	for _, output := range openAIResponse.Output {
		for _, content := range output.Content {
			if content.Text != "" {
				return content.Text, nil
			}
		}
	}

	return "", fmt.Errorf("no output text")
}

func parseOpenAIError(body []byte) string {
	openAIError := &OpenAIErrorResponse{}
	err := json.Unmarshal(body, openAIError)
	if err != nil || openAIError.Error.Message == "" {
		return string(body)
	}

	return openAIError.Error.Message
}
