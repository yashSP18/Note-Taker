package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/yash-gkmit/NOTE-TAKER/helpers"
	"github.com/yash-gkmit/NOTE-TAKER/repo"
	"github.com/yash-gkmit/NOTE-TAKER/services"
	"github.com/yash-gkmit/NOTE-TAKER/types"
)

type NoteHandler struct {
	noteService *services.NoteService
}

func NewNoteHandler(service *services.NoteService) *NoteHandler {
	return &NoteHandler{
		noteService: service,
	}
}

func (h *NoteHandler) CreateNote(w http.ResponseWriter, r *http.Request) {
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

	var noteReq types.CreateNoteReqModel
	if err := json.NewDecoder(r.Body).Decode(&noteReq); err != nil {
		helpers.SendHandlerErrResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	note, err := noteReq.ToNewNote(userID)
	if err != nil {
		helpers.SendHandlerErrResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.noteService.CreateNote(note)
	if err != nil {
		helpers.SendHandlerErrResponse(w, "Failed to create note", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Note created successfully", "noteId": note.NoteID})
}

func (h *NoteHandler) GetAllNote(w http.ResponseWriter, r *http.Request) {

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

	userId, ok := claims["userId"].(string)
	if !ok {
		helpers.SendHandlerErrResponse(w, "Invalid user ID in token", http.StatusUnauthorized)
		return
	}

	notes, err := h.noteService.GetAllNote(r.Context(), userId)
	if err != nil {
		helpers.SendHandlerErrResponse(w, "Failed to fetch notes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notes)
}

func (h *NoteHandler) GetNote(w http.ResponseWriter, r *http.Request) {

	noteId := chi.URLParam(r, "id")

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

	userId, ok := claims["userId"].(string)
	if !ok {
		helpers.SendHandlerErrResponse(w, "Invalid user ID in token", http.StatusUnauthorized)
		return
	}

	note, err := h.noteService.GetNote(r.Context(), userId, noteId)
	if err != nil {
		if errors.Is(err, repo.NoteNotFound) {
			customErr := helpers.NewCustomError(errors.New("note not found"), 404)
			helpers.SendHandlerCustomErrResponse(w, customErr, customErr.StatusCode)
		} else {
			customErr := helpers.NewCustomError(errors.New("failed to fetch note"), 400)
			helpers.SendHandlerCustomErrResponse(w, customErr, customErr.StatusCode)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(note)
}
