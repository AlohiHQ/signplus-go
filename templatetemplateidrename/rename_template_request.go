package templatetemplateidrename

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type RenameTemplateRequest struct {
	Name *param.Nullable[string] `json:"name,omitempty" xml:"name,omitempty"`
}

func (r RenameTemplateRequest) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RenameTemplateRequest to string"
	}
	return string(jsonData)
}

func (r *RenameTemplateRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, r)
}
