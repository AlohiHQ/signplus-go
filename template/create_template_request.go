package template

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type CreateTemplateRequest struct {
	Name *param.Nullable[string] `json:"name,omitempty" xml:"name,omitempty"`
}

func (c CreateTemplateRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateTemplateRequest to string"
	}
	return string(jsonData)
}

func (c *CreateTemplateRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, c)
}
