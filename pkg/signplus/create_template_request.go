package signplus

type CreateTemplateRequest struct {
	Name *string `json:"name,omitempty" required:"true"`
}

func (c *CreateTemplateRequest) SetName(name string) {
	c.Name = &name
}

func (c *CreateTemplateRequest) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}
