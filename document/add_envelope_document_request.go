package document

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type AddEnvelopeDocumentRequest struct {
	File *param.Nullable[[]byte] `json:"file,omitempty" xml:"file,omitempty"`
}

func (a AddEnvelopeDocumentRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddEnvelopeDocumentRequest to string"
	}
	return string(jsonData)
}

func (a *AddEnvelopeDocumentRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, a)
}
