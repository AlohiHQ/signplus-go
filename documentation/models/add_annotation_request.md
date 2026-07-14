# AddAnnotationRequest

**Properties**

| Name        | Type                                                     | Required | Description                                                                                     |
| :---------- | :------------------------------------------------------- | :------- | :---------------------------------------------------------------------------------------------- |
| DocumentID  | string                                                   | ✅       | ID of the document                                                                              |
| Page        | int64                                                    | ✅       | Page number where the annotation is placed                                                      |
| X           | float64                                                  | ✅       | X coordinate of the annotation (in % of the page width from 0 to 100) from the top left corner  |
| Y           | float64                                                  | ✅       | Y coordinate of the annotation (in % of the page height from 0 to 100) from the top left corner |
| Width       | float64                                                  | ✅       | Width of the annotation (in % of the page width from 0 to 100)                                  |
| Height      | float64                                                  | ✅       | Height of the annotation (in % of the page height from 0 to 100)                                |
| Type        | [signplus1.AnnotationType](annotation_type.md)           | ✅       | Type of the annotation                                                                          |
| RecipientID | string                                                   | ❌       | ID of the recipient                                                                             |
| Required    | bool                                                     | ❌       |                                                                                                 |
| Signature   | [signplus1.AnnotationSignature](annotation_signature.md) | ❌       | Signature annotation (null if annotation is not a signature)                                    |
| Initials    | [signplus1.AnnotationInitials](annotation_initials.md)   | ❌       | Initials annotation (null if annotation is not initials)                                        |
| Text        | [signplus1.AnnotationText](annotation_text.md)           | ❌       | Text annotation (null if annotation is not a text)                                              |
| Datetime    | [signplus1.AnnotationDateTime](annotation_date_time.md)  | ❌       | Date annotation (null if annotation is not a date)                                              |
| Checkbox    | [signplus1.AnnotationCheckbox](annotation_checkbox.md)   | ❌       | Checkbox annotation (null if annotation is not a checkbox)                                      |
