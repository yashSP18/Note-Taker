package types

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/yash-gkmit/NOTE-TAKER/models"
)

type CreateNoteReqModel struct {
	Title   string `json:"title" validate:"required,min=3,max=255"`
	Content string `json:"content" validate:"max=5000"`
}

func (model *CreateNoteReqModel) Validate() error {
	validate := validator.New()

	err := validate.Struct(model)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return validationErrors
	}
	return nil
}

func (model *CreateNoteReqModel) ToNewNote(userId string) (*models.NoteModel, error) {
	if err := model.Validate(); err != nil {
		return nil, fmt.Errorf("create note request validation failed: %w", err)
	}

	note := models.NewNote(userId, model.Title, model.Content)
	return note, nil
}
