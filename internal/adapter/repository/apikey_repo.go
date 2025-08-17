package repository

import (
	"connectorapi-go/pkg/config"
)

// Validation of API keys based on configuration
type APIKeyRepository struct {
	keys map[string]*config.APIKeyIncoming
}

// New repository and pre-loads keys into map
func NewAPIKeyRepository(cfg *config.Config) *APIKeyRepository {
	keyMap := make(map[string]*config.APIKeyIncoming)
	for i := range cfg.APIKeys {
		keyMap[cfg.APIKeys[i].Key] = &cfg.APIKeys[i]
	}
	return &APIKeyRepository{keys: keyMap}
}

// Validate checks if an API key is valid, active, and has permission
func (r *APIKeyRepository) Validate(apiKey, method, path string) bool {
	clientKey, exists := r.keys[apiKey]
	if !exists || clientKey.Status != "active" {
		return false // Key does not exist or is inactive
	}

	// Check if the key has permission for the specific METHOD:PATH
	routeKey := method + ":" + path
	for _, p := range clientKey.Permissions {
		if p == routeKey {
			return true // Permission granted
		}
	}

	return false // No permission found
}
