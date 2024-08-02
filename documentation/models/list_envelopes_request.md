# ListEnvelopesRequest

**Properties**

| Name           | Type                        | Required | Description                                       |
| :------------- | :-------------------------- | :------- | :------------------------------------------------ |
| Name           | string                      | ❌       | Name of the envelope                              |
| Tags           | []string                    | ❌       | List of tags                                      |
| Comment        | string                      | ❌       | Comment of the envelope                           |
| Ids            | []string                    | ❌       | List of envelope IDs                              |
| Statuses       | []signplus.EnvelopeStatus   | ❌       | List of envelope statuses                         |
| FolderIds      | []string                    | ❌       | List of folder IDs                                |
| OnlyRootFolder | bool                        | ❌       | Whether to only list envelopes in the root folder |
| DateFrom       | int64                       | ❌       | Unix timestamp of the start date                  |
| DateTo         | int64                       | ❌       | Unix timestamp of the end date                    |
| Uid            | string                      | ❌       | Unique identifier of the user                     |
| First          | int64                       | ❌       |                                                   |
| Last           | int64                       | ❌       |                                                   |
| After          | string                      | ❌       |                                                   |
| Before         | string                      | ❌       |                                                   |
| OrderField     | signplus.EnvelopeOrderField | ❌       | Field to order envelopes by                       |
| Ascending      | bool                        | ❌       | Whether to order envelopes in ascending order     |
| IncludeTrash   | bool                        | ❌       | Whether to include envelopes in the trash         |
