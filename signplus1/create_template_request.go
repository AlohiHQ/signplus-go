package signplus1

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
)

type CreateTemplateRequest struct {
	Name string `json:"name" xml:"name" required:"true" maxLength:"256" minLength:"2" pattern:"^[a-zA-Z0-9][a-zA-Z0-9 ]*[a-zA-Z0-9]$"`
}

func (c CreateTemplateRequest) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateTemplateRequest to string"
	}
	return string(jsonData)
}

func (c *CreateTemplateRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CreateTemplateRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CreateTemplateRequest(tmp)
	return nil
}
