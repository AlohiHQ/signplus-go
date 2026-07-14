package signplus1

import "encoding/json"

type ListTemplateDocumentAnnotationsResponse struct {
	Annotations []Annotation `json:"annotations,omitempty" xml:"annotations,omitempty"`
}

func (l ListTemplateDocumentAnnotationsResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListTemplateDocumentAnnotationsResponse to string"
	}
	return string(jsonData)
}
