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
  ID: signplus.Nullable[string]("string"),
}


addTemplateAnnotationRequestInitials := templatetemplateidannotation.AddTemplateAnnotationRequestInitials{
  ID: signplus.Nullable[string]("string"),
}


textFont2 := templatetemplateidannotation.TextFont2{
  Family: signplus.Nullable[string]("SANS"),
  Italic: signplus.Nullable[bool](true),
  Bold: signplus.Nullable[bool](true),
}

addTemplateAnnotationRequestText := templatetemplateidannotation.AddTemplateAnnotationRequestText{
  Size: signplus.Nullable[float64](float64(6190.822136605691)),
  Color: signplus.Nullable[float64](float64(6489.781325519173)),
  Value: signplus.Nullable[string]("string"),
  Tooltip: signplus.Nullable[string]("string"),
  DynamicFieldName: signplus.Nullable[string]("string"),
  Font: signplus.Nullable[templatetemplateidannotation.TextFont2](textFont2),
}


datetimeFont2 := templatetemplateidannotation.DatetimeFont2{
  Family: signplus.Nullable[string]("SERIF"),
  Italic: signplus.Nullable[bool](true),
  Bold: signplus.Nullable[bool](true),
}

addTemplateAnnotationRequestDatetime := templatetemplateidannotation.AddTemplateAnnotationRequestDatetime{
  Size: signplus.Nullable[float64](float64(3773.1065479576364)),
  Font: signplus.Nullable[templatetemplateidannotation.DatetimeFont2](datetimeFont2),
  Color: signplus.Nullable[string]("string"),
  AutoFill: signplus.Nullable[bool](true),
  Timezone: signplus.Nullable[string]("string"),
  Timestamp: signplus.Nullable[float64](float64(6868)),
  Format: signplus.Nullable[string]("MDY_TEXT_SPACE_SHORT"),
}


addTemplateAnnotationRequestCheckbox := templatetemplateidannotation.AddTemplateAnnotationRequestCheckbox{
  Checked: signplus.Nullable[bool](true),
  Style: signplus.Nullable[string]("SQUARE_CHECK"),
}

request := templatetemplateidannotation.AddTemplateAnnotationRequest{
  DocumentID: signplus.Nullable[string]("string"),
  Page: signplus.Nullable[float64](float64(6387)),
  X: signplus.Nullable[float64](float64(4410.13346533615)),
  Y: signplus.Nullable[float64](float64(5148.888749329143)),
  Width: signplus.Nullable[float64](float64(3756.0248729763225)),
  Height: signplus.Nullable[float64](float64(4178.76189579703)),
  Type: signplus.Nullable[string]("INITIALS"),
  RecipientID: signplus.Nullable[string]("string"),
  Required: signplus.Nullable[bool](true),
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
