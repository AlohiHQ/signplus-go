package signplus

import "encoding/json"

type SetEnvelopeDynamicFieldsRequest struct {
	// List of dynamic fields
	DynamicFields []DynamicField `json:"dynamic_fields,omitempty" required:"true"`
}

func (s *SetEnvelopeDynamicFieldsRequest) GetDynamicFields() []DynamicField {
	if s == nil {
		return nil
	}
	return s.DynamicFields
}

func (s *SetEnvelopeDynamicFieldsRequest) SetDynamicFields(dynamicFields []DynamicField) {
	s.DynamicFields = dynamicFields
}

func (s SetEnvelopeDynamicFieldsRequest) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SetEnvelopeDynamicFieldsRequest to string"
	}
	return string(jsonData)
}
