package signplus1

import "encoding/json"

type SetEnvelopeLegalityLevelRequest struct {
	// Legal level of the envelope (SES is Simple Electronic Signature, QES_EIDAS is Qualified Electronic Signature, QES_ZERTES is Qualified Electronic Signature with Zertes)
	LegalityLevel *EnvelopeLegalityLevel `json:"legality_level,omitempty" xml:"legality_level,omitempty"`
}

func (s SetEnvelopeLegalityLevelRequest) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SetEnvelopeLegalityLevelRequest to string"
	}
	return string(jsonData)
}
