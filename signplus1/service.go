package signplus1

import (
	"context"
	restClient "github.com/alohihq/signplus-go/internal/clients/rest"
	"github.com/alohihq/signplus-go/internal/clients/rest/hooks"
	"github.com/alohihq/signplus-go/internal/clients/rest/httptransport"
	"github.com/alohihq/signplus-go/internal/configmanager"
	"github.com/alohihq/signplus-go/sharedmodels"
	"github.com/alohihq/signplus-go/signplusconfig"
	"time"
)

// Service provides methods to interact with Signplus1-related API endpoints.
// It uses a configuration manager for settings and supports custom hooks for request/response interception.
type Service struct {
	manager                                  *configmanager.ConfigManager
	hook                                     hooks.Hook
	createEnvelopeConfig                     []signplusconfig.RequestOption
	createEnvelopeFromTemplateConfig         []signplusconfig.RequestOption
	listEnvelopesConfig                      []signplusconfig.RequestOption
	getEnvelopeConfig                        []signplusconfig.RequestOption
	deleteEnvelopeConfig                     []signplusconfig.RequestOption
	downloadEnvelopeSignedDocumentsConfig    []signplusconfig.RequestOption
	downloadEnvelopeCertificateConfig        []signplusconfig.RequestOption
	getEnvelopeDocumentConfig                []signplusconfig.RequestOption
	getEnvelopeDocumentsConfig               []signplusconfig.RequestOption
	addEnvelopeDocumentConfig                []signplusconfig.RequestOption
	setEnvelopeDynamicFieldsConfig           []signplusconfig.RequestOption
	addEnvelopeSigningStepsConfig            []signplusconfig.RequestOption
	setEnvelopeAttachmentsSettingsConfig     []signplusconfig.RequestOption
	setEnvelopeAttachmentsPlaceholdersConfig []signplusconfig.RequestOption
	getAttachmentFileConfig                  []signplusconfig.RequestOption
	sendEnvelopeConfig                       []signplusconfig.RequestOption
	duplicateEnvelopeConfig                  []signplusconfig.RequestOption
	voidEnvelopeConfig                       []signplusconfig.RequestOption
	renameEnvelopeConfig                     []signplusconfig.RequestOption
	setEnvelopeCommentConfig                 []signplusconfig.RequestOption
	setEnvelopeNotificationConfig            []signplusconfig.RequestOption
	setEnvelopeExpirationDateConfig          []signplusconfig.RequestOption
	setEnvelopeLegalityLevelConfig           []signplusconfig.RequestOption
	getEnvelopeAnnotationsConfig             []signplusconfig.RequestOption
	getEnvelopeDocumentAnnotationsConfig     []signplusconfig.RequestOption
	addEnvelopeAnnotationConfig              []signplusconfig.RequestOption
	deleteEnvelopeAnnotationConfig           []signplusconfig.RequestOption
	createTemplateConfig                     []signplusconfig.RequestOption
	listTemplatesConfig                      []signplusconfig.RequestOption
	getTemplateConfig                        []signplusconfig.RequestOption
	deleteTemplateConfig                     []signplusconfig.RequestOption
	duplicateTemplateConfig                  []signplusconfig.RequestOption
	addTemplateDocumentConfig                []signplusconfig.RequestOption
	getTemplateDocumentConfig                []signplusconfig.RequestOption
	getTemplateDocumentsConfig               []signplusconfig.RequestOption
	addTemplateSigningStepsConfig            []signplusconfig.RequestOption
	renameTemplateConfig                     []signplusconfig.RequestOption
	setTemplateCommentConfig                 []signplusconfig.RequestOption
	setTemplateNotificationConfig            []signplusconfig.RequestOption
	getTemplateAnnotationsConfig             []signplusconfig.RequestOption
	getDocumentTemplateAnnotationsConfig     []signplusconfig.RequestOption
	addTemplateAnnotationConfig              []signplusconfig.RequestOption
	deleteTemplateAnnotationConfig           []signplusconfig.RequestOption
	setTemplateAttachmentsSettingsConfig     []signplusconfig.RequestOption
	setTemplateAttachmentsPlaceholdersConfig []signplusconfig.RequestOption
	createWebhookConfig                      []signplusconfig.RequestOption
	listWebhooksConfig                       []signplusconfig.RequestOption
	deleteWebhookConfig                      []signplusconfig.RequestOption
}

func NewService() *Service {
	return &Service{
		manager: configmanager.NewConfigManager(signplusconfig.Config{}),
	}
}

// WithConfigManager sets the configuration manager for this service.
// Returns the service instance for method chaining.
func (api *Service) WithConfigManager(manager *configmanager.ConfigManager) *Service {
	api.manager = manager
	return api
}

// WithHook sets a custom hook for request/response interception.
// Returns the service instance for method chaining.
func (api *Service) WithHook(hook hooks.Hook) *Service {
	api.hook = hook
	return api
}

func (api *Service) config() *signplusconfig.Config {
	return api.manager.GetSignplus1()
}

func (api *Service) getHook() hooks.Hook {
	return api.hook
}

func (api *Service) SetBaseURL(baseURL string) {
	config := api.config()
	config.SetBaseURL(baseURL)
}

func (api *Service) SetTimeout(timeout time.Duration) {
	config := api.config()
	config.SetTimeout(timeout)
}

func (api *Service) SetAccessToken(accessToken string) {
	config := api.config()
	config.SetAccessToken(accessToken)
}

// SetCreateEnvelopeConfig sets method-level configuration for CreateEnvelope.
// Options are applied to every future call to CreateEnvelope and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetCreateEnvelopeConfig(opts ...signplusconfig.RequestOption) *Service {
	api.createEnvelopeConfig = opts
	return api
}

// SetCreateEnvelopeFromTemplateConfig sets method-level configuration for CreateEnvelopeFromTemplate.
// Options are applied to every future call to CreateEnvelopeFromTemplate and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetCreateEnvelopeFromTemplateConfig(opts ...signplusconfig.RequestOption) *Service {
	api.createEnvelopeFromTemplateConfig = opts
	return api
}

// SetListEnvelopesConfig sets method-level configuration for ListEnvelopes.
// Options are applied to every future call to ListEnvelopes and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetListEnvelopesConfig(opts ...signplusconfig.RequestOption) *Service {
	api.listEnvelopesConfig = opts
	return api
}

// SetGetEnvelopeConfig sets method-level configuration for GetEnvelope.
// Options are applied to every future call to GetEnvelope and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetEnvelopeConfig(opts ...signplusconfig.RequestOption) *Service {
	api.getEnvelopeConfig = opts
	return api
}

// SetDeleteEnvelopeConfig sets method-level configuration for DeleteEnvelope.
// Options are applied to every future call to DeleteEnvelope and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetDeleteEnvelopeConfig(opts ...signplusconfig.RequestOption) *Service {
	api.deleteEnvelopeConfig = opts
	return api
}

// SetDownloadEnvelopeSignedDocumentsConfig sets method-level configuration for DownloadEnvelopeSignedDocuments.
// Options are applied to every future call to DownloadEnvelopeSignedDocuments and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetDownloadEnvelopeSignedDocumentsConfig(opts ...signplusconfig.RequestOption) *Service {
	api.downloadEnvelopeSignedDocumentsConfig = opts
	return api
}

// SetDownloadEnvelopeCertificateConfig sets method-level configuration for DownloadEnvelopeCertificate.
// Options are applied to every future call to DownloadEnvelopeCertificate and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetDownloadEnvelopeCertificateConfig(opts ...signplusconfig.RequestOption) *Service {
	api.downloadEnvelopeCertificateConfig = opts
	return api
}

// SetGetEnvelopeDocumentConfig sets method-level configuration for GetEnvelopeDocument.
// Options are applied to every future call to GetEnvelopeDocument and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetEnvelopeDocumentConfig(opts ...signplusconfig.RequestOption) *Service {
	api.getEnvelopeDocumentConfig = opts
	return api
}

// SetGetEnvelopeDocumentsConfig sets method-level configuration for GetEnvelopeDocuments.
// Options are applied to every future call to GetEnvelopeDocuments and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetEnvelopeDocumentsConfig(opts ...signplusconfig.RequestOption) *Service {
	api.getEnvelopeDocumentsConfig = opts
	return api
}

// SetAddEnvelopeDocumentConfig sets method-level configuration for AddEnvelopeDocument.
// Options are applied to every future call to AddEnvelopeDocument and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetAddEnvelopeDocumentConfig(opts ...signplusconfig.RequestOption) *Service {
	api.addEnvelopeDocumentConfig = opts
	return api
}

// SetSetEnvelopeDynamicFieldsConfig sets method-level configuration for SetEnvelopeDynamicFields.
// Options are applied to every future call to SetEnvelopeDynamicFields and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetSetEnvelopeDynamicFieldsConfig(opts ...signplusconfig.RequestOption) *Service {
	api.setEnvelopeDynamicFieldsConfig = opts
	return api
}

// SetAddEnvelopeSigningStepsConfig sets method-level configuration for AddEnvelopeSigningSteps.
// Options are applied to every future call to AddEnvelopeSigningSteps and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetAddEnvelopeSigningStepsConfig(opts ...signplusconfig.RequestOption) *Service {
	api.addEnvelopeSigningStepsConfig = opts
	return api
}

// SetSetEnvelopeAttachmentsSettingsConfig sets method-level configuration for SetEnvelopeAttachmentsSettings.
// Options are applied to every future call to SetEnvelopeAttachmentsSettings and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetSetEnvelopeAttachmentsSettingsConfig(opts ...signplusconfig.RequestOption) *Service {
	api.setEnvelopeAttachmentsSettingsConfig = opts
	return api
}

// SetSetEnvelopeAttachmentsPlaceholdersConfig sets method-level configuration for SetEnvelopeAttachmentsPlaceholders.
// Options are applied to every future call to SetEnvelopeAttachmentsPlaceholders and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetSetEnvelopeAttachmentsPlaceholdersConfig(opts ...signplusconfig.RequestOption) *Service {
	api.setEnvelopeAttachmentsPlaceholdersConfig = opts
	return api
}

// SetGetAttachmentFileConfig sets method-level configuration for GetAttachmentFile.
// Options are applied to every future call to GetAttachmentFile and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetAttachmentFileConfig(opts ...signplusconfig.RequestOption) *Service {
	api.getAttachmentFileConfig = opts
	return api
}

// SetSendEnvelopeConfig sets method-level configuration for SendEnvelope.
// Options are applied to every future call to SendEnvelope and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetSendEnvelopeConfig(opts ...signplusconfig.RequestOption) *Service {
	api.sendEnvelopeConfig = opts
	return api
}

// SetDuplicateEnvelopeConfig sets method-level configuration for DuplicateEnvelope.
// Options are applied to every future call to DuplicateEnvelope and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetDuplicateEnvelopeConfig(opts ...signplusconfig.RequestOption) *Service {
	api.duplicateEnvelopeConfig = opts
	return api
}

// SetVoidEnvelopeConfig sets method-level configuration for VoidEnvelope.
// Options are applied to every future call to VoidEnvelope and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetVoidEnvelopeConfig(opts ...signplusconfig.RequestOption) *Service {
	api.voidEnvelopeConfig = opts
	return api
}

// SetRenameEnvelopeConfig sets method-level configuration for RenameEnvelope.
// Options are applied to every future call to RenameEnvelope and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetRenameEnvelopeConfig(opts ...signplusconfig.RequestOption) *Service {
	api.renameEnvelopeConfig = opts
	return api
}

// SetSetEnvelopeCommentConfig sets method-level configuration for SetEnvelopeComment.
// Options are applied to every future call to SetEnvelopeComment and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetSetEnvelopeCommentConfig(opts ...signplusconfig.RequestOption) *Service {
	api.setEnvelopeCommentConfig = opts
	return api
}

// SetSetEnvelopeNotificationConfig sets method-level configuration for SetEnvelopeNotification.
// Options are applied to every future call to SetEnvelopeNotification and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetSetEnvelopeNotificationConfig(opts ...signplusconfig.RequestOption) *Service {
	api.setEnvelopeNotificationConfig = opts
	return api
}

// SetSetEnvelopeExpirationDateConfig sets method-level configuration for SetEnvelopeExpirationDate.
// Options are applied to every future call to SetEnvelopeExpirationDate and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetSetEnvelopeExpirationDateConfig(opts ...signplusconfig.RequestOption) *Service {
	api.setEnvelopeExpirationDateConfig = opts
	return api
}

// SetSetEnvelopeLegalityLevelConfig sets method-level configuration for SetEnvelopeLegalityLevel.
// Options are applied to every future call to SetEnvelopeLegalityLevel and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetSetEnvelopeLegalityLevelConfig(opts ...signplusconfig.RequestOption) *Service {
	api.setEnvelopeLegalityLevelConfig = opts
	return api
}

// SetGetEnvelopeAnnotationsConfig sets method-level configuration for GetEnvelopeAnnotations.
// Options are applied to every future call to GetEnvelopeAnnotations and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetEnvelopeAnnotationsConfig(opts ...signplusconfig.RequestOption) *Service {
	api.getEnvelopeAnnotationsConfig = opts
	return api
}

// SetGetEnvelopeDocumentAnnotationsConfig sets method-level configuration for GetEnvelopeDocumentAnnotations.
// Options are applied to every future call to GetEnvelopeDocumentAnnotations and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetEnvelopeDocumentAnnotationsConfig(opts ...signplusconfig.RequestOption) *Service {
	api.getEnvelopeDocumentAnnotationsConfig = opts
	return api
}

// SetAddEnvelopeAnnotationConfig sets method-level configuration for AddEnvelopeAnnotation.
// Options are applied to every future call to AddEnvelopeAnnotation and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetAddEnvelopeAnnotationConfig(opts ...signplusconfig.RequestOption) *Service {
	api.addEnvelopeAnnotationConfig = opts
	return api
}

// SetDeleteEnvelopeAnnotationConfig sets method-level configuration for DeleteEnvelopeAnnotation.
// Options are applied to every future call to DeleteEnvelopeAnnotation and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetDeleteEnvelopeAnnotationConfig(opts ...signplusconfig.RequestOption) *Service {
	api.deleteEnvelopeAnnotationConfig = opts
	return api
}

// SetCreateTemplateConfig sets method-level configuration for CreateTemplate.
// Options are applied to every future call to CreateTemplate and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetCreateTemplateConfig(opts ...signplusconfig.RequestOption) *Service {
	api.createTemplateConfig = opts
	return api
}

// SetListTemplatesConfig sets method-level configuration for ListTemplates.
// Options are applied to every future call to ListTemplates and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetListTemplatesConfig(opts ...signplusconfig.RequestOption) *Service {
	api.listTemplatesConfig = opts
	return api
}

// SetGetTemplateConfig sets method-level configuration for GetTemplate.
// Options are applied to every future call to GetTemplate and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetTemplateConfig(opts ...signplusconfig.RequestOption) *Service {
	api.getTemplateConfig = opts
	return api
}

// SetDeleteTemplateConfig sets method-level configuration for DeleteTemplate.
// Options are applied to every future call to DeleteTemplate and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetDeleteTemplateConfig(opts ...signplusconfig.RequestOption) *Service {
	api.deleteTemplateConfig = opts
	return api
}

// SetDuplicateTemplateConfig sets method-level configuration for DuplicateTemplate.
// Options are applied to every future call to DuplicateTemplate and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetDuplicateTemplateConfig(opts ...signplusconfig.RequestOption) *Service {
	api.duplicateTemplateConfig = opts
	return api
}

// SetAddTemplateDocumentConfig sets method-level configuration for AddTemplateDocument.
// Options are applied to every future call to AddTemplateDocument and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetAddTemplateDocumentConfig(opts ...signplusconfig.RequestOption) *Service {
	api.addTemplateDocumentConfig = opts
	return api
}

// SetGetTemplateDocumentConfig sets method-level configuration for GetTemplateDocument.
// Options are applied to every future call to GetTemplateDocument and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetTemplateDocumentConfig(opts ...signplusconfig.RequestOption) *Service {
	api.getTemplateDocumentConfig = opts
	return api
}

// SetGetTemplateDocumentsConfig sets method-level configuration for GetTemplateDocuments.
// Options are applied to every future call to GetTemplateDocuments and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetTemplateDocumentsConfig(opts ...signplusconfig.RequestOption) *Service {
	api.getTemplateDocumentsConfig = opts
	return api
}

// SetAddTemplateSigningStepsConfig sets method-level configuration for AddTemplateSigningSteps.
// Options are applied to every future call to AddTemplateSigningSteps and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetAddTemplateSigningStepsConfig(opts ...signplusconfig.RequestOption) *Service {
	api.addTemplateSigningStepsConfig = opts
	return api
}

// SetRenameTemplateConfig sets method-level configuration for RenameTemplate.
// Options are applied to every future call to RenameTemplate and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetRenameTemplateConfig(opts ...signplusconfig.RequestOption) *Service {
	api.renameTemplateConfig = opts
	return api
}

// SetSetTemplateCommentConfig sets method-level configuration for SetTemplateComment.
// Options are applied to every future call to SetTemplateComment and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetSetTemplateCommentConfig(opts ...signplusconfig.RequestOption) *Service {
	api.setTemplateCommentConfig = opts
	return api
}

// SetSetTemplateNotificationConfig sets method-level configuration for SetTemplateNotification.
// Options are applied to every future call to SetTemplateNotification and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetSetTemplateNotificationConfig(opts ...signplusconfig.RequestOption) *Service {
	api.setTemplateNotificationConfig = opts
	return api
}

// SetGetTemplateAnnotationsConfig sets method-level configuration for GetTemplateAnnotations.
// Options are applied to every future call to GetTemplateAnnotations and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetTemplateAnnotationsConfig(opts ...signplusconfig.RequestOption) *Service {
	api.getTemplateAnnotationsConfig = opts
	return api
}

// SetGetDocumentTemplateAnnotationsConfig sets method-level configuration for GetDocumentTemplateAnnotations.
// Options are applied to every future call to GetDocumentTemplateAnnotations and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetGetDocumentTemplateAnnotationsConfig(opts ...signplusconfig.RequestOption) *Service {
	api.getDocumentTemplateAnnotationsConfig = opts
	return api
}

// SetAddTemplateAnnotationConfig sets method-level configuration for AddTemplateAnnotation.
// Options are applied to every future call to AddTemplateAnnotation and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetAddTemplateAnnotationConfig(opts ...signplusconfig.RequestOption) *Service {
	api.addTemplateAnnotationConfig = opts
	return api
}

// SetDeleteTemplateAnnotationConfig sets method-level configuration for DeleteTemplateAnnotation.
// Options are applied to every future call to DeleteTemplateAnnotation and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetDeleteTemplateAnnotationConfig(opts ...signplusconfig.RequestOption) *Service {
	api.deleteTemplateAnnotationConfig = opts
	return api
}

// SetSetTemplateAttachmentsSettingsConfig sets method-level configuration for SetTemplateAttachmentsSettings.
// Options are applied to every future call to SetTemplateAttachmentsSettings and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetSetTemplateAttachmentsSettingsConfig(opts ...signplusconfig.RequestOption) *Service {
	api.setTemplateAttachmentsSettingsConfig = opts
	return api
}

// SetSetTemplateAttachmentsPlaceholdersConfig sets method-level configuration for SetTemplateAttachmentsPlaceholders.
// Options are applied to every future call to SetTemplateAttachmentsPlaceholders and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetSetTemplateAttachmentsPlaceholdersConfig(opts ...signplusconfig.RequestOption) *Service {
	api.setTemplateAttachmentsPlaceholdersConfig = opts
	return api
}

// SetCreateWebhookConfig sets method-level configuration for CreateWebhook.
// Options are applied to every future call to CreateWebhook and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetCreateWebhookConfig(opts ...signplusconfig.RequestOption) *Service {
	api.createWebhookConfig = opts
	return api
}

// SetListWebhooksConfig sets method-level configuration for ListWebhooks.
// Options are applied to every future call to ListWebhooks and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetListWebhooksConfig(opts ...signplusconfig.RequestOption) *Service {
	api.listWebhooksConfig = opts
	return api
}

// SetDeleteWebhookConfig sets method-level configuration for DeleteWebhook.
// Options are applied to every future call to DeleteWebhook and take
// precedence over service-level config. Per-call options still take highest precedence.
func (api *Service) SetDeleteWebhookConfig(opts ...signplusconfig.RequestOption) *Service {
	api.deleteWebhookConfig = opts
	return api
}

// Create new envelope
func (api *Service) CreateEnvelope(ctx context.Context, createEnvelopeRequest CreateEnvelopeRequest, opts ...signplusconfig.RequestOption) (*Envelope, error) {
	config := *api.config()
	for _, opt := range api.createEnvelopeConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/envelope").
		WithConfig(config).
		WithBody(createEnvelopeRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[Envelope, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Create new envelope from template
func (api *Service) CreateEnvelopeFromTemplate(ctx context.Context, templateID string, createEnvelopeFromTemplateRequest CreateEnvelopeFromTemplateRequest, opts ...signplusconfig.RequestOption) (*Envelope, error) {
	config := *api.config()
	for _, opt := range api.createEnvelopeFromTemplateConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/envelope/from_template/{template_id}").
		WithConfig(config).
		WithBody(createEnvelopeFromTemplateRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("template_id", templateID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[Envelope, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// List envelopes
func (api *Service) ListEnvelopes(ctx context.Context, listEnvelopesRequest ListEnvelopesRequest, opts ...signplusconfig.RequestOption) (*ListEnvelopesResponse, error) {
	config := *api.config()
	for _, opt := range api.listEnvelopesConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/envelopes").
		WithConfig(config).
		WithBody(listEnvelopesRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[ListEnvelopesResponse, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Get envelope
func (api *Service) GetEnvelope(ctx context.Context, envelopeID string, opts ...signplusconfig.RequestOption) (*Envelope, error) {
	config := *api.config()
	for _, opt := range api.getEnvelopeConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/envelope/{envelope_id}").
		WithConfig(config).
		AddPathParam("envelope_id", envelopeID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[Envelope, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Delete envelope
func (api *Service) DeleteEnvelope(ctx context.Context, envelopeID string, opts ...signplusconfig.RequestOption) (any, error) {
	config := *api.config()
	for _, opt := range api.deleteEnvelopeConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/envelope/{envelope_id}").
		WithConfig(config).
		AddPathParam("envelope_id", envelopeID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[any, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return resp.Data, nil
}

// Download signed documents for an envelope
func (api *Service) DownloadEnvelopeSignedDocuments(ctx context.Context, envelopeID string, params DownloadEnvelopeSignedDocumentsRequestParams, opts ...signplusconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.downloadEnvelopeSignedDocumentsConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/envelope/{envelope_id}/signed_documents").
		WithConfig(config).
		AddPathParam("envelope_id", envelopeID).
		WithOptions(params).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeBinary).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return resp.Data, nil
}

// Download certificate of completion for an envelope
func (api *Service) DownloadEnvelopeCertificate(ctx context.Context, envelopeID string, opts ...signplusconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.downloadEnvelopeCertificateConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/envelope/{envelope_id}/certificate").
		WithConfig(config).
		AddPathParam("envelope_id", envelopeID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeBinary).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return resp.Data, nil
}

// Get envelope document
func (api *Service) GetEnvelopeDocument(ctx context.Context, envelopeID string, documentID string, opts ...signplusconfig.RequestOption) (*Document, error) {
	config := *api.config()
	for _, opt := range api.getEnvelopeDocumentConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/envelope/{envelope_id}/document/{document_id}").
		WithConfig(config).
		AddPathParam("envelope_id", envelopeID).
		AddPathParam("document_id", documentID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[Document, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Get envelope documents
func (api *Service) GetEnvelopeDocuments(ctx context.Context, envelopeID string, opts ...signplusconfig.RequestOption) (*ListEnvelopeDocumentsResponse, error) {
	config := *api.config()
	for _, opt := range api.getEnvelopeDocumentsConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/envelope/{envelope_id}/documents").
		WithConfig(config).
		AddPathParam("envelope_id", envelopeID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[ListEnvelopeDocumentsResponse, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Add envelope document
func (api *Service) AddEnvelopeDocument(ctx context.Context, envelopeID string, addEnvelopeDocumentRequest AddEnvelopeDocumentRequest, opts ...signplusconfig.RequestOption) (*Document, error) {
	config := *api.config()
	for _, opt := range api.addEnvelopeDocumentConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/envelope/{envelope_id}/document").
		WithConfig(config).
		WithBody(addEnvelopeDocumentRequest).
		AddPathParam("envelope_id", envelopeID).
		WithContentType(httptransport.ContentTypeMultipartFormData).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[Document, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Set envelope dynamic fields
func (api *Service) SetEnvelopeDynamicFields(ctx context.Context, envelopeID string, setEnvelopeDynamicFieldsRequest SetEnvelopeDynamicFieldsRequest, opts ...signplusconfig.RequestOption) (*Envelope, error) {
	config := *api.config()
	for _, opt := range api.setEnvelopeDynamicFieldsConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/envelope/{envelope_id}/dynamic_fields").
		WithConfig(config).
		WithBody(setEnvelopeDynamicFieldsRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("envelope_id", envelopeID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[Envelope, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Add envelope signing steps
func (api *Service) AddEnvelopeSigningSteps(ctx context.Context, envelopeID string, addEnvelopeSigningStepsRequest AddEnvelopeSigningStepsRequest, opts ...signplusconfig.RequestOption) (*Envelope, error) {
	config := *api.config()
	for _, opt := range api.addEnvelopeSigningStepsConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/envelope/{envelope_id}/signing_steps").
		WithConfig(config).
		WithBody(addEnvelopeSigningStepsRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("envelope_id", envelopeID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[Envelope, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Set envelope attachment settings
func (api *Service) SetEnvelopeAttachmentsSettings(ctx context.Context, envelopeID string, setEnvelopeAttachmentsSettingsRequest SetEnvelopeAttachmentsSettingsRequest, opts ...signplusconfig.RequestOption) (*EnvelopeAttachments, error) {
	config := *api.config()
	for _, opt := range api.setEnvelopeAttachmentsSettingsConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/envelope/{envelope_id}/attachments/settings").
		WithConfig(config).
		WithBody(setEnvelopeAttachmentsSettingsRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("envelope_id", envelopeID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[EnvelopeAttachments, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Placeholders to be set, completely replacing the existing ones.
func (api *Service) SetEnvelopeAttachmentsPlaceholders(ctx context.Context, envelopeID string, setEnvelopeAttachmentsPlaceholdersRequest SetEnvelopeAttachmentsPlaceholdersRequest, opts ...signplusconfig.RequestOption) (*EnvelopeAttachments, error) {
	config := *api.config()
	for _, opt := range api.setEnvelopeAttachmentsPlaceholdersConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/envelope/{envelope_id}/attachments/placeholders").
		WithConfig(config).
		WithBody(setEnvelopeAttachmentsPlaceholdersRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("envelope_id", envelopeID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[EnvelopeAttachments, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Get envelope attachment file
func (api *Service) GetAttachmentFile(ctx context.Context, envelopeID string, fileID string, opts ...signplusconfig.RequestOption) ([]byte, error) {
	config := *api.config()
	for _, opt := range api.getAttachmentFileConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/envelope/{envelope_id}/attachments/{file_id}").
		WithConfig(config).
		AddPathParam("envelope_id", envelopeID).
		AddPathParam("file_id", fileID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeBinary).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]byte, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return resp.Data, nil
}

// Send envelope for signature
func (api *Service) SendEnvelope(ctx context.Context, envelopeID string, opts ...signplusconfig.RequestOption) (*Envelope, error) {
	config := *api.config()
	for _, opt := range api.sendEnvelopeConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/envelope/{envelope_id}/send").
		WithConfig(config).
		AddPathParam("envelope_id", envelopeID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[Envelope, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Duplicate envelope
func (api *Service) DuplicateEnvelope(ctx context.Context, envelopeID string, opts ...signplusconfig.RequestOption) (*Envelope, error) {
	config := *api.config()
	for _, opt := range api.duplicateEnvelopeConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/envelope/{envelope_id}/duplicate").
		WithConfig(config).
		AddPathParam("envelope_id", envelopeID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[Envelope, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Void envelope
func (api *Service) VoidEnvelope(ctx context.Context, envelopeID string, opts ...signplusconfig.RequestOption) (*Envelope, error) {
	config := *api.config()
	for _, opt := range api.voidEnvelopeConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/envelope/{envelope_id}/void").
		WithConfig(config).
		AddPathParam("envelope_id", envelopeID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[Envelope, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Rename envelope
func (api *Service) RenameEnvelope(ctx context.Context, envelopeID string, renameEnvelopeRequest RenameEnvelopeRequest, opts ...signplusconfig.RequestOption) (*Envelope, error) {
	config := *api.config()
	for _, opt := range api.renameEnvelopeConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/envelope/{envelope_id}/rename").
		WithConfig(config).
		WithBody(renameEnvelopeRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("envelope_id", envelopeID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[Envelope, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Set envelope comment
func (api *Service) SetEnvelopeComment(ctx context.Context, envelopeID string, setEnvelopeCommentRequest SetEnvelopeCommentRequest, opts ...signplusconfig.RequestOption) (*Envelope, error) {
	config := *api.config()
	for _, opt := range api.setEnvelopeCommentConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/envelope/{envelope_id}/set_comment").
		WithConfig(config).
		WithBody(setEnvelopeCommentRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("envelope_id", envelopeID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[Envelope, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Set envelope notification
func (api *Service) SetEnvelopeNotification(ctx context.Context, envelopeID string, envelopeNotification EnvelopeNotification, opts ...signplusconfig.RequestOption) (*Envelope, error) {
	config := *api.config()
	for _, opt := range api.setEnvelopeNotificationConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/envelope/{envelope_id}/set_notification").
		WithConfig(config).
		WithBody(envelopeNotification).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("envelope_id", envelopeID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[Envelope, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Set envelope expiration date
func (api *Service) SetEnvelopeExpirationDate(ctx context.Context, envelopeID string, setEnvelopeExpirationRequest SetEnvelopeExpirationRequest, opts ...signplusconfig.RequestOption) (*Envelope, error) {
	config := *api.config()
	for _, opt := range api.setEnvelopeExpirationDateConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/envelope/{envelope_id}/set_expiration_date").
		WithConfig(config).
		WithBody(setEnvelopeExpirationRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("envelope_id", envelopeID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[Envelope, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Set envelope legality level
func (api *Service) SetEnvelopeLegalityLevel(ctx context.Context, envelopeID string, setEnvelopeLegalityLevelRequest SetEnvelopeLegalityLevelRequest, opts ...signplusconfig.RequestOption) (*Envelope, error) {
	config := *api.config()
	for _, opt := range api.setEnvelopeLegalityLevelConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/envelope/{envelope_id}/set_legality_level").
		WithConfig(config).
		WithBody(setEnvelopeLegalityLevelRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("envelope_id", envelopeID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[Envelope, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Get envelope annotations
func (api *Service) GetEnvelopeAnnotations(ctx context.Context, envelopeID string, opts ...signplusconfig.RequestOption) ([]Annotation, error) {
	config := *api.config()
	for _, opt := range api.getEnvelopeAnnotationsConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/envelope/{envelope_id}/annotations").
		WithConfig(config).
		AddPathParam("envelope_id", envelopeID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[[]Annotation, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return resp.Data, nil
}

// Get envelope document annotations
func (api *Service) GetEnvelopeDocumentAnnotations(ctx context.Context, envelopeID string, documentID string, opts ...signplusconfig.RequestOption) (*ListEnvelopeDocumentAnnotationsResponse, error) {
	config := *api.config()
	for _, opt := range api.getEnvelopeDocumentAnnotationsConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/envelope/{envelope_id}/annotations/{document_id}").
		WithConfig(config).
		AddPathParam("envelope_id", envelopeID).
		AddPathParam("document_id", documentID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[ListEnvelopeDocumentAnnotationsResponse, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Add envelope annotation
func (api *Service) AddEnvelopeAnnotation(ctx context.Context, envelopeID string, addAnnotationRequest AddAnnotationRequest, opts ...signplusconfig.RequestOption) (*Annotation, error) {
	config := *api.config()
	for _, opt := range api.addEnvelopeAnnotationConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/envelope/{envelope_id}/annotation").
		WithConfig(config).
		WithBody(addAnnotationRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("envelope_id", envelopeID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[Annotation, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Delete envelope annotation
func (api *Service) DeleteEnvelopeAnnotation(ctx context.Context, envelopeID string, annotationID string, opts ...signplusconfig.RequestOption) (any, error) {
	config := *api.config()
	for _, opt := range api.deleteEnvelopeAnnotationConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/envelope/{envelope_id}/annotation/{annotation_id}").
		WithConfig(config).
		AddPathParam("envelope_id", envelopeID).
		AddPathParam("annotation_id", annotationID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[any, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return resp.Data, nil
}

// Create new template
func (api *Service) CreateTemplate(ctx context.Context, createTemplateRequest CreateTemplateRequest, opts ...signplusconfig.RequestOption) (*Template, error) {
	config := *api.config()
	for _, opt := range api.createTemplateConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/template").
		WithConfig(config).
		WithBody(createTemplateRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[Template, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// List templates
func (api *Service) ListTemplates(ctx context.Context, listTemplatesRequest ListTemplatesRequest, opts ...signplusconfig.RequestOption) (*ListTemplatesResponse, error) {
	config := *api.config()
	for _, opt := range api.listTemplatesConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/templates").
		WithConfig(config).
		WithBody(listTemplatesRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[ListTemplatesResponse, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Get template
func (api *Service) GetTemplate(ctx context.Context, templateID string, opts ...signplusconfig.RequestOption) (*Template, error) {
	config := *api.config()
	for _, opt := range api.getTemplateConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/template/{template_id}").
		WithConfig(config).
		AddPathParam("template_id", templateID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[Template, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Delete template
func (api *Service) DeleteTemplate(ctx context.Context, templateID string, opts ...signplusconfig.RequestOption) (any, error) {
	config := *api.config()
	for _, opt := range api.deleteTemplateConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/template/{template_id}").
		WithConfig(config).
		AddPathParam("template_id", templateID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[any, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return resp.Data, nil
}

// Duplicate template
func (api *Service) DuplicateTemplate(ctx context.Context, templateID string, opts ...signplusconfig.RequestOption) (*Template, error) {
	config := *api.config()
	for _, opt := range api.duplicateTemplateConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/template/{template_id}/duplicate").
		WithConfig(config).
		AddPathParam("template_id", templateID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[Template, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Add template document
func (api *Service) AddTemplateDocument(ctx context.Context, templateID string, addTemplateDocumentRequest AddTemplateDocumentRequest, opts ...signplusconfig.RequestOption) (*Document, error) {
	config := *api.config()
	for _, opt := range api.addTemplateDocumentConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/template/{template_id}/document").
		WithConfig(config).
		WithBody(addTemplateDocumentRequest).
		AddPathParam("template_id", templateID).
		WithContentType(httptransport.ContentTypeMultipartFormData).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[Document, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Get template document
func (api *Service) GetTemplateDocument(ctx context.Context, templateID string, documentID string, opts ...signplusconfig.RequestOption) (*Document, error) {
	config := *api.config()
	for _, opt := range api.getTemplateDocumentConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/template/{template_id}/document/{document_id}").
		WithConfig(config).
		AddPathParam("template_id", templateID).
		AddPathParam("document_id", documentID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[Document, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Get template documents
func (api *Service) GetTemplateDocuments(ctx context.Context, templateID string, opts ...signplusconfig.RequestOption) (*ListTemplateDocumentsResponse, error) {
	config := *api.config()
	for _, opt := range api.getTemplateDocumentsConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/template/{template_id}/documents").
		WithConfig(config).
		AddPathParam("template_id", templateID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[ListTemplateDocumentsResponse, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Add template signing steps
func (api *Service) AddTemplateSigningSteps(ctx context.Context, templateID string, addTemplateSigningStepsRequest AddTemplateSigningStepsRequest, opts ...signplusconfig.RequestOption) (*Template, error) {
	config := *api.config()
	for _, opt := range api.addTemplateSigningStepsConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/template/{template_id}/signing_steps").
		WithConfig(config).
		WithBody(addTemplateSigningStepsRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("template_id", templateID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[Template, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Rename template
func (api *Service) RenameTemplate(ctx context.Context, templateID string, renameTemplateRequest RenameTemplateRequest, opts ...signplusconfig.RequestOption) (*Template, error) {
	config := *api.config()
	for _, opt := range api.renameTemplateConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/template/{template_id}/rename").
		WithConfig(config).
		WithBody(renameTemplateRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("template_id", templateID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[Template, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Set template comment
func (api *Service) SetTemplateComment(ctx context.Context, templateID string, setTemplateCommentRequest SetTemplateCommentRequest, opts ...signplusconfig.RequestOption) (*Template, error) {
	config := *api.config()
	for _, opt := range api.setTemplateCommentConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/template/{template_id}/set_comment").
		WithConfig(config).
		WithBody(setTemplateCommentRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("template_id", templateID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[Template, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Set template notification
func (api *Service) SetTemplateNotification(ctx context.Context, templateID string, envelopeNotification EnvelopeNotification, opts ...signplusconfig.RequestOption) (*Template, error) {
	config := *api.config()
	for _, opt := range api.setTemplateNotificationConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/template/{template_id}/set_notification").
		WithConfig(config).
		WithBody(envelopeNotification).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("template_id", templateID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[Template, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Get template annotations
func (api *Service) GetTemplateAnnotations(ctx context.Context, templateID string, opts ...signplusconfig.RequestOption) (*ListTemplateAnnotationsResponse, error) {
	config := *api.config()
	for _, opt := range api.getTemplateAnnotationsConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/template/{template_id}/annotations").
		WithConfig(config).
		AddPathParam("template_id", templateID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[ListTemplateAnnotationsResponse, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Get document template annotations
func (api *Service) GetDocumentTemplateAnnotations(ctx context.Context, templateID string, documentID string, opts ...signplusconfig.RequestOption) (*ListTemplateDocumentAnnotationsResponse, error) {
	config := *api.config()
	for _, opt := range api.getDocumentTemplateAnnotationsConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("GET").
		WithPath("/template/{template_id}/annotations/{document_id}").
		WithConfig(config).
		AddPathParam("template_id", templateID).
		AddPathParam("document_id", documentID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[ListTemplateDocumentAnnotationsResponse, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Add template annotation
func (api *Service) AddTemplateAnnotation(ctx context.Context, templateID string, addAnnotationRequest AddAnnotationRequest, opts ...signplusconfig.RequestOption) (*Annotation, error) {
	config := *api.config()
	for _, opt := range api.addTemplateAnnotationConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/template/{template_id}/annotation").
		WithConfig(config).
		WithBody(addAnnotationRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("template_id", templateID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[Annotation, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Delete template annotation
func (api *Service) DeleteTemplateAnnotation(ctx context.Context, templateID string, annotationID string, opts ...signplusconfig.RequestOption) (any, error) {
	config := *api.config()
	for _, opt := range api.deleteTemplateAnnotationConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/template/{template_id}/annotation/{annotation_id}").
		WithConfig(config).
		AddPathParam("template_id", templateID).
		AddPathParam("annotation_id", annotationID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[any, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return resp.Data, nil
}

// Set template attachment settings
func (api *Service) SetTemplateAttachmentsSettings(ctx context.Context, templateID string, setEnvelopeAttachmentsSettingsRequest SetEnvelopeAttachmentsSettingsRequest, opts ...signplusconfig.RequestOption) (*EnvelopeAttachments, error) {
	config := *api.config()
	for _, opt := range api.setTemplateAttachmentsSettingsConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/template/{template_id}/attachments/settings").
		WithConfig(config).
		WithBody(setEnvelopeAttachmentsSettingsRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("template_id", templateID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[EnvelopeAttachments, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Placeholders to be set, completely replacing the existing ones.
func (api *Service) SetTemplateAttachmentsPlaceholders(ctx context.Context, templateID string, setEnvelopeAttachmentsPlaceholdersRequest SetEnvelopeAttachmentsPlaceholdersRequest, opts ...signplusconfig.RequestOption) (*EnvelopeAttachments, error) {
	config := *api.config()
	for _, opt := range api.setTemplateAttachmentsPlaceholdersConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("PUT").
		WithPath("/template/{template_id}/attachments/placeholders").
		WithConfig(config).
		WithBody(setEnvelopeAttachmentsPlaceholdersRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		AddPathParam("template_id", templateID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[EnvelopeAttachments, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Create webhook
func (api *Service) CreateWebhook(ctx context.Context, createWebhookRequest CreateWebhookRequest, opts ...signplusconfig.RequestOption) (*Webhook, error) {
	config := *api.config()
	for _, opt := range api.createWebhookConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/webhook").
		WithConfig(config).
		WithBody(createWebhookRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[Webhook, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// List webhooks
func (api *Service) ListWebhooks(ctx context.Context, listWebhooksRequest ListWebhooksRequest, opts ...signplusconfig.RequestOption) (*ListWebhooksResponse, error) {
	config := *api.config()
	for _, opt := range api.listWebhooksConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("POST").
		WithPath("/webhooks").
		WithConfig(config).
		WithBody(listWebhooksRequest).
		AddHeader("CONTENT-TYPE", "application/json").
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[ListWebhooksResponse, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return &resp.Data, nil
}

// Delete webhook
func (api *Service) DeleteWebhook(ctx context.Context, webhookID string, opts ...signplusconfig.RequestOption) (any, error) {
	config := *api.config()
	for _, opt := range api.deleteWebhookConfig {
		opt(&config)
	}
	for _, opt := range opts {
		opt(&config)
	}

	httpRequest := httptransport.NewRequestBuilder().WithContext(ctx).
		WithMethod("DELETE").
		WithPath("/webhook/{webhook_id}").
		WithConfig(config).
		AddPathParam("webhook_id", webhookID).
		WithContentType(httptransport.ContentTypeJSON).
		WithResponseContentType(httptransport.ContentTypeJSON).
		WithSecuritySchemes([]httptransport.AuthScheme{httptransport.AuthSchemeBearer}).
		Build()

	httpClient := restClient.NewRestClient[any, []byte](config, api.getHook())
	resp, err := httpClient.Call(*httpRequest)
	if err != nil {
		return nil, sharedmodels.NewSignplusError[[]byte](err)
	}

	return resp.Data, nil
}
