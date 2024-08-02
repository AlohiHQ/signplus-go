package signplusconfig

type Config struct {
	BaseUrl     *string
	AccessToken *string
	HookParams  map[string]string
}

func NewConfig() Config {
	baseUrl := DEFAULT_ENVIRONMENT
	newConfig := Config{
		BaseUrl:    &baseUrl,
		HookParams: make(map[string]string),
	}

	return newConfig
}

func (c *Config) SetBaseUrl(baseUrl string) {
	c.BaseUrl = &baseUrl
}

func (c *Config) GetBaseUrl() string {
	return *c.BaseUrl
}

func (c *Config) SetAccessToken(accessToken string) {
	c.AccessToken = &accessToken
}

func (c *Config) GetAccessToken() string {
	return *c.AccessToken
}
