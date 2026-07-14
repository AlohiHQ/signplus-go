package signplus1

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
)

type SetEnvelopeCommentRequest struct {
	// Comment for the envelope
	Comment string `json:"comment" xml:"comment" required:"true"`
}

func (s SetEnvelopeCommentRequest) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SetEnvelopeCommentRequest to string"
	}
	return string(jsonData)
}

func (s *SetEnvelopeCommentRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SetEnvelopeCommentRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SetEnvelopeCommentRequest(tmp)
	return nil
}
