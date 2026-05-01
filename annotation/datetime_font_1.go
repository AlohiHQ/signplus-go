package annotation

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type DatetimeFont1 struct {
	Family *param.Nullable[string] `json:"family,omitempty" xml:"family,omitempty"`
	Italic *param.Nullable[string] `json:"italic,omitempty" xml:"italic,omitempty"`
	Bold   *param.Nullable[string] `json:"bold,omitempty" xml:"bold,omitempty"`
}

func (d DatetimeFont1) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DatetimeFont1 to string"
	}
	return string(jsonData)
}

func (d *DatetimeFont1) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, d)
}
