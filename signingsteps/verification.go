package signingsteps

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type Verification struct {
	Type  *param.Nullable[string] `json:"type,omitempty" xml:"type,omitempty"`
	Value *param.Nullable[string] `json:"value,omitempty" xml:"value,omitempty"`
}

func (v Verification) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: Verification to string"
	}
	return string(jsonData)
}

func (v *Verification) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, v)
}
