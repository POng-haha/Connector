package handler

import (
	"fmt"
	"net/http"

	"connectorapi-go/internal/core/domain"
	"connectorapi-go/internal/adapter/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"
)

// CollectionService defines the interface
type CollectionService interface {
	CollectionDetail(reqData domain.CollectionDetailRequest, requestID string) (*domain.CollectionDetailResponse, *domain.AppError)
	CollectionLog(reqData domain.CollectionLogRequest, requestID string) (*domain.CollectionLogResponse, *domain.AppError)
}

// CollectionHandler handles all Collection-related API requests
type CollectionHandler struct {
	service   CollectionService
	validator *validator.Validate
	logger	  *zap.SugaredLogger
	repo	  *repository.APIKeyRepository
}

// NewCollectionHandler creates a new instance of CollectionHandler
func NewCollectionHandler(s	CollectionService, logger *zap.SugaredLogger, repo *repository.APIKeyRepository) *CollectionHandler {
	return &CollectionHandler{
		service:   s,
		validator: validator.New(),
		logger: logger,
		repo : repo,
	}
}

// RegisterRoutes registers all routes related to Collection to the router group
func (h *CollectionHandler) RegisterRoutes(rg *gin.RouterGroup) {
	collectionRoutes := rg.Group("/Collection")
	{
		collectionRoutes.POST("/CollectionDetail", h.CollectionDetail)
		collectionRoutes.POST("/CollectionLog", h.CollectionLog)
	}
}

// CollectionDetail handles the POST /Collection/CollectionDetail request
// @Summary      Collection Detail
// @Description  Get CollectionDetail by ID Card
// @Tags         Collection
// @Accept       json
// @Produce      json
// @Param        API-Key   header      string  true  "Client's API Key"
// @Param        RequestID header      string  true  "Client's Request ID"
// @Param        request   body      domain.CollectionDetailRequest  true  "Request Body"
// @Success      200       {object}  domain.CollectionDetailResponse
// @Failure      400       {object}  domain.ErrorResponse
// @Failure      401       {object}  domain.ErrorResponse
// @Failure      500       {object}  domain.ErrorResponse
// @Router       /Collection/CollectionDetail [post]
func (h *CollectionHandler) CollectionDetail(c *gin.Context) {
	var req domain.CollectionDetailRequest
	var domainErr *domain.AppError

	apiKey := c.GetHeader("Api-Key")

	if !repo.Validate(apiKey, c.Request.Method, c.FullPath()) {
		logger.Warnw("Authorization failed", "path", c.FullPath(), "apiKey", apiKey)
		handleErrorResponse(c, domain.ErrUnauthorized)
		return
	}

	apiRequestID	:= c.GetHeader("Api-DeviceOS")
	apiChannel		:= c.GetHeader("Api-Channel")
	apiDeviceOS		:= c.GetHeader("Api-DeviceOS")

	if apiRequestID == "" || apiChannel == "" || apiDeviceOS == "" {
		if apiRequestID == "" {
			domainErr = domain.ErrApiRequestID
		}
		if apiChannel	== "" {
			domainErr = domain.ErrApiChannel
		}
		if apiDeviceOS	== "" {
			domainErr = domain.ErrApiDeviceOS
		}
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleErrorResponse(c, domain.ErrService)
		return
	}

	if  err := h.validator.Struct(req); err != nil {
		appErr := HandleValidationError(err)
		handleErrorResponse(c, appErr)
		return
	}

	response, appErr := h.service.CollectionDetail(req, requestID)
	if appErr != nil {
		handleErrorResponse(c, appErr)
		return
	}

	var responseError *domain.AppError
	if domainErr != nil {
		responseError = domainErr
	}

	if responseError != nil {
		handleErrorResponse(c, responseError)
		return
	}

	c.JSON(http.StatusOK, response)
}

// CollectionLog handles the POST /Collection/CollectionLog request
// @Summary      Collection Log
// @Description  Send Collection Log to sysi
// @Tags         Collection
// @Accept       json
// @Produce      json
// @Param        API-Key   header      string  true  "Client's API Key"
// @Param        RequestID header      string  true  "Client's Request ID"
// @Param        request   body      domain.CollectionLogRequest  true  "Request Body"
// @Success      200       {object}  domain.CollectionLogResponse
// @Failure      400       {object}  domain.ErrorResponse
// @Failure      401       {object}  domain.ErrorResponse
// @Failure      500       {object}  domain.ErrorResponse
// @Router       /Collection/CollectionLog [post]
func (h *CollectionHandler) CollectionLog(c *gin.Context) {
	var req domain.CollectionLogRequest
	var domainErr *domain.AppError

	apiKey := c.GetHeader("Api-Key")

	if !repo.Validate(apiKey, c.Request.Method, c.FullPath()) {
		logger.Warnw("Authorization failed", "path", c.FullPath(), "apiKey", apiKey)
		handleErrorResponse(c, appErr)
		return
	}

	apiRequestID	:= c.GetHeader("Api-DeviceOS")
	apiChannel		:= c.GetHeader("Api-Channel")
	apiDeviceOS		:= c.GetHeader("Api-DeviceOS")

	if apiRequestID == "" || apiChannel == "" || apiDeviceOS == "" {
		if apiRequestID == "" {
			domainErr = domain.ErrApiRequestID
		}
		if apiChannel	== "" {
			domainErr = domain.ErrApiChannel
		}
		if apiDeviceOS	== "" {
			domainErr = domain.ErrApiDeviceOS
		}
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleErrorResponse(c, domain.ErrService)
		return
	}

	if  err := h.validator.Struct(req); err != nil {
		appErr := HandleValidationError(err)
		handleErrorResponse(c, appErr)
		return
	}

	response, appErr := h.service.CollectionLog(req, requestID)
	if appErr != nil {
		handleErrorResponse(c, appErr)
		return
	}

	var responseError *domain.AppError
	if domainErr != nil {
		responseError = domainErr
	}
	
	if responseError != nil {
		handleErrorResponse(c, responseError)
		return
	}

	c.JSON(http.StatusOK, response)
}