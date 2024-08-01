package signplus

type SigningStep struct {
	// List of recipients
	Recipients []Recipient `json:"recipients,omitempty"`
}

func (s *SigningStep) SetRecipients(recipients []Recipient) {
	s.Recipients = recipients
}

func (s *SigningStep) GetRecipients() []Recipient {
	if s == nil {
		return nil
	}
	return s.Recipients
}
