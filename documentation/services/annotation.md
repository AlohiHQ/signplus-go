# Annotation

A list of all methods in the `Annotation` service. Click on the method name to view detailed information about that method.

| Methods                                         | Description             |
| :---------------------------------------------- | :---------------------- |
| [AddEnvelopeAnnotation](#addenvelopeannotation) | Add envelope annotation |

## AddEnvelopeAnnotation

Add envelope annotation

- HTTP Method: `POST`
- Endpoint: `/envelope/{envelope_id}/annotation`

**Parameters**

| Name                         | Type                               | Required | Description                   |
| :--------------------------- | :--------------------------------- | :------- | :---------------------------- |
| ctx                          | Context                            | ✅       | Default go language context   |
| envelopeID                   | string                             | ✅       |                               |
| addEnvelopeAnnotationRequest | AddEnvelopeAnnotationRequest       | ✅       |                               |
| params                       | AddEnvelopeAnnotationRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/annotation"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := annotation.AddEnvelopeAnnotationRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}


addEnvelopeAnnotationRequestSignature := annotation.AddEnvelopeAnnotationRequestSignature{
  ID: signplus.Nullable[string]("<string>"),
}


addEnvelopeAnnotationRequestInitials := annotation.AddEnvelopeAnnotationRequestInitials{
  ID: signplus.Nullable[string]("<string>"),
}


textFont1 := annotation.TextFont1{
  Family: signplus.Nullable[string]("SANS"),
  Italic: signplus.Nullable[string]("<boolean>"),
  Bold: signplus.Nullable[string]("<boolean>"),
}

addEnvelopeAnnotationRequestText := annotation.AddEnvelopeAnnotationRequestText{
  Size: signplus.Nullable[string]("<number>"),
  Color: signplus.Nullable[string]("<number>"),
  Value: signplus.Nullable[string]("<string>"),
  Tooltip: signplus.Nullable[string]("<string>"),
  DynamicFieldName: signplus.Nullable[string]("<string>"),
  Font: signplus.Nullable[annotation.TextFont1](textFont1),
}


datetimeFont1 := annotation.DatetimeFont1{
  Family: signplus.Nullable[string]("UNKNOWN"),
  Italic: signplus.Nullable[string]("<boolean>"),
  Bold: signplus.Nullable[string]("<boolean>"),
}

addEnvelopeAnnotationRequestDatetime := annotation.AddEnvelopeAnnotationRequestDatetime{
  Size: signplus.Nullable[string]("<number>"),
  Font: signplus.Nullable[annotation.DatetimeFont1](datetimeFont1),
  Color: signplus.Nullable[string]("<string>"),
  AutoFill: signplus.Nullable[string]("<boolean>"),
  Timezone: signplus.Nullable[string]("<string>"),
  Timestamp: signplus.Nullable[string]("<integer>"),
  Format: signplus.Nullable[string]("YMD_NUMERIC_SLASH"),
}


addEnvelopeAnnotationRequestCheckbox := annotation.AddEnvelopeAnnotationRequestCheckbox{
  Checked: signplus.Nullable[string]("<boolean>"),
  Style: signplus.Nullable[string]("TIMES_SQUARE"),
}

request := annotation.AddEnvelopeAnnotationRequest{
  DocumentID: signplus.Nullable[string]("<string>"),
  Page: signplus.Nullable[string]("<integer>"),
  X: signplus.Nullable[string]("<float>"),
  Y: signplus.Nullable[string]("<float>"),
  Width: signplus.Nullable[string]("<float>"),
  Height: signplus.Nullable[string]("<float>"),
  Type: signplus.Nullable[string]("INITIALS"),
  RecipientID: signplus.Nullable[string]("<string>"),
  Required: signplus.Nullable[string]("<boolean>"),
  Signature: signplus.Nullable[annotation.AddEnvelopeAnnotationRequestSignature](addEnvelopeAnnotationRequestSignature),
  Initials: signplus.Nullable[annotation.AddEnvelopeAnnotationRequestInitials](addEnvelopeAnnotationRequestInitials),
  Text: signplus.Nullable[annotation.AddEnvelopeAnnotationRequestText](addEnvelopeAnnotationRequestText),
  Datetime: signplus.Nullable[annotation.AddEnvelopeAnnotationRequestDatetime](addEnvelopeAnnotationRequestDatetime),
  Checkbox: signplus.Nullable[annotation.AddEnvelopeAnnotationRequestCheckbox](addEnvelopeAnnotationRequestCheckbox),
}

response, err := client.Annotation.AddEnvelopeAnnotation(context.Background(), "envelope_id", request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
