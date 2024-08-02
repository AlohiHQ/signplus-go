package signplus

type AddEnvelopeDocumentRequest struct {
	// File to upload in binary format
	File *any `json:"file,omitempty"`
}

func (a *AddEnvelopeDocumentRequest) SetFile(file any) {
	a.File = &file
}

func (a *AddEnvelopeDocumentRequest) GetFile() *any {
	if a == nil {
		return nil
	}
	return a.File
}
