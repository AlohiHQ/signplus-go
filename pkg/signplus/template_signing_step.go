package signplus

import "encoding/json"

type TemplateSigningStep struct {
	// List of recipients
	Recipients []TemplateRecipient `json:"recipients,omitempty"`
}

func (t *TemplateSigningStep) GetRecipients() []TemplateRecipient {
	if t == nil {
		return nil
	}
	return t.Recipients
}

func (t *TemplateSigningStep) SetRecipients(recipients []TemplateRecipient) {
	t.Recipients = recipients
}

func (t TemplateSigningStep) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TemplateSigningStep to string"
	}
	return string(jsonData)
}
