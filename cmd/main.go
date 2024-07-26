package main

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
	"github.com/yitsushi/go-misskey/services/notes"
	"github.com/abava00/KakinoTane/internal/ai"
)

func main() {
	client, err := misskey.NewClientWithOptions(
		misskey.WithAPIToken(os.Getenv("MISSKEY_TOKEN")),
		misskey.WithBaseURL("https", os.Getenv("BASE_URL"), ""),
		misskey.WithLogLevel(logrus.DebugLevel),
	)
	if err != nil {
		logrus.Error(err.Error())
	}

	stats, err := client.Meta().Stats()
	if err != nil {
		log.Printf("[Meta] Error happened: %s", err)
		return
	}

	log.Printf("[Stats] Instances:          %d", stats.Instances)
	log.Printf("[Stats] NotesCount:         %d", stats.NotesCount)
	log.Printf("[Stats] UsersCount:         %d", stats.UsersCount)
	log.Printf("[Stats] OriginalNotesCount: %d", stats.OriginalNotesCount)
	log.Printf("[Stats] OriginalUsersCount: %d", stats.OriginalUsersCount)
	text := ""
	openai.NumTokens(text)
	// ExampleService_Create(client, text)
}

func ExampleService_Create(client *misskey.Client, text string) {
	text = openai.Openai(text)
	response, err := client.Notes().Create(notes.CreateRequest{
		Text:       core.NewString(text),
		Visibility: models.VisibilityPublic,
	})
	if err != nil {
		log.Printf("[Notes] Error happened: %s", err)

		return
	}

	log.Println(response.CreatedNote.ID)
}
