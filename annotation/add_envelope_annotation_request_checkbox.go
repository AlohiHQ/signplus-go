package annotation

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type AddEnvelopeAnnotationRequestCheckbox struct {
	Checked *param.Nullable[bool]   `json:"checked,omitempty" xml:"checked,omitempty"`
	Style   *param.Nullable[string] `json:"style,omitempty" xml:"style,omitempty"`
}

func (a AddEnvelopeAnnotationRequestCheckbox) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddEnvelopeAnnotationRequestCheckbox to string"
	}
	return string(jsonData)
}

func (a *AddEnvelopeAnnotationRequestCheckbox) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, a)
}
