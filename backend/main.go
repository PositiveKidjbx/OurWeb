package main

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var databasePath = envOrDefault("OW_DATABASE_PATH", "data/contact_messages.sqlite3")

func main() {
	db, err := initDatabase(databasePath)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer db.Close()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello Gin")
	})

	r.Run()
}

func initDatabase(path string) (*sql.DB, error) {
	_, statErr := os.Stat(path)
	shouldInitialize := errors.Is(statErr, os.ErrNotExist)
	if statErr != nil && !shouldInitialize {
		return nil, statErr
	}

	if shouldInitialize {
		dir := filepath.Dir(path)
		if dir != "." {
			if err := os.MkdirAll(dir, 0755); err != nil {
				return nil, err
			}
		}
	}

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	if shouldInitialize {
		if err := createSchema(db); err != nil {
			db.Close()
			return nil, err
		}
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func createSchema(db *sql.DB) error {
	const schema = `
CREATE TABLE contact_messages (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	email TEXT NOT NULL,
	company TEXT,
	phone TEXT,
	message TEXT NOT NULL,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);`

	_, err := db.Exec(schema)
	return err
}

func envOrDefault(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}
