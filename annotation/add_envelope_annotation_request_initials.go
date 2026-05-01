package annotation

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type AddEnvelopeAnnotationRequestInitials struct {
	ID *param.Nullable[string] `json:"id,omitempty" xml:"id,omitempty"`
}

func (a AddEnvelopeAnnotationRequestInitials) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddEnvelopeAnnotationRequestInitials to string"
	}
	return string(jsonData)
}

func (a *AddEnvelopeAnnotationRequestInitials) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, a)
}
