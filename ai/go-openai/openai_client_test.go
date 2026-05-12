package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGenerateTextSendsResponsesRequest(t *testing.T) {
	var request ResponsesRequest

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST request, got %s", r.Method)
		}

		if r.URL.Path != "/v1/responses" {
			t.Errorf("expected /v1/responses path, got %s", r.URL.Path)
		}

		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("expected application/json content type, got %s", r.Header.Get("Content-Type"))
		}

		if r.Header.Get("Authorization") != "Bearer test-api-key" {
			t.Errorf("unexpected authorization header: %s", r.Header.Get("Authorization"))
		}

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			t.Fatalf("decode request body: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write([]byte(`{
			"id": "resp_123",
			"object": "response",
			"model": "gpt-4o-mini",
			"output": [
				{
					"id": "msg_123",
					"type": "message",
					"status": "completed",
					"content": [
						{
							"type": "output_text",
							"text": "Learn the basics, write small programs, and read standard library docs."
						}
					]
				}
			]
		}`))
		if err != nil {
			t.Fatalf("write response: %v", err)
		}
	}))
	defer server.Close()

	client := &OpenAIClient{
		APIKey:  "test-api-key",
		BaseURL: server.URL,
	}

	response, err := client.GenerateText("", "Keep it short.", "How should I learn Go?")
	if err != nil {
		t.Fatalf("GenerateText returned error: %v", err)
	}

	if response != "Learn the basics, write small programs, and read standard library docs." {
		t.Errorf("unexpected response: %s", response)
	}

	if request.Model != defaultOpenAIModel {
		t.Errorf("expected default model %s, got %s", defaultOpenAIModel, request.Model)
	}

	if request.Instructions != "Keep it short." {
		t.Errorf("unexpected instructions: %s", request.Instructions)
	}

	if request.Input != "How should I learn Go?" {
		t.Errorf("unexpected input: %s", request.Input)
	}

	if request.MaxOutputTokens != 300 {
		t.Errorf("expected max output tokens 300, got %d", request.MaxOutputTokens)
	}
}

func TestGenerateTextUsesProvidedModel(t *testing.T) {
	var request ResponsesRequest

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			t.Fatalf("decode request body: %v", err)
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write([]byte(`{
			"output": [
				{
					"content": [
						{
							"text": "custom model response"
						}
					]
				}
			]
		}`))
		if err != nil {
			t.Fatalf("write response: %v", err)
		}
	}))
	defer server.Close()

	client := &OpenAIClient{
		APIKey:  "test-api-key",
		BaseURL: server.URL,
	}

	_, err := client.GenerateText("gpt-4.1-mini", "", "Say hello")
	if err != nil {
		t.Fatalf("GenerateText returned error: %v", err)
	}

	if request.Model != "gpt-4.1-mini" {
		t.Errorf("expected provided model, got %s", request.Model)
	}
}

func TestGenerateTextRequiresAPIKey(t *testing.T) {
	client := &OpenAIClient{}

	_, err := client.GenerateText("", "", "Say hello")
	if err == nil {
		t.Fatal("expected error")
	}

	if err.Error() != "api key is required" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestGenerateTextRequiresPrompt(t *testing.T) {
	client := &OpenAIClient{APIKey: "test-api-key"}

	_, err := client.GenerateText("", "", "")
	if err == nil {
		t.Fatal("expected error")
	}

	if err.Error() != "prompt is required" {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestGenerateTextReturnsOpenAIErrorMessage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
		_, err := w.Write([]byte(`{
			"error": {
				"message": "Incorrect API key provided",
				"type": "invalid_request_error",
				"code": "invalid_api_key"
			}
		}`))
		if err != nil {
			t.Fatalf("write response: %v", err)
		}
	}))
	defer server.Close()

	client := &OpenAIClient{
		APIKey:  "bad-api-key",
		BaseURL: server.URL,
	}

	_, err := client.GenerateText("", "", "Say hello")
	if err == nil {
		t.Fatal("expected error")
	}

	if !strings.Contains(err.Error(), "Incorrect API key provided") {
		t.Errorf("expected OpenAI error message, got %v", err)
	}
}

func TestGenerateTextReturnsNoOutputTextError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte(`{"output":[]}`))
		if err != nil {
			t.Fatalf("write response: %v", err)
		}
	}))
	defer server.Close()

	client := &OpenAIClient{
		APIKey:  "test-api-key",
		BaseURL: server.URL,
	}

	_, err := client.GenerateText("", "", "Say hello")
	if err == nil {
		t.Fatal("expected error")
	}

	if err.Error() != "no output text" {
		t.Errorf("unexpected error: %v", err)
	}
}
