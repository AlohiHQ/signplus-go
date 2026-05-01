package templatetemplateidsigningsteps

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type AddTemplateSigningStepsRequest struct {
	SigningSteps *param.Nullable[[]AddTemplateSigningStepsRequestSigningSteps] `json:"signing_steps,omitempty" xml:"signing_steps,omitempty"`
}

func (a AddTemplateSigningStepsRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddTemplateSigningStepsRequest to string"
	}
	return string(jsonData)
}

func (a *AddTemplateSigningStepsRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, a)
}
