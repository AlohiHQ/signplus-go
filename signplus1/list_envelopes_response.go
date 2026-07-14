package signplus1

import "encoding/json"

type ListEnvelopesResponse struct {
	// Whether there is a next page
	HasNextPage *bool `json:"has_next_page,omitempty" xml:"has_next_page,omitempty"`
	// Whether there is a previous page
	HasPreviousPage *bool      `json:"has_previous_page,omitempty" xml:"has_previous_page,omitempty"`
	Envelopes       []Envelope `json:"envelopes,omitempty" xml:"envelopes,omitempty"`
}

func (l ListEnvelopesResponse) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListEnvelopesResponse to string"
	}
	return string(jsonData)
}
