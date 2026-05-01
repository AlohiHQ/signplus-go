# Documents

A list of all methods in the `Documents` service. Click on the method name to view detailed information about that method.

| Methods                                       | Description            |
| :-------------------------------------------- | :--------------------- |
| [GetEnvelopeDocuments](#getenvelopedocuments) | Get envelope documents |

## GetEnvelopeDocuments

Get envelope documents

- HTTP Method: `GET`
- Endpoint: `/envelope/{envelope_id}/documents`

**Parameters**

| Name       | Type                              | Required | Description                   |
| :--------- | :-------------------------------- | :------- | :---------------------------- |
| ctx        | Context                           | ✅       | Default go language context   |
| envelopeID | string                            | ✅       |                               |
| params     | GetEnvelopeDocumentsRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/documents"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := documents.GetEnvelopeDocumentsRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}

response, err := client.Documents.GetEnvelopeDocuments(context.Background(), "envelope_id", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
