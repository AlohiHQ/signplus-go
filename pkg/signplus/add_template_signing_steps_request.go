package signplus

import "encoding/json"

type AddTemplateSigningStepsRequest struct {
	// List of signing steps
	SigningSteps []TemplateSigningStep `json:"signing_steps,omitempty" required:"true"`
}

func (a *AddTemplateSigningStepsRequest) GetSigningSteps() []TemplateSigningStep {
	if a == nil {
		return nil
	}
	return a.SigningSteps
}

func (a *AddTemplateSigningStepsRequest) SetSigningSteps(signingSteps []TemplateSigningStep) {
	a.SigningSteps = signingSteps
}

func (a AddTemplateSigningStepsRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddTemplateSigningStepsRequest to string"
	}
	return string(jsonData)
}
