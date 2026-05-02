package main

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/service"
	"backend/internal/store"
)

func TestCreateContactMessage(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tempDir := t.TempDir()
	dbPath := tempDir + "/contact.sqlite3"

	db, err := store.InitSQLite(dbPath)
	if err != nil {
		t.Fatalf("init sqlite: %v", err)
	}
	t.Cleanup(func() {
		_ = db.Close()
	})

	router := gin.New()
	router.POST("/api/contact", handler.NewContactHandler(
		service.NewContactService(
			repository.NewContactRepository(db),
		),
	).Create)

	payload := `{"name":"Alice","email":"alice@example.com","company":"OKIA","phone":"123456789","message":"Hello from contact form"}`
	req := httptest.NewRequest(http.MethodPost, "/api/contact", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("unexpected status code: got %d want %d body=%s", rec.Code, http.StatusCreated, rec.Body.String())
	}

	var count int
	if err := db.QueryRow(`SELECT COUNT(*) FROM contact_messages`).Scan(&count); err != nil {
		t.Fatalf("count rows: %v", err)
	}

	if count != 1 {
		t.Fatalf("unexpected row count: got %d want %d", count, 1)
	}

	var name, email, company, phone, message string
	err = db.QueryRow(`
SELECT name, email, COALESCE(company, ''), COALESCE(phone, ''), message
FROM contact_messages
LIMIT 1
`).Scan(&name, &email, &company, &phone, &message)
	if err != nil {
		t.Fatalf("read inserted row: %v", err)
	}

	if name != "Alice" || email != "alice@example.com" || company != "OKIA" || phone != "123456789" || message != "Hello from contact form" {
		t.Fatalf("unexpected row content: name=%q email=%q company=%q phone=%q message=%q", name, email, company, phone, message)
	}
}

func TestCreateContactMessageRejectsInvalidPayload(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tempDir := t.TempDir()
	dbPath := tempDir + "/contact.sqlite3"

	db, err := store.InitSQLite(dbPath)
	if err != nil {
		t.Fatalf("init sqlite: %v", err)
	}
	t.Cleanup(func() {
		_ = db.Close()
	})

	router := gin.New()
	router.POST("/api/contact", handler.NewContactHandler(
		service.NewContactService(
			repository.NewContactRepository(db),
		),
	).Create)

	payload := `{"name":"","email":"bad","message":""}`
	req := httptest.NewRequest(http.MethodPost, "/api/contact", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("unexpected status code: got %d want %d body=%s", rec.Code, http.StatusBadRequest, rec.Body.String())
	}

	var count int
	if err := db.QueryRow(`SELECT COUNT(*) FROM contact_messages`).Scan(&count); err != nil && err != sql.ErrNoRows {
		t.Fatalf("count rows: %v", err)
	}

	if count != 0 {
		t.Fatalf("unexpected row count after invalid payload: got %d want %d", count, 0)
	}
}
