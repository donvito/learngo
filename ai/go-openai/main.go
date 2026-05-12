package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		panic(errors.New("OPENAI_API_KEY needs to be set as an environment variable"))
	}

	openAIClient := &OpenAIClient{APIKey: apiKey}
	systemPrompt := "you are a helpful assistant. answer clearly and concisely."
	prompt := "write a one sentence explanation of why unit tests are useful in Go"

	response, err := openAIClient.ChatCompletion("", systemPrompt, prompt)
	if err != nil {
		fmt.Println(err)
		return
	}

	if response != nil {
		fmt.Println(*response)
	}
}
