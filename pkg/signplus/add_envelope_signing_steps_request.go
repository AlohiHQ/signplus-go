package signplus

type AddEnvelopeSigningStepsRequest struct {
	// List of signing steps
	SigningSteps []SigningStep `json:"signing_steps,omitempty"`
}

func (a *AddEnvelopeSigningStepsRequest) SetSigningSteps(signingSteps []SigningStep) {
	a.SigningSteps = signingSteps
}

func (a *AddEnvelopeSigningStepsRequest) GetSigningSteps() []SigningStep {
	if a == nil {
		return nil
	}
	return a.SigningSteps
}
