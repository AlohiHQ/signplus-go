package signplus1

import "encoding/json"

type AttachmentPlaceholdersPerRecipient struct {
	// ID of the recipient
	RecipientID *string `json:"recipient_id,omitempty" xml:"recipient_id,omitempty"`
	// Name of the recipient
	RecipientName *string                 `json:"recipient_name,omitempty" xml:"recipient_name,omitempty"`
	Placeholders  []AttachmentPlaceholder `json:"placeholders,omitempty" xml:"placeholders,omitempty"`
}

func (a AttachmentPlaceholdersPerRecipient) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AttachmentPlaceholdersPerRecipient to string"
	}
	return string(jsonData)
}
