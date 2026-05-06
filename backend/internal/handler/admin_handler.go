package handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"backend/internal/service"
)

const adminSessionCookieName = "ow_admin_session"

type AdminConfig struct {
	Username      string
	Password      string
	SessionSecret string
	CookieSecure  bool
}

type AdminHandler struct {
	service *service.ContactService
	config  AdminConfig
	now     func() time.Time
}

type adminLoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewAdminHandler(service *service.ContactService, config AdminConfig) *AdminHandler {
	return &AdminHandler{
		service: service,
		config:  config,
		now:     time.Now,
	}
}

func (h *AdminHandler) Login(c *gin.Context) {
	var input adminLoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json payload",
		})
		return
	}

	if !h.hasCredentials() {
		log.Print("admin login attempted before admin credentials were configured")
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error": "admin login is not configured",
		})
		return
	}

	if !h.credentialsMatch(input.Username, input.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid username or password",
		})
		return
	}

	expiresAt := h.now().Add(24 * time.Hour)
	token := h.signSession(input.Username, expiresAt)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     adminSessionCookieName,
		Value:    token,
		Path:     "/",
		Expires:  expiresAt,
		MaxAge:   int(expiresAt.Sub(h.now()).Seconds()),
		HttpOnly: true,
		Secure:   h.config.CookieSecure,
		SameSite: http.SameSiteLaxMode,
	})

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
	})
}

func (h *AdminHandler) ListMessages(c *gin.Context) {
	page := parsePositiveQueryInt(c, "page", 1)
	pageSize := parsePositiveQueryInt(c, "page_size", 20)

	result, err := h.service.ListMessages(c.Request.Context(), page, pageSize)
	if err != nil {
		log.Printf("admin list messages failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to list contact messages",
		})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *AdminHandler) RequireAdmin(c *gin.Context) {
	token, err := c.Cookie(adminSessionCookieName)
	if err != nil || !h.verifySession(token) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "authentication required",
		})
		c.Abort()
		return
	}

	c.Next()
}

func (h *AdminHandler) hasCredentials() bool {
	return h.config.Username != "" && h.config.Password != "" && h.config.SessionSecret != ""
}

func (h *AdminHandler) credentialsMatch(username string, password string) bool {
	usernameMatches := subtle.ConstantTimeCompare([]byte(username), []byte(h.config.Username)) == 1
	passwordMatches := subtle.ConstantTimeCompare([]byte(password), []byte(h.config.Password)) == 1

	return usernameMatches && passwordMatches
}

func (h *AdminHandler) signSession(username string, expiresAt time.Time) string {
	payload := fmt.Sprintf("%s|%d", username, expiresAt.Unix())
	encodedPayload := base64.RawURLEncoding.EncodeToString([]byte(payload))
	signature := h.sign(encodedPayload)

	return encodedPayload + "." + signature
}

func (h *AdminHandler) verifySession(token string) bool {
	if !h.hasCredentials() {
		return false
	}

	parts := strings.Split(token, ".")
	if len(parts) != 2 {
		return false
	}

	expectedSignature := h.sign(parts[0])
	if subtle.ConstantTimeCompare([]byte(parts[1]), []byte(expectedSignature)) != 1 {
		return false
	}

	payloadBytes, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return false
	}

	payloadParts := strings.Split(string(payloadBytes), "|")
	if len(payloadParts) != 2 || payloadParts[0] != h.config.Username {
		return false
	}

	expiresUnix, err := strconv.ParseInt(payloadParts[1], 10, 64)
	if err != nil {
		return false
	}

	return h.now().Unix() < expiresUnix
}

func (h *AdminHandler) sign(value string) string {
	mac := hmac.New(sha256.New, []byte(h.config.SessionSecret))
	mac.Write([]byte(value))

	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}

func parsePositiveQueryInt(c *gin.Context, key string, fallback int) int {
	value, err := strconv.Atoi(c.Query(key))
	if err != nil || value < 1 {
		return fallback
	}

	return value
}
