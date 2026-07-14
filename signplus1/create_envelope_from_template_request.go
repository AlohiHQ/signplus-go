package signplus1

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
)

type CreateEnvelopeFromTemplateRequest struct {
	// Name of the envelope
	Name string `json:"name" xml:"name" required:"true" maxLength:"256" minLength:"2" pattern:"^[a-zA-Z0-9][a-zA-Z0-9 ]*[a-zA-Z0-9]$"`
	// Comment for the envelope
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// Whether the envelope is created in sandbox mode
	Sandbox *bool `json:"sandbox,omitempty" xml:"sandbox,omitempty"`
}

func (c CreateEnvelopeFromTemplateRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateEnvelopeFromTemplateRequest to string"
	}
	return string(jsonData)
}

func (c *CreateEnvelopeFromTemplateRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreateEnvelopeFromTemplateRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreateEnvelopeFromTemplateRequest(tmp)
	return nil
}
