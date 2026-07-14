package signplus1

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
)

type SetTemplateCommentRequest struct {
	// Comment for the template
	Comment string `json:"comment" xml:"comment" required:"true"`
}

func (s SetTemplateCommentRequest) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SetTemplateCommentRequest to string"
	}
	return string(jsonData)
}

func (s *SetTemplateCommentRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SetTemplateCommentRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SetTemplateCommentRequest(tmp)
	return nil
}
