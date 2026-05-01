package templatetemplateiddocument

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type AddTemplateDocumentRequest struct {
	File *param.Nullable[[]byte] `json:"file,omitempty" xml:"file,omitempty"`
}

func (a AddTemplateDocumentRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddTemplateDocumentRequest to string"
	}
	return string(jsonData)
}

func (a *AddTemplateDocumentRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, a)
}
