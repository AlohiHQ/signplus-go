package signingsteps

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type SigningStepsRecipients1 struct {
	Name         *param.Nullable[string]       `json:"name,omitempty" xml:"name,omitempty"`
	Email        *param.Nullable[string]       `json:"email,omitempty" xml:"email,omitempty"`
	Role         *param.Nullable[string]       `json:"role,omitempty" xml:"role,omitempty"`
	ID           *param.Nullable[string]       `json:"id,omitempty" xml:"id,omitempty"`
	UID          *param.Nullable[string]       `json:"uid,omitempty" xml:"uid,omitempty"`
	Verification *param.Nullable[Verification] `json:"verification,omitempty" xml:"verification,omitempty"`
}

func (s SigningStepsRecipients1) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SigningStepsRecipients1 to string"
	}
	return string(jsonData)
}

func (s *SigningStepsRecipients1) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, s)
}
