package service

import (
	"context"
	"fmt"
	"time"

	"picoapi-go/internal/adapter/client/api"
	dopa "picoapi-go/internal/adapter/client/api/dopa"
	"picoapi-go/internal/core/domain"
	"picoapi-go/pkg/config"
	app_error "picoapi-go/pkg/error"

	"go.uber.org/zap"
)

type DopaService struct {
	config     *config.Config
	logger     *zap.SugaredLogger
	httpClient *api.HTTPClient
	endpoint   string
	certName   string
}

func NewDopaService(cfg *config.Config, logger *zap.SugaredLogger, httpClient *api.HTTPClient, endpoint, certName string) *DopaService {
	return &DopaService{
		config:     cfg,
		logger:     logger,
		httpClient: httpClient,
		endpoint:   endpoint,
		certName:   certName,
	}
}

func (s *DopaService) CheckDOPA(ctx context.Context, req domain.CheckDOPARequest) (*domain.CheckDOPAResponse, error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	serviceName := "CheckDOPA"

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
	if destination.Type != "https" {
		s.logger.Errorw("Destination type is not TCP", "destinationName", destinationName, "type", destination.Type)
		return nil, app_error.ErrConfig
	}

	// 1. สร้าง SOAP request
	soapReq := &dopa.CheckCardByLaser{
		PID:       req.IDCardNo,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		BirthDay:  req.BirthDate,
		Laser:     req.LaserID,
	}

	s.logger.Infow("Send to Dopa request", "service", serviceName, "timestamp", timestamp, "parameter_request", req)

	// 2. Call SOAP
	soapResp := &dopa.CheckCardByLaserResponse{}
	err := s.httpClient.CallSOAP(ctx, s.certName, s.endpoint, "http://tempuri.org/CheckCardByLaser", soapReq, soapResp)
	if err != nil {
		s.logger.Errorw("Dopa service call failed", "error", err)
		return nil, app_error.ErrDopa
	}

	if soapResp.CheckCardByLaserResult == nil {
		s.logger.Error("Dopa response result is nil")
		return nil, app_error.ErrDopa
	}

	result := soapResp.CheckCardByLaserResult
	checkDOPAResponse := struct {
		IsError      bool
		ErrorMessage string
		Code         string
		Desc         string
		Reference    string
	}{
		IsError:      result.IsError,
		ErrorMessage: result.ErrorMessage,
		Code:         fmt.Sprintf("%d", result.Code),
		Desc:         result.Desc,
		Reference:    req.Reference,
	}

	s.logger.Infow("Dopa response", "service", serviceName, "timestamp", timestamp, "data_response", checkDOPAResponse)

	// 3. Business logic
	switch checkDOPAResponse.Code {
	case "0":
		if checkDOPAResponse.ErrorMessage != "สถานะปกติ" {
			return nil, app_error.NewErrorDopaInvalid(checkDOPAResponse.ErrorMessage)
		}
	case "1", "2", "3", "4", "5":
		return nil, app_error.NewErrorDopaInvalid(checkDOPAResponse.ErrorMessage)
	default:
		s.logger.Infow("Unknown code from Dopa", "code", checkDOPAResponse.Code, "message", checkDOPAResponse.ErrorMessage)
		return nil, app_error.ErrDopa
	}

	// 4. Success response
	return &domain.CheckDOPAResponse{
		Result:    true,
		Reference: req.Reference,
	}, nil
}
