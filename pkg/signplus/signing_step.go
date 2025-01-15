package signplus

import (
	"encoding/json"
)

type SigningStep struct {
	// List of recipients
	Recipients []Recipient `json:"recipients,omitempty"`
	touched    map[string]bool
}

func (s *SigningStep) GetRecipients() []Recipient {
	if s == nil {
		return nil
	}
	return s.Recipients
}

func (s *SigningStep) SetRecipients(recipients []Recipient) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Recipients"] = true
	s.Recipients = recipients
}

func (s *SigningStep) SetRecipientsNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["Recipients"] = true
	s.Recipients = nil
}
func (s SigningStep) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if s.touched["Recipients"] && s.Recipients == nil {
		data["recipients"] = nil
	} else if s.Recipients != nil {
		data["recipients"] = s.Recipients
	}

	return json.Marshal(data)
}
