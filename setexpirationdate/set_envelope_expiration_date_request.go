package setexpirationdate

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type SetEnvelopeExpirationDateRequest struct {
	ExpiresAt *param.Nullable[float64] `json:"expires_at,omitempty" xml:"expires_at,omitempty"`
}

func (s SetEnvelopeExpirationDateRequest) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SetEnvelopeExpirationDateRequest to string"
	}
	return string(jsonData)
}

func (s *SetEnvelopeExpirationDateRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, s)
}
