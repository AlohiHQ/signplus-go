package annotation

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type TextFont1 struct {
	Family *param.Nullable[string] `json:"family,omitempty" xml:"family,omitempty"`
	Italic *param.Nullable[string] `json:"italic,omitempty" xml:"italic,omitempty"`
	Bold   *param.Nullable[string] `json:"bold,omitempty" xml:"bold,omitempty"`
}

func (t TextFont1) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TextFont1 to string"
	}
	return string(jsonData)
}

func (t *TextFont1) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, t)
}
