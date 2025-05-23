package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/yash-gkmit/NOTE-TAKER/helpers"
	"github.com/yash-gkmit/NOTE-TAKER/services"
	"github.com/yash-gkmit/NOTE-TAKER/types"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: service,
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var userReqModel types.CreateUserReqModel

	if err := json.NewDecoder(r.Body).Decode(&userReqModel); err != nil {
		helpers.SendHandlerErrResponse(w, "failed to decode: "+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := userReqModel.ToNewUser()
	if err != nil {
		helpers.SendHandlerErrResponse(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	customError := h.authService.Register(r.Context(), user)
	if customError != nil {
		helpers.SendHandlerCustomErrResponse(w, customError, customError.StatusCode)
		return
	}

	response := map[string]interface{}{
		"message": "User registered successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest types.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		helpers.SendHandlerErrResponse(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	err := loginRequest.Validate()
	if err != nil {
		helpers.SendHandlerErrResponse(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	token, customError := h.authService.Login(r.Context(), loginRequest.Email, loginRequest.Password)
	if customError != nil {
		helpers.SendHandlerCustomErrResponse(w, customError, customError.StatusCode)
		return
	}

	response := map[string]interface{}{
		"message": "Login successful",
		"token":   token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	message := h.authService.Logout()

	response := map[string]interface{}{
		"message": message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
