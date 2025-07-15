package signplus

import "encoding/json"

type AddEnvelopeDocumentRequest struct {
	// File to upload in binary format
	File []byte `json:"file,omitempty"`
}

func (a *AddEnvelopeDocumentRequest) GetFile() []byte {
	if a == nil {
		return nil
	}
	return a.File
}

func (a *AddEnvelopeDocumentRequest) SetFile(file []byte) {
	a.File = file
}

func (a AddEnvelopeDocumentRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddEnvelopeDocumentRequest to string"
	}
	return string(jsonData)
}
