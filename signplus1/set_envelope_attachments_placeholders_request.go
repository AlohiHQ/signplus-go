package signplus1

import "encoding/json"

type SetEnvelopeAttachmentsPlaceholdersRequest struct {
	Placeholders []AttachmentPlaceholderRequest `json:"placeholders" xml:"placeholders" required:"true"`
}

func (s SetEnvelopeAttachmentsPlaceholdersRequest) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SetEnvelopeAttachmentsPlaceholdersRequest to string"
	}
	return string(jsonData)
}
