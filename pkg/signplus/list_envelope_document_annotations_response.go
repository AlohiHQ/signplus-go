package signplus

import "encoding/json"

type ListEnvelopeDocumentAnnotationsResponse struct {
	Annotations []Annotation `json:"annotations,omitempty"`
}

func (l *ListEnvelopeDocumentAnnotationsResponse) GetAnnotations() []Annotation {
	if l == nil {
		return nil
	}
	return l.Annotations
}

func (l *ListEnvelopeDocumentAnnotationsResponse) SetAnnotations(annotations []Annotation) {
	l.Annotations = annotations
}

func (l ListEnvelopeDocumentAnnotationsResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListEnvelopeDocumentAnnotationsResponse to string"
	}
	return string(jsonData)
}
