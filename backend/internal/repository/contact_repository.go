package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type ContactMessage struct {
	ID        int64
	Name      string
	Email     string
	Company   string
	Phone     string
	Message   string
	CreatedAt string
}

type ContactRepository struct {
	db *sql.DB
}

func NewContactRepository(db *sql.DB) *ContactRepository {
	return &ContactRepository{db: db}
}

func (r *ContactRepository) Create(ctx context.Context, message ContactMessage) (int64, error) {
	if r == nil || r.db == nil {
		return 0, errors.New("contact repository is not initialized")
	}

	result, err := r.db.ExecContext(
		ctx,
		`INSERT INTO contact_messages (name, email, company, phone, message) VALUES (?, ?, ?, ?, ?)`,
		strings.TrimSpace(message.Name),
		strings.TrimSpace(message.Email),
		strings.TrimSpace(message.Company),
		strings.TrimSpace(message.Phone),
		strings.TrimSpace(message.Message),
	)
	if err != nil {
		return 0, fmt.Errorf("insert contact message: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("resolve inserted id: %w", err)
	}

	return id, nil
}
