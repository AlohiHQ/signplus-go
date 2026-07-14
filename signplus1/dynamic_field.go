package signplus1

import "encoding/json"

type DynamicField struct {
	// Name of the dynamic field
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Value of the dynamic field
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (d DynamicField) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DynamicField to string"
	}
	return string(jsonData)
}
