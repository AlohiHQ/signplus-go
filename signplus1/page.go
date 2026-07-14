package signplus1

import "encoding/json"

type Page struct {
	// Width of the page in pixels
	Width *int64 `json:"width,omitempty" xml:"width,omitempty"`
	// Height of the page in pixels
	Height *int64 `json:"height,omitempty" xml:"height,omitempty"`
}

func (p Page) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: Page to string"
	}
	return string(jsonData)
}
