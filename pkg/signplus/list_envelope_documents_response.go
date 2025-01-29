package signplus

import (
	"encoding/json"
)

type ListEnvelopeDocumentsResponse struct {
	Documents []Document `json:"documents,omitempty"`
	touched   map[string]bool
}

func (l *ListEnvelopeDocumentsResponse) GetDocuments() []Document {
	if l == nil {
		return nil
	}
	return l.Documents
}

func (l *ListEnvelopeDocumentsResponse) SetDocuments(documents []Document) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Documents"] = true
	l.Documents = documents
}

func (l *ListEnvelopeDocumentsResponse) SetDocumentsNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Documents"] = true
	l.Documents = nil
}

func (l ListEnvelopeDocumentsResponse) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if l.touched["Documents"] && l.Documents == nil {
		data["documents"] = nil
	} else if l.Documents != nil {
		data["documents"] = l.Documents
	}

	return json.Marshal(data)
}

func (l ListEnvelopeDocumentsResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListEnvelopeDocumentsResponse to string"
	}
	return string(jsonData)
}
