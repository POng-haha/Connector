package handler

import (
	"net/http"
	
	"connectorapi-go/internal/adapter/utils"
	"connectorapi-go/internal/core/domain"
	app_error "connectorapi-go/pkg/error"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"
)

// CollectionService defines the interface
type CollectionService interface {
	CollectionDetail(c *gin.Context, reqData domain.CollectionDetailRequest) (*domain.CollectionDetailResponse, *app_error.AppError)
	CollectionLog(c *gin.Context, reqData domain.CollectionLogRequest) (*domain.CollectionLogResponse, *app_error.AppError)
}

// CollectionHandler handles all Collection-related API requests
type CollectionHandler struct {
	service CollectionService
	validator *validator.Validate
	logger  *zap.SugaredLogger
	apikey  *utils.APIKeyRepository
}

// NewCollectionHandler creates a new instance of CollectionHandler
func NewCollectionHandler(s CollectionService, logger *zap.SugaredLogger, apikey *utils.APIKeyRepository) *CollectionHandler {
	return &CollectionHandler{
		service: s,
		validator: validator.New(),
		logger:  logger,
		apikey:  apikey,
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

// CollectionDetail
// @Tags         Common
// @Accept       json
// @Produce      json
// @Param        Api-Key              header    string                               false  "API key"
// @Param        request              body      domain.CollectionDetailRequest       false  "Body Request"
// @Success      200  {object}        domain.CollectionDetailResponse
// @Router       /Api/Collection/CollectionDetail [post]
func (h *CollectionHandler) CollectionDetail(c *gin.Context) {
	var req domain.CollectionDetailRequest

	err := ValidateHeadersAndAuth(c, c.Request.Method, c.FullPath(), h.apikey, h.logger)
	if err != nil {
		handleErrorResponse(c, err)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleErrorResponse(c, app_error.ErrService)
		return
	}

	if err := h.validator.Struct(req); err != nil {
    	handleErrorResponse(c, HandleValidationError(err))
    	return
	}

	response, appErr := h.service.CollectionDetail(c, req)
	if appErr != nil {
		handleErrorResponse(c, appErr)
		return
	}

	c.JSON(http.StatusOK, response)
}

// CollectionDetail
// @Tags         Common
// @Accept       json
// @Produce      json
// @Param        Api-Key              header    string                               false  "API key"
// @Param        request              body      domain.CollectionLogRequest       false  "Body Request"
// @Success      200  {object}        domain.CollectionLogResponse
// @Router       /Api/Collection/CollectionLog [post]
func (h *CollectionHandler) CollectionLog(c *gin.Context) {
	var req domain.CollectionLogRequest

	err := ValidateHeadersAndAuth(c, c.Request.Method, c.FullPath(), h.apikey, h.logger)
	if err != nil {
		handleErrorResponse(c, err)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleErrorResponse(c, app_error.ErrInternalServer)
		return
	}

	if err := h.validator.Struct(req); err != nil {
    	handleErrorResponse(c, HandleValidationError(err))
    	return
	}
	
	response, appErr := h.service.CollectionLog(c, req)
	if appErr != nil {
		handleErrorResponse(c, appErr)
		return
	}

	c.JSON(http.StatusOK, response)
}
