package signplus

import (
	"encoding/json"
)

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
	touched      map[string]bool
}

func (l *ListEnvelopesRequest) GetName() *string {
	if l == nil {
		return nil
	}
	return l.Name
}

func (l *ListEnvelopesRequest) SetName(name string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Name"] = true
	l.Name = &name
}

func (l *ListEnvelopesRequest) SetNameNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Name"] = true
	l.Name = nil
}

func (l *ListEnvelopesRequest) GetTags() []string {
	if l == nil {
		return nil
	}
	return l.Tags
}

func (l *ListEnvelopesRequest) SetTags(tags []string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Tags"] = true
	l.Tags = tags
}

func (l *ListEnvelopesRequest) SetTagsNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Tags"] = true
	l.Tags = nil
}

func (l *ListEnvelopesRequest) GetComment() *string {
	if l == nil {
		return nil
	}
	return l.Comment
}

func (l *ListEnvelopesRequest) SetComment(comment string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Comment"] = true
	l.Comment = &comment
}

func (l *ListEnvelopesRequest) SetCommentNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Comment"] = true
	l.Comment = nil
}

func (l *ListEnvelopesRequest) GetIds() []string {
	if l == nil {
		return nil
	}
	return l.Ids
}

func (l *ListEnvelopesRequest) SetIds(ids []string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Ids"] = true
	l.Ids = ids
}

func (l *ListEnvelopesRequest) SetIdsNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Ids"] = true
	l.Ids = nil
}

func (l *ListEnvelopesRequest) GetStatuses() []EnvelopeStatus {
	if l == nil {
		return nil
	}
	return l.Statuses
}

func (l *ListEnvelopesRequest) SetStatuses(statuses []EnvelopeStatus) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Statuses"] = true
	l.Statuses = statuses
}

func (l *ListEnvelopesRequest) SetStatusesNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Statuses"] = true
	l.Statuses = nil
}

func (l *ListEnvelopesRequest) GetFolderIds() []string {
	if l == nil {
		return nil
	}
	return l.FolderIds
}

func (l *ListEnvelopesRequest) SetFolderIds(folderIds []string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["FolderIds"] = true
	l.FolderIds = folderIds
}

func (l *ListEnvelopesRequest) SetFolderIdsNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["FolderIds"] = true
	l.FolderIds = nil
}

func (l *ListEnvelopesRequest) GetOnlyRootFolder() *bool {
	if l == nil {
		return nil
	}
	return l.OnlyRootFolder
}

func (l *ListEnvelopesRequest) SetOnlyRootFolder(onlyRootFolder bool) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["OnlyRootFolder"] = true
	l.OnlyRootFolder = &onlyRootFolder
}

func (l *ListEnvelopesRequest) SetOnlyRootFolderNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["OnlyRootFolder"] = true
	l.OnlyRootFolder = nil
}

func (l *ListEnvelopesRequest) GetDateFrom() *int64 {
	if l == nil {
		return nil
	}
	return l.DateFrom
}

func (l *ListEnvelopesRequest) SetDateFrom(dateFrom int64) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["DateFrom"] = true
	l.DateFrom = &dateFrom
}

func (l *ListEnvelopesRequest) SetDateFromNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["DateFrom"] = true
	l.DateFrom = nil
}

func (l *ListEnvelopesRequest) GetDateTo() *int64 {
	if l == nil {
		return nil
	}
	return l.DateTo
}

func (l *ListEnvelopesRequest) SetDateTo(dateTo int64) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["DateTo"] = true
	l.DateTo = &dateTo
}

func (l *ListEnvelopesRequest) SetDateToNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["DateTo"] = true
	l.DateTo = nil
}

func (l *ListEnvelopesRequest) GetUid() *string {
	if l == nil {
		return nil
	}
	return l.Uid
}

func (l *ListEnvelopesRequest) SetUid(uid string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Uid"] = true
	l.Uid = &uid
}

func (l *ListEnvelopesRequest) SetUidNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Uid"] = true
	l.Uid = nil
}

func (l *ListEnvelopesRequest) GetFirst() *int64 {
	if l == nil {
		return nil
	}
	return l.First
}

func (l *ListEnvelopesRequest) SetFirst(first int64) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["First"] = true
	l.First = &first
}

func (l *ListEnvelopesRequest) SetFirstNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["First"] = true
	l.First = nil
}

func (l *ListEnvelopesRequest) GetLast() *int64 {
	if l == nil {
		return nil
	}
	return l.Last
}

func (l *ListEnvelopesRequest) SetLast(last int64) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Last"] = true
	l.Last = &last
}

func (l *ListEnvelopesRequest) SetLastNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Last"] = true
	l.Last = nil
}

func (l *ListEnvelopesRequest) GetAfter() *string {
	if l == nil {
		return nil
	}
	return l.After
}

func (l *ListEnvelopesRequest) SetAfter(after string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["After"] = true
	l.After = &after
}

func (l *ListEnvelopesRequest) SetAfterNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["After"] = true
	l.After = nil
}

func (l *ListEnvelopesRequest) GetBefore() *string {
	if l == nil {
		return nil
	}
	return l.Before
}

func (l *ListEnvelopesRequest) SetBefore(before string) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Before"] = true
	l.Before = &before
}

func (l *ListEnvelopesRequest) SetBeforeNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Before"] = true
	l.Before = nil
}

func (l *ListEnvelopesRequest) GetOrderField() *EnvelopeOrderField {
	if l == nil {
		return nil
	}
	return l.OrderField
}

func (l *ListEnvelopesRequest) SetOrderField(orderField EnvelopeOrderField) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["OrderField"] = true
	l.OrderField = &orderField
}

func (l *ListEnvelopesRequest) SetOrderFieldNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["OrderField"] = true
	l.OrderField = nil
}

func (l *ListEnvelopesRequest) GetAscending() *bool {
	if l == nil {
		return nil
	}
	return l.Ascending
}

func (l *ListEnvelopesRequest) SetAscending(ascending bool) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Ascending"] = true
	l.Ascending = &ascending
}

func (l *ListEnvelopesRequest) SetAscendingNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["Ascending"] = true
	l.Ascending = nil
}

func (l *ListEnvelopesRequest) GetIncludeTrash() *bool {
	if l == nil {
		return nil
	}
	return l.IncludeTrash
}

func (l *ListEnvelopesRequest) SetIncludeTrash(includeTrash bool) {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["IncludeTrash"] = true
	l.IncludeTrash = &includeTrash
}

func (l *ListEnvelopesRequest) SetIncludeTrashNil() {
	if l.touched == nil {
		l.touched = map[string]bool{}
	}
	l.touched["IncludeTrash"] = true
	l.IncludeTrash = nil
}

func (l ListEnvelopesRequest) MarshalJSON() ([]byte, error) {
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

	if l.touched["Comment"] && l.Comment == nil {
		data["comment"] = nil
	} else if l.Comment != nil {
		data["comment"] = l.Comment
	}

	if l.touched["Ids"] && l.Ids == nil {
		data["ids"] = nil
	} else if l.Ids != nil {
		data["ids"] = l.Ids
	}

	if l.touched["Statuses"] && l.Statuses == nil {
		data["statuses"] = nil
	} else if l.Statuses != nil {
		data["statuses"] = l.Statuses
	}

	if l.touched["FolderIds"] && l.FolderIds == nil {
		data["folder_ids"] = nil
	} else if l.FolderIds != nil {
		data["folder_ids"] = l.FolderIds
	}

	if l.touched["OnlyRootFolder"] && l.OnlyRootFolder == nil {
		data["only_root_folder"] = nil
	} else if l.OnlyRootFolder != nil {
		data["only_root_folder"] = l.OnlyRootFolder
	}

	if l.touched["DateFrom"] && l.DateFrom == nil {
		data["date_from"] = nil
	} else if l.DateFrom != nil {
		data["date_from"] = l.DateFrom
	}

	if l.touched["DateTo"] && l.DateTo == nil {
		data["date_to"] = nil
	} else if l.DateTo != nil {
		data["date_to"] = l.DateTo
	}

	if l.touched["Uid"] && l.Uid == nil {
		data["uid"] = nil
	} else if l.Uid != nil {
		data["uid"] = l.Uid
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

	if l.touched["IncludeTrash"] && l.IncludeTrash == nil {
		data["include_trash"] = nil
	} else if l.IncludeTrash != nil {
		data["include_trash"] = l.IncludeTrash
	}

	return json.Marshal(data)
}

func (l ListEnvelopesRequest) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListEnvelopesRequest to string"
	}
	return string(jsonData)
}
