package services

import (
	"context"
	"fmt"
	"time"

	"github.com/yash-gkmit/NOTE-TAKER/models"
	"github.com/yash-gkmit/NOTE-TAKER/repo"
	"github.com/yash-gkmit/NOTE-TAKER/types"
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

func (s *NoteService) UpdateNote(ctx context.Context, userId, noteId string, NoteRequestModel *types.UpdateNoteReqModel) (*models.NoteModel, error) {
	existingNote, err := s.noteRepo.GetNoteById(userId, noteId)
	if err != nil {
		return nil, err
	}

	if NoteRequestModel.Title != nil {
		existingNote.Title = *NoteRequestModel.Title
	}
	if NoteRequestModel.Content != nil {
		existingNote.Content = *NoteRequestModel.Content
	}
	if NoteRequestModel.Status != nil {
		existingNote.Status = *NoteRequestModel.Status
	}
	existingNote.UpdatedAt = time.Now()

	err = s.noteRepo.UpdateNote(existingNote)
	if err != nil {
		return nil, fmt.Errorf("failed to update Note: %w", err)
	}

	return existingNote, nil
}
