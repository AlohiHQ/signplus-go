package signplus1

import "encoding/json"

type SigningStep struct {
	// List of recipients
	Recipients []Recipient `json:"recipients,omitempty" xml:"recipients,omitempty"`
}

func (s SigningStep) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SigningStep to string"
	}
	return string(jsonData)
}
