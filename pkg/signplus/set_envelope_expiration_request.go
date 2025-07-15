package signplus

import "encoding/json"

type SetEnvelopeExpirationRequest struct {
	// Unix timestamp of the expiration date
	ExpiresAt *int64 `json:"expires_at,omitempty" required:"true"`
}

func (s *SetEnvelopeExpirationRequest) GetExpiresAt() *int64 {
	if s == nil {
		return nil
	}
	return s.ExpiresAt
}

func (s *SetEnvelopeExpirationRequest) SetExpiresAt(expiresAt int64) {
	s.ExpiresAt = &expiresAt
}

func (s SetEnvelopeExpirationRequest) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SetEnvelopeExpirationRequest to string"
	}
	return string(jsonData)
}
