package signplus

type CreateEnvelopeFromTemplateRequest struct {
	// Name of the envelope
	Name *string `json:"name,omitempty"`
	// Comment for the envelope
	Comment *string `json:"comment,omitempty"`
	// Whether the envelope is created in sandbox mode
	Sandbox *bool `json:"sandbox,omitempty"`
}

func (c *CreateEnvelopeFromTemplateRequest) SetName(name string) {
	c.Name = &name
}

func (c *CreateEnvelopeFromTemplateRequest) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *CreateEnvelopeFromTemplateRequest) SetComment(comment string) {
	c.Comment = &comment
}

func (c *CreateEnvelopeFromTemplateRequest) GetComment() *string {
	if c == nil {
		return nil
	}
	return c.Comment
}

func (c *CreateEnvelopeFromTemplateRequest) SetSandbox(sandbox bool) {
	c.Sandbox = &sandbox
}

func (c *CreateEnvelopeFromTemplateRequest) GetSandbox() *bool {
	if c == nil {
		return nil
	}
	return c.Sandbox
}
