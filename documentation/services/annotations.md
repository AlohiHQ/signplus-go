# Annotations

A list of all methods in the `Annotations` service. Click on the method name to view detailed information about that method.

| Methods                                           | Description              |
| :------------------------------------------------ | :----------------------- |
| [GetEnvelopeAnnotations](#getenvelopeannotations) | Get envelope annotations |

## GetEnvelopeAnnotations

Get envelope annotations

- HTTP Method: `GET`
- Endpoint: `/envelope/{envelope_id}/annotations`

**Parameters**

| Name       | Type                                | Required | Description                   |
| :--------- | :---------------------------------- | :------- | :---------------------------- |
| ctx        | Context                             | ✅       | Default go language context   |
| envelopeID | string                              | ✅       |                               |
| params     | GetEnvelopeAnnotationsRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/annotations"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := annotations.GetEnvelopeAnnotationsRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}

response, err := client.Annotations.GetEnvelopeAnnotations(context.Background(), "envelope_id", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
