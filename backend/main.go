package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/service"
	"backend/internal/store"
)

var databasePath = envOrDefault("OW_DATABASE_PATH", "data/contact_messages.sqlite3")
var serverAddr = envOrDefault("PORT", ":8080")

func main() {
	db, err := store.InitSQLite(databasePath)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer db.Close()

	contactRepository := repository.NewContactRepository(db)
	contactService := service.NewContactService(contactRepository)
	contactHandler := handler.NewContactHandler(contactService)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello Gin")
	})
	r.POST("/api/contact", contactHandler.Create)

	if err := r.Run(serverAddr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func envOrDefault(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
