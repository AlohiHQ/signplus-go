package signplus

type ListTemplateDocumentAnnotationsResponse struct {
	Annotations []Annotation `json:"annotations,omitempty"`
}

func (l *ListTemplateDocumentAnnotationsResponse) SetAnnotations(annotations []Annotation) {
	l.Annotations = annotations
}

func (l *ListTemplateDocumentAnnotationsResponse) GetAnnotations() []Annotation {
	if l == nil {
		return nil
	}
	return l.Annotations
}
