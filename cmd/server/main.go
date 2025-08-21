package api

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"

	"picoapi-go/pkg/config"

	"go.uber.org/zap"
)

type HTTPClient struct {
	mu     sync.RWMutex
	client map[string]*http.Client
	logger *zap.SugaredLogger
	config *config.Config
}

func NewHTTPClient(cfg *config.Config, logger *zap.SugaredLogger) *HTTPClient {
	client := make(map[string]*http.Client)
	for name, certCfg := range cfg.Certs {
		clientWithCerts, err := CreateHttpClientWithCert(certCfg.CACertPath)
		if err != nil {
			logger.Errorw("failed to create http client for cert", "cert", name, "error", err)
			continue
		}
		client[name] = clientWithCerts
	}

	return &HTTPClient{
		client: client,
		logger: logger,
		config: cfg,
	}
}

func (c *HTTPClient) CallSOAP(ctx context.Context, certName, endpoint, soapAction string, requestBody interface{}, responseBody interface{}) error {
	xmlBytes, err := xml.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("failed to marshal SOAP request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewReader(xmlBytes))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", soapAction)

	c.logger.Infow("Sending SOAP request", "url", endpoint, "soapAction", soapAction, "requestBody", string(xmlBytes))

	resp, err := c.DoWithCert(certName, endpoint, req)
	if err != nil {
		return fmt.Errorf("failed to send SOAP request: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read SOAP response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("SOAP request failed status=%d body=%s", resp.StatusCode, string(bodyBytes))
	}

	c.logger.Infow("Received SOAP response", "status", resp.StatusCode, "responseBody", string(bodyBytes))

	if err := xml.Unmarshal(bodyBytes, responseBody); err != nil {
		return fmt.Errorf("failed to unmarshal SOAP response: %w", err)
	}
	return nil
}
