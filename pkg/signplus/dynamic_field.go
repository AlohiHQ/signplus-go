package signplus

import (
	"encoding/json"
)

type DynamicField struct {
	// Name of the dynamic field
	Name *string `json:"name,omitempty"`
	// Value of the dynamic field
	Value   *string `json:"value,omitempty"`
	touched map[string]bool
}

func (d *DynamicField) GetName() *string {
	if d == nil {
		return nil
	}
	return d.Name
}

func (d *DynamicField) SetName(name string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Name"] = true
	d.Name = &name
}

func (d *DynamicField) SetNameNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Name"] = true
	d.Name = nil
}

func (d *DynamicField) GetValue() *string {
	if d == nil {
		return nil
	}
	return d.Value
}

func (d *DynamicField) SetValue(value string) {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Value"] = true
	d.Value = &value
}

func (d *DynamicField) SetValueNil() {
	if d.touched == nil {
		d.touched = map[string]bool{}
	}
	d.touched["Value"] = true
	d.Value = nil
}

func (d DynamicField) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if d.touched["Name"] && d.Name == nil {
		data["name"] = nil
	} else if d.Name != nil {
		data["name"] = d.Name
	}

	if d.touched["Value"] && d.Value == nil {
		data["value"] = nil
	} else if d.Value != nil {
		data["value"] = d.Value
	}

	return json.Marshal(data)
}

func (d DynamicField) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DynamicField to string"
	}
	return string(jsonData)
}
