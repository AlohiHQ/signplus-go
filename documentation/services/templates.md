# Templates

A list of all methods in the `Templates` service. Click on the method name to view detailed information about that method.

| Methods                         | Description    |
| :------------------------------ | :------------- |
| [ListTemplates](#listtemplates) | List templates |

## ListTemplates

List templates

- HTTP Method: `POST`
- Endpoint: `/templates`

**Parameters**

| Name                 | Type                       | Required | Description                   |
| :------------------- | :------------------------- | :------- | :---------------------------- |
| ctx                  | Context                    | ✅       | Default go language context   |
| listTemplatesRequest | ListTemplatesRequest       | ✅       |                               |
| params               | ListTemplatesRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/templates"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := templates.ListTemplatesRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}


request := templates.ListTemplatesRequest{
  Name: signplus.Nullable[string]("string"),
  Tags: signplus.Nullable[[]string]([]string{}),
  Ids: signplus.Nullable[[]string]([]string{}),
  First: signplus.Nullable[float64](float64(7297)),
  Last: signplus.Nullable[float64](float64(8379)),
  After: signplus.Nullable[string]("string"),
  Before: signplus.Nullable[string]("string"),
  OrderField: signplus.Nullable[string]("TEMPLATE_NAME"),
  Ascending: signplus.Nullable[bool](true),
}

response, err := client.Templates.ListTemplates(context.Background(), request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
