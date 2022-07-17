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

func (t *NoteUseCase) GetAll() (notes []models.Note, err error) {
	var note []models.Note
	note, err = t.noteRepo.GetAll()

	return note, err
}
