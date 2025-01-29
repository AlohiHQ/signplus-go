package signplus

import (
	"encoding/json"
)

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
	touched   map[string]bool
}

func (l *ListTemplatesRequest) GetName() *string {
	if l == nil {
		return nil
	}
	return l.Name
}

func (l *ListTemplatesRequest) SetName(name string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Name"] = true
	l.Name = &name
}

func (l *ListTemplatesRequest) SetNameNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Name"] = true
	l.Name = nil
}

func (l *ListTemplatesRequest) GetTags() []string {
	if l == nil {
		return nil
	}
	return l.Tags
}

func (l *ListTemplatesRequest) SetTags(tags []string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Tags"] = true
	l.Tags = tags
}

func (l *ListTemplatesRequest) SetTagsNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Tags"] = true
	l.Tags = nil
}

func (l *ListTemplatesRequest) GetIds() []string {
	if l == nil {
		return nil
	}
	return l.Ids
}

func (l *ListTemplatesRequest) SetIds(ids []string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Ids"] = true
	l.Ids = ids
}

func (l *ListTemplatesRequest) SetIdsNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Ids"] = true
	l.Ids = nil
}

func (l *ListTemplatesRequest) GetFirst() *int64 {
	if l == nil {
		return nil
	}
	return l.First
}

func (l *ListTemplatesRequest) SetFirst(first int64) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["First"] = true
	l.First = &first
}

func (l *ListTemplatesRequest) SetFirstNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["First"] = true
	l.First = nil
}

func (l *ListTemplatesRequest) GetLast() *int64 {
	if l == nil {
		return nil
	}
	return l.Last
}

func (l *ListTemplatesRequest) SetLast(last int64) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Last"] = true
	l.Last = &last
}

func (l *ListTemplatesRequest) SetLastNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Last"] = true
	l.Last = nil
}

func (l *ListTemplatesRequest) GetAfter() *string {
	if l == nil {
		return nil
	}
	return l.After
}

func (l *ListTemplatesRequest) SetAfter(after string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["After"] = true
	l.After = &after
}

func (l *ListTemplatesRequest) SetAfterNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["After"] = true
	l.After = nil
}

func (l *ListTemplatesRequest) GetBefore() *string {
	if l == nil {
		return nil
	}
	return l.Before
}

func (l *ListTemplatesRequest) SetBefore(before string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Before"] = true
	l.Before = &before
}

func (l *ListTemplatesRequest) SetBeforeNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Before"] = true
	l.Before = nil
}

func (l *ListTemplatesRequest) GetOrderField() *TemplateOrderField {
	if l == nil {
		return nil
	}
	return l.OrderField
}

func (l *ListTemplatesRequest) SetOrderField(orderField TemplateOrderField) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["OrderField"] = true
	l.OrderField = &orderField
}

func (l *ListTemplatesRequest) SetOrderFieldNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["OrderField"] = true
	l.OrderField = nil
}

func (l *ListTemplatesRequest) GetAscending() *bool {
	if l == nil {
		return nil
	}
	return l.Ascending
}

func (l *ListTemplatesRequest) SetAscending(ascending bool) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Ascending"] = true
	l.Ascending = &ascending
}

func (l *ListTemplatesRequest) SetAscendingNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Ascending"] = true
	l.Ascending = nil
}

func (l ListTemplatesRequest) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if l.touched["Name"] && l.Name == nil {
		data["name"] = nil
	} else if l.Name != nil {
		data["name"] = l.Name
	}

	if l.touched["Tags"] && l.Tags == nil {
		data["tags"] = nil
	} else if l.Tags != nil {
		data["tags"] = l.Tags
	}

	if l.touched["Ids"] && l.Ids == nil {
		data["ids"] = nil
	} else if l.Ids != nil {
		data["ids"] = l.Ids
	}

	if l.touched["First"] && l.First == nil {
		data["first"] = nil
	} else if l.First != nil {
		data["first"] = l.First
	}

	if l.touched["Last"] && l.Last == nil {
		data["last"] = nil
	} else if l.Last != nil {
		data["last"] = l.Last
	}

	if l.touched["After"] && l.After == nil {
		data["after"] = nil
	} else if l.After != nil {
		data["after"] = l.After
	}

	if l.touched["Before"] && l.Before == nil {
		data["before"] = nil
	} else if l.Before != nil {
		data["before"] = l.Before
	}

	if l.touched["OrderField"] && l.OrderField == nil {
		data["order_field"] = nil
	} else if l.OrderField != nil {
		data["order_field"] = l.OrderField
	}

	if l.touched["Ascending"] && l.Ascending == nil {
		data["ascending"] = nil
	} else if l.Ascending != nil {
		data["ascending"] = l.Ascending
	}

	return json.Marshal(data)
}

func (l ListTemplatesRequest) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListTemplatesRequest to string"
	}
	return string(jsonData)
}
