package signplus

import (
	"encoding/json"
)

type ListEnvelopeDocumentAnnotationsResponse struct {
	Annotations []Annotation `json:"annotations,omitempty"`
	touched     map[string]bool
}

func (l *ListEnvelopeDocumentAnnotationsResponse) GetAnnotations() []Annotation {
	if l == nil {
		return nil
	}
	return l.Annotations
}

func (l *ListEnvelopeDocumentAnnotationsResponse) SetAnnotations(annotations []Annotation) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Annotations"] = true
	l.Annotations = annotations
}

func (l *ListEnvelopeDocumentAnnotationsResponse) SetAnnotationsNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Annotations"] = true
	l.Annotations = nil
}
func (l ListEnvelopeDocumentAnnotationsResponse) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if l.touched["Annotations"] && l.Annotations == nil {
		data["annotations"] = nil
	} else if l.Annotations != nil {
		data["annotations"] = l.Annotations
	}

	return json.Marshal(data)
}
