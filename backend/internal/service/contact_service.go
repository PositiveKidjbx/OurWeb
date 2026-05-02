package service

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"strings"
	"unicode/utf8"

	"backend/internal/repository"
)

type ContactInput struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Company string `json:"company"`
	Phone   string `json:"phone"`
	Message string `json:"message"`
}

const (
	// Application-level caps live here; transport-level limits and rate limiting belong at Nginx.
	maxNameLength    = 100
	maxEmailLength   = 254
	maxCompanyLength = 150
	maxPhoneLength   = 32
	maxMessageLength = 5000
)

type ValidationError struct {
	Msg string
}

func (e ValidationError) Error() string {
	return e.Msg
}

type ContactService struct {
	repo *repository.ContactRepository
}

func NewContactService(repo *repository.ContactRepository) *ContactService {
	return &ContactService{repo: repo}
}

func (s *ContactService) Submit(ctx context.Context, input ContactInput) (int64, error) {
	if s == nil || s.repo == nil {
		return 0, errors.New("contact service is not initialized")
	}

	if err := validateContactInput(input); err != nil {
		return 0, err
	}

	id, err := s.repo.Create(ctx, repository.ContactMessage{
		Name:    strings.TrimSpace(input.Name),
		Email:   strings.TrimSpace(input.Email),
		Company: strings.TrimSpace(input.Company),
		Phone:   strings.TrimSpace(input.Phone),
		Message: strings.TrimSpace(input.Message),
	})
	if err != nil {
		return 0, err
	}

	return id, nil
}

func validateContactInput(input ContactInput) error {
	name := strings.TrimSpace(input.Name)
	email := strings.TrimSpace(input.Email)
	company := strings.TrimSpace(input.Company)
	phone := strings.TrimSpace(input.Phone)
	message := strings.TrimSpace(input.Message)

	if name == "" {
		return ValidationError{Msg: "name is required"}
	}
	if runeLen(name) > maxNameLength {
		return ValidationError{Msg: fmt.Sprintf("name must be at most %d characters", maxNameLength)}
	}

	if email == "" {
		return ValidationError{Msg: "email is required"}
	}
	if runeLen(email) > maxEmailLength {
		return ValidationError{Msg: fmt.Sprintf("email must be at most %d characters", maxEmailLength)}
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return ValidationError{Msg: fmt.Sprintf("email is invalid: %v", err)}
	}

	if runeLen(company) > maxCompanyLength {
		return ValidationError{Msg: fmt.Sprintf("company must be at most %d characters", maxCompanyLength)}
	}

	if runeLen(phone) > maxPhoneLength {
		return ValidationError{Msg: fmt.Sprintf("phone must be at most %d characters", maxPhoneLength)}
	}

	if message == "" {
		return ValidationError{Msg: "message is required"}
	}
	if runeLen(message) > maxMessageLength {
		return ValidationError{Msg: fmt.Sprintf("message must be at most %d characters", maxMessageLength)}
	}

	return nil
}

func runeLen(value string) int {
	return utf8.RuneCountInString(value)
}
