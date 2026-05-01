# DynamicFields

A list of all methods in the `DynamicFields` service. Click on the method name to view detailed information about that method.

| Methods                                               | Description                 |
| :---------------------------------------------------- | :-------------------------- |
| [SetEnvelopeDynamicFields](#setenvelopedynamicfields) | Set envelope dynamic fields |

## SetEnvelopeDynamicFields

Set envelope dynamic fields

- HTTP Method: `PUT`
- Endpoint: `/envelope/{envelope_id}/dynamic_fields`

**Parameters**

| Name                            | Type                                  | Required | Description                   |
| :------------------------------ | :------------------------------------ | :------- | :---------------------------- |
| ctx                             | Context                               | ✅       | Default go language context   |
| envelopeID                      | string                                | ✅       |                               |
| setEnvelopeDynamicFieldsRequest | SetEnvelopeDynamicFieldsRequest       | ✅       |                               |
| params                          | SetEnvelopeDynamicFieldsRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/dynamicfields"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := dynamicfields.SetEnvelopeDynamicFieldsRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}


dynamicFields := dynamicfields.DynamicFields{
  Name: signplus.Nullable[string]("<string>"),
  Value: signplus.Nullable[string]("<string>"),
}

request := dynamicfields.SetEnvelopeDynamicFieldsRequest{
  DynamicFields: signplus.Nullable[[]dynamicfields.DynamicFields]([]dynamicfields.DynamicFields{dynamicFields}),
}

response, err := client.DynamicFields.SetEnvelopeDynamicFields(context.Background(), "envelope_id", request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
