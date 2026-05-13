package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	apiBaseUrl = "https://api.groq.com/openai"
	SYSTEM     = "system"
	USER       = "user"

	LLMModelLlama38b       = "llama3-8b-8192"
	LLMModelLlama370b      = "llama3-70b-8192"
	LLMModelMixtral8x7b32k = "mixtral-8x7b-32768"
	LLMModelGemma7b        = "gemma-7b-it"
)

type GroqClient struct {
	ApiKey string
}

type GroqMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatCompletionRequest struct {
	Messages    []GroqMessage `json:"messages"`
	Model       string        `json:"model"`
	Temperature int           `json:"temperature"`
	MaxTokens   int           `json:"max_tokens"`
	TopP        int           `json:"top_p"`
	Stream      bool          `json:"stream"`
	Stop        interface{}   `json:"stop"`
}

type ChatCompletionResponse struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int     `json:"prompt_tokens"`
		PromptTime       float64 `json:"prompt_time"`
		CompletionTokens int     `json:"completion_tokens"`
		CompletionTime   float64 `json:"completion_time"`
		TotalTokens      int     `json:"total_tokens"`
		TotalTime        float64 `json:"total_time"`
	} `json:"usage"`
	SystemFingerprint string `json:"system_fingerprint"`
	XGroq             struct {
		Id string `json:"id"`
	} `json:"x_groq"`
}

func (g *GroqClient) ChatCompletion(llmModel string, systemPrompt string, prompt string) (string, error) {

	llm := llmModel

	if llmModel == "" {
		//default to llama8B
		llm = LLMModelLlama38b
	}
	groqMessages := make([]GroqMessage, 0)

	if systemPrompt != "" {
		systemMessage := GroqMessage{
			Role:    SYSTEM,
			Content: systemPrompt,
		}
		groqMessages = append(groqMessages, systemMessage)
	}

	if prompt != "" {
		userMessage := GroqMessage{
			Role:    USER,
			Content: prompt,
		}
		groqMessages = append(groqMessages, userMessage)
	} else {
		return "", fmt.Errorf("prompt is required")
	}

	chatCompletionRequest := &ChatCompletionRequest{
		Messages:    groqMessages,
		Model:       llm,
		Temperature: 0,
		MaxTokens:   1024,
		TopP:        1,
		Stream:      false,
		Stop:        nil,
	}

	chatCompletionRequestJson, err := json.Marshal(chatCompletionRequest)
	if err != nil {
		return "", err
	}

	//send http post request
	chatCompletionUrl := "/v1/chat/completions"
	finalUrl := fmt.Sprintf("%s%s", apiBaseUrl, chatCompletionUrl)

	req, err := http.NewRequest(http.MethodPost, finalUrl, bytes.NewBuffer(chatCompletionRequestJson))
	if err != nil {
		return "", err
	}

	//set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", g.ApiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("unexpected status code: %d, reason: %s", resp.StatusCode, resp.Status)
	}

	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	chatCompletionResp := &ChatCompletionResponse{}

	err = json.Unmarshal(body, &chatCompletionResp)
	if err != nil {
		return "", err
	}

	var content string
	if chatCompletionResp.Choices != nil && len(chatCompletionResp.Choices) > 0 {
		content = chatCompletionResp.Choices[0].Message.Content
	} else {
		return "", fmt.Errorf("no choices")
	}

	return content, nil
}
