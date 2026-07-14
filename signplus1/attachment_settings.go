package signplus1

import "encoding/json"

type AttachmentSettings struct {
	// Whether the attachment is visible to the recipients
	VisibleToRecipients *bool `json:"visible_to_recipients,omitempty" xml:"visible_to_recipients,omitempty"`
}

func (a AttachmentSettings) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AttachmentSettings to string"
	}
	return string(jsonData)
}
