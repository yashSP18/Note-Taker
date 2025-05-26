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
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewNote(userId, title, content string) *NoteModel {
	return &NoteModel{
		UserID:    userId,
		NoteID:    uuid.New().String(),
		Title:     title,
		Content:   content,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}
