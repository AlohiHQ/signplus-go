package signingsteps

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type AddEnvelopeSigningStepsRequest struct {
	SigningSteps *param.Nullable[[]AddEnvelopeSigningStepsRequestSigningSteps] `json:"signing_steps,omitempty" xml:"signing_steps,omitempty"`
}

func (a AddEnvelopeSigningStepsRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddEnvelopeSigningStepsRequest to string"
	}
	return string(jsonData)
}

func (a *AddEnvelopeSigningStepsRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, a)
}
