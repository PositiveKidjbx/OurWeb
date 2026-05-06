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
	Status    string
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

func (r *ContactRepository) List(ctx context.Context, limit int, offset int) ([]ContactMessage, int, error) {
	if r == nil || r.db == nil {
		return nil, 0, errors.New("contact repository is not initialized")
	}

	var total int
	if err := r.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM contact_messages`).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("count contact messages: %w", err)
	}

	rows, err := r.db.QueryContext(
		ctx,
		`SELECT id, name, email, COALESCE(company, ''), COALESCE(phone, ''), message, status, created_at
FROM contact_messages
ORDER BY created_at DESC, id DESC
LIMIT ? OFFSET ?`,
		limit,
		offset,
	)
	if err != nil {
		return nil, 0, fmt.Errorf("list contact messages: %w", err)
	}
	defer rows.Close()

	messages := make([]ContactMessage, 0)
	for rows.Next() {
		var message ContactMessage
		if err := rows.Scan(
			&message.ID,
			&message.Name,
			&message.Email,
			&message.Company,
			&message.Phone,
			&message.Message,
			&message.Status,
			&message.CreatedAt,
		); err != nil {
			return nil, 0, fmt.Errorf("scan contact message: %w", err)
		}

		messages = append(messages, message)
	}

	if err := rows.Err(); err != nil {
		return nil, 0, fmt.Errorf("iterate contact messages: %w", err)
	}

	return messages, total, nil
}
