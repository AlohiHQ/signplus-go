package signplus

import "encoding/json"

type DynamicField struct {
	// Name of the dynamic field
	Name *string `json:"name,omitempty"`
	// Value of the dynamic field
	Value *string `json:"value,omitempty"`
}

func (d *DynamicField) GetName() *string {
	if d == nil {
		return nil
	}
	return d.Name
}

func (d *DynamicField) SetName(name string) {
	d.Name = &name
}

func (d *DynamicField) GetValue() *string {
	if d == nil {
		return nil
	}
	return d.Value
}

func (d *DynamicField) SetValue(value string) {
	d.Value = &value
}

func (d DynamicField) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DynamicField to string"
	}
	return string(jsonData)
}
