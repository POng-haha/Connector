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

// --- Handler ---

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

	err := h.validateHeadersAndAuth(c, c.Request.Method, c.FullPath())
	if err != nil {
		handleErrorResponse(c, err)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleErrorResponse(c, app_error.ErrService)
		return
	}

	if err := h.validator.Struct(req); err != nil {
    	handleErrorResponse(c, h.HandleValidationError(err))
    	return
	}

	response, appErr := h.service.CollectionDetail(c, req)
	if appErr != nil {
		handleErrorResponse(c, appErr)
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

	err := h.validateHeadersAndAuth(c, c.Request.Method, c.FullPath())
	if err != nil {
		handleErrorResponse(c, err)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleErrorResponse(c, app_error.ErrService)
		return
	}

	if err := h.validator.Struct(req); err != nil {
    handleErrorResponse(c, h.HandleValidationError(err))
    return
	}

	response, appErr := h.service.CollectionLog(c, req)
	if appErr != nil {
		handleErrorResponse(c, appErr)
		return
	}

	c.JSON(http.StatusOK, response)
}
