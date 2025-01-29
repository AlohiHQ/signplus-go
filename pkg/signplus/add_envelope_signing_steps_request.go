package signplus

import (
	"encoding/json"
)

type AddEnvelopeSigningStepsRequest struct {
	// List of signing steps
	SigningSteps []SigningStep `json:"signing_steps,omitempty"`
	touched      map[string]bool
}

func (a *AddEnvelopeSigningStepsRequest) GetSigningSteps() []SigningStep {
	if a == nil {
		return nil
	}
	return a.SigningSteps
}

func (a *AddEnvelopeSigningStepsRequest) SetSigningSteps(signingSteps []SigningStep) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["SigningSteps"] = true
	a.SigningSteps = signingSteps
}

func (a *AddEnvelopeSigningStepsRequest) SetSigningStepsNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["SigningSteps"] = true
	a.SigningSteps = nil
}

func (a AddEnvelopeSigningStepsRequest) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if a.touched["SigningSteps"] && a.SigningSteps == nil {
		data["signing_steps"] = nil
	} else if a.SigningSteps != nil {
		data["signing_steps"] = a.SigningSteps
	}

	return json.Marshal(data)
}

func (a AddEnvelopeSigningStepsRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddEnvelopeSigningStepsRequest to string"
	}
	return string(jsonData)
}
