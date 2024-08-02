package signplus

type ListTemplatesResponse struct {
	// Whether there is a next page
	HasNextPage *bool `json:"has_next_page,omitempty"`
	// Whether there is a previous page
	HasPreviousPage *bool      `json:"has_previous_page,omitempty"`
	Templates       []Template `json:"templates,omitempty"`
}

func (l *ListTemplatesResponse) SetHasNextPage(hasNextPage bool) {
	l.HasNextPage = &hasNextPage
}

func (l *ListTemplatesResponse) GetHasNextPage() *bool {
	if l == nil {
		return nil
	}
	return l.HasNextPage
}

func (l *ListTemplatesResponse) SetHasPreviousPage(hasPreviousPage bool) {
	l.HasPreviousPage = &hasPreviousPage
}

func (l *ListTemplatesResponse) GetHasPreviousPage() *bool {
	if l == nil {
		return nil
	}
	return l.HasPreviousPage
}

func (l *ListTemplatesResponse) SetTemplates(templates []Template) {
	l.Templates = templates
}

func (l *ListTemplatesResponse) GetTemplates() []Template {
	if l == nil {
		return nil
	}
	return l.Templates
}
