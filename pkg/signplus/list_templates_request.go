package signplus

type ListTemplatesRequest struct {
	// Name of the template
	Name *string `json:"name,omitempty"`
	// List of tag templates
	Tags []string `json:"tags,omitempty"`
	// List of templates IDs
	Ids    []string `json:"ids,omitempty"`
	First  *int64   `json:"first,omitempty"`
	Last   *int64   `json:"last,omitempty"`
	After  *string  `json:"after,omitempty"`
	Before *string  `json:"before,omitempty"`
	// Field to order templates by
	OrderField *TemplateOrderField `json:"order_field,omitempty"`
	// Whether to order templates in ascending order
	Ascending *bool `json:"ascending,omitempty"`
}

func (l *ListTemplatesRequest) SetName(name string) {
	l.Name = &name
}

func (l *ListTemplatesRequest) GetName() *string {
	if l == nil {
		return nil
	}
	return l.Name
}

func (l *ListTemplatesRequest) SetTags(tags []string) {
	l.Tags = tags
}

func (l *ListTemplatesRequest) GetTags() []string {
	if l == nil {
		return nil
	}
	return l.Tags
}

func (l *ListTemplatesRequest) SetIds(ids []string) {
	l.Ids = ids
}

func (l *ListTemplatesRequest) GetIds() []string {
	if l == nil {
		return nil
	}
	return l.Ids
}

func (l *ListTemplatesRequest) SetFirst(first int64) {
	l.First = &first
}

func (l *ListTemplatesRequest) GetFirst() *int64 {
	if l == nil {
		return nil
	}
	return l.First
}

func (l *ListTemplatesRequest) SetLast(last int64) {
	l.Last = &last
}

func (l *ListTemplatesRequest) GetLast() *int64 {
	if l == nil {
		return nil
	}
	return l.Last
}

func (l *ListTemplatesRequest) SetAfter(after string) {
	l.After = &after
}

func (l *ListTemplatesRequest) GetAfter() *string {
	if l == nil {
		return nil
	}
	return l.After
}

func (l *ListTemplatesRequest) SetBefore(before string) {
	l.Before = &before
}

func (l *ListTemplatesRequest) GetBefore() *string {
	if l == nil {
		return nil
	}
	return l.Before
}

func (l *ListTemplatesRequest) SetOrderField(orderField TemplateOrderField) {
	l.OrderField = &orderField
}

func (l *ListTemplatesRequest) GetOrderField() *TemplateOrderField {
	if l == nil {
		return nil
	}
	return l.OrderField
}

func (l *ListTemplatesRequest) SetAscending(ascending bool) {
	l.Ascending = &ascending
}

func (l *ListTemplatesRequest) GetAscending() *bool {
	if l == nil {
		return nil
	}
	return l.Ascending
}
