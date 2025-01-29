package signplus

import (
	"encoding/json"
)

type CreateTemplateRequest struct {
	Name    *string `json:"name,omitempty" required:"true"`
	touched map[string]bool
}

func (c *CreateTemplateRequest) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *CreateTemplateRequest) SetName(name string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Name"] = true
	c.Name = &name
}

func (c *CreateTemplateRequest) SetNameNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Name"] = true
	c.Name = nil
}

func (c CreateTemplateRequest) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Name"] && c.Name == nil {
		data["name"] = nil
	} else if c.Name != nil {
		data["name"] = c.Name
	}

	return json.Marshal(data)
}

func (c CreateTemplateRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateTemplateRequest to string"
	}
	return string(jsonData)
}
