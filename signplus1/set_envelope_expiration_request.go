package signplus1

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
)

type SetEnvelopeExpirationRequest struct {
	// Unix timestamp of the expiration date
	ExpiresAt int64 `json:"expires_at" xml:"expires_at" required:"true"`
}

func (s SetEnvelopeExpirationRequest) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SetEnvelopeExpirationRequest to string"
	}
	return string(jsonData)
}

func (s *SetEnvelopeExpirationRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SetEnvelopeExpirationRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SetEnvelopeExpirationRequest(tmp)
	return nil
}
