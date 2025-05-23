package models

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	UserID    string    `json:"user_id"` // Partition Key
	Email     string    `json:"email"`   // GSI for querying by email
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewUser(email, username, password string) *UserModel {
	return &UserModel{
		UserID:    uuid.New().String(),
		Email:     email,
		Username:  username,
		Password:  password,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}
