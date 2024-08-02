package signplus

type ListTemplateDocumentsResponse struct {
	Documents []Document `json:"documents,omitempty"`
}

func (l *ListTemplateDocumentsResponse) SetDocuments(documents []Document) {
	l.Documents = documents
}

func (l *ListTemplateDocumentsResponse) GetDocuments() []Document {
	if l == nil {
		return nil
	}
	return l.Documents
}
