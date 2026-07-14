package signplus1

import "encoding/json"

type SetEnvelopeDynamicFieldsRequest struct {
	// List of dynamic fields
	DynamicFields []DynamicField `json:"dynamic_fields" xml:"dynamic_fields" required:"true"`
}

func (s SetEnvelopeDynamicFieldsRequest) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SetEnvelopeDynamicFieldsRequest to string"
	}
	return string(jsonData)
}
