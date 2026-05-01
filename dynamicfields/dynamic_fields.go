package dynamicfields

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type DynamicFields struct {
	Name  *param.Nullable[string] `json:"name,omitempty" xml:"name,omitempty"`
	Value *param.Nullable[string] `json:"value,omitempty" xml:"value,omitempty"`
}

func (d DynamicFields) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DynamicFields to string"
	}
	return string(jsonData)
}

func (d *DynamicFields) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, d)
}
