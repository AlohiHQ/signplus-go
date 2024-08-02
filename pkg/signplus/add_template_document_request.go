package signplus

type AddTemplateDocumentRequest struct {
	// File to upload in binary format
	File *any `json:"file,omitempty" required:"true"`
}

func (a *AddTemplateDocumentRequest) SetFile(file any) {
	a.File = &file
}

func (a *AddTemplateDocumentRequest) GetFile() *any {
	if a == nil {
		return nil
	}
	return a.File
}
