package signplus

type SetEnvelopeExpirationRequest struct {
	// Unix timestamp of the expiration date
	ExpiresAt *int64 `json:"expires_at,omitempty" required:"true"`
}

func (s *SetEnvelopeExpirationRequest) SetExpiresAt(expiresAt int64) {
	s.ExpiresAt = &expiresAt
}

func (s *SetEnvelopeExpirationRequest) GetExpiresAt() *int64 {
	if s == nil {
		return nil
	}
	return s.ExpiresAt
}
