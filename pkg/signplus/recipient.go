package signplus

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
}

func (r *Recipient) SetId(id string) {
	r.Id = &id
}

func (r *Recipient) GetId() *string {
	if r == nil {
		return nil
	}
	return r.Id
}

func (r *Recipient) SetUid(uid string) {
	r.Uid = &uid
}

func (r *Recipient) GetUid() *string {
	if r == nil {
		return nil
	}
	return r.Uid
}

func (r *Recipient) SetName(name string) {
	r.Name = &name
}

func (r *Recipient) GetName() *string {
	if r == nil {
		return nil
	}
	return r.Name
}

func (r *Recipient) SetEmail(email string) {
	r.Email = &email
}

func (r *Recipient) GetEmail() *string {
	if r == nil {
		return nil
	}
	return r.Email
}

func (r *Recipient) SetRole(role RecipientRole) {
	r.Role = &role
}

func (r *Recipient) GetRole() *RecipientRole {
	if r == nil {
		return nil
	}
	return r.Role
}

func (r *Recipient) SetVerification(verification RecipientVerification) {
	r.Verification = &verification
}

func (r *Recipient) GetVerification() *RecipientVerification {
	if r == nil {
		return nil
	}
	return r.Verification
}
