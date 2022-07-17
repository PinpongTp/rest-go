package usecases

import (
	"pinpong.co/rest-go/models"
	"pinpong.co/rest-go/repositories"
)

type NoteUseCase struct {
	noteRepo repositories.NoteRepository
}

func NewNoteUseCase(noteRepo repositories.NoteRepository) *NoteUseCase {
	return &NoteUseCase{noteRepo}
}

func (t *NoteUseCase) FindAll() (notes []models.Note, err error) {
	var note []models.Note
	note, err = t.noteRepo.FindAll()
	//
	return note, err
}

func (t *NoteUseCase) FindById(id string) (notes models.Note, err error) {
	var note models.Note
	note, err = t.noteRepo.FindByID(id)
	//
	return note, err
}

func (t *NoteUseCase) Create(note models.Note) (notes models.Note, err error) {
	note, err = t.noteRepo.Save(note)
	//
	return note, err
}

func (t *NoteUseCase) Update(note models.Note) (notes models.Note, err error) {
	note, err = t.noteRepo.Save(note)
	//
	return note, err
}

func (t *NoteUseCase) Delete(id string) (err error) {
	err = t.noteRepo.DeleteById(id)
	//
	return err
}
