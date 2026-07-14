package signplus1

import "encoding/json"

type ListEnvelopeDocumentAnnotationsResponse struct {
	Annotations []Annotation `json:"annotations,omitempty" xml:"annotations,omitempty"`
}

func (l ListEnvelopeDocumentAnnotationsResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListEnvelopeDocumentAnnotationsResponse to string"
	}
	return string(jsonData)
}
