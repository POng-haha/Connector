package service

import (
	"connectorapi-go/internal/adapter/client"
	"connectorapi-go/internal/adapter/utils"
	"connectorapi-go/internal/core/domain"
	"connectorapi-go/internal/core/service/format"
	"connectorapi-go/pkg/config"
	"fmt"
	"strings"

	app_error "connectorapi-go/pkg/error"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// TCPSocketClient defines the interface for a TCP socket client
type collectionTCPSocketClient = client.TCPSocketClient

// CollectionService implements the business logic for customer-related features
type CollectionService struct {
	config       *config.Config
	logger       *zap.SugaredLogger
	tcpClient    client.TCPSocketClient
	routes       map[string]config.Route
	destinations map[string]config.Destination
}

// NewCollectionService creates a new instance of CustomerService.
func NewCollectionService(
	cfg *config.Config,
	logger *zap.SugaredLogger,
	tcpClient collectionTCPSocketClient,
	routes map[string]config.Route,
	destinations map[string]config.Destination,
) *CollectionService {
	return &CollectionService{
		config:       cfg,
		logger:       logger,
		tcpClient:    tcpClient,
		routes:       routes,
		destinations: destinations,
	}
}

// CollectionDetail handles the CollectionDetail request
// It sends a request to the TCP service and returns the response.
func (s *CollectionService) CollectionDetail(c *gin.Context, reqData domain.CollectionDetailRequest) (*domain.CollectionDetailResponse, *app_error.AppError) {
	//const routeKey = "POST:/Api/Collection/CollectionDetail"
	routeKey := utils.GetRouteKey(c)
	const destinationName = "systemI"
	//var logLine string

	idValue, _ := c.Get("Api-RequestID")
	apiRequestID, ok := idValue.(string)
	if !ok {
		apiRequestID = ""
	}

	route, ok := s.routes[routeKey]
	if !ok {
		s.logger.Errorw("Route configuration not found for TCP service", "routeKey", routeKey)
		return nil, app_error.ErrConfig
	}

	destination, ok := s.destinations[destinationName]
	if !ok {
		s.logger.Errorw("TCP Destination configuration not found", "destinationName", destinationName)
		return nil, app_error.ErrConfig
	}
	if destination.Type != "tcp" {
		s.logger.Errorw("Destination type is not TCP", "destinationName", destinationName, "type", destination.Type)
		return nil, app_error.ErrConfig
	}

	portList, ok := destination.Ports["CollectionDetail"]
	if !ok || len(portList) == 0 {
		s.logger.Errorw("Invalid port configuration", "port", portList)
		return nil, app_error.ErrConfig
	}
	port := utils.RandomPortFromList(portList)
	if port == "" {
		s.logger.Errorw("Invalid port configuration", "port", portList)
		return nil, app_error.ErrConfig
	}

	tcpAddress := fmt.Sprintf("%s:%s", destination.IP, port)

	formattedRequestID := utils.PadOrTruncate(apiRequestID, 20)
	fixedLengthData := format.FormatCollectionDetailRequest(reqData)

	header := utils.BuildFixedLengthHeader(
		route.System,
		route.Service,
		route.Format,
		formattedRequestID,
		route.RequestLength,
	)

	combinedPayloadString := header + fixedLengthData

	s.logger.Info("Sending TCP request payload : ", combinedPayloadString)

	responseStr, err := s.tcpClient.SendAndReceive(tcpAddress, combinedPayloadString)
	if err != nil {
		s.logger.Errorw("Downstream TCP service call failed", "error", err, "address", tcpAddress)
		return nil, app_error.ErrService
	}

	// Check if the response contains an error
	errorCode := strings.TrimSpace(responseStr[67:73])
	errorMessage := strings.TrimSpace(responseStr[73:123])

	switch errorCode {
	case "SVC105":
		return nil, app_error.ErrRequiedParam
	case "SVC117":
		return nil, app_error.ErrIDCardNotFound
	case "SVC902":
		return nil, app_error.ErrSystemI
	case "SVC203":
		return nil, app_error.ErrSUEInfoNotFound
	default:
    	if errorCode != "" {
        	s.logger.Info("Unknown error code from System I : ",
         	   "code", errorCode,
          	  "message", errorMessage,
        	)
        	return nil, app_error.ErrSystemIUnexpect
    	}
	}

	s.logger.Info("Received downstream TCP response", "response", string(responseStr))

	response, _ := format.FormatCollectionDetailResponse(responseStr)
	return &response, nil
}

// CollectionLog handles the logic for logging collection data via TCP.
// It sends a request to the TCP service and returns the response.
func (s *CollectionService) CollectionLog(c *gin.Context, reqData domain.CollectionLogRequest) (*domain.CollectionLogResponse, *app_error.AppError) {
	//const routeKey = "POST:/Api/Collection/CollectionLog"
	routeKey := utils.GetRouteKey(c)
	const destinationName = "systemI"

	idValue, _ := c.Get("Api-RequestID")
	apiRequestID, ok := idValue.(string)
	if !ok {
		apiRequestID = ""
	}

	route, ok := s.routes[routeKey]
	if !ok {
		s.logger.Errorw("Route configuration not found for TCP service", "routeKey", routeKey)
		return nil, app_error.ErrConfig
	}

	destination, ok := s.destinations[destinationName]
	if !ok {
		s.logger.Errorw("TCP Destination configuration not found", "destinationName", destinationName)
		return nil, app_error.ErrConfig
	}
	if destination.Type != "tcp" {
		s.logger.Errorw("Destination type is not TCP", "destinationName", destinationName, "type", destination.Type)
		return nil, app_error.ErrConfig
	}

	portList, ok := destination.Ports["CollectionDetail"]
	if !ok || len(portList) == 0 {
		s.logger.Errorw("Invalid port configuration", "port", portList)
		return nil, app_error.ErrConfig
	}
	port := utils.RandomPortFromList(portList)
	if port == "" {
		s.logger.Errorw("Invalid port configuration", "port", portList)
		return nil, app_error.ErrConfig
	}

	tcpAddress := fmt.Sprintf("%s:%s", destination.IP, port)

	formattedRequestID := utils.PadOrTruncate(apiRequestID, 20)
	fixedLengthData := format.FormatCollectionLogRequest(reqData)

	header := utils.BuildFixedLengthHeader(
		route.System,
		route.Service,
		route.Format,
		formattedRequestID,
		route.RequestLength,
	)

	combinedPayloadString := header + fixedLengthData
	s.logger.Info("Sending TCP request payload", "payload", combinedPayloadString)

	responseStr, err := s.tcpClient.SendAndReceive(tcpAddress, combinedPayloadString)
	if err != nil {
		s.logger.Errorw("Downstream TCP service call failed", "error", err, "address", tcpAddress)
		return nil, app_error.ErrService
	}

	// Check if the response contains an error
	errorCode := strings.TrimSpace(responseStr[67:73])
	errorMessage := strings.TrimSpace(responseStr[73:123])

	switch errorCode {
	case "SVC216", "SVC235", "SVC342", "SVC343", "SCV344":
		return nil, app_error.ErrRequiedParam
	case "SVC236":
		return nil, app_error.ErrAgmNoNotFound
	case "SVC902":
		return nil, app_error.ErrSystemI
	default:
    	if errorCode != "" {
        	s.logger.Info("Unknown error code from System I : ",
         	   "code", errorCode,
          	  "message", errorMessage,
        	)
        	return nil, app_error.ErrSystemIUnexpect
    	}
	}

	s.logger.Debugw("Received downstream TCP response", "response", string(responseStr))

	response, _:= format.FormatCollectionLogResponse(responseStr)
	return &response, nil
}
