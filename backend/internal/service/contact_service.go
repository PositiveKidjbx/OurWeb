package service

import (
	"context"
	"errors"
	"fmt"
	"net/mail"
	"strings"

	"backend/internal/repository"
)

type ContactInput struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Company string `json:"company"`
	Phone   string `json:"phone"`
	Message string `json:"message"`
}

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
	if strings.TrimSpace(input.Name) == "" {
		return ValidationError{Msg: "name is required"}
	}

	email := strings.TrimSpace(input.Email)
	if email == "" {
		return ValidationError{Msg: "email is required"}
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return ValidationError{Msg: fmt.Sprintf("email is invalid: %v", err)}
	}

	if strings.TrimSpace(input.Message) == "" {
		return ValidationError{Msg: "message is required"}
	}

	return nil
}
