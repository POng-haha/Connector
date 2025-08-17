package handler

import (
	"fmt"
	"net/http"
	"time"

	"connectorapi-go/internal/core/domain"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func handleErrorResponse(c *gin.Context, appErr *domain.AppError) {
	statusCode := http.StatusInternalServerError
	switch appErr.Code {
	case domain.ErrValidation.Code:
		statusCode = http.StatusBadRequest
	case domain.ErrUnauthorized.Code:
		statusCode = http.StatusUnauthorized
	case domain.ErrForbidden.Code:
		statusCode = http.StatusForbidden
	case domain.ErrNotFound.Code:
		statusCode = http.StatusNotFound
	case domain.ErrServiceDown.Code:
		statusCode = http.StatusBadGateway
	}

	errResponse := domain.ErrorResponse{
		ErrorCode:    appErr.Code,
		ErrorMessage: appErr.Message,
		Status:       statusCode,
		Timestamp:    time.Now().UTC(),
		RequestID:    c.GetString(RequestID),
	}

	c.JSON(statusCode, errResponse)
}

func formatValidationErrors(err error) []domain.ValidationErrorDetail {
	var validationErrors []domain.ValidationErrorDetail

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range ve {
			validationErrors = append(validationErrors, domain.ValidationErrorDetail{
				Field:   fieldErr.Field(),
				Tag:     fieldErr.Tag(),
				Message: fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", fieldErr.Field(), fieldErr.Tag()),
			})
		}
	}
	return validationErrors
}
