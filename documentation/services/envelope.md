# Envelope

A list of all methods in the `Envelope` service. Click on the method name to view detailed information about that method.

| Methods                           | Description         |
| :-------------------------------- | :------------------ |
| [CreateEnvelope](#createenvelope) | Create new envelope |

## CreateEnvelope

Create new envelope

- HTTP Method: `POST`
- Endpoint: `/envelope`

**Parameters**

| Name                  | Type                        | Required | Description                   |
| :-------------------- | :-------------------------- | :------- | :---------------------------- |
| ctx                   | Context                     | ✅       | Default go language context   |
| createEnvelopeRequest | CreateEnvelopeRequest       | ✅       |                               |
| params                | CreateEnvelopeRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/envelope"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := envelope.CreateEnvelopeRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}


request := envelope.CreateEnvelopeRequest{
  Name: signplus.Nullable[string]("PNeiZUQ9LB9"),
  LegalityLevel: signplus.Nullable[string]("QES_ZERTES"),
  ExpiresAt: signplus.Nullable[string]("<integer>"),
  Comment: signplus.Nullable[string]("<string>"),
  Sandbox: signplus.Nullable[bool](true),
}

response, err := client.Envelope.CreateEnvelope(context.Background(), request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
