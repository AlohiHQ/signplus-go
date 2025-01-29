package signplus

import (
	"encoding/json"
)

type SetEnvelopeDynamicFieldsRequest struct {
	// List of dynamic fields
	DynamicFields []DynamicField `json:"dynamic_fields,omitempty" required:"true"`
	touched       map[string]bool
}

func (s *SetEnvelopeDynamicFieldsRequest) GetDynamicFields() []DynamicField {
	if s == nil {
		return nil
	}
	return s.DynamicFields
}

func (s *SetEnvelopeDynamicFieldsRequest) SetDynamicFields(dynamicFields []DynamicField) {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["DynamicFields"] = true
	s.DynamicFields = dynamicFields
}

func (s *SetEnvelopeDynamicFieldsRequest) SetDynamicFieldsNil() {
	if s.touched == nil {
		s.touched = map[string]bool{}
	}
	s.touched["DynamicFields"] = true
	s.DynamicFields = nil
}

func (s SetEnvelopeDynamicFieldsRequest) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if s.touched["DynamicFields"] && s.DynamicFields == nil {
		data["dynamic_fields"] = nil
	} else if s.DynamicFields != nil {
		data["dynamic_fields"] = s.DynamicFields
	}

	return json.Marshal(data)
}

func (s SetEnvelopeDynamicFieldsRequest) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SetEnvelopeDynamicFieldsRequest to string"
	}
	return string(jsonData)
}
