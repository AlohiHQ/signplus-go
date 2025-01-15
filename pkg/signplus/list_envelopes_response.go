package signplus

import (
	"encoding/json"
)

type ListEnvelopesResponse struct {
	// Whether there is a next page
	HasNextPage *bool `json:"has_next_page,omitempty"`
	// Whether there is a previous page
	HasPreviousPage *bool      `json:"has_previous_page,omitempty"`
	Envelopes       []Envelope `json:"envelopes,omitempty"`
	touched         map[string]bool
}

func (l *ListEnvelopesResponse) GetHasNextPage() *bool {
	if l == nil {
		return nil
	}
	return l.HasNextPage
}

func (l *ListEnvelopesResponse) SetHasNextPage(hasNextPage bool) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["HasNextPage"] = true
	l.HasNextPage = &hasNextPage
}

func (l *ListEnvelopesResponse) SetHasNextPageNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["HasNextPage"] = true
	l.HasNextPage = nil
}

func (l *ListEnvelopesResponse) GetHasPreviousPage() *bool {
	if l == nil {
		return nil
	}
	return l.HasPreviousPage
}

func (l *ListEnvelopesResponse) SetHasPreviousPage(hasPreviousPage bool) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["HasPreviousPage"] = true
	l.HasPreviousPage = &hasPreviousPage
}

func (l *ListEnvelopesResponse) SetHasPreviousPageNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["HasPreviousPage"] = true
	l.HasPreviousPage = nil
}

func (l *ListEnvelopesResponse) GetEnvelopes() []Envelope {
	if l == nil {
		return nil
	}
	return l.Envelopes
}

func (l *ListEnvelopesResponse) SetEnvelopes(envelopes []Envelope) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Envelopes"] = true
	l.Envelopes = envelopes
}

func (l *ListEnvelopesResponse) SetEnvelopesNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Envelopes"] = true
	l.Envelopes = nil
}
func (l ListEnvelopesResponse) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if l.touched["HasNextPage"] && l.HasNextPage == nil {
		data["has_next_page"] = nil
	} else if l.HasNextPage != nil {
		data["has_next_page"] = l.HasNextPage
	}

	if l.touched["HasPreviousPage"] && l.HasPreviousPage == nil {
		data["has_previous_page"] = nil
	} else if l.HasPreviousPage != nil {
		data["has_previous_page"] = l.HasPreviousPage
	}

	if l.touched["Envelopes"] && l.Envelopes == nil {
		data["envelopes"] = nil
	} else if l.Envelopes != nil {
		data["envelopes"] = l.Envelopes
	}

	return json.Marshal(data)
}
