# SetExpirationDate

A list of all methods in the `SetExpirationDate` service. Click on the method name to view detailed information about that method.

| Methods                                                 | Description                  |
| :------------------------------------------------------ | :--------------------------- |
| [SetEnvelopeExpirationDate](#setenvelopeexpirationdate) | Set envelope expiration date |

## SetEnvelopeExpirationDate

Set envelope expiration date

- HTTP Method: `PUT`
- Endpoint: `/envelope/{envelope_id}/set_expiration_date`

**Parameters**

| Name                             | Type                                   | Required | Description                   |
| :------------------------------- | :------------------------------------- | :------- | :---------------------------- |
| ctx                              | Context                                | ✅       | Default go language context   |
| envelopeID                       | string                                 | ✅       |                               |
| setEnvelopeExpirationDateRequest | SetEnvelopeExpirationDateRequest       | ✅       |                               |
| params                           | SetEnvelopeExpirationDateRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/setexpirationdate"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := setexpirationdate.SetEnvelopeExpirationDateRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}


request := setexpirationdate.SetEnvelopeExpirationDateRequest{
  ExpiresAt: signplus.Nullable[float64](float64(5958)),
}

response, err := client.SetExpirationDate.SetEnvelopeExpirationDate(context.Background(), "envelope_id", request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
