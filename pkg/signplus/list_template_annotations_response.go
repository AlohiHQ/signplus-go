package signplus

import "encoding/json"

type ListTemplateAnnotationsResponse struct {
	Annotations []Annotation `json:"annotations,omitempty"`
}

func (l *ListTemplateAnnotationsResponse) GetAnnotations() []Annotation {
	if l == nil {
		return nil
	}
	return l.Annotations
}

func (l *ListTemplateAnnotationsResponse) SetAnnotations(annotations []Annotation) {
	l.Annotations = annotations
}

func (l ListTemplateAnnotationsResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListTemplateAnnotationsResponse to string"
	}
	return string(jsonData)
}
