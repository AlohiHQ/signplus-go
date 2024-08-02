# ListTemplatesRequest

**Properties**

| Name       | Type                        | Required | Description                                   |
| :--------- | :-------------------------- | :------- | :-------------------------------------------- |
| Name       | string                      | ❌       | Name of the template                          |
| Tags       | []string                    | ❌       | List of tag templates                         |
| Ids        | []string                    | ❌       | List of templates IDs                         |
| First      | int64                       | ❌       |                                               |
| Last       | int64                       | ❌       |                                               |
| After      | string                      | ❌       |                                               |
| Before     | string                      | ❌       |                                               |
| OrderField | signplus.TemplateOrderField | ❌       | Field to order templates by                   |
| Ascending  | bool                        | ❌       | Whether to order templates in ascending order |
