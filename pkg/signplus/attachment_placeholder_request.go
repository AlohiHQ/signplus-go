package signplus

import "encoding/json"

type AttachmentPlaceholderRequest struct {
	// ID of the recipient
	RecipientId *string `json:"recipient_id,omitempty" required:"true"`
	// ID of the attachment placeholder
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty" required:"true"`
	// Hint of the attachment placeholder
	Hint *string `json:"hint,omitempty"`
	// Whether the attachment placeholder is required
	Required *bool `json:"required,omitempty" required:"true"`
	Multiple *bool `json:"multiple,omitempty" required:"true"`
}

func (a *AttachmentPlaceholderRequest) GetRecipientId() *string {
	if a == nil {
		return nil
	}
	return a.RecipientId
}

func (a *AttachmentPlaceholderRequest) SetRecipientId(recipientId string) {
	a.RecipientId = &recipientId
}

func (a *AttachmentPlaceholderRequest) GetId() *string {
	if a == nil {
		return nil
	}
	return a.Id
}

func (a *AttachmentPlaceholderRequest) SetId(id string) {
	a.Id = &id
}

func (a *AttachmentPlaceholderRequest) GetName() *string {
	if a == nil {
		return nil
	}
	return a.Name
}

func (a *AttachmentPlaceholderRequest) SetName(name string) {
	a.Name = &name
}

func (a *AttachmentPlaceholderRequest) GetHint() *string {
	if a == nil {
		return nil
	}
	return a.Hint
}

func (a *AttachmentPlaceholderRequest) SetHint(hint string) {
	a.Hint = &hint
}

func (a *AttachmentPlaceholderRequest) GetRequired() *bool {
	if a == nil {
		return nil
	}
	return a.Required
}

func (a *AttachmentPlaceholderRequest) SetRequired(required bool) {
	a.Required = &required
}

func (a *AttachmentPlaceholderRequest) GetMultiple() *bool {
	if a == nil {
		return nil
	}
	return a.Multiple
}

func (a *AttachmentPlaceholderRequest) SetMultiple(multiple bool) {
	a.Multiple = &multiple
}

func (a AttachmentPlaceholderRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AttachmentPlaceholderRequest to string"
	}
	return string(jsonData)
}
