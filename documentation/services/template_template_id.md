# TemplateTemplateID

A list of all methods in the `TemplateTemplateID` service. Click on the method name to view detailed information about that method.

| Methods                           | Description     |
| :-------------------------------- | :-------------- |
| [GetTemplate](#gettemplate)       | Get template    |
| [DeleteTemplate](#deletetemplate) | Delete template |

## GetTemplate

Get template

- HTTP Method: `GET`
- Endpoint: `/template/{template_id}`

**Parameters**

| Name       | Type                     | Required | Description                   |
| :--------- | :----------------------- | :------- | :---------------------------- |
| ctx        | Context                  | ✅       | Default go language context   |
| templateID | string                   | ✅       |                               |
| params     | GetTemplateRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/templatetemplateid"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := templatetemplateid.GetTemplateRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}

response, err := client.TemplateTemplateID.GetTemplate(context.Background(), "template_id", params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```

## DeleteTemplate

Delete template

- HTTP Method: `DELETE`
- Endpoint: `/template/{template_id}`

**Parameters**

| Name       | Type    | Required | Description                 |
| :--------- | :------ | :------- | :-------------------------- |
| ctx        | Context | ✅       | Default go language context |
| templateID | string  | ✅       |                             |

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

response, err := client.TemplateTemplateID.DeleteTemplate(context.Background(), "template_id")
if err != nil {
  panic(err)
}

fmt.Println(response)
```
