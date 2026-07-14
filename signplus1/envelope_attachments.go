package signplus1

import "encoding/json"

type EnvelopeAttachments struct {
	Settings   *AttachmentSettings                  `json:"settings,omitempty" xml:"settings,omitempty"`
	Recipients []AttachmentPlaceholdersPerRecipient `json:"recipients,omitempty" xml:"recipients,omitempty"`
}

func (e EnvelopeAttachments) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EnvelopeAttachments to string"
	}
	return string(jsonData)
}
