package signplus1

import "encoding/json"

type AddEnvelopeDocumentRequest struct {
	// File to upload in binary format
	File []byte `json:"file,omitempty" xml:"file,omitempty"`
}

func (a AddEnvelopeDocumentRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddEnvelopeDocumentRequest to string"
	}
	return string(jsonData)
}
