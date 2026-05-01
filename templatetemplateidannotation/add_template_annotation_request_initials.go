package templatetemplateidannotation

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type AddTemplateAnnotationRequestInitials struct {
	ID *param.Nullable[string] `json:"id,omitempty" xml:"id,omitempty"`
}

func (a AddTemplateAnnotationRequestInitials) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddTemplateAnnotationRequestInitials to string"
	}
	return string(jsonData)
}

func (a *AddTemplateAnnotationRequestInitials) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, a)
}
