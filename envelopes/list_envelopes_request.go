package envelopes

import (
	"encoding/json"
	"github.com/alohihq/signplus-go/internal/unmarshal"
	"github.com/alohihq/signplus-go/param"
)

type ListEnvelopesRequest struct {
	Name           *param.Nullable[string]   `json:"name,omitempty" xml:"name,omitempty"`
	Tags           *param.Nullable[[]string] `json:"tags,omitempty" xml:"tags,omitempty"`
	Comment        *param.Nullable[string]   `json:"comment,omitempty" xml:"comment,omitempty"`
	Ids            *param.Nullable[[]string] `json:"ids,omitempty" xml:"ids,omitempty"`
	Statuses       *param.Nullable[[]string] `json:"statuses,omitempty" xml:"statuses,omitempty"`
	FolderIds      *param.Nullable[[]string] `json:"folder_ids,omitempty" xml:"folder_ids,omitempty"`
	OnlyRootFolder *param.Nullable[string]   `json:"only_root_folder,omitempty" xml:"only_root_folder,omitempty"`
	DateFrom       *param.Nullable[string]   `json:"date_from,omitempty" xml:"date_from,omitempty"`
	DateTo         *param.Nullable[string]   `json:"date_to,omitempty" xml:"date_to,omitempty"`
	UID            *param.Nullable[string]   `json:"uid,omitempty" xml:"uid,omitempty"`
	First          *param.Nullable[string]   `json:"first,omitempty" xml:"first,omitempty"`
	Last           *param.Nullable[string]   `json:"last,omitempty" xml:"last,omitempty"`
	After          *param.Nullable[string]   `json:"after,omitempty" xml:"after,omitempty"`
	Before         *param.Nullable[string]   `json:"before,omitempty" xml:"before,omitempty"`
	OrderField     *param.Nullable[string]   `json:"order_field,omitempty" xml:"order_field,omitempty"`
	Ascending      *param.Nullable[string]   `json:"ascending,omitempty" xml:"ascending,omitempty"`
	IncludeTrash   *param.Nullable[string]   `json:"include_trash,omitempty" xml:"include_trash,omitempty"`
}

func (l ListEnvelopesRequest) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListEnvelopesRequest to string"
	}
	return string(jsonData)
}

func (l *ListEnvelopesRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, l)
}
