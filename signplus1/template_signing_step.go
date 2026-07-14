package signplus1

import "encoding/json"

type TemplateSigningStep struct {
	// List of recipients
	Recipients []TemplateRecipient `json:"recipients,omitempty" xml:"recipients,omitempty"`
}

func (t TemplateSigningStep) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TemplateSigningStep to string"
	}
	return string(jsonData)
}
