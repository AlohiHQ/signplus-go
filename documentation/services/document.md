# Document

A list of all methods in the `Document` service. Click on the method name to view detailed information about that method.

| Methods                                     | Description           |
| :------------------------------------------ | :-------------------- |
| [AddEnvelopeDocument](#addenvelopedocument) | Add envelope document |

## AddEnvelopeDocument

Add envelope document

- HTTP Method: `POST`
- Endpoint: `/envelope/{envelope_id}/document`

**Parameters**

| Name                       | Type                             | Required | Description                   |
| :------------------------- | :------------------------------- | :------- | :---------------------------- |
| ctx                        | Context                          | ✅       | Default go language context   |
| envelopeID                 | string                           | ✅       |                               |
| addEnvelopeDocumentRequest | AddEnvelopeDocumentRequest       | ✅       |                               |
| params                     | AddEnvelopeDocumentRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/document"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := document.AddEnvelopeDocumentRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}


request := document.AddEnvelopeDocumentRequest{
  File: signplus.Nullable[[]byte]([]byte{}),
}

response, err := client.Document.AddEnvelopeDocument(context.Background(), "envelope_id", request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
