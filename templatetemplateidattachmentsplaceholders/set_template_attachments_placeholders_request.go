package templatetemplateidattachmentsplaceholders

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type SetTemplateAttachmentsPlaceholdersRequest struct {
	Placeholders *param.Nullable[[]SetTemplateAttachmentsPlaceholdersRequestPlaceholders] `json:"placeholders,omitempty" xml:"placeholders,omitempty"`
}

func (s SetTemplateAttachmentsPlaceholdersRequest) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SetTemplateAttachmentsPlaceholdersRequest to string"
	}
	return string(jsonData)
}

func (s *SetTemplateAttachmentsPlaceholdersRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, s)
}
