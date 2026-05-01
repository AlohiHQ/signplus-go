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
  ID: signplus.Nullable[string]("string"),
}


addEnvelopeAnnotationRequestInitials := annotation.AddEnvelopeAnnotationRequestInitials{
  ID: signplus.Nullable[string]("string"),
}


textFont1 := annotation.TextFont1{
  Family: signplus.Nullable[string]("SANS"),
  Italic: signplus.Nullable[bool](true),
  Bold: signplus.Nullable[bool](true),
}

addEnvelopeAnnotationRequestText := annotation.AddEnvelopeAnnotationRequestText{
  Size: signplus.Nullable[float64](float64(6190.822136605691)),
  Color: signplus.Nullable[float64](float64(6489.781325519173)),
  Value: signplus.Nullable[string]("string"),
  Tooltip: signplus.Nullable[string]("string"),
  DynamicFieldName: signplus.Nullable[string]("string"),
  Font: signplus.Nullable[annotation.TextFont1](textFont1),
}


datetimeFont1 := annotation.DatetimeFont1{
  Family: signplus.Nullable[string]("SERIF"),
  Italic: signplus.Nullable[bool](true),
  Bold: signplus.Nullable[bool](true),
}

addEnvelopeAnnotationRequestDatetime := annotation.AddEnvelopeAnnotationRequestDatetime{
  Size: signplus.Nullable[float64](float64(3773.1065479576364)),
  Font: signplus.Nullable[annotation.DatetimeFont1](datetimeFont1),
  Color: signplus.Nullable[string]("string"),
  AutoFill: signplus.Nullable[bool](true),
  Timezone: signplus.Nullable[string]("string"),
  Timestamp: signplus.Nullable[float64](float64(6868)),
  Format: signplus.Nullable[string]("MDY_TEXT_SPACE_SHORT"),
}


addEnvelopeAnnotationRequestCheckbox := annotation.AddEnvelopeAnnotationRequestCheckbox{
  Checked: signplus.Nullable[bool](true),
  Style: signplus.Nullable[string]("SQUARE_CHECK"),
}

request := annotation.AddEnvelopeAnnotationRequest{
  DocumentID: signplus.Nullable[string]("string"),
  Page: signplus.Nullable[float64](float64(6387)),
  X: signplus.Nullable[float64](float64(4410.13346533615)),
  Y: signplus.Nullable[float64](float64(5148.888749329143)),
  Width: signplus.Nullable[float64](float64(3756.0248729763225)),
  Height: signplus.Nullable[float64](float64(4178.76189579703)),
  Type: signplus.Nullable[string]("INITIALS"),
  RecipientID: signplus.Nullable[string]("string"),
  Required: signplus.Nullable[bool](true),
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
