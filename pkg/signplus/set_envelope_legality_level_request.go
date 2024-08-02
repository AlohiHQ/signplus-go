package signplus

type SetEnvelopeLegalityLevelRequest struct {
	// Legal level of the envelope (SES is Simple Electronic Signature, QES_EIDAS is Qualified Electronic Signature, QES_ZERTES is Qualified Electronic Signature with Zertes)
	LegalityLevel *EnvelopeLegalityLevel `json:"legality_level,omitempty"`
}

func (s *SetEnvelopeLegalityLevelRequest) SetLegalityLevel(legalityLevel EnvelopeLegalityLevel) {
	s.LegalityLevel = &legalityLevel
}

func (s *SetEnvelopeLegalityLevelRequest) GetLegalityLevel() *EnvelopeLegalityLevel {
	if s == nil {
		return nil
	}
	return s.LegalityLevel
}
