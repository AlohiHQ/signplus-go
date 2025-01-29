package signplus

import (
	"encoding/json"
)

type ListTemplateAnnotationsResponse struct {
	Annotations []Annotation `json:"annotations,omitempty"`
	touched     map[string]bool
}

func (l *ListTemplateAnnotationsResponse) GetAnnotations() []Annotation {
	if l == nil {
		return nil
	}
	return l.Annotations
}

func (l *ListTemplateAnnotationsResponse) SetAnnotations(annotations []Annotation) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Annotations"] = true
	l.Annotations = annotations
}

func (l *ListTemplateAnnotationsResponse) SetAnnotationsNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Annotations"] = true
	l.Annotations = nil
}

func (l ListTemplateAnnotationsResponse) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if l.touched["Annotations"] && l.Annotations == nil {
		data["annotations"] = nil
	} else if l.Annotations != nil {
		data["annotations"] = l.Annotations
	}

	return json.Marshal(data)
}

func (l ListTemplateAnnotationsResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListTemplateAnnotationsResponse to string"
	}
	return string(jsonData)
}
