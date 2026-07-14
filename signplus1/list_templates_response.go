package signplus1

import "encoding/json"

type ListTemplatesResponse struct {
	// Whether there is a next page
	HasNextPage *bool `json:"has_next_page,omitempty" xml:"has_next_page,omitempty"`
	// Whether there is a previous page
	HasPreviousPage *bool      `json:"has_previous_page,omitempty" xml:"has_previous_page,omitempty"`
	Templates       []Template `json:"templates,omitempty" xml:"templates,omitempty"`
}

func (l ListTemplatesResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListTemplatesResponse to string"
	}
	return string(jsonData)
}
