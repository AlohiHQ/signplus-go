package signplus

type AddTemplateSigningStepsRequest struct {
	// List of signing steps
	SigningSteps []TemplateSigningStep `json:"signing_steps,omitempty" required:"true"`
}

func (a *AddTemplateSigningStepsRequest) SetSigningSteps(signingSteps []TemplateSigningStep) {
	a.SigningSteps = signingSteps
}

func (a *AddTemplateSigningStepsRequest) GetSigningSteps() []TemplateSigningStep {
	if a == nil {
		return nil
	}
	return a.SigningSteps
}
