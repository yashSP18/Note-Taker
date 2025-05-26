package routes

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/go-chi/chi/v5"
	"github.com/yash-gkmit/NOTE-TAKER/handlers"
	"github.com/yash-gkmit/NOTE-TAKER/repo"
)

func UserRoutes(ddb *dynamodb.DynamoDB) func(router chi.Router) {

	userRepo := repo.NewUserRepo(ddb)
	userHandler := handlers.NewUserHandler(userRepo)

	return func(r chi.Router) {
		r.Get("/me", userHandler.GetMe)
		r.Get("/", userHandler.GetAllUsers)
		r.Delete("/{id}", userHandler.DeleteUser)
	}
}
