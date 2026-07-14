package signplus1

import "encoding/json"

type ListTemplatesRequest struct {
	// Name of the template
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// List of tag templates
	Tags []string `json:"tags,omitempty" xml:"tags,omitempty"`
	// List of templates IDs
	Ids    []string `json:"ids,omitempty" xml:"ids,omitempty"`
	First  *int64   `json:"first,omitempty" xml:"first,omitempty"`
	Last   *int64   `json:"last,omitempty" xml:"last,omitempty"`
	After  *string  `json:"after,omitempty" xml:"after,omitempty"`
	Before *string  `json:"before,omitempty" xml:"before,omitempty"`
	// Field to order templates by
	OrderField *TemplateOrderField `json:"order_field,omitempty" xml:"order_field,omitempty"`
	// Whether to order templates in ascending order
	Ascending *bool `json:"ascending,omitempty" xml:"ascending,omitempty"`
}

func (l ListTemplatesRequest) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListTemplatesRequest to string"
	}
	return string(jsonData)
}
