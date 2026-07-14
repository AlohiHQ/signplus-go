package signplus1

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
)

type Recipient struct {
	// Unique identifier of the recipient
	ID *string `json:"id,omitempty" xml:"id,omitempty"`
	// Unique identifier of the user associated with the recipient
	UID *string `json:"uid,omitempty" xml:"uid,omitempty"`
	// Name of the recipient
	Name string `json:"name" xml:"name" required:"true"`
	// Email of the recipient
	Email string `json:"email" xml:"email" required:"true"`
	// Role of the recipient (SIGNER signs the document, RECEIVES_COPY receives a copy of the document, IN_PERSON_SIGNER signs the document in person, SENDER sends the document)
	Role         RecipientRole          `json:"role" xml:"role" required:"true"`
	Verification *RecipientVerification `json:"verification,omitempty" xml:"verification,omitempty"`
}

func (r Recipient) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: Recipient to string"
	}
	return string(jsonData)
}

func (r *Recipient) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias Recipient
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = Recipient(tmp)
	return nil
}
