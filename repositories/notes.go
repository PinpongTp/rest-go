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

// findAll
func (t *NoteRepository) FindAll() (note []models.Note, err error) {
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
	return Datas, nil
}

// findById
func (t *NoteRepository) FindByID(id string) (note models.Note, err error) {
	Datas := models.Note{}
	q := t.conn.Client.Collection("notes").Where("Id", "==", id)
	item := q.Documents(t.conn.Ctx)
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
		Datas = Data
	}
	return Datas, nil
}

// save
func (t *NoteRepository) Save(note models.Note) (Datas models.Note, err error) {
	if len(note.Id) == 0 {

		// create
		q := t.conn.Client.Collection("notes")
		_, _, err = q.Add(t.conn.Ctx, &note)
		if err != nil {
			log.Fatalln(err)
		}
		return note, nil
	} else {

		// update
		var refId string
		q := t.conn.Client.Collection("notes").Where("Id", "==", note.Id)
		item := q.Documents(t.conn.Ctx)
		for {
			doc, err := item.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Fatalln(err)
			}
			refId = doc.Ref.ID
		}
		_, err = t.conn.Client.Collection("notes").Doc(refId).Set(t.conn.Ctx, note)
		if err != nil {
			log.Fatalln(err)
		}
		return note, nil
	}
}

// delete
func (t *NoteRepository) DeleteById(id string) (err error) {
	var refId string
	q := t.conn.Client.Collection("notes").Where("Id", "==", id)
	item := q.Documents(t.conn.Ctx)
	for {
		doc, err := item.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		refId = doc.Ref.ID
	}
	_, err = t.conn.Client.Collection("notes").Doc(refId).Delete(t.conn.Ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}
