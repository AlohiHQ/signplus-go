package templateid

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type CreateEnvelopeFromTemplateRequest struct {
	Name    *param.Nullable[string] `json:"name,omitempty" xml:"name,omitempty"`
	Comment *param.Nullable[string] `json:"comment,omitempty" xml:"comment,omitempty"`
	Sandbox *param.Nullable[bool]   `json:"sandbox,omitempty" xml:"sandbox,omitempty"`
}

func (c CreateEnvelopeFromTemplateRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateEnvelopeFromTemplateRequest to string"
	}
	return string(jsonData)
}

func (c *CreateEnvelopeFromTemplateRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, c)
}
