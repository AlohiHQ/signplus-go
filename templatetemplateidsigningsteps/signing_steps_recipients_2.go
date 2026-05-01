package templatetemplateidsigningsteps

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type SigningStepsRecipients2 struct {
	ID    *param.Nullable[string] `json:"id,omitempty" xml:"id,omitempty"`
	UID   *param.Nullable[string] `json:"uid,omitempty" xml:"uid,omitempty"`
	Name  *param.Nullable[string] `json:"name,omitempty" xml:"name,omitempty"`
	Email *param.Nullable[string] `json:"email,omitempty" xml:"email,omitempty"`
	Role  *param.Nullable[string] `json:"role,omitempty" xml:"role,omitempty"`
}

func (s SigningStepsRecipients2) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SigningStepsRecipients2 to string"
	}
	return string(jsonData)
}

func (s *SigningStepsRecipients2) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, s)
}
