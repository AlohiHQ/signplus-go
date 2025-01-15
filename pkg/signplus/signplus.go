package signplus

import (
	"github.com/alohihq/signplus-go/internal/configmanager"
	"github.com/alohihq/signplus-go/pkg/signplus"
	"github.com/alohihq/signplus-go/pkg/signplusconfig"
	"time"
)

type Signplus struct {
	Signplus *signplus.SignplusService
	manager  *configmanager.ConfigManager
}

func NewSignplus(config signplusconfig.Config) *Signplus {
	signplus := signplus.NewSignplusService()

	manager := configmanager.NewConfigManager(config)
	signplus.WithConfigManager(manager)

	return &Signplus{
		Signplus: signplus,
		manager:  manager,
	}
}

func (s *Signplus) SetBaseUrl(baseUrl string) {
	s.manager.SetBaseUrl(baseUrl)
}

func (s *Signplus) SetTimeout(timeout time.Duration) {
	s.manager.SetTimeout(timeout)
}

func (s *Signplus) SetAccessToken(accessToken string) {
	s.manager.SetAccessToken(accessToken)
}

// c029837e0e474b76bc487506e8799df5e3335891efe4fb02bda7a1441840310c
