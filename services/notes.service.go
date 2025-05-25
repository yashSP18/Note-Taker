package services

import (
	"context"

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

func (s *NoteService) GetAllNote(ctx context.Context, userId string) ([]*models.NoteModel, error) {
	notes, err := s.noteRepo.GetAllNote(userId)
	if err != nil {
		return nil, err
	}
	return notes, nil
}

func (s *NoteService) GetNote(ctx context.Context, userId, noteId string) (*models.NoteModel, error) {
	note, err := s.noteRepo.GetNoteById(userId, noteId)
	if err != nil {
		return nil, err
	}
	return note, nil
}
