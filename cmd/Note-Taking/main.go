package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/yash-gkmit/NOTE-TAKER/config"
)

func main() {
	fmt.Println("Main started")
	// load envs
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	config.Init()
	serverConfig := config.GetInstance()

	app := NewApp(serverConfig)

	// Create tables before starting the app
	// db.CreateDynamodbTables(serverConfig)

	fmt.Println("App created") // <-- add this

	err = app.Start()
	if err != nil {
		fmt.Println("App failed to start:", err)
	}
}
