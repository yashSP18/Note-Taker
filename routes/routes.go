package routes

import (
	"net/http"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func NewRoutes(ddb *dynamodb.DynamoDB) http.Handler {
	router := chi.NewRouter()

	cors := cors.AllowAll()

	router.Use(cors.Handler)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`Hello World`))
	})

	return router
}
