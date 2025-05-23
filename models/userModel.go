package models

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	UserID    string    `json:"userId"` // Partition Key
	Email     string    `json:"email"`  // GSI for querying by email
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewUser(email, password string) *UserModel {
	return &UserModel{
		UserID:    uuid.New().String(),
		Email:     email,
		Password:  password,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}
