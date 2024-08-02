package signplus

type ListTemplateAnnotationsResponse struct {
	Annotations []Annotation `json:"annotations,omitempty"`
}

func (l *ListTemplateAnnotationsResponse) SetAnnotations(annotations []Annotation) {
	l.Annotations = annotations
}

func (l *ListTemplateAnnotationsResponse) GetAnnotations() []Annotation {
	if l == nil {
		return nil
	}
	return l.Annotations
}
