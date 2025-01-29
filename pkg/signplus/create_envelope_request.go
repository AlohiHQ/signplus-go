package signplus

import (
	"encoding/json"
)

type CreateEnvelopeRequest struct {
	// Name of the envelope
	Name *string `json:"name,omitempty" required:"true"`
	// Legal level of the envelope (SES is Simple Electronic Signature, QES_EIDAS is Qualified Electronic Signature, QES_ZERTES is Qualified Electronic Signature with Zertes)
	LegalityLevel *EnvelopeLegalityLevel `json:"legality_level,omitempty" required:"true"`
	// Unix timestamp of the expiration date
	ExpiresAt *int64 `json:"expires_at,omitempty"`
	// Comment for the envelope
	Comment *string `json:"comment,omitempty"`
	// Whether the envelope is created in sandbox mode
	Sandbox *bool `json:"sandbox,omitempty"`
	touched map[string]bool
}

func (c *CreateEnvelopeRequest) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *CreateEnvelopeRequest) SetName(name string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Name"] = true
	c.Name = &name
}

func (c *CreateEnvelopeRequest) SetNameNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Name"] = true
	c.Name = nil
}

func (c *CreateEnvelopeRequest) GetLegalityLevel() *EnvelopeLegalityLevel {
	if c == nil {
		return nil
	}
	return c.LegalityLevel
}

func (c *CreateEnvelopeRequest) SetLegalityLevel(legalityLevel EnvelopeLegalityLevel) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["LegalityLevel"] = true
	c.LegalityLevel = &legalityLevel
}

func (c *CreateEnvelopeRequest) SetLegalityLevelNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["LegalityLevel"] = true
	c.LegalityLevel = nil
}

func (c *CreateEnvelopeRequest) GetExpiresAt() *int64 {
	if c == nil {
		return nil
	}
	return c.ExpiresAt
}

func (c *CreateEnvelopeRequest) SetExpiresAt(expiresAt int64) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["ExpiresAt"] = true
	c.ExpiresAt = &expiresAt
}

func (c *CreateEnvelopeRequest) SetExpiresAtNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["ExpiresAt"] = true
	c.ExpiresAt = nil
}

func (c *CreateEnvelopeRequest) GetComment() *string {
	if c == nil {
		return nil
	}
	return c.Comment
}

func (c *CreateEnvelopeRequest) SetComment(comment string) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Comment"] = true
	c.Comment = &comment
}

func (c *CreateEnvelopeRequest) SetCommentNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Comment"] = true
	c.Comment = nil
}

func (c *CreateEnvelopeRequest) GetSandbox() *bool {
	if c == nil {
		return nil
	}
	return c.Sandbox
}

func (c *CreateEnvelopeRequest) SetSandbox(sandbox bool) {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Sandbox"] = true
	c.Sandbox = &sandbox
}

func (c *CreateEnvelopeRequest) SetSandboxNil() {
	if c.touched == nil {
		c.touched = map[string]bool{}
	}
	c.touched["Sandbox"] = true
	c.Sandbox = nil
}

func (c CreateEnvelopeRequest) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if c.touched["Name"] && c.Name == nil {
		data["name"] = nil
	} else if c.Name != nil {
		data["name"] = c.Name
	}

	if c.touched["LegalityLevel"] && c.LegalityLevel == nil {
		data["legality_level"] = nil
	} else if c.LegalityLevel != nil {
		data["legality_level"] = c.LegalityLevel
	}

	if c.touched["ExpiresAt"] && c.ExpiresAt == nil {
		data["expires_at"] = nil
	} else if c.ExpiresAt != nil {
		data["expires_at"] = c.ExpiresAt
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

func (c CreateEnvelopeRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateEnvelopeRequest to string"
	}
	return string(jsonData)
}
