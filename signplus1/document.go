package signplus1

import "encoding/json"

type Document struct {
	// Unique identifier of the document
	ID *string `json:"id,omitempty" xml:"id,omitempty"`
	// Name of the document
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Filename of the document
	Filename *string `json:"filename,omitempty" xml:"filename,omitempty"`
	// Number of pages in the document
	PageCount *int64 `json:"page_count,omitempty" xml:"page_count,omitempty"`
	// List of pages in the document
	Pages []Page `json:"pages,omitempty" xml:"pages,omitempty"`
}

func (d Document) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: Document to string"
	}
	return string(jsonData)
}
