package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/internal/service"
)

type ContactHandler struct {
	service *service.ContactService
}

func NewContactHandler(service *service.ContactService) *ContactHandler {
	return &ContactHandler{service: service}
}

func (h *ContactHandler) Create(c *gin.Context) {
	var input service.ContactInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json payload",
		})
		return
	}

	_, err := h.service.Submit(c.Request.Context(), input)
	if err != nil {
		if validationErr, ok := errors.AsType[service.ValidationError](err); ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": validationErr.Error(),
			})
			return
		}

		log.Printf("contact submission failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to save contact message",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "contact message saved",
	})
}
