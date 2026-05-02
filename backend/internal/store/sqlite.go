package store

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func InitSQLite(path string) (*sql.DB, error) {
	if path == "" {
		return nil, errors.New("database path is required")
	}

	if err := ensureDirectory(path); err != nil {
		return nil, err
	}

	_, statErr := os.Stat(path)
	shouldInitialize := errors.Is(statErr, os.ErrNotExist)
	if statErr != nil && !shouldInitialize {
		return nil, statErr
	}

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	if err := ensureSchema(db); err != nil {
		_ = db.Close()
		return nil, err
	}

	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, err
	}

	return db, nil
}

func ensureDirectory(path string) error {
	dir := filepath.Dir(path)
	if dir == "." {
		return nil
	}

	return os.MkdirAll(dir, 0755)
}

func ensureSchema(db *sql.DB) error {
	const schema = `
CREATE TABLE IF NOT EXISTS contact_messages (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	email TEXT NOT NULL,
	company TEXT,
	phone TEXT,
	message TEXT NOT NULL,
	created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);`

	if _, err := db.Exec(schema); err != nil {
		return fmt.Errorf("create contact_messages table: %w", err)
	}

	return nil
}
