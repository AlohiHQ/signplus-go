package signplus

import (
	"encoding/json"
)

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
	Role    *TemplateRecipientRole `json:"role,omitempty"`
	touched map[string]bool
}

func (t *TemplateRecipient) GetId() *string {
	if t == nil {
		return nil
	}
	return t.Id
}

func (t *TemplateRecipient) SetId(id string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Id"] = true
	t.Id = &id
}

func (t *TemplateRecipient) SetIdNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Id"] = true
	t.Id = nil
}

func (t *TemplateRecipient) GetUid() *string {
	if t == nil {
		return nil
	}
	return t.Uid
}

func (t *TemplateRecipient) SetUid(uid string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Uid"] = true
	t.Uid = &uid
}

func (t *TemplateRecipient) SetUidNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Uid"] = true
	t.Uid = nil
}

func (t *TemplateRecipient) GetName() *string {
	if t == nil {
		return nil
	}
	return t.Name
}

func (t *TemplateRecipient) SetName(name string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Name"] = true
	t.Name = &name
}

func (t *TemplateRecipient) SetNameNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Name"] = true
	t.Name = nil
}

func (t *TemplateRecipient) GetEmail() *string {
	if t == nil {
		return nil
	}
	return t.Email
}

func (t *TemplateRecipient) SetEmail(email string) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Email"] = true
	t.Email = &email
}

func (t *TemplateRecipient) SetEmailNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Email"] = true
	t.Email = nil
}

func (t *TemplateRecipient) GetRole() *TemplateRecipientRole {
	if t == nil {
		return nil
	}
	return t.Role
}

func (t *TemplateRecipient) SetRole(role TemplateRecipientRole) {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Role"] = true
	t.Role = &role
}

func (t *TemplateRecipient) SetRoleNil() {
	if t.touched == nil {
		t.touched = map[string]bool{}
	}
	t.touched["Role"] = true
	t.Role = nil
}
func (t TemplateRecipient) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if t.touched["Id"] && t.Id == nil {
		data["id"] = nil
	} else if t.Id != nil {
		data["id"] = t.Id
	}

	if t.touched["Uid"] && t.Uid == nil {
		data["uid"] = nil
	} else if t.Uid != nil {
		data["uid"] = t.Uid
	}

	if t.touched["Name"] && t.Name == nil {
		data["name"] = nil
	} else if t.Name != nil {
		data["name"] = t.Name
	}

	if t.touched["Email"] && t.Email == nil {
		data["email"] = nil
	} else if t.Email != nil {
		data["email"] = t.Email
	}

	if t.touched["Role"] && t.Role == nil {
		data["role"] = nil
	} else if t.Role != nil {
		data["role"] = t.Role
	}

	return json.Marshal(data)
}
