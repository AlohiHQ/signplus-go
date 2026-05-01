# TemplateTemplateIDAttachmentsPlaceholders

A list of all methods in the `TemplateTemplateIDAttachmentsPlaceholders` service. Click on the method name to view detailed information about that method.

| Methods                                                                   | Description                                                     |
| :------------------------------------------------------------------------ | :-------------------------------------------------------------- |
| [SetTemplateAttachmentsPlaceholders](#settemplateattachmentsplaceholders) | Placeholders to be set, completely replacing the existing ones. |

## SetTemplateAttachmentsPlaceholders

Placeholders to be set, completely replacing the existing ones.

- HTTP Method: `PUT`
- Endpoint: `/template/{template_id}/attachments/placeholders`

**Parameters**

| Name                                      | Type                                            | Required | Description                   |
| :---------------------------------------- | :---------------------------------------------- | :------- | :---------------------------- |
| ctx                                       | Context                                         | ✅       | Default go language context   |
| templateID                                | string                                          | ✅       |                               |
| setTemplateAttachmentsPlaceholdersRequest | SetTemplateAttachmentsPlaceholdersRequest       | ✅       |                               |
| params                                    | SetTemplateAttachmentsPlaceholdersRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/templatetemplateidattachmentsplaceholders"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := templatetemplateidattachmentsplaceholders.SetTemplateAttachmentsPlaceholdersRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}


setTemplateAttachmentsPlaceholdersRequestPlaceholders := templatetemplateidattachmentsplaceholders.SetTemplateAttachmentsPlaceholdersRequestPlaceholders{
  RecipientID: signplus.Nullable[string]("<string>"),
  Name: signplus.Nullable[string]("<string>"),
  Required: signplus.Nullable[string]("<boolean>"),
  Multiple: signplus.Nullable[string]("<boolean>"),
  ID: signplus.Nullable[string]("<string>"),
  Hint: signplus.Nullable[string]("<string>"),
}

request := templatetemplateidattachmentsplaceholders.SetTemplateAttachmentsPlaceholdersRequest{
  Placeholders: signplus.Nullable[[]templatetemplateidattachmentsplaceholders.SetTemplateAttachmentsPlaceholdersRequestPlaceholders]([]templatetemplateidattachmentsplaceholders.SetTemplateAttachmentsPlaceholdersRequestPlaceholders{setTemplateAttachmentsPlaceholdersRequestPlaceholders}),
}

response, err := client.TemplateTemplateIDAttachmentsPlaceholders.SetTemplateAttachmentsPlaceholders(context.Background(), "template_id", request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
