package signplus

import "encoding/json"

type ListEnvelopeDocumentsResponse struct {
	Documents []Document `json:"documents,omitempty"`
}

func (l *ListEnvelopeDocumentsResponse) GetDocuments() []Document {
	if l == nil {
		return nil
	}
	return l.Documents
}

func (l *ListEnvelopeDocumentsResponse) SetDocuments(documents []Document) {
	l.Documents = documents
}

func (l ListEnvelopeDocumentsResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListEnvelopeDocumentsResponse to string"
	}
	return string(jsonData)
}
