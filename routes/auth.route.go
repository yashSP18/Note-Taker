package routes

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/go-chi/chi/v5"
	"github.com/yash-gkmit/NOTE-TAKER/handlers"
	"github.com/yash-gkmit/NOTE-TAKER/middlewares"
	"github.com/yash-gkmit/NOTE-TAKER/repo"
	"github.com/yash-gkmit/NOTE-TAKER/services"
)

func AuthRoutes(ddb *dynamodb.DynamoDB) func(router chi.Router) {

	userRepo := repo.NewUserRepo(ddb)
	userService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(userService)

	return func(r chi.Router) {
		r.Post("/register", authHandler.Register)
		r.Post("/login", authHandler.Login)
		r.With(middlewares.AuthMiddleware).Delete("/logout", authHandler.Logout)
	}
}
