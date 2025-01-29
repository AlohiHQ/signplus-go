package signplus

import (
	"encoding/json"
)

type TemplateSigningStep struct {
	// List of recipients
	Recipients []TemplateRecipient `json:"recipients,omitempty"`
	touched    map[string]bool
}

func (t *TemplateSigningStep) GetRecipients() []TemplateRecipient {
	if t == nil {
		return nil
	}
	return t.Recipients
}

func (t *TemplateSigningStep) SetRecipients(recipients []TemplateRecipient) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Recipients"] = true
	t.Recipients = recipients
}

func (t *TemplateSigningStep) SetRecipientsNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Recipients"] = true
	t.Recipients = nil
}

func (t TemplateSigningStep) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if t.touched["Recipients"] && t.Recipients == nil {
		data["recipients"] = nil
	} else if t.Recipients != nil {
		data["recipients"] = t.Recipients
	}

	return json.Marshal(data)
}

func (t TemplateSigningStep) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TemplateSigningStep to string"
	}
	return string(jsonData)
}
