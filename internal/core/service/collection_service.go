package service

import (
	"encoding/json"
	"connectorapi-go/internal/core/client"
	"connectorapi-go/internal/core/domain"
	"connectorapi-go/pkg/config"
	"connectorapi-go/pkg/format"

	"go.uber.org/zap"
)

// CollectionService implements the business logic for customer-related features
type CollectionService struct {
	config    *config.Config
	logger    *zap.SugaredLogger
	tcpClient *client.TCPSocketClient
}

// NewCollectionService creates a new instance of CustomerService.
func NewCollectionService(cfg *config.Config, logger *zap.SugaredLogger, tcpClient *client.TCPSocketClient) *CollectionService {
	return &CollectionService{
		config:    cfg,
		logger:    logger,
		tcpClient: tcpClient,
	}
}

// CollectionDetail handles the logic for retrieving customer information via TCP.
func (s *CollectionService) CollectionDetail(reqData domain.CollectionDetailRequest, requestID string) (*domain.CollectionDetailResponse, *domain.AppError) {
	const routeKey = "POST:/Api/Collection/CollectionDetail"
	const destinationName = "systemi"

	route, ok := s.config.Routes[routeKey]
	if !ok {
		s.logger.Errorw("Route configuration not found for TCP service", "routeKey", routeKey)
		return nil, domain.ErrConfig
	}

	destination, ok := s.config.Destinations[destinationName]
	if !ok {
		s.logger.Errorw("TCP Destination configuration not found", "destinationName", destinationName)
		return nil, domain.ErrConfig
	}
	if destination.Type != "tcp" {
		s.logger.Errorw("Destination type is not TCP", "destinationName", destinationName, "type", destination.Type)
		return nil, domain.ErrConfig
	}
	tcpAddress := destination.Host

	formattedRequestID := format.PadOrTruncate(requestID, 20)
	fixedLengthData := format.FormatCollectionDetailRequestFixedLength(reqData)

	header := format.BuildFixedLengthHeader(
		route.System,
		route.Service,
		route.Format,
		formattedRequestID,
		fixedLengthData,
	)

	combinedPayloadString := header + fixedLengthData
	requestPayload := []byte(combinedPayloadString)

	s.logger.Debugw("Sending TCP request payload", "payload", combinedPayloadString)

	responseBytes, err := s.tcpClient.SendAndReceive(tcpAddress, requestPayload)
	if err != nil {
		s.logger.Errorw("Downstream TCP service call failed", "error", err, "address", tcpAddress)
		return nil, domain.ErrServiceDown
	}

	s.logger.Debugw("Received downstream TCP response", "response", string(responseBytes))

	var CollectionDetailResponse domain.CollectionDetailResponse
	if err := json.Unmarshal(responseBytes, &CollectionDetailResponse); err != nil {
		s.logger.Errorw("Failed to decode downstream TCP response", "error", err)
		return nil, domain.ErrInternalServer
	}
	return &CollectionDetailResponse, nil
}

// CollectionLog handles the logic for updating a customer's address via TCP.
func (s *CollectionService) CollectionLog(reqData domain.CollectionLogRequest, requestID string) (*domain.CollectionLogResponse, *domain.AppError) {
	const routeKey = "POST:/Api/Collection/CollectionLog"
	const destinationName = "systemi"

	route, ok := s.config.Routes[routeKey]
	if !ok {
		s.logger.Errorw("Route configuration not found for TCP service", "routeKey", routeKey)
		return nil, domain.ErrConfig
	}

	destination, ok := s.config.Destinations[destinationName]
	if !ok {
		s.logger.Errorw("TCP Destination configuration not found", "destinationName", destinationName)
		return nil, domain.ErrConfig
	}
	if destination.Type != "tcp" {
		s.logger.Errorw("Destination type is not TCP", "destinationName", destinationName, "type", destination.Type)
		return nil, domain.ErrConfig
	}
	tcpAddress := destination.Host

	formattedRequestID := format.PadOrTruncate(requestID, 20)
	fixedLengthData := format.FormatCollectionLogRequestFixedLength(reqData)

	header := format.BuildFixedLengthHeader(
		route.System,
		route.Service,
		route.Format,
		formattedRequestID,
		fixedLengthData,
	)

	combinedPayloadString := header + fixedLengthData
	requestPayload := []byte(combinedPayloadString)

	s.logger.Debugw("Sending TCP request payload", "payload", combinedPayloadString)

	responseBytes, err := s.tcpClient.SendAndReceive(tcpAddress, requestPayload)
	if err != nil {
		s.logger.Errorw("Downstream TCP service call failed", "error", err, "address", tcpAddress)
		return nil, domain.ErrServiceDown
	}

	s.logger.Debugw("Received downstream TCP response", "response", string(responseBytes))

	var CollectionLogResponse domain.CollectionLogResponse
	if err := json.Unmarshal(responseBytes, &CollectionLogResponse); err != nil {
		s.logger.Errorw("Failed to decode downstream TCP response", "error", err)
		return nil, domain.ErrInternalServer
	}
	return &CollectionLogResponse, nil
}