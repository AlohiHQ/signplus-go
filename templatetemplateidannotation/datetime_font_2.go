package templatetemplateidannotation

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type DatetimeFont2 struct {
	Family *param.Nullable[string] `json:"family,omitempty" xml:"family,omitempty"`
	Italic *param.Nullable[bool]   `json:"italic,omitempty" xml:"italic,omitempty"`
	Bold   *param.Nullable[bool]   `json:"bold,omitempty" xml:"bold,omitempty"`
}

func (d DatetimeFont2) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DatetimeFont2 to string"
	}
	return string(jsonData)
}

func (d *DatetimeFont2) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, d)
}
