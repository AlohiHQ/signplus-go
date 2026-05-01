# Send

A list of all methods in the `Send` service. Click on the method name to view detailed information about that method.

| Methods                       | Description                 |
| :---------------------------- | :-------------------------- |
| [SendEnvelope](#sendenvelope) | Send envelope for signature |

## SendEnvelope

Send envelope for signature

- HTTP Method: `POST`
- Endpoint: `/envelope/{envelope_id}/send`

**Parameters**

| Name       | Type                      | Required | Description                   |
| :--------- | :------------------------ | :------- | :---------------------------- |
| ctx        | Context                   | ✅       | Default go language context   |
| envelopeID | string                    | ✅       |                               |
| params     | SendEnvelopeRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/send"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := send.SendEnvelopeRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}

response, err := client.Send.SendEnvelope(context.Background(), "envelope_id", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
