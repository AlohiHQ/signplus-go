# Duplicate

A list of all methods in the `Duplicate` service. Click on the method name to view detailed information about that method.

| Methods                                 | Description        |
| :-------------------------------------- | :----------------- |
| [DuplicateEnvelope](#duplicateenvelope) | Duplicate envelope |

## DuplicateEnvelope

Duplicate envelope

- HTTP Method: `POST`
- Endpoint: `/envelope/{envelope_id}/duplicate`

**Parameters**

| Name       | Type                           | Required | Description                   |
| :--------- | :----------------------------- | :------- | :---------------------------- |
| ctx        | Context                        | ✅       | Default go language context   |
| envelopeID | string                         | ✅       |                               |
| params     | DuplicateEnvelopeRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/duplicate"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := duplicate.DuplicateEnvelopeRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}

response, err := client.Duplicate.DuplicateEnvelope(context.Background(), "envelope_id", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
