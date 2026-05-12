package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		err := errors.New("OPENAI_API_KEY need to be set as an environment variable")
		panic(err)
	}

	openAIClient := &OpenAIClient{APIKey: apiKey}
	instructions := "You are a helpful assistant. Keep the answer short and practical."
	prompt := "Give me three tips for learning Go."

	response, err := openAIClient.GenerateText(defaultOpenAIModel, instructions, prompt)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(response)
}
