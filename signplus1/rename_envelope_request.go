package signplus1

import "encoding/json"

type RenameEnvelopeRequest struct {
	// Name of the envelope
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (r RenameEnvelopeRequest) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RenameEnvelopeRequest to string"
	}
	return string(jsonData)
}
