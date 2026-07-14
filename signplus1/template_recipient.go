package signplus1

import "encoding/json"

type TemplateRecipient struct {
	// Unique identifier of the recipient
	ID *string `json:"id,omitempty" xml:"id,omitempty"`
	// Unique identifier of the user associated with the recipient
	UID *string `json:"uid,omitempty" xml:"uid,omitempty"`
	// Name of the recipient
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Email of the recipient
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// Role of the recipient (SIGNER signs the document, RECEIVES_COPY receives a copy of the document, IN_PERSON_SIGNER signs the document in person, SENDER sends the document)
	Role *TemplateRecipientRole `json:"role,omitempty" xml:"role,omitempty"`
}

func (t TemplateRecipient) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TemplateRecipient to string"
	}
	return string(jsonData)
}
