package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/yash-gkmit/NOTE-TAKER/config"
	"github.com/yash-gkmit/NOTE-TAKER/routes"
)

type App struct {
	router http.Handler
	ddb    *dynamodb.DynamoDB
	config *config.Config
}

func NewApp(configI *config.Config) *App {

	awsSession := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:   &configI.AwsConfig.Region,
			Endpoint: &configI.DynamoEndpoint,
			Credentials: credentials.NewStaticCredentials(
				configI.AwsConfig.AccessKey,
				configI.AwsConfig.SecretKey,
				"",
			),
		},
	}))

	app := &App{
		ddb:    dynamodb.New(awsSession),
		config: configI,
	}

	app.loadRoutes()

	return app
}

func (app *App) Start() error {

	// Test DynamoDB connection
	_, err := app.ddb.ListTables(&dynamodb.ListTablesInput{Limit: aws.Int64(1)})
	if err != nil {
		return fmt.Errorf("error connecting db: %w", err)
	}

	// Print before starting server
	fmt.Println(" Starting HTTP server on :3000")

	server := &http.Server{
		Addr:    ":3000",
		Handler: app.router,
	}

	// create table script
	// scripts.CreateDynamodbTables(app.ddb)

	// This blocks unless the server crashes
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println(" Error running server:", err)
		return fmt.Errorf("error running server: %w", err)
	}

	return nil
}

func (app *App) loadRoutes() {
	app.router = routes.NewRoutes(app.ddb)
}
