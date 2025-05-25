package models

import (
	"time"

	"github.com/google/uuid"
)

type NoteModel struct {
	NoteID    string    `json:"noteId"` // Partition Key
	UserID    string    `json:"userId"` // GSI for user-wise filtering
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewNote(userId, title, content, status string) *NoteModel {
	return &NoteModel{
		UserID:    userId,
		NoteID:    uuid.New().String(),
		Title:     title,
		Content:   content,
		Status:    status,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}
