package types

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/yash-gkmit/NOTE-TAKER/models"
)

type CreateUserReqModel struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (model *CreateUserReqModel) Validate() error {
	validate := validator.New()
	err := validate.Struct(model)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return validationErrors
	}
	return nil
}

func (model *CreateUserReqModel) ToNewUser() (*models.UserModel, error) {
	if err := model.Validate(); err != nil {
		return nil, fmt.Errorf("put pass request model validation failed: %w", err)
	}

	user := models.NewUser(model.Email, model.Password)

	return user, nil
}

type UserResponseModel struct {
	UserId    string    `json:"userId"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
