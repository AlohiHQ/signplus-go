package signplus1

import "encoding/json"

type ListTemplateAnnotationsResponse struct {
	Annotations []Annotation `json:"annotations,omitempty" xml:"annotations,omitempty"`
}

func (l ListTemplateAnnotationsResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListTemplateAnnotationsResponse to string"
	}
	return string(jsonData)
}
