package main

import (
	"github.com/lucaspichi06/strings-reverter/src/api/app"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := app.Start()
	if err := router.Run(":" + port); err != nil {
		log.Println("[main][message: error running server]", err)
	}
}

