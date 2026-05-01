package templatetemplateidannotation

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type AddTemplateAnnotationRequestSignature struct {
	ID *param.Nullable[string] `json:"id,omitempty" xml:"id,omitempty"`
}

func (a AddTemplateAnnotationRequestSignature) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddTemplateAnnotationRequestSignature to string"
	}
	return string(jsonData)
}

func (a *AddTemplateAnnotationRequestSignature) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, a)
}
