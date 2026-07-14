package signplus1

import "encoding/json"

type AddEnvelopeSigningStepsRequest struct {
	// List of signing steps
	SigningSteps []SigningStep `json:"signing_steps,omitempty" xml:"signing_steps,omitempty"`
}

func (a AddEnvelopeSigningStepsRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddEnvelopeSigningStepsRequest to string"
	}
	return string(jsonData)
}
