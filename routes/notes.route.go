package routes

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/go-chi/chi/v5"
	"github.com/yash-gkmit/NOTE-TAKER/handlers"
	"github.com/yash-gkmit/NOTE-TAKER/middlewares"
	"github.com/yash-gkmit/NOTE-TAKER/repo"
	"github.com/yash-gkmit/NOTE-TAKER/services"
)

func NoteRoutes(ddb *dynamodb.DynamoDB) func(router chi.Router) {

	noteRepo := repo.NewNoteRepo(ddb)
	noteService := services.NewNoteService(noteRepo)
	noteHandler := handlers.NewNoteHandler(noteService)

	return func(r chi.Router) {
		r.Use(middlewares.AuthMiddleware)
		r.Post("/", noteHandler.CreateNote)
		r.Get("/", noteHandler.GetAllNote)
		r.Get("/{id}", noteHandler.GetNote)
		r.Patch("/{id}", noteHandler.UpdateNote)
	}
}
