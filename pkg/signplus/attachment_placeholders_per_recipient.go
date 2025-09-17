package signplus

import "encoding/json"

type AttachmentPlaceholdersPerRecipient struct {
	// ID of the recipient
	RecipientId *string `json:"recipient_id,omitempty"`
	// Name of the recipient
	RecipientName *string                 `json:"recipient_name,omitempty"`
	Placeholders  []AttachmentPlaceholder `json:"placeholders,omitempty"`
}

func (a *AttachmentPlaceholdersPerRecipient) GetRecipientId() *string {
	if a == nil {
		return nil
	}
	return a.RecipientId
}

func (a *AttachmentPlaceholdersPerRecipient) SetRecipientId(recipientId string) {
	a.RecipientId = &recipientId
}

func (a *AttachmentPlaceholdersPerRecipient) GetRecipientName() *string {
	if a == nil {
		return nil
	}
	return a.RecipientName
}

func (a *AttachmentPlaceholdersPerRecipient) SetRecipientName(recipientName string) {
	a.RecipientName = &recipientName
}

func (a *AttachmentPlaceholdersPerRecipient) GetPlaceholders() []AttachmentPlaceholder {
	if a == nil {
		return nil
	}
	return a.Placeholders
}

func (a *AttachmentPlaceholdersPerRecipient) SetPlaceholders(placeholders []AttachmentPlaceholder) {
	a.Placeholders = placeholders
}

func (a AttachmentPlaceholdersPerRecipient) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AttachmentPlaceholdersPerRecipient to string"
	}
	return string(jsonData)
}
