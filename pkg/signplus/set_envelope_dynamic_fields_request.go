package signplus

type SetEnvelopeDynamicFieldsRequest struct {
	// List of dynamic fields
	DynamicFields []DynamicField `json:"dynamic_fields,omitempty" required:"true"`
}

func (s *SetEnvelopeDynamicFieldsRequest) SetDynamicFields(dynamicFields []DynamicField) {
	s.DynamicFields = dynamicFields
}

func (s *SetEnvelopeDynamicFieldsRequest) GetDynamicFields() []DynamicField {
	if s == nil {
		return nil
	}
	return s.DynamicFields
}
