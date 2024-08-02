package configmanager

import "github.com/alohihq/signplus-go/pkg/signplusconfig"

type ConfigManager struct {
	Signplus signplusconfig.Config
}

func NewConfigManager(config signplusconfig.Config) *ConfigManager {
	return &ConfigManager{
		Signplus: config,
	}
}

func (c *ConfigManager) SetBaseUrl(baseUrl string) {
	c.Signplus.SetBaseUrl(baseUrl)
}

func (c *ConfigManager) SetAccessToken(accessToken string) {
	c.Signplus.SetAccessToken(accessToken)
}

func (c *ConfigManager) UpdateAccessToken(originalValue string, newValue string) {

	if c.Signplus.AccessToken != nil && *c.Signplus.AccessToken == originalValue {
		c.Signplus.SetAccessToken(newValue)
	}
}

func (c *ConfigManager) GetSignplus() *signplusconfig.Config {
	return &c.Signplus
}
