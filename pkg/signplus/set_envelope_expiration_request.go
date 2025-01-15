package signplus

import (
	"encoding/json"
)

type SetEnvelopeExpirationRequest struct {
	// Unix timestamp of the expiration date
	ExpiresAt *int64 `json:"expires_at,omitempty" required:"true"`
	touched   map[string]bool
}

func (s *SetEnvelopeExpirationRequest) GetExpiresAt() *int64 {
	if s == nil {
		return nil
	}
	return s.ExpiresAt
}

func (s *SetEnvelopeExpirationRequest) SetExpiresAt(expiresAt int64) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["ExpiresAt"] = true
	s.ExpiresAt = &expiresAt
}

func (s *SetEnvelopeExpirationRequest) SetExpiresAtNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["ExpiresAt"] = true
	s.ExpiresAt = nil
}
func (s SetEnvelopeExpirationRequest) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if s.touched["ExpiresAt"] && s.ExpiresAt == nil {
		data["expires_at"] = nil
	} else if s.ExpiresAt != nil {
		data["expires_at"] = s.ExpiresAt
	}

	return json.Marshal(data)
}
