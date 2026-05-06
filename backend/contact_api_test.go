package main

import (
	"database/sql"
	"encoding/json"
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

	payload := `{"name":"Alice","email":"alice@example.com","company":"PROJECT EYEWEAR","phone":"123456789","message":"Hello from contact form"}`
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

	if name != "Alice" || email != "alice@example.com" || company != "PROJECT EYEWEAR" || phone != "123456789" || message != "Hello from contact form" {
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

func TestAdminLoginSetsSessionCookie(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.POST("/api/admin/login", handler.NewAdminHandler(nil, handler.AdminConfig{
		Username:      "admin",
		Password:      "secret",
		SessionSecret: "test-session-secret",
	}).Login)

	payload := `{"username":"admin","password":"secret"}`
	req := httptest.NewRequest(http.MethodPost, "/api/admin/login", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("unexpected status code: got %d want %d body=%s", rec.Code, http.StatusOK, rec.Body.String())
	}

	cookies := rec.Result().Cookies()
	if len(cookies) != 1 {
		t.Fatalf("unexpected cookie count: got %d want 1", len(cookies))
	}

	cookie := cookies[0]
	if cookie.Name != "ow_admin_session" || cookie.Value == "" {
		t.Fatalf("unexpected session cookie: name=%q value=%q", cookie.Name, cookie.Value)
	}

	if !cookie.HttpOnly {
		t.Fatal("session cookie should be HttpOnly")
	}
}

func TestAdminLoginRejectsInvalidCredentials(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	router.POST("/api/admin/login", handler.NewAdminHandler(nil, handler.AdminConfig{
		Username:      "admin",
		Password:      "secret",
		SessionSecret: "test-session-secret",
	}).Login)

	payload := `{"username":"admin","password":"wrong"}`
	req := httptest.NewRequest(http.MethodPost, "/api/admin/login", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("unexpected status code: got %d want %d body=%s", rec.Code, http.StatusUnauthorized, rec.Body.String())
	}
}

func TestAdminMessagesRequiresAuthentication(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db := initTestDB(t)
	adminHandler := handler.NewAdminHandler(
		service.NewContactService(repository.NewContactRepository(db)),
		handler.AdminConfig{
			Username:      "admin",
			Password:      "secret",
			SessionSecret: "test-session-secret",
		},
	)

	router := gin.New()
	group := router.Group("/api/admin")
	group.Use(adminHandler.RequireAdmin)
	group.GET("/messages", adminHandler.ListMessages)

	req := httptest.NewRequest(http.MethodGet, "/api/admin/messages", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("unexpected status code: got %d want %d body=%s", rec.Code, http.StatusUnauthorized, rec.Body.String())
	}
}

func TestAdminMessagesReturnsPaginatedMessages(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db := initTestDB(t)
	insertContactMessage(t, db, "Alice")
	insertContactMessage(t, db, "Bob")
	insertContactMessage(t, db, "Carol")

	adminHandler := handler.NewAdminHandler(
		service.NewContactService(repository.NewContactRepository(db)),
		handler.AdminConfig{
			Username:      "admin",
			Password:      "secret",
			SessionSecret: "test-session-secret",
		},
	)

	router := gin.New()
	router.POST("/api/admin/login", adminHandler.Login)
	group := router.Group("/api/admin")
	group.Use(adminHandler.RequireAdmin)
	group.GET("/messages", adminHandler.ListMessages)

	loginReq := httptest.NewRequest(http.MethodPost, "/api/admin/login", strings.NewReader(`{"username":"admin","password":"secret"}`))
	loginReq.Header.Set("Content-Type", "application/json")
	loginRec := httptest.NewRecorder()
	router.ServeHTTP(loginRec, loginReq)
	if loginRec.Code != http.StatusOK {
		t.Fatalf("login failed: status=%d body=%s", loginRec.Code, loginRec.Body.String())
	}

	req := httptest.NewRequest(http.MethodGet, "/api/admin/messages?page=2&page_size=2", nil)
	req.AddCookie(loginRec.Result().Cookies()[0])
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("unexpected status code: got %d want %d body=%s", rec.Code, http.StatusOK, rec.Body.String())
	}

	var response struct {
		Messages []struct {
			Name   string `json:"name"`
			Status string `json:"status"`
		} `json:"messages"`
		Page       int `json:"page"`
		PageSize   int `json:"page_size"`
		Total      int `json:"total"`
		TotalPages int `json:"total_pages"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if response.Page != 2 || response.PageSize != 2 || response.Total != 3 || response.TotalPages != 2 {
		t.Fatalf("unexpected pagination: page=%d page_size=%d total=%d total_pages=%d", response.Page, response.PageSize, response.Total, response.TotalPages)
	}

	if len(response.Messages) != 1 {
		t.Fatalf("unexpected message count: got %d want 1", len(response.Messages))
	}

	if response.Messages[0].Name != "Alice" || response.Messages[0].Status != "new" {
		t.Fatalf("unexpected message: name=%q status=%q", response.Messages[0].Name, response.Messages[0].Status)
	}
}

func initTestDB(t *testing.T) *sql.DB {
	t.Helper()

	tempDir := t.TempDir()
	dbPath := tempDir + "/contact.sqlite3"

	db, err := store.InitSQLite(dbPath)
	if err != nil {
		t.Fatalf("init sqlite: %v", err)
	}
	t.Cleanup(func() {
		_ = db.Close()
	})

	return db
}

func insertContactMessage(t *testing.T, db *sql.DB, name string) {
	t.Helper()

	_, err := db.Exec(
		`INSERT INTO contact_messages (name, email, company, phone, message) VALUES (?, ?, ?, ?, ?)`,
		name,
		strings.ToLower(name)+"@example.com",
		"PROJECT EYEWEAR",
		"123456789",
		"Hello from "+name,
	)
	if err != nil {
		t.Fatalf("insert contact message: %v", err)
	}
}
