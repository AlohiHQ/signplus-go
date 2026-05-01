# Rename

A list of all methods in the `Rename` service. Click on the method name to view detailed information about that method.

| Methods                           | Description     |
| :-------------------------------- | :-------------- |
| [RenameEnvelope](#renameenvelope) | Rename envelope |

## RenameEnvelope

Rename envelope

- HTTP Method: `PUT`
- Endpoint: `/envelope/{envelope_id}/rename`

**Parameters**

| Name                  | Type                        | Required | Description                   |
| :-------------------- | :-------------------------- | :------- | :---------------------------- |
| ctx                   | Context                     | ✅       | Default go language context   |
| envelopeID            | string                      | ✅       |                               |
| renameEnvelopeRequest | RenameEnvelopeRequest       | ✅       |                               |
| params                | RenameEnvelopeRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/rename"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := rename.RenameEnvelopeRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}


request := rename.RenameEnvelopeRequest{
  Name: signplus.Nullable[string]("<string>"),
}

response, err := client.Rename.RenameEnvelope(context.Background(), "envelope_id", request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
