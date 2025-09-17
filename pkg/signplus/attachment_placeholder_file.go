package signplus

import "encoding/json"

type AttachmentPlaceholderFile struct {
	// ID of the file
	Id *string `json:"id,omitempty"`
	// Name of the file
	Name *string `json:"name,omitempty"`
	// Size of the file in bytes
	Size *int64 `json:"size,omitempty"`
	// MIME type of the file
	Mimetype *string `json:"mimetype,omitempty"`
}

func (a *AttachmentPlaceholderFile) GetId() *string {
	if a == nil {
		return nil
	}
	return a.Id
}

func (a *AttachmentPlaceholderFile) SetId(id string) {
	a.Id = &id
}

func (a *AttachmentPlaceholderFile) GetName() *string {
	if a == nil {
		return nil
	}
	return a.Name
}

func (a *AttachmentPlaceholderFile) SetName(name string) {
	a.Name = &name
}

func (a *AttachmentPlaceholderFile) GetSize() *int64 {
	if a == nil {
		return nil
	}
	return a.Size
}

func (a *AttachmentPlaceholderFile) SetSize(size int64) {
	a.Size = &size
}

func (a *AttachmentPlaceholderFile) GetMimetype() *string {
	if a == nil {
		return nil
	}
	return a.Mimetype
}

func (a *AttachmentPlaceholderFile) SetMimetype(mimetype string) {
	a.Mimetype = &mimetype
}

func (a AttachmentPlaceholderFile) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AttachmentPlaceholderFile to string"
	}
	return string(jsonData)
}
