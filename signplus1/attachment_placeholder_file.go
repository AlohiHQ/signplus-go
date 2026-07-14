package signplus1

import "encoding/json"

type AttachmentPlaceholderFile struct {
	// ID of the file
	ID *string `json:"id,omitempty" xml:"id,omitempty"`
	// Name of the file
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Size of the file in bytes
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// MIME type of the file
	Mimetype *string `json:"mimetype,omitempty" xml:"mimetype,omitempty"`
}

func (a AttachmentPlaceholderFile) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AttachmentPlaceholderFile to string"
	}
	return string(jsonData)
}
