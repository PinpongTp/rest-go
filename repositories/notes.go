package repositories

import (
	"log"

	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/iterator"
	"pinpong.co/rest-go/models"
	fb "pinpong.co/rest-go/util"
)

type NoteRepository struct {
	conn fb.Firebase
}

func NewNoteRepository() *NoteRepository {
	conn := fb.NewFirebase()
	return &NoteRepository{*conn}
}

func (t *NoteRepository) GetAll() (note *[]models.Note, err error) {
	Datas := []models.Note{}

	item := t.conn.Client.Collection("notes").Documents(t.conn.Ctx)
	for {
		Data := models.Note{}
		doc, err := item.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		mapstructure.Decode(doc.Data(), &Data)
		Datas = append(Datas, Data)
	}
	return &Datas, nil
}
