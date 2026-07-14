package signplus1

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
)

type AddTemplateDocumentRequest struct {
	// File to upload in binary format
	File []byte `json:"file" xml:"file" required:"true"`
}

func (a AddTemplateDocumentRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddTemplateDocumentRequest to string"
	}
	return string(jsonData)
}

func (a *AddTemplateDocumentRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AddTemplateDocumentRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AddTemplateDocumentRequest(tmp)
	return nil
}
