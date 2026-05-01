# Void

A list of all methods in the `Void` service. Click on the method name to view detailed information about that method.

| Methods                       | Description   |
| :---------------------------- | :------------ |
| [VoidEnvelope](#voidenvelope) | Void envelope |

## VoidEnvelope

Void envelope

- HTTP Method: `PUT`
- Endpoint: `/envelope/{envelope_id}/void`

**Parameters**

| Name       | Type                      | Required | Description                   |
| :--------- | :------------------------ | :------- | :---------------------------- |
| ctx        | Context                   | ✅       | Default go language context   |
| envelopeID | string                    | ✅       |                               |
| params     | VoidEnvelopeRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/void"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := void.VoidEnvelopeRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}

response, err := client.Void.VoidEnvelope(context.Background(), "envelope_id", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
