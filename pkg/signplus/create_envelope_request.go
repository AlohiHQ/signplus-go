package signplus

import "encoding/json"

type CreateEnvelopeRequest struct {
	// Name of the envelope
	Name *string `json:"name,omitempty" required:"true" maxLength:"256" minLength:"2" pattern:"^[a-zA-Z0-9][a-zA-Z0-9 ]*[a-zA-Z0-9]$"`
	// Legal level of the envelope (SES is Simple Electronic Signature, QES_EIDAS is Qualified Electronic Signature, QES_ZERTES is Qualified Electronic Signature with Zertes)
	LegalityLevel *EnvelopeLegalityLevel `json:"legality_level,omitempty" required:"true"`
	// Unix timestamp of the expiration date
	ExpiresAt *int64 `json:"expires_at,omitempty"`
	// Comment for the envelope
	Comment *string `json:"comment,omitempty"`
	// Whether the envelope is created in sandbox mode
	Sandbox *bool `json:"sandbox,omitempty"`
}

func (c *CreateEnvelopeRequest) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *CreateEnvelopeRequest) SetName(name string) {
	c.Name = &name
}

func (c *CreateEnvelopeRequest) GetLegalityLevel() *EnvelopeLegalityLevel {
	if c == nil {
		return nil
	}
	return c.LegalityLevel
}

func (c *CreateEnvelopeRequest) SetLegalityLevel(legalityLevel EnvelopeLegalityLevel) {
	c.LegalityLevel = &legalityLevel
}

func (c *CreateEnvelopeRequest) GetExpiresAt() *int64 {
	if c == nil {
		return nil
	}
	return c.ExpiresAt
}

func (c *CreateEnvelopeRequest) SetExpiresAt(expiresAt int64) {
	c.ExpiresAt = &expiresAt
}

func (c *CreateEnvelopeRequest) GetComment() *string {
	if c == nil {
		return nil
	}
	return c.Comment
}

func (c *CreateEnvelopeRequest) SetComment(comment string) {
	c.Comment = &comment
}

func (c *CreateEnvelopeRequest) GetSandbox() *bool {
	if c == nil {
		return nil
	}
	return c.Sandbox
}

func (c *CreateEnvelopeRequest) SetSandbox(sandbox bool) {
	c.Sandbox = &sandbox
}

func (c CreateEnvelopeRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateEnvelopeRequest to string"
	}
	return string(jsonData)
}
