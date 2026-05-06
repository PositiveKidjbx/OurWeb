package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/service"
	"backend/internal/store"
)

var databasePath = envOrDefault("OW_DATABASE_PATH", "data/contact_messages.sqlite3")
var serverAddr = envOrDefault("PORT", ":8080")
var frontendDir = envOrDefault("OW_FRONTEND_DIR", "..")

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
	r.LoadHTMLGlob(filepath.Join(frontendDir, "templates", "*"))
	r.Static("/static", filepath.Join(frontendDir, "public"))

	api := r.Group("/api")
	api.POST("/contact", contactHandler.Create)

	registerPageRoute(r, "/", "index.html", "Home")
	registerPageRoute(r, "/contact", "contact.html", "Home")
	registerPageRoute(r, "/sustainability", "sustainability.html", "Sustainability")
	registerPageRoute(r, "/awards", "awards.html", "Our Awards")
	registerPageRoute(r, "/career", "career.html", "Career")
	registerPageRoute(r, "/terms", "terms.html", "Terms")
	registerPageRoute(r, "/privacy", "privacy.html", "Privacy")
	registerPageRoute(r, "/news", "news.html", "News")

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"title": "Page Not Found",
		})
	})

	if err := r.Run(serverAddr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func registerPageRoute(r *gin.Engine, path string, templateName string, title string) {
	r.GET(path, func(c *gin.Context) {
		c.HTML(http.StatusOK, templateName, gin.H{
			"title": title,
		})
	})
}

func envOrDefault(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
