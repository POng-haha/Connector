package handler

import (
	"fmt"
	"net/http"

	"connectorapi-go/internal/core/domain"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// CustomerService defines the interface
type CustomerService interface {
	GetCustomerInfo(reqData domain.GetCustomerInfoRequest, requestID string) (*domain.GetCustomerInfoResponse, *domain.AppError)
	UpdateAddress(reqData domain.UpdateAddressRequest, requestID string) (*domain.UpdateAddressResponse, *domain.AppError)
	GetCustomerInfo004(reqData domain.GetCustomerInfo004Request, requestID string) (*domain.GetCustomerInfo004Response, *domain.AppError)
}

// CustomerHandler handles all customer-related API requests
type CustomerHandler struct {
	service   CustomerService
	validator *validator.Validate
}

// NewCustomerHandler creates a new instance of CustomerHandler
func NewCustomerHandler(s CustomerService) *CustomerHandler {
	return &CustomerHandler{
		service:   s,
		validator: validator.New(),
	}
}

// RegisterRoutes registers all routes related to customers to the router group
func (h *CustomerHandler) RegisterRoutes(rg *gin.RouterGroup) {
	customerRoutes := rg.Group("/customer")
	{
		customerRoutes.POST("/getcustomerinfo", h.GetCustomerInfo)
		customerRoutes.POST("/updateaddress", h.UpdateAddress)
		customerRoutes.POST("/getcustomerinfo004", h.GetCustomerInfo004)
	}
}

// GetCustomerInfo handles the POST /customer/getcustomerinfo request
// @Summary      Get Customer Information
// @Description  retrieves customer details based on customer ID and channel
// @Tags         Customer
// @Accept       json
// @Produce      json
// @Param        API-Key   header      string  true  "Client's API Key"
// @Param        RequestID header      string  true  "Client's Request ID"
// @Param        request   body      domain.GetCustomerInfoRequest  true  "Request Body"
// @Success      200       {object}  domain.GetCustomerInfoResponse
// @Failure      400       {object}  domain.ErrorResponse
// @Failure      401       {object}  domain.ErrorResponse
// @Failure      500       {object}  domain.ErrorResponse
// @Router       /customer/getcustomerinfo [post]
func (h *CustomerHandler) GetCustomerInfo(c *gin.Context) {
	var req domain.GetCustomerInfoRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrors := formatValidationErrors(err)
		appErr := &domain.AppError{
			Code:    domain.ErrValidation.Code,
			Message: domain.ErrValidation.Message,
			Err:     fmt.Errorf("%v", validationErrors),
		}
		handleErrorResponse(c, appErr)
		return
	}

	requestID := c.GetString(RequestID)
	if requestID == "" {
		// This case should ideally not happen if RequestIDMiddleware always sets it,
		// but it's good for robustness.
		appErr := &domain.AppError{
			Code:    domain.ErrInternalServer.Code, // Or ErrConfig if RequestID constant is not resolved
			Message: "RequestID not found in context after middleware",
		}
		handleErrorResponse(c, appErr)
		return
	}

	customerInfo, appErr := h.service.GetCustomerInfo(req, requestID)
	if appErr != nil {
		handleErrorResponse(c, appErr)
		return
	}

	c.JSON(http.StatusOK, customerInfo)
}

// UpdateAddress handles the POST /customer/updateaddress request
// @Summary      Update Customer Address
// @Description  update customer address from request
// @Tags         Customer
// @Accept       json
// @Produce      json
// @Param        API-Key   header      string  true  "Client's API Key"
// @Param        RequestID header      string  true  "Client's Request ID"
// @Param        request   body      domain.UpdateAddressRequest  true  "Request Body"
// @Success      200       {object}  domain.UpdateAddressResponse
// @Failure      400       {object}  domain.ErrorResponse
// @Failure      401       {object}  domain.ErrorResponse
// @Failure      500       {object}  domain.ErrorResponse
// @Router       /customer/updateaddress [post]
func (h *CustomerHandler) UpdateAddress(c *gin.Context) {
	var req domain.UpdateAddressRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrors := formatValidationErrors(err)
		appErr := &domain.AppError{
			Code:    domain.ErrValidation.Code,
			Message: domain.ErrValidation.Message,
			Err:     fmt.Errorf("%v", validationErrors),
		}
		handleErrorResponse(c, appErr)
		return
	}

	requestID := c.GetString(RequestID)
	if requestID == "" {
		// This case should ideally not happen if RequestIDMiddleware always sets it,
		// but it's good for robustness.
		appErr := &domain.AppError{
			Code:    domain.ErrInternalServer.Code, // Or ErrConfig if RequestID constant is not resolved
			Message: "RequestID not found in context after middleware",
		}
		handleErrorResponse(c, appErr)
		return
	}

	response, appErr := h.service.UpdateAddress(req, requestID)
	if appErr != nil {
		handleErrorResponse(c, appErr)
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetCustomerInfo004 handles the POST /customer/getcustomerinfo004 request
// @Summary      Get Customer Information format 004
// @Description  retrieves customer details based on customer ID and channel
// @Tags         Customer
// @Accept       json
// @Produce      json
// @Param        API-Key   header      string  true  "Client's API Key"
// @Param        RequestID header      string  true  "Client's Request ID"
// @Param        request   body      domain.GetCustomerInfoRequest  true  "Request Body"
// @Success      200       {object}  domain.GetCustomerInfoResponse
// @Failure      400       {object}  domain.ErrorResponse
// @Failure      401       {object}  domain.ErrorResponse
// @Failure      500       {object}  domain.ErrorResponse
// @Router       /customer/getcustomerinfo [post]
func (h *CustomerHandler) GetCustomerInfo004(c *gin.Context) {
	var req domain.GetCustomerInfo004Request

	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrors := formatValidationErrors(err)
		appErr := &domain.AppError{
			Code:    domain.ErrValidation.Code,
			Message: domain.ErrValidation.Message,
			Err:     fmt.Errorf("%v", validationErrors),
		}
		handleErrorResponse(c, appErr)
		return
	}

	requestID := c.GetString(RequestID)
	if requestID == "" {
		// This case should ideally not happen if RequestIDMiddleware always sets it,
		// but it's good for robustness.
		appErr := &domain.AppError{
			Code:    domain.ErrInternalServer.Code, // Or ErrConfig if RequestID constant is not resolved
			Message: "RequestID not found in context after middleware",
		}
		handleErrorResponse(c, appErr)
		return
	}

	customerInfo004, appErr := h.service.GetCustomerInfo004(req, requestID)
	if appErr != nil {
		handleErrorResponse(c, appErr)
		return
	}

	c.JSON(http.StatusOK, customerInfo004)
}