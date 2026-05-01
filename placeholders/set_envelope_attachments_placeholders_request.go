package placeholders

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type SetEnvelopeAttachmentsPlaceholdersRequest struct {
	Placeholders *param.Nullable[[]SetEnvelopeAttachmentsPlaceholdersRequestPlaceholders] `json:"placeholders,omitempty" xml:"placeholders,omitempty"`
}

func (s SetEnvelopeAttachmentsPlaceholdersRequest) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SetEnvelopeAttachmentsPlaceholdersRequest to string"
	}
	return string(jsonData)
}

func (s *SetEnvelopeAttachmentsPlaceholdersRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, s)
}
