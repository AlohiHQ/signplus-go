package signplus

import (
	"encoding/json"
)

type AddTemplateSigningStepsRequest struct {
	// List of signing steps
	SigningSteps []TemplateSigningStep `json:"signing_steps,omitempty" required:"true"`
	touched      map[string]bool
}

func (a *AddTemplateSigningStepsRequest) GetSigningSteps() []TemplateSigningStep {
	if a == nil {
		return nil
	}
	return a.SigningSteps
}

func (a *AddTemplateSigningStepsRequest) SetSigningSteps(signingSteps []TemplateSigningStep) {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["SigningSteps"] = true
	a.SigningSteps = signingSteps
}

func (a *AddTemplateSigningStepsRequest) SetSigningStepsNil() {
	if a.touched == nil {
		a.touched = map[string]bool{}
	}
	a.touched["SigningSteps"] = true
	a.SigningSteps = nil
}

func (a AddTemplateSigningStepsRequest) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if a.touched["SigningSteps"] && a.SigningSteps == nil {
		data["signing_steps"] = nil
	} else if a.SigningSteps != nil {
		data["signing_steps"] = a.SigningSteps
	}

	return json.Marshal(data)
}

func (a AddTemplateSigningStepsRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddTemplateSigningStepsRequest to string"
	}
	return string(jsonData)
}
