package templatetemplateidsigningsteps

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type AddTemplateSigningStepsRequestSigningSteps struct {
	Recipients *param.Nullable[[]SigningStepsRecipients2] `json:"recipients,omitempty" xml:"recipients,omitempty"`
}

func (a AddTemplateSigningStepsRequestSigningSteps) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddTemplateSigningStepsRequestSigningSteps to string"
	}
	return string(jsonData)
}

func (a *AddTemplateSigningStepsRequestSigningSteps) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, a)
}
