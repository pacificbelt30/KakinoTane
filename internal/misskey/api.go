package misskey

import (
	"log"
	"os"
	"github.com/yitsushi/go-misskey"
	"github.com/yitsushi/go-misskey/core"
	"github.com/yitsushi/go-misskey/models"
	"github.com/yitsushi/go-misskey/services/notes"
	"github.com/sirupsen/logrus"
)

type Misskey struct {
	client *misskey.Client
}

func NewMisskeyClient() (*Misskey, error) {
	client, err := misskey.NewClientWithOptions(
		misskey.WithAPIToken(os.Getenv("MISSKEY_TOKEN")),
		misskey.WithBaseURL("https", os.Getenv("BASE_URL"), ""),
		misskey.WithLogLevel(logrus.DebugLevel),
	)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return &Misskey{client}, nil
}

func (mi *Misskey) GetStats() error {
	stats, err := mi.client.Meta().Stats()
	if err != nil {
		log.Printf("[Meta] Error happened: %s", err)
		return err
	}

	log.Printf("[Stats] Instances:          %d", stats.Instances)
	log.Printf("[Stats] NotesCount:         %d", stats.NotesCount)
	log.Printf("[Stats] UsersCount:         %d", stats.UsersCount)
	log.Printf("[Stats] OriginalNotesCount: %d", stats.OriginalNotesCount)
	log.Printf("[Stats] OriginalUsersCount: %d", stats.OriginalUsersCount)
	return nil
}

func (mi *Misskey) CreateNote(text string) error {
	response, err := mi.client.Notes().Create(notes.CreateRequest{
		Text:       core.NewString(text),
		Visibility: models.VisibilityPublic,
	})
	if err != nil {
		log.Printf("[Notes] Error happened: %s", err)

		return err
	}

	log.Println(response.CreatedNote.ID)
	return nil
}

