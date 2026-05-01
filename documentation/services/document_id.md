# DocumentID

A list of all methods in the `DocumentID` service. Click on the method name to view detailed information about that method.

| Methods                                     | Description           |
| :------------------------------------------ | :-------------------- |
| [GetEnvelopeDocument](#getenvelopedocument) | Get envelope document |

## GetEnvelopeDocument

Get envelope document

- HTTP Method: `GET`
- Endpoint: `/envelope/{envelope_id}/document/{document_id}`

**Parameters**

| Name       | Type                             | Required | Description                   |
| :--------- | :------------------------------- | :------- | :---------------------------- |
| ctx        | Context                          | ✅       | Default go language context   |
| envelopeID | string                           | ✅       |                               |
| documentID | string                           | ✅       |                               |
| params     | GetEnvelopeDocumentRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/documentid"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := documentid.GetEnvelopeDocumentRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}

response, err := client.DocumentID.GetEnvelopeDocument(context.Background(), "envelope_id", "document_id", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
