package usecases

import (
	"pinpong.co/rest-go/models"
	"pinpong.co/rest-go/repositories"
)

type noteUseCase struct {
	noteRepo repositories.NoteRepository
}

func NewNoteUseCase() *noteUseCase {
	noteRepo := repositories.NewNoteRepository()
	return &noteUseCase{*noteRepo}
}

func (t *noteUseCase) GetAll() (notes []models.Note, err error) {
	var note []models.Note
	note, err = t.noteRepo.GetAll()

	return note, err
}
