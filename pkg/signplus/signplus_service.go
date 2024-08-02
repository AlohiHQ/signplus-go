package signplus

import (
	"context"
	restClient "github.com/alohihq/signplus-go/internal/clients/rest"
	"github.com/alohihq/signplus-go/internal/clients/rest/httptransport"
	"github.com/alohihq/signplus-go/internal/configmanager"
	"github.com/alohihq/signplus-go/pkg/shared"
	"github.com/alohihq/signplus-go/pkg/signplusconfig"
)

type SignplusService struct {
	manager *configmanager.ConfigManager
}

func NewSignplusService(manager *configmanager.ConfigManager) *SignplusService {
	return &SignplusService{
		manager: manager,
	}
}

func (api *SignplusService) getConfig() *signplusconfig.Config {
	return api.manager.GetSignplus()
}

func (api *SignplusService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *SignplusService) SetAccessToken(accessToken string) {
	config := api.getConfig()
	config.SetAccessToken(accessToken)
}

// Create new envelope
func (api *SignplusService) CreateEnvelope(ctx context.Context, createEnvelopeRequest CreateEnvelopeRequest) (*shared.SignplusResponse[Envelope], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Envelope](config)

	request := httptransport.NewRequest(ctx, "POST", "/envelope", config)

	request.Body = createEnvelopeRequest

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[Envelope](err)
	}

	return shared.NewSignplusResponse[Envelope](resp), nil
}

// Create new envelope from template
func (api *SignplusService) CreateEnvelopeFromTemplate(ctx context.Context, templateId string, createEnvelopeFromTemplateRequest CreateEnvelopeFromTemplateRequest) (*shared.SignplusResponse[Envelope], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Envelope](config)

	request := httptransport.NewRequest(ctx, "POST", "/envelope/from_template/{template_id}", config)

	request.Body = createEnvelopeFromTemplateRequest

	request.SetPathParam("template_id", templateId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[Envelope](err)
	}

	return shared.NewSignplusResponse[Envelope](resp), nil
}

// List envelopes
func (api *SignplusService) ListEnvelopes(ctx context.Context, listEnvelopesRequest ListEnvelopesRequest) (*shared.SignplusResponse[ListEnvelopesResponse], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[ListEnvelopesResponse](config)

	request := httptransport.NewRequest(ctx, "POST", "/envelopes", config)

	request.Body = listEnvelopesRequest

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[ListEnvelopesResponse](err)
	}

	return shared.NewSignplusResponse[ListEnvelopesResponse](resp), nil
}

// Get envelope
func (api *SignplusService) GetEnvelope(ctx context.Context, envelopeId string) (*shared.SignplusResponse[Envelope], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Envelope](config)

	request := httptransport.NewRequest(ctx, "GET", "/envelope/{envelope_id}", config)

	request.SetPathParam("envelope_id", envelopeId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[Envelope](err)
	}

	return shared.NewSignplusResponse[Envelope](resp), nil
}

// Delete envelope
func (api *SignplusService) DeleteEnvelope(ctx context.Context, envelopeId string) (*shared.SignplusResponse[any], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[any](config)

	request := httptransport.NewRequest(ctx, "DELETE", "/envelope/{envelope_id}", config)

	request.SetPathParam("envelope_id", envelopeId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[any](err)
	}

	return shared.NewSignplusResponse[any](resp), nil
}

// Get envelope document
func (api *SignplusService) GetEnvelopeDocument(ctx context.Context, envelopeId string, documentId string) (*shared.SignplusResponse[Document], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Document](config)

	request := httptransport.NewRequest(ctx, "GET", "/envelope/{envelope_id}/document/{document_id}", config)

	request.SetPathParam("envelope_id", envelopeId)
	request.SetPathParam("document_id", documentId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[Document](err)
	}

	return shared.NewSignplusResponse[Document](resp), nil
}

// Get envelope documents
func (api *SignplusService) GetEnvelopeDocuments(ctx context.Context, envelopeId string) (*shared.SignplusResponse[ListEnvelopeDocumentsResponse], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[ListEnvelopeDocumentsResponse](config)

	request := httptransport.NewRequest(ctx, "GET", "/envelope/{envelope_id}/documents", config)

	request.SetPathParam("envelope_id", envelopeId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[ListEnvelopeDocumentsResponse](err)
	}

	return shared.NewSignplusResponse[ListEnvelopeDocumentsResponse](resp), nil
}

// Add envelope document
func (api *SignplusService) AddEnvelopeDocument(ctx context.Context, envelopeId string, addEnvelopeDocumentRequest AddEnvelopeDocumentRequest) (*shared.SignplusResponse[Document], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Document](config)

	request := httptransport.NewRequest(ctx, "POST", "/envelope/{envelope_id}/document", config)

	request.Body = addEnvelopeDocumentRequest

	request.SetPathParam("envelope_id", envelopeId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[Document](err)
	}

	return shared.NewSignplusResponse[Document](resp), nil
}

// Set envelope dynamic fields
func (api *SignplusService) SetEnvelopeDynamicFields(ctx context.Context, envelopeId string, setEnvelopeDynamicFieldsRequest SetEnvelopeDynamicFieldsRequest) (*shared.SignplusResponse[Envelope], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Envelope](config)

	request := httptransport.NewRequest(ctx, "PUT", "/envelope/{envelope_id}/dynamic_fields", config)

	request.Body = setEnvelopeDynamicFieldsRequest

	request.SetPathParam("envelope_id", envelopeId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[Envelope](err)
	}

	return shared.NewSignplusResponse[Envelope](resp), nil
}

// Add envelope signing steps
func (api *SignplusService) AddEnvelopeSigningSteps(ctx context.Context, envelopeId string, addEnvelopeSigningStepsRequest AddEnvelopeSigningStepsRequest) (*shared.SignplusResponse[Envelope], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Envelope](config)

	request := httptransport.NewRequest(ctx, "POST", "/envelope/{envelope_id}/signing_steps", config)

	request.Body = addEnvelopeSigningStepsRequest

	request.SetPathParam("envelope_id", envelopeId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[Envelope](err)
	}

	return shared.NewSignplusResponse[Envelope](resp), nil
}

// Send envelope for signature
func (api *SignplusService) SendEnvelope(ctx context.Context, envelopeId string) (*shared.SignplusResponse[Envelope], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Envelope](config)

	request := httptransport.NewRequest(ctx, "POST", "/envelope/{envelope_id}/send", config)

	request.SetPathParam("envelope_id", envelopeId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[Envelope](err)
	}

	return shared.NewSignplusResponse[Envelope](resp), nil
}

// Duplicate envelope
func (api *SignplusService) DuplicateEnvelope(ctx context.Context, envelopeId string) (*shared.SignplusResponse[Envelope], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Envelope](config)

	request := httptransport.NewRequest(ctx, "POST", "/envelope/{envelope_id}/duplicate", config)

	request.SetPathParam("envelope_id", envelopeId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[Envelope](err)
	}

	return shared.NewSignplusResponse[Envelope](resp), nil
}

// Void envelope
func (api *SignplusService) VoidEnvelope(ctx context.Context, envelopeId string) (*shared.SignplusResponse[Envelope], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Envelope](config)

	request := httptransport.NewRequest(ctx, "PUT", "/envelope/{envelope_id}/void", config)

	request.SetPathParam("envelope_id", envelopeId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[Envelope](err)
	}

	return shared.NewSignplusResponse[Envelope](resp), nil
}

// Rename envelope
func (api *SignplusService) RenameEnvelope(ctx context.Context, envelopeId string, renameEnvelopeRequest RenameEnvelopeRequest) (*shared.SignplusResponse[Envelope], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Envelope](config)

	request := httptransport.NewRequest(ctx, "PUT", "/envelope/{envelope_id}/rename", config)

	request.Body = renameEnvelopeRequest

	request.SetPathParam("envelope_id", envelopeId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[Envelope](err)
	}

	return shared.NewSignplusResponse[Envelope](resp), nil
}

// Set envelope comment
func (api *SignplusService) SetEnvelopeComment(ctx context.Context, envelopeId string, setEnvelopeCommentRequest SetEnvelopeCommentRequest) (*shared.SignplusResponse[Envelope], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Envelope](config)

	request := httptransport.NewRequest(ctx, "PUT", "/envelope/{envelope_id}/set_comment", config)

	request.Body = setEnvelopeCommentRequest

	request.SetPathParam("envelope_id", envelopeId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[Envelope](err)
	}

	return shared.NewSignplusResponse[Envelope](resp), nil
}

// Set envelope notification
func (api *SignplusService) SetEnvelopeNotification(ctx context.Context, envelopeId string, envelopeNotification EnvelopeNotification) (*shared.SignplusResponse[Envelope], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Envelope](config)

	request := httptransport.NewRequest(ctx, "PUT", "/envelope/{envelope_id}/set_notification", config)

	request.Body = envelopeNotification

	request.SetPathParam("envelope_id", envelopeId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[Envelope](err)
	}

	return shared.NewSignplusResponse[Envelope](resp), nil
}

// Set envelope expiration date
func (api *SignplusService) SetEnvelopeExpirationDate(ctx context.Context, envelopeId string, setEnvelopeExpirationRequest SetEnvelopeExpirationRequest) (*shared.SignplusResponse[Envelope], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Envelope](config)

	request := httptransport.NewRequest(ctx, "PUT", "/envelope/{envelope_id}/set_expiration_date", config)

	request.Body = setEnvelopeExpirationRequest

	request.SetPathParam("envelope_id", envelopeId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[Envelope](err)
	}

	return shared.NewSignplusResponse[Envelope](resp), nil
}

// Set envelope legality level
func (api *SignplusService) SetEnvelopeLegalityLevel(ctx context.Context, envelopeId string, setEnvelopeLegalityLevelRequest SetEnvelopeLegalityLevelRequest) (*shared.SignplusResponse[Envelope], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Envelope](config)

	request := httptransport.NewRequest(ctx, "PUT", "/envelope/{envelope_id}/set_legality_level", config)

	request.Body = setEnvelopeLegalityLevelRequest

	request.SetPathParam("envelope_id", envelopeId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[Envelope](err)
	}

	return shared.NewSignplusResponse[Envelope](resp), nil
}

// Get envelope annotations
func (api *SignplusService) GetEnvelopeAnnotations(ctx context.Context, envelopeId string) (*shared.SignplusResponse[[]Annotation], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[[]Annotation](config)

	request := httptransport.NewRequest(ctx, "GET", "/envelope/{envelope_id}/annotations", config)

	request.SetPathParam("envelope_id", envelopeId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[[]Annotation](err)
	}

	return shared.NewSignplusResponse[[]Annotation](resp), nil
}

// Get envelope document annotations
func (api *SignplusService) GetEnvelopeDocumentAnnotations(ctx context.Context, envelopeId string, documentId string) (*shared.SignplusResponse[ListEnvelopeDocumentAnnotationsResponse], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[ListEnvelopeDocumentAnnotationsResponse](config)

	request := httptransport.NewRequest(ctx, "GET", "/envelope/{envelope_id}/annotations/{document_id}", config)

	request.SetPathParam("envelope_id", envelopeId)
	request.SetPathParam("document_id", documentId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[ListEnvelopeDocumentAnnotationsResponse](err)
	}

	return shared.NewSignplusResponse[ListEnvelopeDocumentAnnotationsResponse](resp), nil
}

// Add envelope annotation
func (api *SignplusService) AddEnvelopeAnnotation(ctx context.Context, envelopeId string, addAnnotationRequest AddAnnotationRequest) (*shared.SignplusResponse[Annotation], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Annotation](config)

	request := httptransport.NewRequest(ctx, "POST", "/envelope/{envelope_id}/annotation", config)

	request.Body = addAnnotationRequest

	request.SetPathParam("envelope_id", envelopeId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[Annotation](err)
	}

	return shared.NewSignplusResponse[Annotation](resp), nil
}

// Delete envelope annotation
func (api *SignplusService) DeleteEnvelopeAnnotation(ctx context.Context, envelopeId string, annotationId string) (*shared.SignplusResponse[any], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[any](config)

	request := httptransport.NewRequest(ctx, "DELETE", "/envelope/{envelope_id}/annotation/{annotation_id}", config)

	request.SetPathParam("envelope_id", envelopeId)
	request.SetPathParam("annotation_id", annotationId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[any](err)
	}

	return shared.NewSignplusResponse[any](resp), nil
}

// Create new template
func (api *SignplusService) CreateTemplate(ctx context.Context, createTemplateRequest CreateTemplateRequest) (*shared.SignplusResponse[Template], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Template](config)

	request := httptransport.NewRequest(ctx, "POST", "/template", config)

	request.Body = createTemplateRequest

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[Template](err)
	}

	return shared.NewSignplusResponse[Template](resp), nil
}

// List templates
func (api *SignplusService) ListTemplates(ctx context.Context, listTemplatesRequest ListTemplatesRequest) (*shared.SignplusResponse[ListTemplatesResponse], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[ListTemplatesResponse](config)

	request := httptransport.NewRequest(ctx, "POST", "/templates", config)

	request.Body = listTemplatesRequest

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[ListTemplatesResponse](err)
	}

	return shared.NewSignplusResponse[ListTemplatesResponse](resp), nil
}

// Get template
func (api *SignplusService) GetTemplate(ctx context.Context, templateId string) (*shared.SignplusResponse[Template], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Template](config)

	request := httptransport.NewRequest(ctx, "GET", "/template/{template_id}", config)

	request.SetPathParam("template_id", templateId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[Template](err)
	}

	return shared.NewSignplusResponse[Template](resp), nil
}

// Delete template
func (api *SignplusService) DeleteTemplate(ctx context.Context, templateId string) (*shared.SignplusResponse[any], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[any](config)

	request := httptransport.NewRequest(ctx, "DELETE", "/template/{template_id}", config)

	request.SetPathParam("template_id", templateId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[any](err)
	}

	return shared.NewSignplusResponse[any](resp), nil
}

// Duplicate template
func (api *SignplusService) DuplicateTemplate(ctx context.Context, templateId string) (*shared.SignplusResponse[Template], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Template](config)

	request := httptransport.NewRequest(ctx, "POST", "/template/{template_id}/duplicate", config)

	request.SetPathParam("template_id", templateId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[Template](err)
	}

	return shared.NewSignplusResponse[Template](resp), nil
}

// Add template document
func (api *SignplusService) AddTemplateDocument(ctx context.Context, templateId string, addTemplateDocumentRequest AddTemplateDocumentRequest) (*shared.SignplusResponse[Document], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Document](config)

	request := httptransport.NewRequest(ctx, "POST", "/template/{template_id}/document", config)

	request.Body = addTemplateDocumentRequest

	request.SetPathParam("template_id", templateId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[Document](err)
	}

	return shared.NewSignplusResponse[Document](resp), nil
}

// Get template document
func (api *SignplusService) GetTemplateDocument(ctx context.Context, templateId string, documentId string) (*shared.SignplusResponse[Document], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Document](config)

	request := httptransport.NewRequest(ctx, "GET", "/template/{template_id}/document/{document_id}", config)

	request.SetPathParam("template_id", templateId)
	request.SetPathParam("document_id", documentId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[Document](err)
	}

	return shared.NewSignplusResponse[Document](resp), nil
}

// Get template documents
func (api *SignplusService) GetTemplateDocuments(ctx context.Context, templateId string) (*shared.SignplusResponse[ListTemplateDocumentsResponse], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[ListTemplateDocumentsResponse](config)

	request := httptransport.NewRequest(ctx, "GET", "/template/{template_id}/documents", config)

	request.SetPathParam("template_id", templateId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[ListTemplateDocumentsResponse](err)
	}

	return shared.NewSignplusResponse[ListTemplateDocumentsResponse](resp), nil
}

// Add template signing steps
func (api *SignplusService) AddTemplateSigningSteps(ctx context.Context, templateId string, addTemplateSigningStepsRequest AddTemplateSigningStepsRequest) (*shared.SignplusResponse[Template], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Template](config)

	request := httptransport.NewRequest(ctx, "POST", "/template/{template_id}/signing_steps", config)

	request.Body = addTemplateSigningStepsRequest

	request.SetPathParam("template_id", templateId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[Template](err)
	}

	return shared.NewSignplusResponse[Template](resp), nil
}

// Rename template
func (api *SignplusService) RenameTemplate(ctx context.Context, templateId string, renameTemplateRequest RenameTemplateRequest) (*shared.SignplusResponse[Template], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Template](config)

	request := httptransport.NewRequest(ctx, "PUT", "/template/{template_id}/rename", config)

	request.Body = renameTemplateRequest

	request.SetPathParam("template_id", templateId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[Template](err)
	}

	return shared.NewSignplusResponse[Template](resp), nil
}

// Set template comment
func (api *SignplusService) SetTemplateComment(ctx context.Context, templateId string, setTemplateCommentRequest SetTemplateCommentRequest) (*shared.SignplusResponse[Template], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Template](config)

	request := httptransport.NewRequest(ctx, "PUT", "/template/{template_id}/set_comment", config)

	request.Body = setTemplateCommentRequest

	request.SetPathParam("template_id", templateId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[Template](err)
	}

	return shared.NewSignplusResponse[Template](resp), nil
}

// Set template notification
func (api *SignplusService) SetTemplateNotification(ctx context.Context, templateId string, envelopeNotification EnvelopeNotification) (*shared.SignplusResponse[Template], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Template](config)

	request := httptransport.NewRequest(ctx, "PUT", "/template/{template_id}/set_notification", config)

	request.Body = envelopeNotification

	request.SetPathParam("template_id", templateId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[Template](err)
	}

	return shared.NewSignplusResponse[Template](resp), nil
}

// Get template annotations
func (api *SignplusService) GetTemplateAnnotations(ctx context.Context, templateId string) (*shared.SignplusResponse[ListTemplateAnnotationsResponse], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[ListTemplateAnnotationsResponse](config)

	request := httptransport.NewRequest(ctx, "GET", "/template/{template_id}/annotations", config)

	request.SetPathParam("template_id", templateId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[ListTemplateAnnotationsResponse](err)
	}

	return shared.NewSignplusResponse[ListTemplateAnnotationsResponse](resp), nil
}

// Get document template annotations
func (api *SignplusService) GetDocumentTemplateAnnotations(ctx context.Context, templateId string, documentId string) (*shared.SignplusResponse[ListTemplateDocumentAnnotationsResponse], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[ListTemplateDocumentAnnotationsResponse](config)

	request := httptransport.NewRequest(ctx, "GET", "/template/{template_id}/annotations/{document_id}", config)

	request.SetPathParam("template_id", templateId)
	request.SetPathParam("document_id", documentId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[ListTemplateDocumentAnnotationsResponse](err)
	}

	return shared.NewSignplusResponse[ListTemplateDocumentAnnotationsResponse](resp), nil
}

// Add template annotation
func (api *SignplusService) AddTemplateAnnotation(ctx context.Context, templateId string, addAnnotationRequest AddAnnotationRequest) (*shared.SignplusResponse[Annotation], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Annotation](config)

	request := httptransport.NewRequest(ctx, "POST", "/template/{template_id}/annotation", config)

	request.Body = addAnnotationRequest

	request.SetPathParam("template_id", templateId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[Annotation](err)
	}

	return shared.NewSignplusResponse[Annotation](resp), nil
}

// Delete template annotation
func (api *SignplusService) DeleteTemplateAnnotation(ctx context.Context, templateId string, annotationId string) (*shared.SignplusResponse[any], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[any](config)

	request := httptransport.NewRequest(ctx, "DELETE", "/template/{template_id}/annotation/{annotation_id}", config)

	request.SetPathParam("template_id", templateId)
	request.SetPathParam("annotation_id", annotationId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[any](err)
	}

	return shared.NewSignplusResponse[any](resp), nil
}

// Create webhook
func (api *SignplusService) CreateWebhook(ctx context.Context, createWebhookRequest CreateWebhookRequest) (*shared.SignplusResponse[Webhook], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[Webhook](config)

	request := httptransport.NewRequest(ctx, "POST", "/webhook", config)

	request.Body = createWebhookRequest

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[Webhook](err)
	}

	return shared.NewSignplusResponse[Webhook](resp), nil
}

// List webhooks
func (api *SignplusService) ListWebhooks(ctx context.Context, listWebhooksRequest ListWebhooksRequest) (*shared.SignplusResponse[ListWebhooksResponse], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[ListWebhooksResponse](config)

	request := httptransport.NewRequest(ctx, "POST", "/webhooks", config)

	request.Body = listWebhooksRequest

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[ListWebhooksResponse](err)
	}

	return shared.NewSignplusResponse[ListWebhooksResponse](resp), nil
}

// Delete webhook
func (api *SignplusService) DeleteWebhook(ctx context.Context, webhookId string) (*shared.SignplusResponse[any], *shared.SignplusError) {
	config := *api.getConfig()

	client := restClient.NewRestClient[any](config)

	request := httptransport.NewRequest(ctx, "DELETE", "/webhook/{webhook_id}", config)

	request.SetPathParam("webhook_id", webhookId)

	resp, err := client.Call(request)
	if err != nil {
		return nil, shared.NewSignplusError[any](err)
	}

	return shared.NewSignplusResponse[any](resp), nil
}
