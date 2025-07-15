package signplus

import "encoding/json"

type AddTemplateDocumentRequest struct {
	// File to upload in binary format
	File []byte `json:"file,omitempty" required:"true"`
}

func (a *AddTemplateDocumentRequest) GetFile() []byte {
	if a == nil {
		return nil
	}
	return a.File
}

func (a *AddTemplateDocumentRequest) SetFile(file []byte) {
	a.File = file
}

func (a AddTemplateDocumentRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddTemplateDocumentRequest to string"
	}
	return string(jsonData)
}
