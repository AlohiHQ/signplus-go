# TemplateTemplateIDAnnotation

A list of all methods in the `TemplateTemplateIDAnnotation` service. Click on the method name to view detailed information about that method.

| Methods                                         | Description             |
| :---------------------------------------------- | :---------------------- |
| [AddTemplateAnnotation](#addtemplateannotation) | Add template annotation |

## AddTemplateAnnotation

Add template annotation

- HTTP Method: `POST`
- Endpoint: `/template/{template_id}/annotation`

**Parameters**

| Name                         | Type                               | Required | Description                   |
| :--------------------------- | :--------------------------------- | :------- | :---------------------------- |
| ctx                          | Context                            | ✅       | Default go language context   |
| templateID                   | string                             | ✅       |                               |
| addTemplateAnnotationRequest | AddTemplateAnnotationRequest       | ✅       |                               |
| params                       | AddTemplateAnnotationRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/templatetemplateidannotation"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := templatetemplateidannotation.AddTemplateAnnotationRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}


addTemplateAnnotationRequestSignature := templatetemplateidannotation.AddTemplateAnnotationRequestSignature{
  ID: signplus.Nullable[string]("<string>"),
}


addTemplateAnnotationRequestInitials := templatetemplateidannotation.AddTemplateAnnotationRequestInitials{
  ID: signplus.Nullable[string]("<string>"),
}


textFont2 := templatetemplateidannotation.TextFont2{
  Family: signplus.Nullable[string]("SANS"),
  Italic: signplus.Nullable[string]("<boolean>"),
  Bold: signplus.Nullable[string]("<boolean>"),
}

addTemplateAnnotationRequestText := templatetemplateidannotation.AddTemplateAnnotationRequestText{
  Size: signplus.Nullable[string]("<number>"),
  Color: signplus.Nullable[string]("<number>"),
  Value: signplus.Nullable[string]("<string>"),
  Tooltip: signplus.Nullable[string]("<string>"),
  DynamicFieldName: signplus.Nullable[string]("<string>"),
  Font: signplus.Nullable[templatetemplateidannotation.TextFont2](textFont2),
}


datetimeFont2 := templatetemplateidannotation.DatetimeFont2{
  Family: signplus.Nullable[string]("UNKNOWN"),
  Italic: signplus.Nullable[string]("<boolean>"),
  Bold: signplus.Nullable[string]("<boolean>"),
}

addTemplateAnnotationRequestDatetime := templatetemplateidannotation.AddTemplateAnnotationRequestDatetime{
  Size: signplus.Nullable[string]("<number>"),
  Font: signplus.Nullable[templatetemplateidannotation.DatetimeFont2](datetimeFont2),
  Color: signplus.Nullable[string]("<string>"),
  AutoFill: signplus.Nullable[string]("<boolean>"),
  Timezone: signplus.Nullable[string]("<string>"),
  Timestamp: signplus.Nullable[string]("<integer>"),
  Format: signplus.Nullable[string]("YMD_NUMERIC_SLASH"),
}


addTemplateAnnotationRequestCheckbox := templatetemplateidannotation.AddTemplateAnnotationRequestCheckbox{
  Checked: signplus.Nullable[string]("<boolean>"),
  Style: signplus.Nullable[string]("TIMES_SQUARE"),
}

request := templatetemplateidannotation.AddTemplateAnnotationRequest{
  DocumentID: signplus.Nullable[string]("<string>"),
  Page: signplus.Nullable[string]("<integer>"),
  X: signplus.Nullable[string]("<float>"),
  Y: signplus.Nullable[string]("<float>"),
  Width: signplus.Nullable[string]("<float>"),
  Height: signplus.Nullable[string]("<float>"),
  Type: signplus.Nullable[string]("INITIALS"),
  RecipientID: signplus.Nullable[string]("<string>"),
  Required: signplus.Nullable[string]("<boolean>"),
  Signature: signplus.Nullable[templatetemplateidannotation.AddTemplateAnnotationRequestSignature](addTemplateAnnotationRequestSignature),
  Initials: signplus.Nullable[templatetemplateidannotation.AddTemplateAnnotationRequestInitials](addTemplateAnnotationRequestInitials),
  Text: signplus.Nullable[templatetemplateidannotation.AddTemplateAnnotationRequestText](addTemplateAnnotationRequestText),
  Datetime: signplus.Nullable[templatetemplateidannotation.AddTemplateAnnotationRequestDatetime](addTemplateAnnotationRequestDatetime),
  Checkbox: signplus.Nullable[templatetemplateidannotation.AddTemplateAnnotationRequestCheckbox](addTemplateAnnotationRequestCheckbox),
}

response, err := client.TemplateTemplateIDAnnotation.AddTemplateAnnotation(context.Background(), "template_id", request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
