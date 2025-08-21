package handler

import (
	// "fmt"
	"net/http"
	"time"

	"picoapi-go/internal/core/domain"
	"picoapi-go/internal/adapter/utils"
	"picoapi-go/pkg/config"
	app_error "picoapi-go/pkg/error"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"
)

// DopaService defines the interface
type DopaService interface {
	CheckDOPA(c *gin.Context, checkDOPARq domain.CheckDOPARequest) domain.CheckDOPAResponse
}

// DopaHandler handles all customer-related API requests
type DopaHandler struct {
	service   DopaService
	validator *validator.Validate
	logger    *zap.SugaredLogger
	config    *config.Config
	repo 	  *utils.APIKeyRepository
}

// NewDopaHandler creates a new instance of CommonHandler
func NewDopaHandler(s DopaService, logger *zap.SugaredLogger, cfg *config.Config, repo *utils.APIKeyRepository) *DopaHandler {
	return &DopaHandler{
		service:   s,
		validator: validator.New(),
		logger:    logger,
		config:    cfg,
		repo:	   repo,
	}
}

// RegisterRoutes registers all routes related to customers to the router group
func (h *DopaHandler) RegisterRoutes(rg *gin.RouterGroup) {
	customerRoutes := rg.Group("/Customer")
	{
		customerRoutes.POST("/CheckDOPA", h.CheckDOPA)
	}
}

// CheckDOPA godoc
// @Tags         customer
// @Accept       json
// @Produce      json
// @Param        Api-Key              header    string                               false  "API key"
// @Param        Api-DeviceOS         header    string                               false  "Device OS"
// @Param        request              body      domain.CheckDOPARequest              false  "Body Request"
// @Success      200  {object}        domain.CheckDOPAResponse
// @Router       /api/Customer/CheckDOPA [post]
func (h *DopaHandler) CheckDOPA(c *gin.Context) {
	timeNow := time.Now()
	var logList []string
	var req domain.CheckDOPARequest
	serviceName := "CheckDOPA"

	appLogger.Info("Request from client",
    zap.Any(serviceName, "Data request", req),
	)

	apiKey := c.GetHeader("Api-Key")
	if !h.repo.Validate(apiKey, c.Request.Method, c.FullPath()) {
		h.logger.Errorw("Authorization failed", "path", c.FullPath(), "apiKey", apiKey)
		handleErrorResponse(c, app_error.ErrUnauthorized)
		return
	}

	apiDeviceOS := c.GetHeader("Api-DeviceOS")
	if apiDeviceOS == "" {
		handleErrorResponse(c, app_error.ErrApiDeivceOS)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleErrorResponse(c, app_error.ErrService)
		return
	}
	if err := h.validator.Struct(req); err != nil {
		appErr := HandleValidationError(err)
		handleErrorResponse(c, appErr)
		if appErr.ErrorCode == "SYS500" {
			return
		}
		return
	}

	CheckDOPAResponse := h.service.CheckDOPARequest(c, req)
	if CheckDOPAResponse.AppError != nil {
		handleErrorResponse(c, CheckDOPAResponse.AppError)
		return
	}

	appLogger.Info("Response to client",
    zap.Any(serviceName, "Data response", CheckDOPAResponse),
	)

	var responseError *app_error.AppError
	if  CheckDOPAResponse.DomainError != nil {
		responseError = CheckDOPAResponse.DomainError
	}
	if responseError != nil {
		handleErrorResponse(c, responseError)
		return
	}

	c.JSON(http.StatusOK, CheckDOPAResponse.Response)
}
