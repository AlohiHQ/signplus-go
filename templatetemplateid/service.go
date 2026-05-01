package templatetemplateid

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

// Service provides methods to interact with TemplateTemplateID-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type Service struct {
	manager              *configmanager.ConfigManager
	hook                 hooks.Hook
	getTemplateConfig    []signplusconfig.RequestOption
	deleteTemplateConfig []signplusconfig.RequestOption
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
	return api.manager.GetTemplateTemplateID()
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

// SetGetTemplateConfig sets method-level configuration for GetTemplate.
// Options are applied to every future call to GetTemplate and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetTemplateConfig(opts ...signplusconfig.RequestOption) *Service {
	api.getTemplateConfig = opts
	return api
}

// SetDeleteTemplateConfig sets method-level configuration for DeleteTemplate.
// Options are applied to every future call to DeleteTemplate and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetDeleteTemplateConfig(opts ...signplusconfig.RequestOption) *Service {
	api.deleteTemplateConfig = opts
	return api
}

// Get template
func (api *Service) GetTemplate(ctx context.Context, templateID string, params GetTemplateRequestParams, opts ...signplusconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getTemplateConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/template/{template_id}").
		WithConfig(config).
		AddPathParam("template_id", templateID).
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

// Delete template
func (api *Service) DeleteTemplate(ctx context.Context, templateID string, opts ...signplusconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.deleteTemplateConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/template/{template_id}").
		WithConfig(config).
		AddPathParam("template_id", templateID).
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
