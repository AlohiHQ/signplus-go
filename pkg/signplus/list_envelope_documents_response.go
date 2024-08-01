package signplus

type ListEnvelopeDocumentsResponse struct {
	Documents []Document `json:"documents,omitempty"`
}

func (l *ListEnvelopeDocumentsResponse) SetDocuments(documents []Document) {
	l.Documents = documents
}

func (l *ListEnvelopeDocumentsResponse) GetDocuments() []Document {
	if l == nil {
		return nil
	}
	return l.Documents
}
