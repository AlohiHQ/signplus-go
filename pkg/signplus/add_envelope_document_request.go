package signplus

import (
	"encoding/json"
)

type AddEnvelopeDocumentRequest struct {
	// File to upload in binary format
	File    *any `json:"file,omitempty"`
	touched map[string]bool
}

func (a *AddEnvelopeDocumentRequest) GetFile() *any {
	if a == nil {
		return nil
	}
	return a.File
}

func (a *AddEnvelopeDocumentRequest) SetFile(file any) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["File"] = true
	a.File = &file
}

func (a *AddEnvelopeDocumentRequest) SetFileNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["File"] = true
	a.File = nil
}

func (a AddEnvelopeDocumentRequest) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if a.touched["File"] && a.File == nil {
		data["file"] = nil
	} else if a.File != nil {
		data["file"] = a.File
	}

	return json.Marshal(data)
}

func (a AddEnvelopeDocumentRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddEnvelopeDocumentRequest to string"
	}
	return string(jsonData)
}
