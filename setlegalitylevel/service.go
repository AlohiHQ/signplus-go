package setlegalitylevel

import (
	"context"
	restClient "github.com/alohihq/signplus-go/internal/clients/rest"
	"github.com/alohihq/signplus-go/internal/clients/rest/hooks"
	"github.com/alohihq/signplus-go/internal/clients/rest/httptransport"
	"github.com/alohihq/signplus-go/internal/configmanager"
	"github.com/alohihq/signplus-go/sharedmodels"
	"github.com/alohihq/signplus-go/signplusconfig"
	"time"
)

// Service provides methods to interact with SetLegalityLevel-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type Service struct {
	manager                        *configmanager.ConfigManager
	hook                           hooks.Hook
	setEnvelopeLegalityLevelConfig []signplusconfig.RequestOption
}

func NewService() *Service {
	return &Service{
		manager: configmanager.NewConfigManager(signplusconfig.Config{}),
	}
}

// WithConfigManager sets the configuration manager for this service.
// Returns the service instance for method chaining.
func (api *Service) WithConfigManager(manager *configmanager.ConfigManager) *Service {
	api.manager = manager
	return api
}

// WithHook sets a custom hook for request/response interception.
// Returns the service instance for method chaining.
func (api *Service) WithHook(hook hooks.Hook) *Service {
	api.hook = hook
	return api
}

func (api *Service) config() *signplusconfig.Config {
	return api.manager.GetSetLegalityLevel()
}

func (api *Service) getHook() hooks.Hook {
	return api.hook
}

func (api *Service) SetBaseURL(baseURL string) {
	config := api.config()
	config.SetBaseURL(baseURL)
}

func (api *Service) SetTimeout(timeout time.Duration) {
	config := api.config()
	config.SetTimeout(timeout)
}

func (api *Service) SetAccessToken(accessToken string) {
	config := api.config()
	config.SetAccessToken(accessToken)
}

// SetSetEnvelopeLegalityLevelConfig sets method-level configuration for SetEnvelopeLegalityLevel.
// Options are applied to every future call to SetEnvelopeLegalityLevel and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetSetEnvelopeLegalityLevelConfig(opts ...signplusconfig.RequestOption) *Service {
	api.setEnvelopeLegalityLevelConfig = opts
	return api
}

// Set envelope legality level
func (api *Service) SetEnvelopeLegalityLevel(ctx context.Context, envelopeID string, setEnvelopeLegalityLevelRequest SetEnvelopeLegalityLevelRequest, params SetEnvelopeLegalityLevelRequestParams, opts ...signplusconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.setEnvelopeLegalityLevelConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/envelope/{envelope_id}/set_legality_level").
		WithConfig(config).
		WithBody(setEnvelopeLegalityLevelRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("envelope_id", envelopeID).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return resp.Data, nil
}
