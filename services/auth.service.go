package services

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/yash-gkmit/NOTE-TAKER/helpers"
	"github.com/yash-gkmit/NOTE-TAKER/models"
	"github.com/yash-gkmit/NOTE-TAKER/repo"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repo.UserRepo
}

func NewAuthService(repo *repo.UserRepo) *AuthService {
	return &AuthService{
		userRepo: repo,
	}
}

func (s *AuthService) Register(ctx context.Context, user *models.UserModel) *helpers.CustomError {

	// email should be unique for the user
	userExist, err := s.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		log.Printf("failed to query user by email: %v", err)
		return helpers.NewCustomError(fmt.Errorf("internal error while checking user existence: %v", err), 500)
	}

	if userExist != nil {
		return helpers.NewCustomError(fmt.Errorf("user already exists"), 409)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("failed to hash password: %v", err)
		return helpers.NewCustomError(fmt.Errorf("internal error while hashing password: %v", err), 500)
	}

	user.Password = string(hashedPassword)

	if err := s.userRepo.CreateUser(user); err != nil {
		log.Printf("failed to create user in repository: %v", err)
		return helpers.NewCustomError(fmt.Errorf("error creating user: %v", err), 500)
	}

	return nil
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, *helpers.CustomError) {

	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return "", helpers.NewCustomError(errors.New("Invalid Credentials!"), 401)
	}

	if user == nil {
		return "", helpers.NewCustomError(errors.New("Invalid Credentials!"), 401)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", helpers.NewCustomError(errors.New("Invalid Credentials!"), 401)
	}

	token, err := helpers.GenerateJWT(user.UserID, user.Email)
	if err != nil {
		return "", helpers.NewCustomError(errors.New("failed to generate token"), 500)
	}

	return token, nil
}

func (s *AuthService) Logout() string {
	return "Logged out successfully!"
}
