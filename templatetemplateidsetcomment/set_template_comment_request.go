package templatetemplateidsetcomment

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type SetTemplateCommentRequest struct {
	Comment *param.Nullable[string] `json:"comment,omitempty" xml:"comment,omitempty"`
}

func (s SetTemplateCommentRequest) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SetTemplateCommentRequest to string"
	}
	return string(jsonData)
}

func (s *SetTemplateCommentRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, s)
}
