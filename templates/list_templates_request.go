package templates

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type ListTemplatesRequest struct {
	Name       *param.Nullable[string]   `json:"name,omitempty" xml:"name,omitempty"`
	Tags       *param.Nullable[[]string] `json:"tags,omitempty" xml:"tags,omitempty"`
	Ids        *param.Nullable[[]string] `json:"ids,omitempty" xml:"ids,omitempty"`
	First      *param.Nullable[float64]  `json:"first,omitempty" xml:"first,omitempty"`
	Last       *param.Nullable[float64]  `json:"last,omitempty" xml:"last,omitempty"`
	After      *param.Nullable[string]   `json:"after,omitempty" xml:"after,omitempty"`
	Before     *param.Nullable[string]   `json:"before,omitempty" xml:"before,omitempty"`
	OrderField *param.Nullable[string]   `json:"order_field,omitempty" xml:"order_field,omitempty"`
	Ascending  *param.Nullable[bool]     `json:"ascending,omitempty" xml:"ascending,omitempty"`
}

func (l ListTemplatesRequest) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListTemplatesRequest to string"
	}
	return string(jsonData)
}

func (l *ListTemplatesRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, l)
}
