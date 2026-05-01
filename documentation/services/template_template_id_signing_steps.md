# TemplateTemplateIDSigningSteps

A list of all methods in the `TemplateTemplateIDSigningSteps` service. Click on the method name to view detailed information about that method.

| Methods                                             | Description                |
| :-------------------------------------------------- | :------------------------- |
| [AddTemplateSigningSteps](#addtemplatesigningsteps) | Add template signing steps |

## AddTemplateSigningSteps

Add template signing steps

- HTTP Method: `POST`
- Endpoint: `/template/{template_id}/signing_steps`

**Parameters**

| Name                           | Type                                 | Required | Description                   |
| :----------------------------- | :----------------------------------- | :------- | :---------------------------- |
| ctx                            | Context                              | ✅       | Default go language context   |
| templateID                     | string                               | ✅       |                               |
| addTemplateSigningStepsRequest | AddTemplateSigningStepsRequest       | ✅       |                               |
| params                         | AddTemplateSigningStepsRequestParams | ✅       | Additional request parameters |

**Return Type**

`[]byte`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "context"
  "github.com/alohihq/signplus-go"
  "github.com/alohihq/signplus-go/templatetemplateidsigningsteps"
)

config := signplus.NewConfig()
config.SetAccessToken("ACCESS_TOKEN")
client := signplus.NewSignplus(config)


params := templatetemplateidsigningsteps.AddTemplateSigningStepsRequestParams{
  Accept: signplus.Nullable[string]("application/json"),
}


signingStepsRecipients2 := templatetemplateidsigningsteps.SigningStepsRecipients2{
  ID: signplus.Nullable[string]("string"),
  UID: signplus.Nullable[string]("string"),
  Name: signplus.Nullable[string]("string"),
  Email: signplus.Nullable[string]("string"),
  Role: signplus.Nullable[string]("RECEIVES_COPY"),
}

addTemplateSigningStepsRequestSigningSteps := templatetemplateidsigningsteps.AddTemplateSigningStepsRequestSigningSteps{
  Recipients: signplus.Nullable[[]templatetemplateidsigningsteps.SigningStepsRecipients2]([]templatetemplateidsigningsteps.SigningStepsRecipients2{signingStepsRecipients2}),
}

request := templatetemplateidsigningsteps.AddTemplateSigningStepsRequest{
  SigningSteps: signplus.Nullable[[]templatetemplateidsigningsteps.AddTemplateSigningStepsRequestSigningSteps]([]templatetemplateidsigningsteps.AddTemplateSigningStepsRequestSigningSteps{addTemplateSigningStepsRequestSigningSteps}),
}

response, err := client.TemplateTemplateIDSigningSteps.AddTemplateSigningSteps(context.Background(), "template_id", request, params)
if err != nil {
  panic(err)
}

fmt.Println(response)
```
