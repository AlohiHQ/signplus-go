package signplus1

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
)

type RenameTemplateRequest struct {
	// Name of the template
	Name string `json:"name" xml:"name" required:"true"`
}

func (r RenameTemplateRequest) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RenameTemplateRequest to string"
	}
	return string(jsonData)
}

func (r *RenameTemplateRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias RenameTemplateRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = RenameTemplateRequest(tmp)
	return nil
}
