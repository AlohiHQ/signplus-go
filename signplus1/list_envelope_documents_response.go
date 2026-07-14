package signplus1

import "encoding/json"

type ListEnvelopeDocumentsResponse struct {
	Documents []Document `json:"documents,omitempty" xml:"documents,omitempty"`
}

func (l ListEnvelopeDocumentsResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListEnvelopeDocumentsResponse to string"
	}
	return string(jsonData)
}
