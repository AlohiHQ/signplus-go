# SetLegalityLevel

A list of all methods in the `SetLegalityLevel` service. Click on the method name to view detailed information about that method.

| Methods                                               | Description                 |
| :---------------------------------------------------- | :-------------------------- |
| [SetEnvelopeLegalityLevel](#setenvelopelegalitylevel) | Set envelope legality level |

## SetEnvelopeLegalityLevel

Set envelope legality level

- HTTP Method: `PUT`
- Endpoint: `/envelope/{envelope_id}/set_legality_level`

**Parameters**

| Name                            | Type                                  | Required | Description                   |
| :------------------------------ | :------------------------------------ | :------- | :---------------------------- |
| ctx                             | Context                               | ✅       | Default go language context   |
| envelopeID                      | string                                | ✅       |                               |
| setEnvelopeLegalityLevelRequest | SetEnvelopeLegalityLevelRequest       | ✅       |                               |
| params                          | SetEnvelopeLegalityLevelRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/setlegalitylevel"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := setlegalitylevel.SetEnvelopeLegalityLevelRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}


request := setlegalitylevel.SetEnvelopeLegalityLevelRequest{
  LegalityLevel: signplus.Nullable[string]("QES_EIDAS"),
}

response, err := client.SetLegalityLevel.SetEnvelopeLegalityLevel(context.Background(), "envelope_id", request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
