package signplus

import (
	"encoding/json"
)

type Recipient struct {
	// Unique identifier of the recipient
	Id *string `json:"id,omitempty"`
	// Unique identifier of the user associated with the recipient
	Uid *string `json:"uid,omitempty"`
	// Name of the recipient
	Name *string `json:"name,omitempty" required:"true"`
	// Email of the recipient
	Email *string `json:"email,omitempty" required:"true"`
	// Role of the recipient (SIGNER signs the document, RECEIVES_COPY receives a copy of the document, IN_PERSON_SIGNER signs the document in person, SENDER sends the document)
	Role         *RecipientRole         `json:"role,omitempty" required:"true"`
	Verification *RecipientVerification `json:"verification,omitempty"`
	touched      map[string]bool
}

func (r *Recipient) GetId() *string {
	if r == nil {
		return nil
	}
	return r.Id
}

func (r *Recipient) SetId(id string) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Id"] = true
	r.Id = &id
}

func (r *Recipient) SetIdNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Id"] = true
	r.Id = nil
}

func (r *Recipient) GetUid() *string {
	if r == nil {
		return nil
	}
	return r.Uid
}

func (r *Recipient) SetUid(uid string) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Uid"] = true
	r.Uid = &uid
}

func (r *Recipient) SetUidNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Uid"] = true
	r.Uid = nil
}

func (r *Recipient) GetName() *string {
	if r == nil {
		return nil
	}
	return r.Name
}

func (r *Recipient) SetName(name string) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Name"] = true
	r.Name = &name
}

func (r *Recipient) SetNameNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Name"] = true
	r.Name = nil
}

func (r *Recipient) GetEmail() *string {
	if r == nil {
		return nil
	}
	return r.Email
}

func (r *Recipient) SetEmail(email string) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Email"] = true
	r.Email = &email
}

func (r *Recipient) SetEmailNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Email"] = true
	r.Email = nil
}

func (r *Recipient) GetRole() *RecipientRole {
	if r == nil {
		return nil
	}
	return r.Role
}

func (r *Recipient) SetRole(role RecipientRole) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Role"] = true
	r.Role = &role
}

func (r *Recipient) SetRoleNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Role"] = true
	r.Role = nil
}

func (r *Recipient) GetVerification() *RecipientVerification {
	if r == nil {
		return nil
	}
	return r.Verification
}

func (r *Recipient) SetVerification(verification RecipientVerification) {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Verification"] = true
	r.Verification = &verification
}

func (r *Recipient) SetVerificationNil() {
	if r.touched == nil {
		r.touched = map[string]bool{}
	}
	r.touched["Verification"] = true
	r.Verification = nil
}
func (r Recipient) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if r.touched["Id"] && r.Id == nil {
		data["id"] = nil
	} else if r.Id != nil {
		data["id"] = r.Id
	}

	if r.touched["Uid"] && r.Uid == nil {
		data["uid"] = nil
	} else if r.Uid != nil {
		data["uid"] = r.Uid
	}

	if r.touched["Name"] && r.Name == nil {
		data["name"] = nil
	} else if r.Name != nil {
		data["name"] = r.Name
	}

	if r.touched["Email"] && r.Email == nil {
		data["email"] = nil
	} else if r.Email != nil {
		data["email"] = r.Email
	}

	if r.touched["Role"] && r.Role == nil {
		data["role"] = nil
	} else if r.Role != nil {
		data["role"] = r.Role
	}

	if r.touched["Verification"] && r.Verification == nil {
		data["verification"] = nil
	} else if r.Verification != nil {
		data["verification"] = r.Verification
	}

	return json.Marshal(data)
}
