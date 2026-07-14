package signplus1

import "encoding/json"

type ListEnvelopesRequest struct {
	// Name of the envelope
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// List of tags
	Tags []string `json:"tags,omitempty" xml:"tags,omitempty"`
	// Comment of the envelope
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// List of envelope IDs
	Ids []string `json:"ids,omitempty" xml:"ids,omitempty"`
	// List of envelope statuses
	Statuses []EnvelopeStatus `json:"statuses,omitempty" xml:"statuses,omitempty"`
	// List of folder IDs
	FolderIds []string `json:"folder_ids,omitempty" xml:"folder_ids,omitempty"`
	// Whether to only list envelopes in the root folder
	OnlyRootFolder *bool `json:"only_root_folder,omitempty" xml:"only_root_folder,omitempty"`
	// Unix timestamp of the start date
	DateFrom *int64 `json:"date_from,omitempty" xml:"date_from,omitempty"`
	// Unix timestamp of the end date
	DateTo *int64 `json:"date_to,omitempty" xml:"date_to,omitempty"`
	// Unique identifier of the user
	UID    *string `json:"uid,omitempty" xml:"uid,omitempty"`
	First  *int64  `json:"first,omitempty" xml:"first,omitempty"`
	Last   *int64  `json:"last,omitempty" xml:"last,omitempty"`
	After  *string `json:"after,omitempty" xml:"after,omitempty"`
	Before *string `json:"before,omitempty" xml:"before,omitempty"`
	// Field to order envelopes by
	OrderField *EnvelopeOrderField `json:"order_field,omitempty" xml:"order_field,omitempty"`
	// Whether to order envelopes in ascending order
	Ascending *bool `json:"ascending,omitempty" xml:"ascending,omitempty"`
	// Whether to include envelopes in the trash
	IncludeTrash *bool `json:"include_trash,omitempty" xml:"include_trash,omitempty"`
}

func (l ListEnvelopesRequest) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListEnvelopesRequest to string"
	}
	return string(jsonData)
}
