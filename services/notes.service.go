package services

import (
	"github.com/yash-gkmit/NOTE-TAKER/models"
	"github.com/yash-gkmit/NOTE-TAKER/repo"
)

type NoteService struct {
	noteRepo *repo.NoteRepo
}

func NewNoteService(noteRepo *repo.NoteRepo) *NoteService {
	return &NoteService{
		noteRepo: noteRepo,
	}
}

func (s *NoteService) CreateNote(note *models.NoteModel) error {
	return s.noteRepo.CreateNote(note)
}
