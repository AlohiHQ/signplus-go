package signplus

type TemplateRecipient struct {
	// Unique identifier of the recipient
	Id *string `json:"id,omitempty"`
	// Unique identifier of the user associated with the recipient
	Uid *string `json:"uid,omitempty"`
	// Name of the recipient
	Name *string `json:"name,omitempty"`
	// Email of the recipient
	Email *string `json:"email,omitempty"`
	// Role of the recipient (SIGNER signs the document, RECEIVES_COPY receives a copy of the document, IN_PERSON_SIGNER signs the document in person, SENDER sends the document)
	Role *TemplateRecipientRole `json:"role,omitempty"`
}

func (t *TemplateRecipient) SetId(id string) {
	t.Id = &id
}

func (t *TemplateRecipient) GetId() *string {
	if t == nil {
		return nil
	}
	return t.Id
}

func (t *TemplateRecipient) SetUid(uid string) {
	t.Uid = &uid
}

func (t *TemplateRecipient) GetUid() *string {
	if t == nil {
		return nil
	}
	return t.Uid
}

func (t *TemplateRecipient) SetName(name string) {
	t.Name = &name
}

func (t *TemplateRecipient) GetName() *string {
	if t == nil {
		return nil
	}
	return t.Name
}

func (t *TemplateRecipient) SetEmail(email string) {
	t.Email = &email
}

func (t *TemplateRecipient) GetEmail() *string {
	if t == nil {
		return nil
	}
	return t.Email
}

func (t *TemplateRecipient) SetRole(role TemplateRecipientRole) {
	t.Role = &role
}

func (t *TemplateRecipient) GetRole() *TemplateRecipientRole {
	if t == nil {
		return nil
	}
	return t.Role
}
