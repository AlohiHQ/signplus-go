package signplus

import "encoding/json"

type EnvelopeAttachments struct {
	Settings   *AttachmentSettings                  `json:"settings,omitempty"`
	Recipients []AttachmentPlaceholdersPerRecipient `json:"recipients,omitempty"`
}

func (e *EnvelopeAttachments) GetSettings() *AttachmentSettings {
	if e == nil {
		return nil
	}
	return e.Settings
}

func (e *EnvelopeAttachments) SetSettings(settings AttachmentSettings) {
	e.Settings = &settings
}

func (e *EnvelopeAttachments) GetRecipients() []AttachmentPlaceholdersPerRecipient {
	if e == nil {
		return nil
	}
	return e.Recipients
}

func (e *EnvelopeAttachments) SetRecipients(recipients []AttachmentPlaceholdersPerRecipient) {
	e.Recipients = recipients
}

func (e EnvelopeAttachments) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EnvelopeAttachments to string"
	}
	return string(jsonData)
}
