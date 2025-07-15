package signplus

import "encoding/json"

type AddEnvelopeSigningStepsRequest struct {
	// List of signing steps
	SigningSteps []SigningStep `json:"signing_steps,omitempty"`
}

func (a *AddEnvelopeSigningStepsRequest) GetSigningSteps() []SigningStep {
	if a == nil {
		return nil
	}
	return a.SigningSteps
}

func (a *AddEnvelopeSigningStepsRequest) SetSigningSteps(signingSteps []SigningStep) {
	a.SigningSteps = signingSteps
}

func (a AddEnvelopeSigningStepsRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddEnvelopeSigningStepsRequest to string"
	}
	return string(jsonData)
}
