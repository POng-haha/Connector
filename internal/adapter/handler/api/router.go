package handler

import (
	"fmt"
	"net/http"
	"time"

	"connectorapi-go/internal/adapter/repository"
	"connectorapi-go/internal/core/domain"
	"connectorapi-go/pkg/logger"
	"connectorapi-go/pkg/metrics"

	_ "connectorapi-go/docs"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

const apiRequestID 	= "Api-RequestID"
const apiKey 		= "Api-Key"
const apiLanguage	= "Api-Language"
const apiDeviceOS 	= "Api-DeviceOS"
const apiChannel 	= "Api-Channel"

// SetupRouter
func SetupRouter(
	appLogger *zap.SugaredLogger,
	repo *repository.APIKeyRepository,
	customerHandler *CustomerHandler,
	collectionHandler *CollectionHandler,
) *gin.Engine {
	router := gin.New()

	// --- Global Middlewares ---
	router.Use(ApiRequestIDMiddleware())
	router.Use(ApiKeyMiddleware())
	router.Use(ApiLanguageMiddleware())
	router.Use(ApiDeviceOSMiddleware())
	router.Use(ApiChannelMiddleware())

	router.Use(logger.GinLogger(appLogger, apiRequestID, apiKey, apiLanguage, apiDeviceOS, apiChannel))
	router.Use(PrometheusMiddleware())
	router.Use(gin.Recovery())

	// --- Public API Group (/Api) ---
	api := router.Group("/Api")
	{
		api.GET("/healthz", HealthCheck)
		api.GET("/metrics", gin.WrapH(promhttp.Handler()))
		api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// --- Sub API Group  ---
	apiAuth.api.Group("")
	{
		customerHandler.RegisterRoutes(apiAuth)
		collectionHandler.RegisterRoutes(apiAuth)
	}

	return router
}

// --- Middlewares Definitions ---

// RequestIDMiddleware checks for an incoming X-Request-ID, RequestID header
func ApiRequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := c.GetHeader("X-Request-ID")

		if reqID == "" {
			reqID = c.GetHeader("Api-RequestID")
		}

		if reqID == "" {
			prefix := "RQ"
			now := time.Now()
			datetime := now.Format("20060102150405") //format YYYYMMDDhhmmss
			rand.Seed(time.now().Uninano()) // seed random with time now
			runningNo := rand.Intn(1000) // random 0-9999
			runningStr := fmt.Sprintf("%04d", runningNo)

			reqID = prefix + datetime + runningStr
		}

		c.Set(RequestID, reqID)
		c.Header("X-Request-ID", reqID)
		c.Header("Api-RequestID", reqID)

		c.Next()
	}
}

func ApiKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("X-Key")
		if key == "" {
			key = c.GetHeader("Api-Key")
		}

		c.Set(apiKey, key)
		c.Header("X-Key", key)
		c.Header("Api-Key", key)

		c.Next()
	}
}

func ApiLanguageMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		language := c.GetHeader("X-Language")
		if language == "" {
			language = c.GetHeader("Api-Language")
		}
		if language == "" {
			language = "EN"
		}

		c.Set(apiLanguage, language)
		c.Header("X-Language", language)
		c.Header("Api-Language", language)

		c.Next()
	}
}

func ApiDeviceOSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		deviceOS := c.GetHeader("X-Device-OS")
		if deviceOS == "" {
			deviceOS = c.GetHeader("Api-DeviceOS")
		}

		c.Set(apiDeviceOS, deviceOS)
		c.Header("X-DeviceOS", deviceOS)
		c.Header("Api-DeviceOS", deviceOS)

		c.Next()
	}
}

func ApiChannelMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		channel := c.GetHeader("X-Channel")
		if channel == "" {
			channel = c.GetHeader("Api-Channel")
		}

		c.Set(apiChannel, channel)
		c.Header("X-Channel", channel)
		c.Header("Api-Channel", channel)

		c.Next()
	}
}

// PrometheusMiddleware
func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		status := fmt.Sprintf("%d", c.Writer.Status())
		path := c.FullPath()
		method := c.Request.Method
		metrics.HttpRequestsTotal.With(prometheus.Labels{"method": method, "path": path, "status": status}).Inc()
		metrics.HttpRequestDuration.With(prometheus.Labels{"method": method, "path": path, "status": status}).Observe(time.Since(start).Seconds())
	}
}

// HealthCheck provides a simple health check endpoint.
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
