package signplus

import (
	"encoding/json"
)

type ListTemplatesResponse struct {
	// Whether there is a next page
	HasNextPage *bool `json:"has_next_page,omitempty"`
	// Whether there is a previous page
	HasPreviousPage *bool      `json:"has_previous_page,omitempty"`
	Templates       []Template `json:"templates,omitempty"`
	touched         map[string]bool
}

func (l *ListTemplatesResponse) GetHasNextPage() *bool {
	if l == nil {
		return nil
	}
	return l.HasNextPage
}

func (l *ListTemplatesResponse) SetHasNextPage(hasNextPage bool) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["HasNextPage"] = true
	l.HasNextPage = &hasNextPage
}

func (l *ListTemplatesResponse) SetHasNextPageNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["HasNextPage"] = true
	l.HasNextPage = nil
}

func (l *ListTemplatesResponse) GetHasPreviousPage() *bool {
	if l == nil {
		return nil
	}
	return l.HasPreviousPage
}

func (l *ListTemplatesResponse) SetHasPreviousPage(hasPreviousPage bool) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["HasPreviousPage"] = true
	l.HasPreviousPage = &hasPreviousPage
}

func (l *ListTemplatesResponse) SetHasPreviousPageNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["HasPreviousPage"] = true
	l.HasPreviousPage = nil
}

func (l *ListTemplatesResponse) GetTemplates() []Template {
	if l == nil {
		return nil
	}
	return l.Templates
}

func (l *ListTemplatesResponse) SetTemplates(templates []Template) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Templates"] = true
	l.Templates = templates
}

func (l *ListTemplatesResponse) SetTemplatesNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Templates"] = true
	l.Templates = nil
}
func (l ListTemplatesResponse) MarshalJSON() ([]byte, error) {
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

	if l.touched["Templates"] && l.Templates == nil {
		data["templates"] = nil
	} else if l.Templates != nil {
		data["templates"] = l.Templates
	}

	return json.Marshal(data)
}
