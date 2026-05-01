package templatetemplateidannotation

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type AddTemplateAnnotationRequestText struct {
	Size             *param.Nullable[float64]   `json:"size,omitempty" xml:"size,omitempty"`
	Color            *param.Nullable[float64]   `json:"color,omitempty" xml:"color,omitempty"`
	Value            *param.Nullable[string]    `json:"value,omitempty" xml:"value,omitempty"`
	Tooltip          *param.Nullable[string]    `json:"tooltip,omitempty" xml:"tooltip,omitempty"`
	DynamicFieldName *param.Nullable[string]    `json:"dynamic_field_name,omitempty" xml:"dynamic_field_name,omitempty"`
	Font             *param.Nullable[TextFont2] `json:"font,omitempty" xml:"font,omitempty"`
}

func (a AddTemplateAnnotationRequestText) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddTemplateAnnotationRequestText to string"
	}
	return string(jsonData)
}

func (a *AddTemplateAnnotationRequestText) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, a)
}
