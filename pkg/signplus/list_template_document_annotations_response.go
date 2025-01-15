package signplus

import (
	"encoding/json"
)

type ListTemplateDocumentAnnotationsResponse struct {
	Annotations []Annotation `json:"annotations,omitempty"`
	touched     map[string]bool
}

func (l *ListTemplateDocumentAnnotationsResponse) GetAnnotations() []Annotation {
	if l == nil {
		return nil
	}
	return l.Annotations
}

func (l *ListTemplateDocumentAnnotationsResponse) SetAnnotations(annotations []Annotation) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Annotations"] = true
	l.Annotations = annotations
}

func (l *ListTemplateDocumentAnnotationsResponse) SetAnnotationsNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Annotations"] = true
	l.Annotations = nil
}
func (l ListTemplateDocumentAnnotationsResponse) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if l.touched["Annotations"] && l.Annotations == nil {
		data["annotations"] = nil
	} else if l.Annotations != nil {
		data["annotations"] = l.Annotations
	}

	return json.Marshal(data)
}
