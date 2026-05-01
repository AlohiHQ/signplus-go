package setcomment

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type SetEnvelopeCommentRequest struct {
	Comment *param.Nullable[string] `json:"comment,omitempty" xml:"comment,omitempty"`
}

func (s SetEnvelopeCommentRequest) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SetEnvelopeCommentRequest to string"
	}
	return string(jsonData)
}

func (s *SetEnvelopeCommentRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, s)
}
