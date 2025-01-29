package signplus

import (
	"encoding/json"
)

type CreateEnvelopeFromTemplateRequest struct {
	// Name of the envelope
	Name *string `json:"name,omitempty" required:"true"`
	// Comment for the envelope
	Comment *string `json:"comment,omitempty"`
	// Whether the envelope is created in sandbox mode
	Sandbox *bool `json:"sandbox,omitempty"`
	touched map[string]bool
}

func (c *CreateEnvelopeFromTemplateRequest) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *CreateEnvelopeFromTemplateRequest) SetName(name string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Name"] = true
	c.Name = &name
}

func (c *CreateEnvelopeFromTemplateRequest) SetNameNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Name"] = true
	c.Name = nil
}

func (c *CreateEnvelopeFromTemplateRequest) GetComment() *string {
	if c == nil {
		return nil
	}
	return c.Comment
}

func (c *CreateEnvelopeFromTemplateRequest) SetComment(comment string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Comment"] = true
	c.Comment = &comment
}

func (c *CreateEnvelopeFromTemplateRequest) SetCommentNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Comment"] = true
	c.Comment = nil
}

func (c *CreateEnvelopeFromTemplateRequest) GetSandbox() *bool {
	if c == nil {
		return nil
	}
	return c.Sandbox
}

func (c *CreateEnvelopeFromTemplateRequest) SetSandbox(sandbox bool) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Sandbox"] = true
	c.Sandbox = &sandbox
}

func (c *CreateEnvelopeFromTemplateRequest) SetSandboxNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Sandbox"] = true
	c.Sandbox = nil
}

func (c CreateEnvelopeFromTemplateRequest) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Name"] && c.Name == nil {
		data["name"] = nil
	} else if c.Name != nil {
		data["name"] = c.Name
	}

	if c.touched["Comment"] && c.Comment == nil {
		data["comment"] = nil
	} else if c.Comment != nil {
		data["comment"] = c.Comment
	}

	if c.touched["Sandbox"] && c.Sandbox == nil {
		data["sandbox"] = nil
	} else if c.Sandbox != nil {
		data["sandbox"] = c.Sandbox
	}

	return json.Marshal(data)
}

func (c CreateEnvelopeFromTemplateRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateEnvelopeFromTemplateRequest to string"
	}
	return string(jsonData)
}
