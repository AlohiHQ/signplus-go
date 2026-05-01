# TemplateTemplateIDAnnotationAnnotationID

A list of all methods in the `TemplateTemplateIDAnnotationAnnotationID` service. Click on the method name to view detailed information about that method.

| Methods                                               | Description                |
| :---------------------------------------------------- | :------------------------- |
| [DeleteTemplateAnnotation](#deletetemplateannotation) | Delete template annotation |

## DeleteTemplateAnnotation

Delete template annotation

- HTTP Method: `DELETE`
- Endpoint: `/template/{template_id}/annotation/{annotation_id}`

**Parameters**

| Name         | Type    | Required | Description                 |
| :----------- | :------ | :------- | :-------------------------- |
| ctx          | Context | ✅       | Default go language context |
| templateID   | string  | ✅       |                             |
| annotationID | string  | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

response, err := client.TemplateTemplateIDAnnotationAnnotationID.DeleteTemplateAnnotation(context.Background(), "template_id", "annotation_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```
