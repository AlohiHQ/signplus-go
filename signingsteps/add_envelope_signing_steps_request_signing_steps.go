package signingsteps

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type AddEnvelopeSigningStepsRequestSigningSteps struct {
	Recipients *param.Nullable[[]SigningStepsRecipients1] `json:"recipients,omitempty" xml:"recipients,omitempty"`
}

func (a AddEnvelopeSigningStepsRequestSigningSteps) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddEnvelopeSigningStepsRequestSigningSteps to string"
	}
	return string(jsonData)
}

func (a *AddEnvelopeSigningStepsRequestSigningSteps) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, a)
}
