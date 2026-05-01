package annotation

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type AddEnvelopeAnnotationRequestSignature struct {
	ID *param.Nullable[string] `json:"id,omitempty" xml:"id,omitempty"`
}

func (a AddEnvelopeAnnotationRequestSignature) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddEnvelopeAnnotationRequestSignature to string"
	}
	return string(jsonData)
}

func (a *AddEnvelopeAnnotationRequestSignature) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, a)
}
