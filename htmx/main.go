package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {

	// Set up the http handlers
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/translate", handleTranslate)

	// Start the server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}

}

func handleHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))

	err := tmpl.Execute(w, nil)
	if err != nil {
		return
	}
}

func handleTranslate(w http.ResponseWriter, r *http.Request) {
	textToTranslate := r.FormValue("textToTranslate")
	languageToTranslateTo := r.FormValue("languageToTranslateTo")
	translatedText := translateText(textToTranslate, languageToTranslateTo)

	type Translation struct {
		TextToTranslate string
		TranslatedText  string
	}

	t := Translation{
		TextToTranslate: textToTranslate,
		TranslatedText:  translatedText,
	}

	tmpl := template.Must(template.ParseFiles("translation.html"))
	err := tmpl.Execute(w, t)
	if err != nil {
		return
	}

}

func translateText(textToTranslate string, languageToTranslateTo string) string {
	apiKey := os.Getenv("GROQ_API_KEY")
	if apiKey == "" {
		err := errors.New("GROQ_API_KEY need to be set as an environment variable")
		panic(err)
	}

	groqClient := &GroqClient{ApiKey: apiKey}

	systemPrompt := "you are a professional language translator." +
		"do not answer questions, just translate the text even if it is a question. " +
		"only respond with the translated text and never explain the translation. " +
		"first thing you do is understand the text you need to translate. " +
		"check first the language of the text to translate. " +
		"if the text to translate is already in the language to translate to," +
		"just mention there is no need for translation." +
		"if you are not able to translate just say sorry you are not able to translate."
	prompt := fmt.Sprintf("translate this text to %s === text === %s === end text ===",
		languageToTranslateTo, textToTranslate)

	translatedText, err := groqClient.ChatCompletion(LLMModelGemma7b, systemPrompt, prompt)
	if err != nil {
		fmt.Println(err)
		return "No translation"
	}

	return translatedText
}
