package signplus

type DynamicField struct {
	// Name of the dynamic field
	Name *string `json:"name,omitempty"`
	// Value of the dynamic field
	Value *string `json:"value,omitempty"`
}

func (d *DynamicField) SetName(name string) {
	d.Name = &name
}

func (d *DynamicField) GetName() *string {
	if d == nil {
		return nil
	}
	return d.Name
}

func (d *DynamicField) SetValue(value string) {
	d.Value = &value
}

func (d *DynamicField) GetValue() *string {
	if d == nil {
		return nil
	}
	return d.Value
}
