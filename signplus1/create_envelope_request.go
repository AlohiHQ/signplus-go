package signplus1

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
)

type CreateEnvelopeRequest struct {
	// Name of the envelope
	Name string `json:"name" xml:"name" required:"true" maxLength:"256" minLength:"2" pattern:"^[a-zA-Z0-9][a-zA-Z0-9 ]*[a-zA-Z0-9]$"`
	// Legal level of the envelope (SES is Simple Electronic Signature, QES_EIDAS is Qualified Electronic Signature, QES_ZERTES is Qualified Electronic Signature with Zertes)
	LegalityLevel EnvelopeLegalityLevel `json:"legality_level" xml:"legality_level" required:"true"`
	// Unix timestamp of the expiration date
	ExpiresAt *int64 `json:"expires_at,omitempty" xml:"expires_at,omitempty"`
	// Comment for the envelope
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// Whether the envelope is created in sandbox mode
	Sandbox *bool `json:"sandbox,omitempty" xml:"sandbox,omitempty"`
}

func (c CreateEnvelopeRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateEnvelopeRequest to string"
	}
	return string(jsonData)
}

func (c *CreateEnvelopeRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreateEnvelopeRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreateEnvelopeRequest(tmp)
	return nil
}
