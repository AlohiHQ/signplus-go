package signplus

import (
	"github.com/alohihq/signplus-go/internal/clients/rest/hooks"
	"github.com/alohihq/signplus-go/internal/configmanager"
	"github.com/alohihq/signplus-go/signplus1"
	"time"
)

// Signplus is the main SDK client that provides access to all service endpoints.
// It manages configuration, authentication, and service instances with centralized settings.
type Signplus struct {
	Signplus1 *signplus1.Service
	manager   *configmanager.ConfigManager
}

func NewSignplus(config Config) *Signplus {
	signplus1 := signplus1.NewService()

	manager := configmanager.NewConfigManager(config)
	hook := hooks.NewDefaultHook()
	signplus1.WithConfigManager(manager)
	signplus1.WithHook(hook)

	return &Signplus{
		Signplus1: signplus1,
		manager:   manager,
	}
}

func (s *Signplus) SetBaseURL(baseURL string) {
	s.manager.SetBaseURL(baseURL)
}

func (s *Signplus) SetTimeout(timeout time.Duration) {
	s.manager.SetTimeout(timeout)
}

func (s *Signplus) SetAccessToken(accessToken string) {
	s.manager.SetAccessToken(accessToken)
}

// SetEnvironment configures the SDK to use the specified environment's base URL.
func (s *Signplus) SetEnvironment(environment Environment) {
	s.manager.SetBaseURL(string(environment))
}

// c029837e0e474b76bc487506e8799df5e3335891efe4fb02bda7a1441840310c
