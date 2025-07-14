package signplus

import "encoding/json"

type SigningStep struct {
	// List of recipients
	Recipients []Recipient `json:"recipients,omitempty"`
}

func (s *SigningStep) GetRecipients() []Recipient {
	if s == nil {
		return nil
	}
	return s.Recipients
}

func (s *SigningStep) SetRecipients(recipients []Recipient) {
	s.Recipients = recipients
}

func (s SigningStep) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SigningStep to string"
	}
	return string(jsonData)
}
