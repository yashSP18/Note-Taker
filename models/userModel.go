package models

type UserModel struct {
	UserID   string `json:"user_id"` // Partition Key
	Email    string `json:"email"`   // GSI for querying by email
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUser(userId, email, username, password string) *UserModel {
	return &UserModel{
		UserID:   userId,
		Email:    email,
		Username: username,
		Password: password,
	}
}
