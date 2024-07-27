package main

import (
	"log"
	"github.com/abava00/KakinoTane/internal/ai"
	"github.com/abava00/KakinoTane/internal/misskey"
)

func main() {
	client, _ := misskey.NewMisskeyClient()
	client.GetStats()
	pre_text := "Good Morning"
	openai.NumTokens(pre_text)
	text := openai.Openai(pre_text)
	openai.NumTokens(text)
	err := client.CreateNote(text)
	if err != nil {
		log.Fatalf("Failed to create note")
	}
}

