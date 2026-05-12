package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		err := errors.New("OPENAI_API_KEY needs to be set as an environment variable")
		panic(err)
	}

	openAIClient := &OpenAIClient{APIKey: apiKey}

	systemPrompt := "You are a helpful assistant. Keep your answer concise."
	prompt := "Explain why Go is a good language for building API clients."

	response, err := openAIClient.ChatCompletion(defaultOpenAIModel, systemPrompt, prompt)
	if err != nil {
		fmt.Println(err)
		return
	}

	if response != nil {
		fmt.Println(*response)
	}
}
