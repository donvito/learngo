package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestChatCompletionSendsRequestAndReturnsContent(t *testing.T) {
	var request ChatCompletionRequest

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected method %s, got %s", http.MethodPost, r.Method)
		}

		if r.URL.Path != "/v1/chat/completions" {
			t.Errorf("expected path /v1/chat/completions, got %s", r.URL.Path)
		}

		if got := r.Header.Get("Authorization"); got != "Bearer test-api-key" {
			t.Errorf("expected authorization bearer token, got %q", got)
		}

		if got := r.Header.Get("Content-Type"); got != "application/json" {
			t.Errorf("expected application/json content type, got %q", got)
		}

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{
			"id": "chatcmpl-test",
			"object": "chat.completion",
			"created": 1710000000,
			"model": "gpt-4o-mini",
			"choices": [
				{
					"index": 0,
					"message": {
						"role": "assistant",
						"content": "Go is concise and ships with a strong standard library."
					},
					"finish_reason": "stop"
				}
			]
		}`))
	}))
	defer server.Close()

	client := &OpenAIClient{
		APIKey:  "test-api-key",
		BaseURL: server.URL,
	}

	content, err := client.ChatCompletion("", "You are concise.", "Why use Go for API clients?")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if content == nil {
		t.Fatal("expected content, got nil")
	}

	if *content != "Go is concise and ships with a strong standard library." {
		t.Errorf("unexpected content: %q", *content)
	}

	if request.Model != defaultOpenAIModel {
		t.Errorf("expected default model %q, got %q", defaultOpenAIModel, request.Model)
	}

	if len(request.Messages) != 2 {
		t.Fatalf("expected 2 messages, got %d", len(request.Messages))
	}

	if request.Messages[0].Role != roleSystem || request.Messages[0].Content != "You are concise." {
		t.Errorf("unexpected system message: %+v", request.Messages[0])
	}

	if request.Messages[1].Role != roleUser || request.Messages[1].Content != "Why use Go for API clients?" {
		t.Errorf("unexpected user message: %+v", request.Messages[1])
	}
}

func TestChatCompletionRequiresAPIKey(t *testing.T) {
	client := &OpenAIClient{}

	content, err := client.ChatCompletion(defaultOpenAIModel, "", "Hello")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if content != nil {
		t.Fatalf("expected nil content, got %q", *content)
	}

	if !strings.Contains(err.Error(), "api key is required") {
		t.Errorf("expected api key error, got %v", err)
	}
}

func TestChatCompletionRequiresPrompt(t *testing.T) {
	client := &OpenAIClient{APIKey: "test-api-key"}

	content, err := client.ChatCompletion(defaultOpenAIModel, "", " ")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if content != nil {
		t.Fatalf("expected nil content, got %q", *content)
	}

	if !strings.Contains(err.Error(), "prompt is required") {
		t.Errorf("expected prompt error, got %v", err)
	}
}

func TestChatCompletionReturnsErrorForUnexpectedStatus(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "bad request", http.StatusBadRequest)
	}))
	defer server.Close()

	client := &OpenAIClient{
		APIKey:  "test-api-key",
		BaseURL: server.URL,
	}

	content, err := client.ChatCompletion(defaultOpenAIModel, "", "Hello")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if content != nil {
		t.Fatalf("expected nil content, got %q", *content)
	}

	if !strings.Contains(err.Error(), "unexpected status code: 400") {
		t.Errorf("expected status code error, got %v", err)
	}
}

func TestChatCompletionReturnsErrorWhenChoicesAreMissing(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"choices":[]}`))
	}))
	defer server.Close()

	client := &OpenAIClient{
		APIKey:  "test-api-key",
		BaseURL: server.URL,
	}

	content, err := client.ChatCompletion(defaultOpenAIModel, "", "Hello")
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if content != nil {
		t.Fatalf("expected nil content, got %q", *content)
	}

	if !strings.Contains(err.Error(), "no choices returned") {
		t.Errorf("expected no choices error, got %v", err)
	}
}
