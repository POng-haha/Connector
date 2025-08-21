package handler

import (
	"fmt"
	"net/http"
	"strings"

	app_error "connectorapi-go/pkg/error"
	"connectorapi-go/internal/adapter/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"
)

// --- Helper struct & function ---
type apiHeaders struct {
	APIKey    string
	RequestID string
	Channel   string
	DeviceOS  string
}

func handleErrorResponse(c *gin.Context, appErr *app_error.AppError) {
	statusCode := http.StatusInternalServerError
	switch appErr.Code {
	case app_error.ErrValidation.Code,
		app_error.ErrRequiedParam.Code,
		app_error.ErrApiChannel.Code,
		app_error.ErrApiRequestID.Code,
		app_error.ErrApiDeviceOS.Code,
		app_error.ErrIDCardNotFound.Code,
		app_error.ErrSUEInfoNotFound.Code,
		app_error.ErrAgmNoNotFound.Code,
		app_error.ErrService.Code,
		app_error.ErrSystemI.Code,
		app_error.ErrSystemIUnexpect.Code:
		statusCode = http.StatusBadRequest
	case app_error.ErrUnauthorized.Code:
		statusCode = http.StatusUnauthorized
	case app_error.ErrForbidden.Code:
		statusCode = http.StatusForbidden
	case app_error.ErrNotFound.Code:
		statusCode = http.StatusNotFound
	case app_error.ErrService.Code:
		statusCode = http.StatusBadGateway
	case app_error.ErrTimeOut.Code:
		statusCode = http.StatusGatewayTimeout
	}

	errResponse := app_error.ErrorResponse{
		ErrorCode:    appErr.Code,
		ErrorMessage: appErr.Message,
	}

	c.JSON(statusCode, errResponse)
}

func getAPIHeaders(c *gin.Context) apiHeaders {
	return apiHeaders{
		APIKey:    c.GetHeader("Api-Key"),
		RequestID: c.GetHeader("Api-RequestID"),
		Channel:   c.GetHeader("Api-Channel"),
		DeviceOS:  c.GetHeader("Api-DeviceOS"),
	}
}

func ValidateHeadersAndAuth(c *gin.Context, method string, path string, apiKeyRepo *utils.APIKeyRepository, logger *zap.SugaredLogger) *app_error.AppError {
	headers := getAPIHeaders(c)

	if !apiKeyRepo.Validate(headers.APIKey, method, path) {
		logger.Warnw("Authorization failed", "path", path, "apiKey", headers.APIKey)
		return app_error.ErrUnauthorized
	}

	if headers.RequestID == "" || len(headers.RequestID) > 20 {
		return app_error.ErrApiRequestID
	}
	if headers.Channel == "" {
		return app_error.ErrApiChannel
	}
	if headers.DeviceOS == "" {
		return app_error.ErrApiDeviceOS
	}

	return nil
}

func formatValidationErrors(err error) []app_error.ValidationErrorDetail {
	var validationErrors []app_error.ValidationErrorDetail

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range ve {
			validationErrors = append(validationErrors, app_error.ValidationErrorDetail{
				Field:   fieldErr.Field(),
				Tag:     fieldErr.Tag(),
				Message: fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", fieldErr.Field(), fieldErr.Tag()),
			})
		}
	}
	return validationErrors
}

func HandleValidationError(err error) *app_error.AppError {
	validationErrors := formatValidationErrors(err)

	var missingFields []string
	var lengthExceededFields []string

	for _, ve := range validationErrors {
		switch ve.Tag {
		case "required":
			missingFields = append(missingFields, ve.Field)
		case "max", "lte":
			lengthExceededFields = append(lengthExceededFields, ve.Field)
		}
	}

	if len(missingFields) > 0 {
		return &app_error.AppError{
			Code:    app_error.ErrRequiedParam.Code,
			Message: app_error.ErrRequiedParam.Message + "(" + strings.Join(missingFields, ", ") + ")",
			Err:     fmt.Errorf("required fields: %v", missingFields),
		}
	}

	if len(lengthExceededFields) > 0 {
		return &app_error.AppError{
			Code:    app_error.ErrInternalLength.Code,
			Message: app_error.ErrInternalLength.Message + " (" + strings.Join(lengthExceededFields, ", ") + ")",
			Err:     fmt.Errorf("max length fields: %v", lengthExceededFields),
		}
	}

	return &app_error.AppError{
		Code:    app_error.ErrService.Code,
		Message: app_error.ErrService.Message,
		Err:     fmt.Errorf("%v", validationErrors),
	}
}
