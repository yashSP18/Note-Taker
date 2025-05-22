package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Main started")
	// load envs
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	app := NewApp()

	fmt.Println("App created") // <-- add this

	err = app.Start()
	if err != nil {
		fmt.Println("App failed to start:", err)
	}
}
