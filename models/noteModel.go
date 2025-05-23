package models

type NoteModel struct {
	NoteID  string `json:"noteId"` // Partition Key
	UserID  string `json:"userId"` // GSI for user-wise filtering
	Title   string `json:"title"`
	Content string `json:"content"`
}

func NewNote(noteId, userId, title, content string) *NoteModel {
	return &NoteModel{
		UserID:  userId,
		NoteID:  noteId,
		Title:   title,
		Content: content,
	}
}
