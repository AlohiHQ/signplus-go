package annotation

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type AddEnvelopeAnnotationRequestText struct {
	Size             *param.Nullable[string]    `json:"size,omitempty" xml:"size,omitempty"`
	Color            *param.Nullable[string]    `json:"color,omitempty" xml:"color,omitempty"`
	Value            *param.Nullable[string]    `json:"value,omitempty" xml:"value,omitempty"`
	Tooltip          *param.Nullable[string]    `json:"tooltip,omitempty" xml:"tooltip,omitempty"`
	DynamicFieldName *param.Nullable[string]    `json:"dynamic_field_name,omitempty" xml:"dynamic_field_name,omitempty"`
	Font             *param.Nullable[TextFont1] `json:"font,omitempty" xml:"font,omitempty"`
}

func (a AddEnvelopeAnnotationRequestText) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddEnvelopeAnnotationRequestText to string"
	}
	return string(jsonData)
}

func (a *AddEnvelopeAnnotationRequestText) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, a)
}
