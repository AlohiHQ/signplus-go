package signplus

import "encoding/json"

type AttachmentPlaceholder struct {
	// ID of the recipient
	RecipientId *string `json:"recipient_id,omitempty"`
	// ID of the attachment placeholder
	Id *string `json:"id,omitempty"`
	// Name of the attachment placeholder
	Name *string `json:"name,omitempty"`
	// Hint of the attachment placeholder
	Hint *string `json:"hint,omitempty"`
	// Whether the attachment placeholder is required
	Required *bool `json:"required,omitempty"`
	// Whether the attachment placeholder can have multiple files
	Multiple *bool                       `json:"multiple,omitempty"`
	Files    []AttachmentPlaceholderFile `json:"files,omitempty"`
}

func (a *AttachmentPlaceholder) GetRecipientId() *string {
	if a == nil {
		return nil
	}
	return a.RecipientId
}

func (a *AttachmentPlaceholder) SetRecipientId(recipientId string) {
	a.RecipientId = &recipientId
}

func (a *AttachmentPlaceholder) GetId() *string {
	if a == nil {
		return nil
	}
	return a.Id
}

func (a *AttachmentPlaceholder) SetId(id string) {
	a.Id = &id
}

func (a *AttachmentPlaceholder) GetName() *string {
	if a == nil {
		return nil
	}
	return a.Name
}

func (a *AttachmentPlaceholder) SetName(name string) {
	a.Name = &name
}

func (a *AttachmentPlaceholder) GetHint() *string {
	if a == nil {
		return nil
	}
	return a.Hint
}

func (a *AttachmentPlaceholder) SetHint(hint string) {
	a.Hint = &hint
}

func (a *AttachmentPlaceholder) GetRequired() *bool {
	if a == nil {
		return nil
	}
	return a.Required
}

func (a *AttachmentPlaceholder) SetRequired(required bool) {
	a.Required = &required
}

func (a *AttachmentPlaceholder) GetMultiple() *bool {
	if a == nil {
		return nil
	}
	return a.Multiple
}

func (a *AttachmentPlaceholder) SetMultiple(multiple bool) {
	a.Multiple = &multiple
}

func (a *AttachmentPlaceholder) GetFiles() []AttachmentPlaceholderFile {
	if a == nil {
		return nil
	}
	return a.Files
}

func (a *AttachmentPlaceholder) SetFiles(files []AttachmentPlaceholderFile) {
	a.Files = files
}

func (a AttachmentPlaceholder) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AttachmentPlaceholder to string"
	}
	return string(jsonData)
}
