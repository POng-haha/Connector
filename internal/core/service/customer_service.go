package service

import (
	"encoding/json"
	"connectorapi-go/internal/core/client"
	"connectorapi-go/internal/core/domain"
	"connectorapi-go/pkg/config"
	"connectorapi-go/pkg/format"

	"go.uber.org/zap"
)

// TCPSocketClient defines the interface for a TCP socket client
type TCPSocketClient = client.TCPSocketClient

// CustomerService implements the business logic for customer-related features
type CustomerService struct {
	config    *config.Config
	logger    *zap.SugaredLogger
	tcpClient TCPSocketClient // Use TCP client
}

// NewCustomerService creates a new instance of CustomerService.
func NewCustomerService(cfg *config.Config, logger *zap.SugaredLogger, tcpClient TCPSocketClient) *CustomerService {
	return &CustomerService{
		config:    cfg,
		logger:    logger,
		tcpClient: tcpClient,
	}
}

// GetCustomerInfo handles the logic for retrieving customer information via TCP.
func (s *CustomerService) GetCustomerInfo(reqData domain.GetCustomerInfoRequest, requestID string) (*domain.GetCustomerInfoResponse, *domain.AppError) {
	const routeKey = "POST:/api/customer/getcustomerinfo"
	const destinationName = "systemi"

	s.logger.Infow("Executing GetCustomerInfo service via TCP", "idcardno", reqData.IDCardNo)

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
	fixedLengthData := format.FormatGetCustomerInfoRequestFixedLength(reqData)

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

	var getCustomerInfoResponse domain.GetCustomerInfoResponse
	if err := json.Unmarshal(responseBytes, &getCustomerInfoResponse); err != nil {
		s.logger.Errorw("Failed to decode downstream TCP response", "error", err)
		return nil, domain.ErrInternalServer
	}
	return &getCustomerInfoResponse, nil
}

// UpdateAddress handles the logic for updating a customer's address via TCP.
func (s *CustomerService) UpdateAddress(reqData domain.UpdateAddressRequest, requestID string) (*domain.UpdateAddressResponse, *domain.AppError) {
	const routeKey = "POST:/api/customer/updateaddress"
	const destinationName = "systemi"

	s.logger.Infow("Executing UpdateAddress service", "customerId", reqData.CustomerID)

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
	fixedLengthData := format.FormatUpdateAddressRequestFixedLength(reqData)

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

	var updateAddressResponse domain.UpdateAddressResponse
	if err := json.Unmarshal(responseBytes, &updateAddressResponse); err != nil {
		s.logger.Errorw("Failed to decode downstream TCP response for UpdateAddress", "error", err)
		return nil, domain.ErrInternalServer
	}
	return &updateAddressResponse, nil
}

// GetCustomerInfo004 handles the logic for retrieving customer information via TCP.
func (s *CustomerService) GetCustomerInfo004(reqData domain.GetCustomerInfo004Request, requestID string) (*domain.GetCustomerInfo004Response, *domain.AppError) {
	const routeKey = "POST:/api/customer/getcustomerinfo004"
	const destinationName = "systemi"

	s.logger.Infow("Executing GetCustomerInfo004 service via TCP", "IDCardNo", reqData.IDCardNo)

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
	fixedLengthData := format.FormatGetCustomerInfo004RequestFixedLength(reqData)

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

	var getCustomerInfo004Response domain.GetCustomerInfo004Response
	if err := json.Unmarshal(responseBytes, &getCustomerInfo004Response); err != nil {
		s.logger.Errorw("Failed to decode downstream TCP response", "error", err)
		return nil, domain.ErrInternalServer
	}
	return &getCustomerInfo004Response, nil
}