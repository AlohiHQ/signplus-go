package templatetemplateidannotation

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type AddTemplateAnnotationRequestCheckbox struct {
	Checked *param.Nullable[string] `json:"checked,omitempty" xml:"checked,omitempty"`
	Style   *param.Nullable[string] `json:"style,omitempty" xml:"style,omitempty"`
}

func (a AddTemplateAnnotationRequestCheckbox) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddTemplateAnnotationRequestCheckbox to string"
	}
	return string(jsonData)
}

func (a *AddTemplateAnnotationRequestCheckbox) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, a)
}
