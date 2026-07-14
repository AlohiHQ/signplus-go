package signplus1

import "encoding/json"

type AttachmentPlaceholder struct {
	// ID of the recipient
	RecipientID *string `json:"recipient_id,omitempty" xml:"recipient_id,omitempty"`
	// ID of the attachment placeholder
	ID *string `json:"id,omitempty" xml:"id,omitempty"`
	// Name of the attachment placeholder
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Hint of the attachment placeholder
	Hint *string `json:"hint,omitempty" xml:"hint,omitempty"`
	// Whether the attachment placeholder is required
	Required *bool `json:"required,omitempty" xml:"required,omitempty"`
	// Whether the attachment placeholder can have multiple files
	Multiple *bool                       `json:"multiple,omitempty" xml:"multiple,omitempty"`
	Files    []AttachmentPlaceholderFile `json:"files,omitempty" xml:"files,omitempty"`
}

func (a AttachmentPlaceholder) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AttachmentPlaceholder to string"
	}
	return string(jsonData)
}
