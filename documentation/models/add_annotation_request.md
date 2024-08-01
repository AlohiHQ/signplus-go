# AddAnnotationRequest

**Properties**

| Name        | Type                         | Required | Description                                                                                     |
| :---------- | :--------------------------- | :------- | :---------------------------------------------------------------------------------------------- |
| DocumentId  | string                       | ✅       | ID of the document                                                                              |
| Page        | int64                        | ✅       | Page number where the annotation is placed                                                      |
| X           | float64                      | ✅       | X coordinate of the annotation (in % of the page width from 0 to 100) from the top left corner  |
| Y           | float64                      | ✅       | Y coordinate of the annotation (in % of the page height from 0 to 100) from the top left corner |
| Width       | float64                      | ✅       | Width of the annotation (in % of the page width from 0 to 100)                                  |
| Height      | float64                      | ✅       | Height of the annotation (in % of the page height from 0 to 100)                                |
| Type\_      | signplus.AnnotationType      | ✅       | Type of the annotation                                                                          |
| RecipientId | string                       | ❌       | ID of the recipient                                                                             |
| Required    | bool                         | ❌       |                                                                                                 |
| Signature   | signplus.AnnotationSignature | ❌       | Signature annotation (null if annotation is not a signature)                                    |
| Initials    | signplus.AnnotationInitials  | ❌       | Initials annotation (null if annotation is not initials)                                        |
| Text        | signplus.AnnotationText      | ❌       | Text annotation (null if annotation is not a text)                                              |
| Datetime    | signplus.AnnotationDateTime  | ❌       | Date annotation (null if annotation is not a date)                                              |
| Checkbox    | signplus.AnnotationCheckbox  | ❌       | Checkbox annotation (null if annotation is not a checkbox)                                      |
