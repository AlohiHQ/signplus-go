package rename

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type RenameEnvelopeRequest struct {
	Name *param.Nullable[string] `json:"name,omitempty" xml:"name,omitempty"`
}

func (r RenameEnvelopeRequest) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RenameEnvelopeRequest to string"
	}
	return string(jsonData)
}

func (r *RenameEnvelopeRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, r)
}
