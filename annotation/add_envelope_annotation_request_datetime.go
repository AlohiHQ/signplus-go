package annotation

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type AddEnvelopeAnnotationRequestDatetime struct {
	Size      *param.Nullable[float64]       `json:"size,omitempty" xml:"size,omitempty"`
	Font      *param.Nullable[DatetimeFont1] `json:"font,omitempty" xml:"font,omitempty"`
	Color     *param.Nullable[string]        `json:"color,omitempty" xml:"color,omitempty"`
	AutoFill  *param.Nullable[bool]          `json:"auto_fill,omitempty" xml:"auto_fill,omitempty"`
	Timezone  *param.Nullable[string]        `json:"timezone,omitempty" xml:"timezone,omitempty"`
	Timestamp *param.Nullable[float64]       `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	Format    *param.Nullable[string]        `json:"format,omitempty" xml:"format,omitempty"`
}

func (a AddEnvelopeAnnotationRequestDatetime) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AddEnvelopeAnnotationRequestDatetime to string"
	}
	return string(jsonData)
}

func (a *AddEnvelopeAnnotationRequestDatetime) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, a)
}
