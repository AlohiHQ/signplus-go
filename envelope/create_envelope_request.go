package envelope

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type CreateEnvelopeRequest struct {
	Name          *param.Nullable[string] `json:"name,omitempty" xml:"name,omitempty"`
	LegalityLevel *param.Nullable[string] `json:"legality_level,omitempty" xml:"legality_level,omitempty"`
	ExpiresAt     *param.Nullable[string] `json:"expires_at,omitempty" xml:"expires_at,omitempty"`
	Comment       *param.Nullable[string] `json:"comment,omitempty" xml:"comment,omitempty"`
	Sandbox       *param.Nullable[bool]   `json:"sandbox,omitempty" xml:"sandbox,omitempty"`
}

func (c CreateEnvelopeRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateEnvelopeRequest to string"
	}
	return string(jsonData)
}

func (c *CreateEnvelopeRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, c)
}
