package signplus

import (
	"encoding/json"
)

type SetEnvelopeLegalityLevelRequest struct {
	// Legal level of the envelope (SES is Simple Electronic Signature, QES_EIDAS is Qualified Electronic Signature, QES_ZERTES is Qualified Electronic Signature with Zertes)
	LegalityLevel *EnvelopeLegalityLevel `json:"legality_level,omitempty"`
	touched       map[string]bool
}

func (s *SetEnvelopeLegalityLevelRequest) GetLegalityLevel() *EnvelopeLegalityLevel {
	if s == nil {
		return nil
	}
	return s.LegalityLevel
}

func (s *SetEnvelopeLegalityLevelRequest) SetLegalityLevel(legalityLevel EnvelopeLegalityLevel) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["LegalityLevel"] = true
	s.LegalityLevel = &legalityLevel
}

func (s *SetEnvelopeLegalityLevelRequest) SetLegalityLevelNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["LegalityLevel"] = true
	s.LegalityLevel = nil
}

func (s SetEnvelopeLegalityLevelRequest) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if s.touched["LegalityLevel"] && s.LegalityLevel == nil {
		data["legality_level"] = nil
	} else if s.LegalityLevel != nil {
		data["legality_level"] = s.LegalityLevel
	}

	return json.Marshal(data)
}

func (s SetEnvelopeLegalityLevelRequest) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SetEnvelopeLegalityLevelRequest to string"
	}
	return string(jsonData)
}
