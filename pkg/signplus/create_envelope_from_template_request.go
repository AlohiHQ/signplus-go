package signplus

import "encoding/json"

type CreateEnvelopeFromTemplateRequest struct {
	// Name of the envelope
	Name *string `json:"name,omitempty" required:"true" maxLength:"256" minLength:"2" pattern:"^[a-zA-Z0-9][a-zA-Z0-9 ]*[a-zA-Z0-9]$"`
	// Comment for the envelope
	Comment *string `json:"comment,omitempty"`
	// Whether the envelope is created in sandbox mode
	Sandbox *bool `json:"sandbox,omitempty"`
}

func (c *CreateEnvelopeFromTemplateRequest) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *CreateEnvelopeFromTemplateRequest) SetName(name string) {
	c.Name = &name
}

func (c *CreateEnvelopeFromTemplateRequest) GetComment() *string {
	if c == nil {
		return nil
	}
	return c.Comment
}

func (c *CreateEnvelopeFromTemplateRequest) SetComment(comment string) {
	c.Comment = &comment
}

func (c *CreateEnvelopeFromTemplateRequest) GetSandbox() *bool {
	if c == nil {
		return nil
	}
	return c.Sandbox
}

func (c *CreateEnvelopeFromTemplateRequest) SetSandbox(sandbox bool) {
	c.Sandbox = &sandbox
}

func (c CreateEnvelopeFromTemplateRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateEnvelopeFromTemplateRequest to string"
	}
	return string(jsonData)
}
