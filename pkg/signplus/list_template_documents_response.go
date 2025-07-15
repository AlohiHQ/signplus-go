package signplus

import "encoding/json"

type ListTemplateDocumentsResponse struct {
	Documents []Document `json:"documents,omitempty"`
}

func (l *ListTemplateDocumentsResponse) GetDocuments() []Document {
	if l == nil {
		return nil
	}
	return l.Documents
}

func (l *ListTemplateDocumentsResponse) SetDocuments(documents []Document) {
	l.Documents = documents
}

func (l ListTemplateDocumentsResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListTemplateDocumentsResponse to string"
	}
	return string(jsonData)
}
