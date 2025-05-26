package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/go-chi/chi/v5"
	"github.com/yash-gkmit/NOTE-TAKER/helpers"
	"github.com/yash-gkmit/NOTE-TAKER/repo"
	"github.com/yash-gkmit/NOTE-TAKER/types"
)

type UserHandler struct {
	userRepo *repo.UserRepo
}

func NewUserHandler(repo *repo.UserRepo) *UserHandler {
	return &UserHandler{
		userRepo: repo,
	}
}

func (h *UserHandler) GetMe(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		helpers.SendHandlerErrResponse(w, "Authorization header missing", http.StatusUnauthorized)
		return
	}

	tokenParts := strings.Split(authHeader, "Bearer ")
	if len(tokenParts) != 2 {
		helpers.SendHandlerErrResponse(w, "Invalid token format", http.StatusUnauthorized)
		return
	}

	claims, err := helpers.VerifyJWT(tokenParts[1])
	if err != nil {
		helpers.SendHandlerErrResponse(w, "Invalid or expired token", http.StatusUnauthorized)
		return
	}

	userID, ok := claims["userId"].(string)
	if !ok {
		helpers.SendHandlerErrResponse(w, "Invalid user ID in token", http.StatusUnauthorized)
		return
	}

	user, err := h.userRepo.GetUserById(userID)
	if err != nil {
		helpers.SendHandlerErrResponse(w, "User not found", http.StatusNotFound)
		return
	}

	response := types.UserResponseModel{
		UserId: user.UserID,
		Email:  user.Email,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// GetAllUsers fetches all registered users and returns basic user details.
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userRepo.GetAllUsers()
	if err != nil {
		helpers.SendHandlerErrResponse(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	// Create a response slice excluding sensitive data
	var response []types.UserResponseModel
	for _, user := range users {
		response = append(response, types.UserResponseModel{
			UserId: user.UserID,
			Email:  user.Email,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	if userId == "" {
		helpers.SendHandlerErrResponse(w, "Missing userId in the request path", http.StatusBadRequest)
		return
	}

	err := h.userRepo.DeleteUser(userId)
	if err != nil {
		if strings.Contains(err.Error(), dynamodb.ErrCodeConditionalCheckFailedException) {
			helpers.SendHandlerErrResponse(w, "User not found", http.StatusNotFound)
		} else {
			helpers.SendHandlerErrResponse(w, "Failed to delete user", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User deleted successfully",
	})
}
