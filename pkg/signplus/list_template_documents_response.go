package signplus

import (
	"encoding/json"
)

type ListTemplateDocumentsResponse struct {
	Documents []Document `json:"documents,omitempty"`
	touched   map[string]bool
}

func (l *ListTemplateDocumentsResponse) GetDocuments() []Document {
	if l == nil {
		return nil
	}
	return l.Documents
}

func (l *ListTemplateDocumentsResponse) SetDocuments(documents []Document) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Documents"] = true
	l.Documents = documents
}

func (l *ListTemplateDocumentsResponse) SetDocumentsNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Documents"] = true
	l.Documents = nil
}

func (l ListTemplateDocumentsResponse) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if l.touched["Documents"] && l.Documents == nil {
		data["documents"] = nil
	} else if l.Documents != nil {
		data["documents"] = l.Documents
	}

	return json.Marshal(data)
}

func (l ListTemplateDocumentsResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListTemplateDocumentsResponse to string"
	}
	return string(jsonData)
}
