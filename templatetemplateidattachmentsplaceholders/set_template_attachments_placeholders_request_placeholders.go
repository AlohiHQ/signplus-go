package templatetemplateidattachmentsplaceholders

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type SetTemplateAttachmentsPlaceholdersRequestPlaceholders struct {
	RecipientID *param.Nullable[string] `json:"recipient_id,omitempty" xml:"recipient_id,omitempty"`
	Name        *param.Nullable[string] `json:"name,omitempty" xml:"name,omitempty"`
	Required    *param.Nullable[bool]   `json:"required,omitempty" xml:"required,omitempty"`
	Multiple    *param.Nullable[bool]   `json:"multiple,omitempty" xml:"multiple,omitempty"`
	ID          *param.Nullable[string] `json:"id,omitempty" xml:"id,omitempty"`
	Hint        *param.Nullable[string] `json:"hint,omitempty" xml:"hint,omitempty"`
}

func (s SetTemplateAttachmentsPlaceholdersRequestPlaceholders) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SetTemplateAttachmentsPlaceholdersRequestPlaceholders to string"
	}
	return string(jsonData)
}

func (s *SetTemplateAttachmentsPlaceholdersRequestPlaceholders) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, s)
}
