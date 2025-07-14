package signplus

import "encoding/json"

type ListTemplateDocumentAnnotationsResponse struct {
	Annotations []Annotation `json:"annotations,omitempty"`
}

func (l *ListTemplateDocumentAnnotationsResponse) GetAnnotations() []Annotation {
	if l == nil {
		return nil
	}
	return l.Annotations
}

func (l *ListTemplateDocumentAnnotationsResponse) SetAnnotations(annotations []Annotation) {
	l.Annotations = annotations
}

func (l ListTemplateDocumentAnnotationsResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListTemplateDocumentAnnotationsResponse to string"
	}
	return string(jsonData)
}
