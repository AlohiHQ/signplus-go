package signplus1

import "encoding/json"

type ListTemplateDocumentsResponse struct {
	Documents []Document `json:"documents,omitempty" xml:"documents,omitempty"`
}

func (l ListTemplateDocumentsResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListTemplateDocumentsResponse to string"
	}
	return string(jsonData)
}
