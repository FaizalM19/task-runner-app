package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"task-runner/internal/auth"
	"task-runner/internal/storage"
	"task-runner/internal/task"
	"task-runner/pkg/config"
	"task-runner/pkg/server"
)

func init() {
	// Manually load environment variables from .env file
	err := config.LoadEnv(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	dbPath := "db/database.db"
	port := os.Getenv("port")

	if dbPath == "" {
		log.Fatal("DB_PATH not set in .env file")
	}
	db, err := storage.InitDB(dbPath)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	defer db.Close()

	// Initialize services
	authService := auth.NewAuthService()
	taskService := task.NewTaskService(storage.NewStorage(db))

	server := server.NewServer(authService, taskService)

	// Start the server on the configured port
	fmt.Printf("Starting server on :%s...\n", port)
	if err := http.ListenAndServe(":"+port, server); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
