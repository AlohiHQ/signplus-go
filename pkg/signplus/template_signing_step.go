package signplus

type TemplateSigningStep struct {
	// List of recipients
	Recipients []TemplateRecipient `json:"recipients,omitempty"`
}

func (t *TemplateSigningStep) SetRecipients(recipients []TemplateRecipient) {
	t.Recipients = recipients
}

func (t *TemplateSigningStep) GetRecipients() []TemplateRecipient {
	if t == nil {
		return nil
	}
	return t.Recipients
}
