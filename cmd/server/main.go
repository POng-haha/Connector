// @title           Template Go API
// @version         1.0
// @description     This is a sample API documentation for our Go API application.
// @termsOfService  http://swagger.io/terms/

// @contact.name   SYE
// @contact.url    https://aeon.co.th
// @contact.email  sye@aeon.co.th

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host            localhost:8080
// @BasePath        /api

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"

	tcp_client_adapter "connectorapi-go/internal/core/client"
	handler_adapter "connectorapi-go/internal/adapter/handler/api"
	repo_adapter "connectorapi-go/internal/adapter/repository"
	service_core "connectorapi-go/internal/core/service"
	"connectorapi-go/pkg/config"
	"connectorapi-go/pkg/logger"
	"connectorapi-go/pkg/metrics"
)

// main is the primary function that starts the API Gateway
func main() {
	cfg, err := config.Load("./configs/config.yaml")
	if err != nil {
		log.Fatalf("FATAL: Failed to load configuration: %v", err)
	}

	appLogger := logger.New(cfg.Logger.Level)
	defer appLogger.Sync()
	appLogger.Info("Logger initialized")

	gin.SetMode(cfg.Server.Mode)
	appLogger.Infow("Gin mode set", "mode", cfg.Server.Mode)

	appLogger.Info("Initializing dependencies...")
	metrics.Init()

	// --- Adapters ---
	apiKeyRepo := repo_adapter.NewAPIKeyRepository(cfg)

	// --- TCP Socket Client Initialization ---
	// Instantiate the BasicTCPSocketClient with appropriate timeouts.
	// You might want to make these timeouts configurable via config.yaml.
	tcpClient := tcp_client_adapter.NewBasicTCPSocketClient(
		5*time.Second,  // Dial Timeout (e.g., 5 seconds to establish connection)
		10*time.Second, // Read/Write Timeout (e.g., 10 seconds for data transfer)
	)
	appLogger.Info("TCP Socket Client initialized")

	// --- Core Services ---
	// Each service handles a specific domain of business logic
	// Now, inject the new tcpClient into CustomerService.
	customerService := service_core.NewCustomerService(cfg, appLogger, tcpClient)
	collectionService := service_core.NewCollectionrService(cfg, appLogger, tcpClient)
	appLogger.Info("Customer Service initialized with TCP client")

	// --- Handlers (API Layer) ---
	// Handlers are the entry point for API requests and depend on services
	customerHandler := handler_adapter.NewCustomerHandler(customerService)
	collectionHandler := handler_adapter.NewCollectionHandler(collectionService)

	appLogger.Info("Setting up router...")
	router := handler_adapter.SetupRouter(appLogger, apiKeyRepo, customerHandler, collectionHandler)

	serverAddress := fmt.Sprintf(":%s", cfg.Server.Port)
	appLogger.Infow("Starting server", "address", serverAddress)
	if err := router.Run(serverAddress); err != nil {
		appLogger.Fatalw("Failed to start server", "error", err)
	}
}
