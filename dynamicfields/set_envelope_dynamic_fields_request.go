package dynamicfields

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type SetEnvelopeDynamicFieldsRequest struct {
	DynamicFields *param.Nullable[[]DynamicFields] `json:"dynamic_fields,omitempty" xml:"dynamic_fields,omitempty"`
}

func (s SetEnvelopeDynamicFieldsRequest) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SetEnvelopeDynamicFieldsRequest to string"
	}
	return string(jsonData)
}

func (s *SetEnvelopeDynamicFieldsRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, s)
}
