package setlegalitylevel

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type SetEnvelopeLegalityLevelRequest struct {
	LegalityLevel *param.Nullable[string] `json:"legality_level,omitempty" xml:"legality_level,omitempty"`
}

func (s SetEnvelopeLegalityLevelRequest) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SetEnvelopeLegalityLevelRequest to string"
	}
	return string(jsonData)
}

func (s *SetEnvelopeLegalityLevelRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, s)
}
