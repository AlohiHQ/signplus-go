# EnvelopeID

A list of all methods in the `EnvelopeID` service. Click on the method name to view detailed information about that method.

| Methods                           | Description     |
| :-------------------------------- | :-------------- |
| [GetEnvelope](#getenvelope)       | Get envelope    |
| [DeleteEnvelope](#deleteenvelope) | Delete envelope |

## GetEnvelope

Get envelope

- HTTP Method: `GET`
- Endpoint: `/envelope/{envelope_id}`

**Parameters**

| Name       | Type                     | Required | Description                   |
| :--------- | :----------------------- | :------- | :---------------------------- |
| ctx        | Context                  | ✅       | Default go language context   |
| envelopeID | string                   | ✅       |                               |
| params     | GetEnvelopeRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/envelopeid"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := envelopeid.GetEnvelopeRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}

response, err := client.EnvelopeID.GetEnvelope(context.Background(), "envelope_id", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteEnvelope

Delete envelope

- HTTP Method: `DELETE`
- Endpoint: `/envelope/{envelope_id}`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| envelopeID | string  | ✅       |                             |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)

response, err := client.EnvelopeID.DeleteEnvelope(context.Background(), "envelope_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```
