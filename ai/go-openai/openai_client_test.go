package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestChatCompletionSendsRequestAndReturnsContent(t *testing.T) {
	t.Parallel()

	requestChecked := false
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/chat/completions" {
			t.Errorf("expected path /v1/chat/completions, got %s", r.URL.Path)
		}

		if r.Method != http.MethodPost {
			t.Errorf("expected method POST, got %s", r.Method)
		}

		if got := r.Header.Get("Content-Type"); got != "application/json" {
			t.Errorf("expected content type application/json, got %s", got)
		}

		if got := r.Header.Get("Authorization"); got != "Bearer test-api-key" {
			t.Errorf("expected authorization header Bearer test-api-key, got %s", got)
		}

		var req ChatCompletionRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}

		if req.Model != "gpt-4.1-mini" {
			t.Errorf("expected model gpt-4.1-mini, got %s", req.Model)
		}

		if req.Temperature != 0 {
			t.Errorf("expected temperature 0, got %f", req.Temperature)
		}

		if req.MaxTokens != 1024 {
			t.Errorf("expected max tokens 1024, got %d", req.MaxTokens)
		}

		expectedMessages := []OpenAIMessage{
			{Role: roleSystem, Content: "be brief"},
			{Role: roleUser, Content: "say hello"},
		}

		if len(req.Messages) != len(expectedMessages) {
			t.Fatalf("expected %d messages, got %d", len(expectedMessages), len(req.Messages))
		}

		for i, expected := range expectedMessages {
			if req.Messages[i] != expected {
				t.Errorf("message %d expected %+v, got %+v", i, expected, req.Messages[i])
			}
		}

		requestChecked = true
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"choices":[{"message":{"role":"assistant","content":"hello from openai"}}]}`))
	}))
	defer server.Close()

	client := &OpenAIClient{
		APIKey:  "test-api-key",
		BaseURL: server.URL + "/",
	}

	content, err := client.ChatCompletion("gpt-4.1-mini", "be brief", "say hello")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if content == nil {
		t.Fatal("expected content, got nil")
	}

	if *content != "hello from openai" {
		t.Errorf("expected response content, got %q", *content)
	}

	if !requestChecked {
		t.Fatal("server did not check the request")
	}
}

func TestChatCompletionUsesDefaultModel(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req ChatCompletionRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}

		if req.Model != defaultOpenAIModel {
			t.Errorf("expected default model %s, got %s", defaultOpenAIModel, req.Model)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"choices":[{"message":{"role":"assistant","content":"default model response"}}]}`))
	}))
	defer server.Close()

	client := &OpenAIClient{
		APIKey:  "test-api-key",
		BaseURL: server.URL,
	}

	content, err := client.ChatCompletion("", "", "hello")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if content == nil || *content != "default model response" {
		t.Fatalf("expected default model response, got %v", content)
	}
}

func TestChatCompletionRequiresPrompt(t *testing.T) {
	t.Parallel()

	client := &OpenAIClient{APIKey: "test-api-key"}

	content, err := client.ChatCompletion("", "", "   ")
	if err == nil {
		t.Fatal("expected prompt validation error")
	}

	if content != nil {
		t.Fatalf("expected nil content, got %v", *content)
	}

	if err.Error() != "prompt is required" {
		t.Fatalf("expected prompt is required error, got %v", err)
	}
}

func TestChatCompletionReturnsAPIError(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte(`{"error":{"message":"invalid api key"}}`))
	}))
	defer server.Close()

	client := &OpenAIClient{
		APIKey:  "bad-api-key",
		BaseURL: server.URL,
	}

	content, err := client.ChatCompletion("", "", "hello")
	if err == nil {
		t.Fatal("expected api error")
	}

	if content != nil {
		t.Fatalf("expected nil content, got %v", *content)
	}

	if !strings.Contains(err.Error(), "unexpected status code: 401") {
		t.Fatalf("expected status code in error, got %v", err)
	}
}

func TestChatCompletionReturnsNoChoicesError(t *testing.T) {
	t.Parallel()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"choices":[]}`))
	}))
	defer server.Close()

	client := &OpenAIClient{
		APIKey:  "test-api-key",
		BaseURL: server.URL,
	}

	content, err := client.ChatCompletion("", "", "hello")
	if err == nil {
		t.Fatal("expected no choices error")
	}

	if content != nil {
		t.Fatalf("expected nil content, got %v", *content)
	}

	if err.Error() != "no choices" {
		t.Fatalf("expected no choices error, got %v", err)
	}
}
