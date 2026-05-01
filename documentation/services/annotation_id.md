# AnnotationID

A list of all methods in the `AnnotationID` service. Click on the method name to view detailed information about that method.

| Methods                                               | Description                |
| :---------------------------------------------------- | :------------------------- |
| [DeleteEnvelopeAnnotation](#deleteenvelopeannotation) | Delete envelope annotation |

## DeleteEnvelopeAnnotation

Delete envelope annotation

- HTTP Method: `DELETE`
- Endpoint: `/envelope/{envelope_id}/annotation/{annotation_id}`

**Parameters**

| Name         | Type    | Required | Description                 |
| :----------- | :------ | :------- | :-------------------------- |
| ctx          | Context | ✅       | Default go language context |
| envelopeID   | string  | ✅       |                             |
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

response, err := client.AnnotationID.DeleteEnvelopeAnnotation(context.Background(), "envelope_id", "annotation_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```
