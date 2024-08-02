package signplus

type ListEnvelopesRequest struct {
	// Name of the envelope
	Name *string `json:"name,omitempty"`
	// List of tags
	Tags []string `json:"tags,omitempty"`
	// Comment of the envelope
	Comment *string `json:"comment,omitempty"`
	// List of envelope IDs
	Ids []string `json:"ids,omitempty"`
	// List of envelope statuses
	Statuses []EnvelopeStatus `json:"statuses,omitempty"`
	// List of folder IDs
	FolderIds []string `json:"folder_ids,omitempty"`
	// Whether to only list envelopes in the root folder
	OnlyRootFolder *bool `json:"only_root_folder,omitempty"`
	// Unix timestamp of the start date
	DateFrom *int64 `json:"date_from,omitempty"`
	// Unix timestamp of the end date
	DateTo *int64 `json:"date_to,omitempty"`
	// Unique identifier of the user
	Uid    *string `json:"uid,omitempty"`
	First  *int64  `json:"first,omitempty"`
	Last   *int64  `json:"last,omitempty"`
	After  *string `json:"after,omitempty"`
	Before *string `json:"before,omitempty"`
	// Field to order envelopes by
	OrderField *EnvelopeOrderField `json:"order_field,omitempty"`
	// Whether to order envelopes in ascending order
	Ascending *bool `json:"ascending,omitempty"`
	// Whether to include envelopes in the trash
	IncludeTrash *bool `json:"include_trash,omitempty"`
}

func (l *ListEnvelopesRequest) SetName(name string) {
	l.Name = &name
}

func (l *ListEnvelopesRequest) GetName() *string {
	if l == nil {
		return nil
	}
	return l.Name
}

func (l *ListEnvelopesRequest) SetTags(tags []string) {
	l.Tags = tags
}

func (l *ListEnvelopesRequest) GetTags() []string {
	if l == nil {
		return nil
	}
	return l.Tags
}

func (l *ListEnvelopesRequest) SetComment(comment string) {
	l.Comment = &comment
}

func (l *ListEnvelopesRequest) GetComment() *string {
	if l == nil {
		return nil
	}
	return l.Comment
}

func (l *ListEnvelopesRequest) SetIds(ids []string) {
	l.Ids = ids
}

func (l *ListEnvelopesRequest) GetIds() []string {
	if l == nil {
		return nil
	}
	return l.Ids
}

func (l *ListEnvelopesRequest) SetStatuses(statuses []EnvelopeStatus) {
	l.Statuses = statuses
}

func (l *ListEnvelopesRequest) GetStatuses() []EnvelopeStatus {
	if l == nil {
		return nil
	}
	return l.Statuses
}

func (l *ListEnvelopesRequest) SetFolderIds(folderIds []string) {
	l.FolderIds = folderIds
}

func (l *ListEnvelopesRequest) GetFolderIds() []string {
	if l == nil {
		return nil
	}
	return l.FolderIds
}

func (l *ListEnvelopesRequest) SetOnlyRootFolder(onlyRootFolder bool) {
	l.OnlyRootFolder = &onlyRootFolder
}

func (l *ListEnvelopesRequest) GetOnlyRootFolder() *bool {
	if l == nil {
		return nil
	}
	return l.OnlyRootFolder
}

func (l *ListEnvelopesRequest) SetDateFrom(dateFrom int64) {
	l.DateFrom = &dateFrom
}

func (l *ListEnvelopesRequest) GetDateFrom() *int64 {
	if l == nil {
		return nil
	}
	return l.DateFrom
}

func (l *ListEnvelopesRequest) SetDateTo(dateTo int64) {
	l.DateTo = &dateTo
}

func (l *ListEnvelopesRequest) GetDateTo() *int64 {
	if l == nil {
		return nil
	}
	return l.DateTo
}

func (l *ListEnvelopesRequest) SetUid(uid string) {
	l.Uid = &uid
}

func (l *ListEnvelopesRequest) GetUid() *string {
	if l == nil {
		return nil
	}
	return l.Uid
}

func (l *ListEnvelopesRequest) SetFirst(first int64) {
	l.First = &first
}

func (l *ListEnvelopesRequest) GetFirst() *int64 {
	if l == nil {
		return nil
	}
	return l.First
}

func (l *ListEnvelopesRequest) SetLast(last int64) {
	l.Last = &last
}

func (l *ListEnvelopesRequest) GetLast() *int64 {
	if l == nil {
		return nil
	}
	return l.Last
}

func (l *ListEnvelopesRequest) SetAfter(after string) {
	l.After = &after
}

func (l *ListEnvelopesRequest) GetAfter() *string {
	if l == nil {
		return nil
	}
	return l.After
}

func (l *ListEnvelopesRequest) SetBefore(before string) {
	l.Before = &before
}

func (l *ListEnvelopesRequest) GetBefore() *string {
	if l == nil {
		return nil
	}
	return l.Before
}

func (l *ListEnvelopesRequest) SetOrderField(orderField EnvelopeOrderField) {
	l.OrderField = &orderField
}

func (l *ListEnvelopesRequest) GetOrderField() *EnvelopeOrderField {
	if l == nil {
		return nil
	}
	return l.OrderField
}

func (l *ListEnvelopesRequest) SetAscending(ascending bool) {
	l.Ascending = &ascending
}

func (l *ListEnvelopesRequest) GetAscending() *bool {
	if l == nil {
		return nil
	}
	return l.Ascending
}

func (l *ListEnvelopesRequest) SetIncludeTrash(includeTrash bool) {
	l.IncludeTrash = &includeTrash
}

func (l *ListEnvelopesRequest) GetIncludeTrash() *bool {
	if l == nil {
		return nil
	}
	return l.IncludeTrash
}
