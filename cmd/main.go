package main

import (
	"log"

	"github.com/Marketen/lido-index/internal/app"
)

func main() {
    // Call the initialization function from app.go
    // This starts all services and the HTTP server
    app.InitializeApp()

    // Optionally, include logging to indicate the application has started successfully
    log.Println("Application started successfully")
}
