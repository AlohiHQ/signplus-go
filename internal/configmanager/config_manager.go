package configmanager

import (
	"github.com/alohihq/signplus-go/signplusconfig"
	"sync"
	"time"
)

// ConfigManager manages configuration across all services with synchronized updates.
// Provides centralized configuration management and OAuth token handling for multiple services.
type ConfigManager struct {
	signplus1 signplusconfig.Config
	// mu protects concurrent reads and writes to access/refresh token fields.
	// Guarded operations: GetAccessToken, GetRefreshToken, UpdateAccessToken.
	mu sync.RWMutex
}

// NewConfigManager creates a new configuration manager with the provided config and optional OAuth token service.
// Initializes service-specific configs and sets up OAuth token management if enabled.
func NewConfigManager(config signplusconfig.Config) *ConfigManager {
	return &ConfigManager{
		signplus1: config,
	}
}

// SetBaseURL updates the BaseURL configuration parameter across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetBaseURL(baseURL string) {
	c.signplus1.SetBaseURL(baseURL)
}

// SetTimeout updates the Timeout configuration parameter across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetTimeout(timeout time.Duration) {
	c.signplus1.SetTimeout(timeout)
}

// SetAccessToken updates the AccessToken configuration parameter across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetAccessToken(accessToken string) {
	c.signplus1.SetAccessToken(accessToken)
}

// SetRetryConfig updates the retry configuration across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetRetryConfig(retry signplusconfig.RetryConfig) {
	c.signplus1.SetRetryConfig(retry)
}

// UpdateAccessToken replaces an access token across all services that use the original value.
// Used for token refresh to update all service configurations simultaneously.
// Write-locked so concurrent GetAccessToken reads see a consistent value.
func (c *ConfigManager) UpdateAccessToken(originalValue string, newValue string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.signplus1.AccessToken != nil && *c.signplus1.AccessToken == originalValue {
		c.signplus1.SetAccessToken(newValue)
	}
}

// GetSignplus1 returns the configuration for the Signplus1 service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetSignplus1() *signplusconfig.Config {
	return &c.signplus1
}

// GetBaseURL returns the currently configured base URL.
// All services share the same base URL; this reads it from the first service's config.
func (c *ConfigManager) GetBaseURL() string {
	return c.signplus1.BaseURL
}
