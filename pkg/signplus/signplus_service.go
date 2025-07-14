package signplus

import (
	"context"
	restClient "github.com/alohihq/signplus-go/internal/clients/rest"
	"github.com/alohihq/signplus-go/internal/clients/rest/httptransport"
	"github.com/alohihq/signplus-go/internal/configmanager"
	"github.com/alohihq/signplus-go/pkg/shared"
	"github.com/alohihq/signplus-go/pkg/signplusconfig"
	"time"
)

type SignplusService struct {
	manager *configmanager.ConfigManager
}

func NewSignplusService() *SignplusService {
	return &SignplusService{
		manager: configmanager.NewConfigManager(signplusconfig.Config{}),
	}
}

func (api *SignplusService) WithConfigManager(manager *configmanager.ConfigManager) *SignplusService {
	api.manager = manager
	return api
}

func (api *SignplusService) getConfig() *signplusconfig.Config {
	return api.manager.GetSignplus()
}

func (api *SignplusService) SetBaseUrl(baseUrl string) {
	config := api.getConfig()
	config.SetBaseUrl(baseUrl)
}

func (api *SignplusService) SetTimeout(timeout time.Duration) {
	config := api.getConfig()
	config.SetTimeout(timeout)
}

func (api *SignplusService) SetAccessToken(accessToken string) {
	config := api.getConfig()
	config.SetAccessToken(accessToken)
}

// Create new envelope
func (api *SignplusService) CreateEnvelope(ctx context.Context, createEnvelopeRequest CreateEnvelopeRequest) (*shared.SignplusResponse[Envelope], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/envelope").
		WithConfig(config).
		WithBody(createEnvelopeRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Envelope](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[Envelope](err)
	}

	return shared.NewSignplusResponse[Envelope](resp), nil
}

// Create new envelope from template
func (api *SignplusService) CreateEnvelopeFromTemplate(ctx context.Context, templateId string, createEnvelopeFromTemplateRequest CreateEnvelopeFromTemplateRequest) (*shared.SignplusResponse[Envelope], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/envelope/from_template/{template_id}").
		WithConfig(config).
		WithBody(createEnvelopeFromTemplateRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("template_id", templateId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Envelope](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[Envelope](err)
	}

	return shared.NewSignplusResponse[Envelope](resp), nil
}

// List envelopes
func (api *SignplusService) ListEnvelopes(ctx context.Context, listEnvelopesRequest ListEnvelopesRequest) (*shared.SignplusResponse[ListEnvelopesResponse], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/envelopes").
		WithConfig(config).
		WithBody(listEnvelopesRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ListEnvelopesResponse](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[ListEnvelopesResponse](err)
	}

	return shared.NewSignplusResponse[ListEnvelopesResponse](resp), nil
}

// Get envelope
func (api *SignplusService) GetEnvelope(ctx context.Context, envelopeId string) (*shared.SignplusResponse[Envelope], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/envelope/{envelope_id}").
		WithConfig(config).
		AddPathParam("envelope_id", envelopeId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Envelope](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[Envelope](err)
	}

	return shared.NewSignplusResponse[Envelope](resp), nil
}

// Delete envelope
func (api *SignplusService) DeleteEnvelope(ctx context.Context, envelopeId string) (*shared.SignplusResponse[any], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/envelope/{envelope_id}").
		WithConfig(config).
		AddPathParam("envelope_id", envelopeId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[any](err)
	}

	return shared.NewSignplusResponse[any](resp), nil
}

// Download signed documents for an envelope
func (api *SignplusService) DownloadEnvelopeSignedDocuments(ctx context.Context, envelopeId string, params DownloadEnvelopeSignedDocumentsRequestParams) (*shared.SignplusResponse[[]byte], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/envelope/{envelope_id}/signed_documents").
		WithConfig(config).
		AddPathParam("envelope_id", envelopeId).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[[]byte](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[[]byte](err)
	}

	return shared.NewSignplusResponse[[]byte](resp), nil
}

// Download certificate of completion for an envelope
func (api *SignplusService) DownloadEnvelopeCertificate(ctx context.Context, envelopeId string) (*shared.SignplusResponse[[]byte], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/envelope/{envelope_id}/certificate").
		WithConfig(config).
		AddPathParam("envelope_id", envelopeId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[[]byte](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[[]byte](err)
	}

	return shared.NewSignplusResponse[[]byte](resp), nil
}

// Get envelope document
func (api *SignplusService) GetEnvelopeDocument(ctx context.Context, envelopeId string, documentId string) (*shared.SignplusResponse[Document], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/envelope/{envelope_id}/document/{document_id}").
		WithConfig(config).
		AddPathParam("envelope_id", envelopeId).
		AddPathParam("document_id", documentId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Document](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[Document](err)
	}

	return shared.NewSignplusResponse[Document](resp), nil
}

// Get envelope documents
func (api *SignplusService) GetEnvelopeDocuments(ctx context.Context, envelopeId string) (*shared.SignplusResponse[ListEnvelopeDocumentsResponse], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/envelope/{envelope_id}/documents").
		WithConfig(config).
		AddPathParam("envelope_id", envelopeId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ListEnvelopeDocumentsResponse](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[ListEnvelopeDocumentsResponse](err)
	}

	return shared.NewSignplusResponse[ListEnvelopeDocumentsResponse](resp), nil
}

// Add envelope document
func (api *SignplusService) AddEnvelopeDocument(ctx context.Context, envelopeId string, addEnvelopeDocumentRequest AddEnvelopeDocumentRequest) (*shared.SignplusResponse[Document], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/envelope/{envelope_id}/document").
		WithConfig(config).
		WithBody(addEnvelopeDocumentRequest).
		AddPathParam("envelope_id", envelopeId).
		WithContentType(httptransport.ContentTypeMultipartFormData).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Document](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[Document](err)
	}

	return shared.NewSignplusResponse[Document](resp), nil
}

// Set envelope dynamic fields
func (api *SignplusService) SetEnvelopeDynamicFields(ctx context.Context, envelopeId string, setEnvelopeDynamicFieldsRequest SetEnvelopeDynamicFieldsRequest) (*shared.SignplusResponse[Envelope], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/envelope/{envelope_id}/dynamic_fields").
		WithConfig(config).
		WithBody(setEnvelopeDynamicFieldsRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("envelope_id", envelopeId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Envelope](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[Envelope](err)
	}

	return shared.NewSignplusResponse[Envelope](resp), nil
}

// Add envelope signing steps
func (api *SignplusService) AddEnvelopeSigningSteps(ctx context.Context, envelopeId string, addEnvelopeSigningStepsRequest AddEnvelopeSigningStepsRequest) (*shared.SignplusResponse[Envelope], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/envelope/{envelope_id}/signing_steps").
		WithConfig(config).
		WithBody(addEnvelopeSigningStepsRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("envelope_id", envelopeId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Envelope](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[Envelope](err)
	}

	return shared.NewSignplusResponse[Envelope](resp), nil
}

// Send envelope for signature
func (api *SignplusService) SendEnvelope(ctx context.Context, envelopeId string) (*shared.SignplusResponse[Envelope], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/envelope/{envelope_id}/send").
		WithConfig(config).
		AddPathParam("envelope_id", envelopeId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Envelope](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[Envelope](err)
	}

	return shared.NewSignplusResponse[Envelope](resp), nil
}

// Duplicate envelope
func (api *SignplusService) DuplicateEnvelope(ctx context.Context, envelopeId string) (*shared.SignplusResponse[Envelope], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/envelope/{envelope_id}/duplicate").
		WithConfig(config).
		AddPathParam("envelope_id", envelopeId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Envelope](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[Envelope](err)
	}

	return shared.NewSignplusResponse[Envelope](resp), nil
}

// Void envelope
func (api *SignplusService) VoidEnvelope(ctx context.Context, envelopeId string) (*shared.SignplusResponse[Envelope], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/envelope/{envelope_id}/void").
		WithConfig(config).
		AddPathParam("envelope_id", envelopeId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Envelope](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[Envelope](err)
	}

	return shared.NewSignplusResponse[Envelope](resp), nil
}

// Rename envelope
func (api *SignplusService) RenameEnvelope(ctx context.Context, envelopeId string, renameEnvelopeRequest RenameEnvelopeRequest) (*shared.SignplusResponse[Envelope], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/envelope/{envelope_id}/rename").
		WithConfig(config).
		WithBody(renameEnvelopeRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("envelope_id", envelopeId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Envelope](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[Envelope](err)
	}

	return shared.NewSignplusResponse[Envelope](resp), nil
}

// Set envelope comment
func (api *SignplusService) SetEnvelopeComment(ctx context.Context, envelopeId string, setEnvelopeCommentRequest SetEnvelopeCommentRequest) (*shared.SignplusResponse[Envelope], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/envelope/{envelope_id}/set_comment").
		WithConfig(config).
		WithBody(setEnvelopeCommentRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("envelope_id", envelopeId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Envelope](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[Envelope](err)
	}

	return shared.NewSignplusResponse[Envelope](resp), nil
}

// Set envelope notification
func (api *SignplusService) SetEnvelopeNotification(ctx context.Context, envelopeId string, envelopeNotification EnvelopeNotification) (*shared.SignplusResponse[Envelope], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/envelope/{envelope_id}/set_notification").
		WithConfig(config).
		WithBody(envelopeNotification).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("envelope_id", envelopeId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Envelope](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[Envelope](err)
	}

	return shared.NewSignplusResponse[Envelope](resp), nil
}

// Set envelope expiration date
func (api *SignplusService) SetEnvelopeExpirationDate(ctx context.Context, envelopeId string, setEnvelopeExpirationRequest SetEnvelopeExpirationRequest) (*shared.SignplusResponse[Envelope], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/envelope/{envelope_id}/set_expiration_date").
		WithConfig(config).
		WithBody(setEnvelopeExpirationRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("envelope_id", envelopeId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Envelope](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[Envelope](err)
	}

	return shared.NewSignplusResponse[Envelope](resp), nil
}

// Set envelope legality level
func (api *SignplusService) SetEnvelopeLegalityLevel(ctx context.Context, envelopeId string, setEnvelopeLegalityLevelRequest SetEnvelopeLegalityLevelRequest) (*shared.SignplusResponse[Envelope], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/envelope/{envelope_id}/set_legality_level").
		WithConfig(config).
		WithBody(setEnvelopeLegalityLevelRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("envelope_id", envelopeId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Envelope](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[Envelope](err)
	}

	return shared.NewSignplusResponse[Envelope](resp), nil
}

// Get envelope annotations
func (api *SignplusService) GetEnvelopeAnnotations(ctx context.Context, envelopeId string) (*shared.SignplusResponse[[]Annotation], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/envelope/{envelope_id}/annotations").
		WithConfig(config).
		AddPathParam("envelope_id", envelopeId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[[]Annotation](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[[]Annotation](err)
	}

	return shared.NewSignplusResponse[[]Annotation](resp), nil
}

// Get envelope document annotations
func (api *SignplusService) GetEnvelopeDocumentAnnotations(ctx context.Context, envelopeId string, documentId string) (*shared.SignplusResponse[ListEnvelopeDocumentAnnotationsResponse], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/envelope/{envelope_id}/annotations/{document_id}").
		WithConfig(config).
		AddPathParam("envelope_id", envelopeId).
		AddPathParam("document_id", documentId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ListEnvelopeDocumentAnnotationsResponse](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[ListEnvelopeDocumentAnnotationsResponse](err)
	}

	return shared.NewSignplusResponse[ListEnvelopeDocumentAnnotationsResponse](resp), nil
}

// Add envelope annotation
func (api *SignplusService) AddEnvelopeAnnotation(ctx context.Context, envelopeId string, addAnnotationRequest AddAnnotationRequest) (*shared.SignplusResponse[Annotation], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/envelope/{envelope_id}/annotation").
		WithConfig(config).
		WithBody(addAnnotationRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("envelope_id", envelopeId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Annotation](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[Annotation](err)
	}

	return shared.NewSignplusResponse[Annotation](resp), nil
}

// Delete envelope annotation
func (api *SignplusService) DeleteEnvelopeAnnotation(ctx context.Context, envelopeId string, annotationId string) (*shared.SignplusResponse[any], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/envelope/{envelope_id}/annotation/{annotation_id}").
		WithConfig(config).
		AddPathParam("envelope_id", envelopeId).
		AddPathParam("annotation_id", annotationId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[any](err)
	}

	return shared.NewSignplusResponse[any](resp), nil
}

// Create new template
func (api *SignplusService) CreateTemplate(ctx context.Context, createTemplateRequest CreateTemplateRequest) (*shared.SignplusResponse[Template], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/template").
		WithConfig(config).
		WithBody(createTemplateRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Template](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[Template](err)
	}

	return shared.NewSignplusResponse[Template](resp), nil
}

// List templates
func (api *SignplusService) ListTemplates(ctx context.Context, listTemplatesRequest ListTemplatesRequest) (*shared.SignplusResponse[ListTemplatesResponse], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/templates").
		WithConfig(config).
		WithBody(listTemplatesRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ListTemplatesResponse](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[ListTemplatesResponse](err)
	}

	return shared.NewSignplusResponse[ListTemplatesResponse](resp), nil
}

// Get template
func (api *SignplusService) GetTemplate(ctx context.Context, templateId string) (*shared.SignplusResponse[Template], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/template/{template_id}").
		WithConfig(config).
		AddPathParam("template_id", templateId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Template](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[Template](err)
	}

	return shared.NewSignplusResponse[Template](resp), nil
}

// Delete template
func (api *SignplusService) DeleteTemplate(ctx context.Context, templateId string) (*shared.SignplusResponse[any], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/template/{template_id}").
		WithConfig(config).
		AddPathParam("template_id", templateId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[any](err)
	}

	return shared.NewSignplusResponse[any](resp), nil
}

// Duplicate template
func (api *SignplusService) DuplicateTemplate(ctx context.Context, templateId string) (*shared.SignplusResponse[Template], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/template/{template_id}/duplicate").
		WithConfig(config).
		AddPathParam("template_id", templateId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Template](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[Template](err)
	}

	return shared.NewSignplusResponse[Template](resp), nil
}

// Add template document
func (api *SignplusService) AddTemplateDocument(ctx context.Context, templateId string, addTemplateDocumentRequest AddTemplateDocumentRequest) (*shared.SignplusResponse[Document], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/template/{template_id}/document").
		WithConfig(config).
		WithBody(addTemplateDocumentRequest).
		AddPathParam("template_id", templateId).
		WithContentType(httptransport.ContentTypeMultipartFormData).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Document](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[Document](err)
	}

	return shared.NewSignplusResponse[Document](resp), nil
}

// Get template document
func (api *SignplusService) GetTemplateDocument(ctx context.Context, templateId string, documentId string) (*shared.SignplusResponse[Document], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/template/{template_id}/document/{document_id}").
		WithConfig(config).
		AddPathParam("template_id", templateId).
		AddPathParam("document_id", documentId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Document](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[Document](err)
	}

	return shared.NewSignplusResponse[Document](resp), nil
}

// Get template documents
func (api *SignplusService) GetTemplateDocuments(ctx context.Context, templateId string) (*shared.SignplusResponse[ListTemplateDocumentsResponse], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/template/{template_id}/documents").
		WithConfig(config).
		AddPathParam("template_id", templateId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ListTemplateDocumentsResponse](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[ListTemplateDocumentsResponse](err)
	}

	return shared.NewSignplusResponse[ListTemplateDocumentsResponse](resp), nil
}

// Add template signing steps
func (api *SignplusService) AddTemplateSigningSteps(ctx context.Context, templateId string, addTemplateSigningStepsRequest AddTemplateSigningStepsRequest) (*shared.SignplusResponse[Template], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/template/{template_id}/signing_steps").
		WithConfig(config).
		WithBody(addTemplateSigningStepsRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("template_id", templateId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Template](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[Template](err)
	}

	return shared.NewSignplusResponse[Template](resp), nil
}

// Rename template
func (api *SignplusService) RenameTemplate(ctx context.Context, templateId string, renameTemplateRequest RenameTemplateRequest) (*shared.SignplusResponse[Template], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/template/{template_id}/rename").
		WithConfig(config).
		WithBody(renameTemplateRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("template_id", templateId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Template](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[Template](err)
	}

	return shared.NewSignplusResponse[Template](resp), nil
}

// Set template comment
func (api *SignplusService) SetTemplateComment(ctx context.Context, templateId string, setTemplateCommentRequest SetTemplateCommentRequest) (*shared.SignplusResponse[Template], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/template/{template_id}/set_comment").
		WithConfig(config).
		WithBody(setTemplateCommentRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("template_id", templateId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Template](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[Template](err)
	}

	return shared.NewSignplusResponse[Template](resp), nil
}

// Set template notification
func (api *SignplusService) SetTemplateNotification(ctx context.Context, templateId string, envelopeNotification EnvelopeNotification) (*shared.SignplusResponse[Template], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/template/{template_id}/set_notification").
		WithConfig(config).
		WithBody(envelopeNotification).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("template_id", templateId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Template](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[Template](err)
	}

	return shared.NewSignplusResponse[Template](resp), nil
}

// Get template annotations
func (api *SignplusService) GetTemplateAnnotations(ctx context.Context, templateId string) (*shared.SignplusResponse[ListTemplateAnnotationsResponse], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/template/{template_id}/annotations").
		WithConfig(config).
		AddPathParam("template_id", templateId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ListTemplateAnnotationsResponse](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[ListTemplateAnnotationsResponse](err)
	}

	return shared.NewSignplusResponse[ListTemplateAnnotationsResponse](resp), nil
}

// Get document template annotations
func (api *SignplusService) GetDocumentTemplateAnnotations(ctx context.Context, templateId string, documentId string) (*shared.SignplusResponse[ListTemplateDocumentAnnotationsResponse], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/template/{template_id}/annotations/{document_id}").
		WithConfig(config).
		AddPathParam("template_id", templateId).
		AddPathParam("document_id", documentId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ListTemplateDocumentAnnotationsResponse](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[ListTemplateDocumentAnnotationsResponse](err)
	}

	return shared.NewSignplusResponse[ListTemplateDocumentAnnotationsResponse](resp), nil
}

// Add template annotation
func (api *SignplusService) AddTemplateAnnotation(ctx context.Context, templateId string, addAnnotationRequest AddAnnotationRequest) (*shared.SignplusResponse[Annotation], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/template/{template_id}/annotation").
		WithConfig(config).
		WithBody(addAnnotationRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("template_id", templateId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Annotation](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[Annotation](err)
	}

	return shared.NewSignplusResponse[Annotation](resp), nil
}

// Delete template annotation
func (api *SignplusService) DeleteTemplateAnnotation(ctx context.Context, templateId string, annotationId string) (*shared.SignplusResponse[any], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/template/{template_id}/annotation/{annotation_id}").
		WithConfig(config).
		AddPathParam("template_id", templateId).
		AddPathParam("annotation_id", annotationId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[any](err)
	}

	return shared.NewSignplusResponse[any](resp), nil
}

// Create webhook
func (api *SignplusService) CreateWebhook(ctx context.Context, createWebhookRequest CreateWebhookRequest) (*shared.SignplusResponse[Webhook], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/webhook").
		WithConfig(config).
		WithBody(createWebhookRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[Webhook](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[Webhook](err)
	}

	return shared.NewSignplusResponse[Webhook](resp), nil
}

// List webhooks
func (api *SignplusService) ListWebhooks(ctx context.Context, listWebhooksRequest ListWebhooksRequest) (*shared.SignplusResponse[ListWebhooksResponse], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/webhooks").
		WithConfig(config).
		WithBody(listWebhooksRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[ListWebhooksResponse](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[ListWebhooksResponse](err)
	}

	return shared.NewSignplusResponse[ListWebhooksResponse](resp), nil
}

// Delete webhook
func (api *SignplusService) DeleteWebhook(ctx context.Context, webhookId string) (*shared.SignplusResponse[any], *shared.SignplusError) {
	config := *api.getConfig()

	request := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/webhook/{webhook_id}").
		WithConfig(config).
		AddPathParam("webhook_id", webhookId).
		WithContentType(httptransport.ContentTypeJson).
		WithResponseContentType(httptransport.ContentTypeJson).
		Build()

	client := restClient.NewRestClient[any](config)
	resp, err := client.Call(*request)
	if err != nil {
		return nil, shared.NewSignplusError[any](err)
	}

	return shared.NewSignplusResponse[any](resp), nil
}
