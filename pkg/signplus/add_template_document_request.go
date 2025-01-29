package signplus

import (
	"encoding/json"
)

type AddTemplateDocumentRequest struct {
	// File to upload in binary format
	File    *any `json:"file,omitempty" required:"true"`
	touched map[string]bool
}

func (a *AddTemplateDocumentRequest) GetFile() *any {
	if a == nil {
		return nil
	}
	return a.File
}

func (a *AddTemplateDocumentRequest) SetFile(file any) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["File"] = true
	a.File = &file
}

func (a *AddTemplateDocumentRequest) SetFileNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["File"] = true
	a.File = nil
}

func (a AddTemplateDocumentRequest) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if a.touched["File"] && a.File == nil {
		data["file"] = nil
	} else if a.File != nil {
		data["file"] = a.File
	}

	return json.Marshal(data)
}

func (a AddTemplateDocumentRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddTemplateDocumentRequest to string"
	}
	return string(jsonData)
}
