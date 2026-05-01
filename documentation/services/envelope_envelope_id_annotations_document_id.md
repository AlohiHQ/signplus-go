# EnvelopeEnvelopeIDAnnotationsDocumentID

A list of all methods in the `EnvelopeEnvelopeIDAnnotationsDocumentID` service. Click on the method name to view detailed information about that method.

| Methods                                                           | Description                       |
| :---------------------------------------------------------------- | :-------------------------------- |
| [GetEnvelopeDocumentAnnotations](#getenvelopedocumentannotations) | Get envelope document annotations |

## GetEnvelopeDocumentAnnotations

Get envelope document annotations

- HTTP Method: `GET`
- Endpoint: `/envelope/{envelope_id}/annotations/{document_id}`

**Parameters**

| Name       | Type                                        | Required | Description                   |
| :--------- | :------------------------------------------ | :------- | :---------------------------- |
| ctx        | Context                                     | ✅       | Default go language context   |
| envelopeID | string                                      | ✅       |                               |
| documentID | string                                      | ✅       |                               |
| params     | GetEnvelopeDocumentAnnotationsRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/envelopeenvelopeidannotationsdocumentid"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := envelopeenvelopeidannotationsdocumentid.GetEnvelopeDocumentAnnotationsRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}

response, err := client.EnvelopeEnvelopeIDAnnotationsDocumentID.GetEnvelopeDocumentAnnotations(context.Background(), "envelope_id", "document_id", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
