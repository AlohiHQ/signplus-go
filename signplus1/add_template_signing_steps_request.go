package signplus1

import "encoding/json"

type AddTemplateSigningStepsRequest struct {
	// List of signing steps
	SigningSteps []TemplateSigningStep `json:"signing_steps" xml:"signing_steps" required:"true"`
}

func (a AddTemplateSigningStepsRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddTemplateSigningStepsRequest to string"
	}
	return string(jsonData)
}
