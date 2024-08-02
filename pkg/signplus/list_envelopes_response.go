package signplus

type ListEnvelopesResponse struct {
	// Whether there is a next page
	HasNextPage *bool `json:"has_next_page,omitempty"`
	// Whether there is a previous page
	HasPreviousPage *bool      `json:"has_previous_page,omitempty"`
	Envelopes       []Envelope `json:"envelopes,omitempty"`
}

func (l *ListEnvelopesResponse) SetHasNextPage(hasNextPage bool) {
	l.HasNextPage = &hasNextPage
}

func (l *ListEnvelopesResponse) GetHasNextPage() *bool {
	if l == nil {
		return nil
	}
	return l.HasNextPage
}

func (l *ListEnvelopesResponse) SetHasPreviousPage(hasPreviousPage bool) {
	l.HasPreviousPage = &hasPreviousPage
}

func (l *ListEnvelopesResponse) GetHasPreviousPage() *bool {
	if l == nil {
		return nil
	}
	return l.HasPreviousPage
}

func (l *ListEnvelopesResponse) SetEnvelopes(envelopes []Envelope) {
	l.Envelopes = envelopes
}

func (l *ListEnvelopesResponse) GetEnvelopes() []Envelope {
	if l == nil {
		return nil
	}
	return l.Envelopes
}
