package signplus

import "encoding/json"

type SetEnvelopeAttachmentsSettingsRequest struct {
	Settings *AttachmentSettings `json:"settings,omitempty" required:"true"`
}

func (s *SetEnvelopeAttachmentsSettingsRequest) GetSettings() *AttachmentSettings {
	if s == nil {
		return nil
	}
	return s.Settings
}

func (s *SetEnvelopeAttachmentsSettingsRequest) SetSettings(settings AttachmentSettings) {
	s.Settings = &settings
}

func (s SetEnvelopeAttachmentsSettingsRequest) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SetEnvelopeAttachmentsSettingsRequest to string"
	}
	return string(jsonData)
}
