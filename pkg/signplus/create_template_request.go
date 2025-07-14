package signplus

import "encoding/json"

type CreateTemplateRequest struct {
	Name *string `json:"name,omitempty" required:"true" maxLength:"256" minLength:"2" pattern:"^[a-zA-Z0-9][a-zA-Z0-9 ]*[a-zA-Z0-9]$"`
}

func (c *CreateTemplateRequest) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *CreateTemplateRequest) SetName(name string) {
	c.Name = &name
}

func (c CreateTemplateRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateTemplateRequest to string"
	}
	return string(jsonData)
}
