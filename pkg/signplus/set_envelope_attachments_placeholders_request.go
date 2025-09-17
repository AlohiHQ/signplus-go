package signplus

import "encoding/json"

type SetEnvelopeAttachmentsPlaceholdersRequest struct {
	Placeholders []AttachmentPlaceholderRequest `json:"placeholders,omitempty" required:"true"`
}

func (s *SetEnvelopeAttachmentsPlaceholdersRequest) GetPlaceholders() []AttachmentPlaceholderRequest {
	if s == nil {
		return nil
	}
	return s.Placeholders
}

func (s *SetEnvelopeAttachmentsPlaceholdersRequest) SetPlaceholders(placeholders []AttachmentPlaceholderRequest) {
	s.Placeholders = placeholders
}

func (s SetEnvelopeAttachmentsPlaceholdersRequest) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SetEnvelopeAttachmentsPlaceholdersRequest to string"
	}
	return string(jsonData)
}
