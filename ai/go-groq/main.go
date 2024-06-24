package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	apiKey := os.Getenv("GROQ_API_KEY")
	if apiKey == "" {
		err := errors.New("GROQ_API_KEY need to be set as an environment variable")
		panic(err)
	}

	groqClient := &GroqClient{ApiKey: apiKey}
	textToTranslate := "Kim Ji-won was born on October 19, 1992, in Geumcheon District, Seoul, South Korea, and has an elder sister who is two years older than her. While still a teenager in 2007, she was scouted on the street and signed with an entertainment agency, she subsequently became a trainee for over three years while preparing for her debut. During her first year of junior high school, she spent six months to a year studying in Chicago, Illinois, United States, where her maternal relatives lived."

	systemPrompt := "you are a professional language translator. " +
		"only respond with the translated text. do not explain."
	prompt := fmt.Sprintf("translate this text to tagalog: %s", textToTranslate)

	translatedText, err := groqClient.ChatCompletion(LLMModelLlama370b, systemPrompt, prompt)
	if err != nil {
		fmt.Println(err)
	}

	if translatedText != nil {
		fmt.Println(*translatedText)
	}
}
