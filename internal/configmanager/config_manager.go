package configmanager

import (
	"github.com/alohihq/signplus-go/signplusconfig"
	"sync"
	"time"
)

// ConfigManager manages configuration across all services with synchronized updates.
// Provides centralized configuration management and OAuth token handling for multiple services.
type ConfigManager struct {
	templateID                                signplusconfig.Config
	signedDocuments                           signplusconfig.Config
	certificate                               signplusconfig.Config
	documentID                                signplusconfig.Config
	document                                  signplusconfig.Config
	documents                                 signplusconfig.Config
	dynamicFields                             signplusconfig.Config
	signingSteps                              signplusconfig.Config
	settings                                  signplusconfig.Config
	placeholders                              signplusconfig.Config
	fileID                                    signplusconfig.Config
	send                                      signplusconfig.Config
	duplicate                                 signplusconfig.Config
	void                                      signplusconfig.Config
	rename                                    signplusconfig.Config
	setComment                                signplusconfig.Config
	setNotification                           signplusconfig.Config
	setExpirationDate                         signplusconfig.Config
	setLegalityLevel                          signplusconfig.Config
	envelopeEnvelopeIDAnnotationsDocumentID   signplusconfig.Config
	annotations                               signplusconfig.Config
	annotationID                              signplusconfig.Config
	annotation                                signplusconfig.Config
	envelopeID                                signplusconfig.Config
	envelope                                  signplusconfig.Config
	envelopes                                 signplusconfig.Config
	templateTemplateIDDuplicate               signplusconfig.Config
	templateTemplateIDDocumentDocumentID      signplusconfig.Config
	templateTemplateIDDocument                signplusconfig.Config
	templateTemplateIDDocuments               signplusconfig.Config
	templateTemplateIDSigningSteps            signplusconfig.Config
	templateTemplateIDRename                  signplusconfig.Config
	templateTemplateIDSetComment              signplusconfig.Config
	templateTemplateIDSetNotification         signplusconfig.Config
	templateTemplateIDAnnotationsDocumentID   signplusconfig.Config
	templateTemplateIDAnnotations             signplusconfig.Config
	templateTemplateIDAnnotationAnnotationID  signplusconfig.Config
	templateTemplateIDAnnotation              signplusconfig.Config
	templateTemplateIDAttachmentsSettings     signplusconfig.Config
	templateTemplateIDAttachmentsPlaceholders signplusconfig.Config
	templateTemplateID                        signplusconfig.Config
	template                                  signplusconfig.Config
	templates                                 signplusconfig.Config
	webhookID                                 signplusconfig.Config
	webhook                                   signplusconfig.Config
	webhooks                                  signplusconfig.Config
	// mu protects concurrent reads and writes to access/refresh token fields.
	// Guarded operations: GetAccessToken, GetRefreshToken, UpdateAccessToken.
	mu sync.RWMutex
}

// NewConfigManager creates a new configuration manager with the provided config and optional OAuth token service.
// Initializes service-specific configs and sets up OAuth token management if enabled.
func NewConfigManager(config signplusconfig.Config) *ConfigManager {
	return &ConfigManager{
		templateID:                               config,
		signedDocuments:                          config,
		certificate:                              config,
		documentID:                               config,
		document:                                 config,
		documents:                                config,
		dynamicFields:                            config,
		signingSteps:                             config,
		settings:                                 config,
		placeholders:                             config,
		fileID:                                   config,
		send:                                     config,
		duplicate:                                config,
		void:                                     config,
		rename:                                   config,
		setComment:                               config,
		setNotification:                          config,
		setExpirationDate:                        config,
		setLegalityLevel:                         config,
		envelopeEnvelopeIDAnnotationsDocumentID:  config,
		annotations:                              config,
		annotationID:                             config,
		annotation:                               config,
		envelopeID:                               config,
		envelope:                                 config,
		envelopes:                                config,
		templateTemplateIDDuplicate:              config,
		templateTemplateIDDocumentDocumentID:     config,
		templateTemplateIDDocument:               config,
		templateTemplateIDDocuments:              config,
		templateTemplateIDSigningSteps:           config,
		templateTemplateIDRename:                 config,
		templateTemplateIDSetComment:             config,
		templateTemplateIDSetNotification:        config,
		templateTemplateIDAnnotationsDocumentID:  config,
		templateTemplateIDAnnotations:            config,
		templateTemplateIDAnnotationAnnotationID: config,
		templateTemplateIDAnnotation:             config,
		templateTemplateIDAttachmentsSettings:    config,
		templateTemplateIDAttachmentsPlaceholders: config,
		templateTemplateID:                        config,
		template:                                  config,
		templates:                                 config,
		webhookID:                                 config,
		webhook:                                   config,
		webhooks:                                  config,
	}
}

// SetBaseURL updates the BaseURL configuration parameter across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetBaseURL(baseURL string) {
	c.templateID.SetBaseURL(baseURL)
	c.signedDocuments.SetBaseURL(baseURL)
	c.certificate.SetBaseURL(baseURL)
	c.documentID.SetBaseURL(baseURL)
	c.document.SetBaseURL(baseURL)
	c.documents.SetBaseURL(baseURL)
	c.dynamicFields.SetBaseURL(baseURL)
	c.signingSteps.SetBaseURL(baseURL)
	c.settings.SetBaseURL(baseURL)
	c.placeholders.SetBaseURL(baseURL)
	c.fileID.SetBaseURL(baseURL)
	c.send.SetBaseURL(baseURL)
	c.duplicate.SetBaseURL(baseURL)
	c.void.SetBaseURL(baseURL)
	c.rename.SetBaseURL(baseURL)
	c.setComment.SetBaseURL(baseURL)
	c.setNotification.SetBaseURL(baseURL)
	c.setExpirationDate.SetBaseURL(baseURL)
	c.setLegalityLevel.SetBaseURL(baseURL)
	c.envelopeEnvelopeIDAnnotationsDocumentID.SetBaseURL(baseURL)
	c.annotations.SetBaseURL(baseURL)
	c.annotationID.SetBaseURL(baseURL)
	c.annotation.SetBaseURL(baseURL)
	c.envelopeID.SetBaseURL(baseURL)
	c.envelope.SetBaseURL(baseURL)
	c.envelopes.SetBaseURL(baseURL)
	c.templateTemplateIDDuplicate.SetBaseURL(baseURL)
	c.templateTemplateIDDocumentDocumentID.SetBaseURL(baseURL)
	c.templateTemplateIDDocument.SetBaseURL(baseURL)
	c.templateTemplateIDDocuments.SetBaseURL(baseURL)
	c.templateTemplateIDSigningSteps.SetBaseURL(baseURL)
	c.templateTemplateIDRename.SetBaseURL(baseURL)
	c.templateTemplateIDSetComment.SetBaseURL(baseURL)
	c.templateTemplateIDSetNotification.SetBaseURL(baseURL)
	c.templateTemplateIDAnnotationsDocumentID.SetBaseURL(baseURL)
	c.templateTemplateIDAnnotations.SetBaseURL(baseURL)
	c.templateTemplateIDAnnotationAnnotationID.SetBaseURL(baseURL)
	c.templateTemplateIDAnnotation.SetBaseURL(baseURL)
	c.templateTemplateIDAttachmentsSettings.SetBaseURL(baseURL)
	c.templateTemplateIDAttachmentsPlaceholders.SetBaseURL(baseURL)
	c.templateTemplateID.SetBaseURL(baseURL)
	c.template.SetBaseURL(baseURL)
	c.templates.SetBaseURL(baseURL)
	c.webhookID.SetBaseURL(baseURL)
	c.webhook.SetBaseURL(baseURL)
	c.webhooks.SetBaseURL(baseURL)
}

// SetTimeout updates the Timeout configuration parameter across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetTimeout(timeout time.Duration) {
	c.templateID.SetTimeout(timeout)
	c.signedDocuments.SetTimeout(timeout)
	c.certificate.SetTimeout(timeout)
	c.documentID.SetTimeout(timeout)
	c.document.SetTimeout(timeout)
	c.documents.SetTimeout(timeout)
	c.dynamicFields.SetTimeout(timeout)
	c.signingSteps.SetTimeout(timeout)
	c.settings.SetTimeout(timeout)
	c.placeholders.SetTimeout(timeout)
	c.fileID.SetTimeout(timeout)
	c.send.SetTimeout(timeout)
	c.duplicate.SetTimeout(timeout)
	c.void.SetTimeout(timeout)
	c.rename.SetTimeout(timeout)
	c.setComment.SetTimeout(timeout)
	c.setNotification.SetTimeout(timeout)
	c.setExpirationDate.SetTimeout(timeout)
	c.setLegalityLevel.SetTimeout(timeout)
	c.envelopeEnvelopeIDAnnotationsDocumentID.SetTimeout(timeout)
	c.annotations.SetTimeout(timeout)
	c.annotationID.SetTimeout(timeout)
	c.annotation.SetTimeout(timeout)
	c.envelopeID.SetTimeout(timeout)
	c.envelope.SetTimeout(timeout)
	c.envelopes.SetTimeout(timeout)
	c.templateTemplateIDDuplicate.SetTimeout(timeout)
	c.templateTemplateIDDocumentDocumentID.SetTimeout(timeout)
	c.templateTemplateIDDocument.SetTimeout(timeout)
	c.templateTemplateIDDocuments.SetTimeout(timeout)
	c.templateTemplateIDSigningSteps.SetTimeout(timeout)
	c.templateTemplateIDRename.SetTimeout(timeout)
	c.templateTemplateIDSetComment.SetTimeout(timeout)
	c.templateTemplateIDSetNotification.SetTimeout(timeout)
	c.templateTemplateIDAnnotationsDocumentID.SetTimeout(timeout)
	c.templateTemplateIDAnnotations.SetTimeout(timeout)
	c.templateTemplateIDAnnotationAnnotationID.SetTimeout(timeout)
	c.templateTemplateIDAnnotation.SetTimeout(timeout)
	c.templateTemplateIDAttachmentsSettings.SetTimeout(timeout)
	c.templateTemplateIDAttachmentsPlaceholders.SetTimeout(timeout)
	c.templateTemplateID.SetTimeout(timeout)
	c.template.SetTimeout(timeout)
	c.templates.SetTimeout(timeout)
	c.webhookID.SetTimeout(timeout)
	c.webhook.SetTimeout(timeout)
	c.webhooks.SetTimeout(timeout)
}

// SetAccessToken updates the AccessToken configuration parameter across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetAccessToken(accessToken string) {
	c.templateID.SetAccessToken(accessToken)
	c.signedDocuments.SetAccessToken(accessToken)
	c.certificate.SetAccessToken(accessToken)
	c.documentID.SetAccessToken(accessToken)
	c.document.SetAccessToken(accessToken)
	c.documents.SetAccessToken(accessToken)
	c.dynamicFields.SetAccessToken(accessToken)
	c.signingSteps.SetAccessToken(accessToken)
	c.settings.SetAccessToken(accessToken)
	c.placeholders.SetAccessToken(accessToken)
	c.fileID.SetAccessToken(accessToken)
	c.send.SetAccessToken(accessToken)
	c.duplicate.SetAccessToken(accessToken)
	c.void.SetAccessToken(accessToken)
	c.rename.SetAccessToken(accessToken)
	c.setComment.SetAccessToken(accessToken)
	c.setNotification.SetAccessToken(accessToken)
	c.setExpirationDate.SetAccessToken(accessToken)
	c.setLegalityLevel.SetAccessToken(accessToken)
	c.envelopeEnvelopeIDAnnotationsDocumentID.SetAccessToken(accessToken)
	c.annotations.SetAccessToken(accessToken)
	c.annotationID.SetAccessToken(accessToken)
	c.annotation.SetAccessToken(accessToken)
	c.envelopeID.SetAccessToken(accessToken)
	c.envelope.SetAccessToken(accessToken)
	c.envelopes.SetAccessToken(accessToken)
	c.templateTemplateIDDuplicate.SetAccessToken(accessToken)
	c.templateTemplateIDDocumentDocumentID.SetAccessToken(accessToken)
	c.templateTemplateIDDocument.SetAccessToken(accessToken)
	c.templateTemplateIDDocuments.SetAccessToken(accessToken)
	c.templateTemplateIDSigningSteps.SetAccessToken(accessToken)
	c.templateTemplateIDRename.SetAccessToken(accessToken)
	c.templateTemplateIDSetComment.SetAccessToken(accessToken)
	c.templateTemplateIDSetNotification.SetAccessToken(accessToken)
	c.templateTemplateIDAnnotationsDocumentID.SetAccessToken(accessToken)
	c.templateTemplateIDAnnotations.SetAccessToken(accessToken)
	c.templateTemplateIDAnnotationAnnotationID.SetAccessToken(accessToken)
	c.templateTemplateIDAnnotation.SetAccessToken(accessToken)
	c.templateTemplateIDAttachmentsSettings.SetAccessToken(accessToken)
	c.templateTemplateIDAttachmentsPlaceholders.SetAccessToken(accessToken)
	c.templateTemplateID.SetAccessToken(accessToken)
	c.template.SetAccessToken(accessToken)
	c.templates.SetAccessToken(accessToken)
	c.webhookID.SetAccessToken(accessToken)
	c.webhook.SetAccessToken(accessToken)
	c.webhooks.SetAccessToken(accessToken)
}

// SetRetryConfig updates the retry configuration across all services.
// Changes are applied synchronously to all registered service configurations.
func (c *ConfigManager) SetRetryConfig(retry signplusconfig.RetryConfig) {
	c.templateID.SetRetryConfig(retry)
	c.signedDocuments.SetRetryConfig(retry)
	c.certificate.SetRetryConfig(retry)
	c.documentID.SetRetryConfig(retry)
	c.document.SetRetryConfig(retry)
	c.documents.SetRetryConfig(retry)
	c.dynamicFields.SetRetryConfig(retry)
	c.signingSteps.SetRetryConfig(retry)
	c.settings.SetRetryConfig(retry)
	c.placeholders.SetRetryConfig(retry)
	c.fileID.SetRetryConfig(retry)
	c.send.SetRetryConfig(retry)
	c.duplicate.SetRetryConfig(retry)
	c.void.SetRetryConfig(retry)
	c.rename.SetRetryConfig(retry)
	c.setComment.SetRetryConfig(retry)
	c.setNotification.SetRetryConfig(retry)
	c.setExpirationDate.SetRetryConfig(retry)
	c.setLegalityLevel.SetRetryConfig(retry)
	c.envelopeEnvelopeIDAnnotationsDocumentID.SetRetryConfig(retry)
	c.annotations.SetRetryConfig(retry)
	c.annotationID.SetRetryConfig(retry)
	c.annotation.SetRetryConfig(retry)
	c.envelopeID.SetRetryConfig(retry)
	c.envelope.SetRetryConfig(retry)
	c.envelopes.SetRetryConfig(retry)
	c.templateTemplateIDDuplicate.SetRetryConfig(retry)
	c.templateTemplateIDDocumentDocumentID.SetRetryConfig(retry)
	c.templateTemplateIDDocument.SetRetryConfig(retry)
	c.templateTemplateIDDocuments.SetRetryConfig(retry)
	c.templateTemplateIDSigningSteps.SetRetryConfig(retry)
	c.templateTemplateIDRename.SetRetryConfig(retry)
	c.templateTemplateIDSetComment.SetRetryConfig(retry)
	c.templateTemplateIDSetNotification.SetRetryConfig(retry)
	c.templateTemplateIDAnnotationsDocumentID.SetRetryConfig(retry)
	c.templateTemplateIDAnnotations.SetRetryConfig(retry)
	c.templateTemplateIDAnnotationAnnotationID.SetRetryConfig(retry)
	c.templateTemplateIDAnnotation.SetRetryConfig(retry)
	c.templateTemplateIDAttachmentsSettings.SetRetryConfig(retry)
	c.templateTemplateIDAttachmentsPlaceholders.SetRetryConfig(retry)
	c.templateTemplateID.SetRetryConfig(retry)
	c.template.SetRetryConfig(retry)
	c.templates.SetRetryConfig(retry)
	c.webhookID.SetRetryConfig(retry)
	c.webhook.SetRetryConfig(retry)
	c.webhooks.SetRetryConfig(retry)
}

// UpdateAccessToken replaces an access token across all services that use the original value.
// Used for token refresh to update all service configurations simultaneously.
// Write-locked so concurrent GetAccessToken reads see a consistent value.
func (c *ConfigManager) UpdateAccessToken(originalValue string, newValue string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.templateID.AccessToken != nil && *c.templateID.AccessToken == originalValue {
		c.templateID.SetAccessToken(newValue)
	}

	if c.signedDocuments.AccessToken != nil && *c.signedDocuments.AccessToken == originalValue {
		c.signedDocuments.SetAccessToken(newValue)
	}

	if c.certificate.AccessToken != nil && *c.certificate.AccessToken == originalValue {
		c.certificate.SetAccessToken(newValue)
	}

	if c.documentID.AccessToken != nil && *c.documentID.AccessToken == originalValue {
		c.documentID.SetAccessToken(newValue)
	}

	if c.document.AccessToken != nil && *c.document.AccessToken == originalValue {
		c.document.SetAccessToken(newValue)
	}

	if c.documents.AccessToken != nil && *c.documents.AccessToken == originalValue {
		c.documents.SetAccessToken(newValue)
	}

	if c.dynamicFields.AccessToken != nil && *c.dynamicFields.AccessToken == originalValue {
		c.dynamicFields.SetAccessToken(newValue)
	}

	if c.signingSteps.AccessToken != nil && *c.signingSteps.AccessToken == originalValue {
		c.signingSteps.SetAccessToken(newValue)
	}

	if c.settings.AccessToken != nil && *c.settings.AccessToken == originalValue {
		c.settings.SetAccessToken(newValue)
	}

	if c.placeholders.AccessToken != nil && *c.placeholders.AccessToken == originalValue {
		c.placeholders.SetAccessToken(newValue)
	}

	if c.fileID.AccessToken != nil && *c.fileID.AccessToken == originalValue {
		c.fileID.SetAccessToken(newValue)
	}

	if c.send.AccessToken != nil && *c.send.AccessToken == originalValue {
		c.send.SetAccessToken(newValue)
	}

	if c.duplicate.AccessToken != nil && *c.duplicate.AccessToken == originalValue {
		c.duplicate.SetAccessToken(newValue)
	}

	if c.void.AccessToken != nil && *c.void.AccessToken == originalValue {
		c.void.SetAccessToken(newValue)
	}

	if c.rename.AccessToken != nil && *c.rename.AccessToken == originalValue {
		c.rename.SetAccessToken(newValue)
	}

	if c.setComment.AccessToken != nil && *c.setComment.AccessToken == originalValue {
		c.setComment.SetAccessToken(newValue)
	}

	if c.setNotification.AccessToken != nil && *c.setNotification.AccessToken == originalValue {
		c.setNotification.SetAccessToken(newValue)
	}

	if c.setExpirationDate.AccessToken != nil && *c.setExpirationDate.AccessToken == originalValue {
		c.setExpirationDate.SetAccessToken(newValue)
	}

	if c.setLegalityLevel.AccessToken != nil && *c.setLegalityLevel.AccessToken == originalValue {
		c.setLegalityLevel.SetAccessToken(newValue)
	}

	if c.envelopeEnvelopeIDAnnotationsDocumentID.AccessToken != nil && *c.envelopeEnvelopeIDAnnotationsDocumentID.AccessToken == originalValue {
		c.envelopeEnvelopeIDAnnotationsDocumentID.SetAccessToken(newValue)
	}

	if c.annotations.AccessToken != nil && *c.annotations.AccessToken == originalValue {
		c.annotations.SetAccessToken(newValue)
	}

	if c.annotationID.AccessToken != nil && *c.annotationID.AccessToken == originalValue {
		c.annotationID.SetAccessToken(newValue)
	}

	if c.annotation.AccessToken != nil && *c.annotation.AccessToken == originalValue {
		c.annotation.SetAccessToken(newValue)
	}

	if c.envelopeID.AccessToken != nil && *c.envelopeID.AccessToken == originalValue {
		c.envelopeID.SetAccessToken(newValue)
	}

	if c.envelope.AccessToken != nil && *c.envelope.AccessToken == originalValue {
		c.envelope.SetAccessToken(newValue)
	}

	if c.envelopes.AccessToken != nil && *c.envelopes.AccessToken == originalValue {
		c.envelopes.SetAccessToken(newValue)
	}

	if c.templateTemplateIDDuplicate.AccessToken != nil && *c.templateTemplateIDDuplicate.AccessToken == originalValue {
		c.templateTemplateIDDuplicate.SetAccessToken(newValue)
	}

	if c.templateTemplateIDDocumentDocumentID.AccessToken != nil && *c.templateTemplateIDDocumentDocumentID.AccessToken == originalValue {
		c.templateTemplateIDDocumentDocumentID.SetAccessToken(newValue)
	}

	if c.templateTemplateIDDocument.AccessToken != nil && *c.templateTemplateIDDocument.AccessToken == originalValue {
		c.templateTemplateIDDocument.SetAccessToken(newValue)
	}

	if c.templateTemplateIDDocuments.AccessToken != nil && *c.templateTemplateIDDocuments.AccessToken == originalValue {
		c.templateTemplateIDDocuments.SetAccessToken(newValue)
	}

	if c.templateTemplateIDSigningSteps.AccessToken != nil && *c.templateTemplateIDSigningSteps.AccessToken == originalValue {
		c.templateTemplateIDSigningSteps.SetAccessToken(newValue)
	}

	if c.templateTemplateIDRename.AccessToken != nil && *c.templateTemplateIDRename.AccessToken == originalValue {
		c.templateTemplateIDRename.SetAccessToken(newValue)
	}

	if c.templateTemplateIDSetComment.AccessToken != nil && *c.templateTemplateIDSetComment.AccessToken == originalValue {
		c.templateTemplateIDSetComment.SetAccessToken(newValue)
	}

	if c.templateTemplateIDSetNotification.AccessToken != nil && *c.templateTemplateIDSetNotification.AccessToken == originalValue {
		c.templateTemplateIDSetNotification.SetAccessToken(newValue)
	}

	if c.templateTemplateIDAnnotationsDocumentID.AccessToken != nil && *c.templateTemplateIDAnnotationsDocumentID.AccessToken == originalValue {
		c.templateTemplateIDAnnotationsDocumentID.SetAccessToken(newValue)
	}

	if c.templateTemplateIDAnnotations.AccessToken != nil && *c.templateTemplateIDAnnotations.AccessToken == originalValue {
		c.templateTemplateIDAnnotations.SetAccessToken(newValue)
	}

	if c.templateTemplateIDAnnotationAnnotationID.AccessToken != nil && *c.templateTemplateIDAnnotationAnnotationID.AccessToken == originalValue {
		c.templateTemplateIDAnnotationAnnotationID.SetAccessToken(newValue)
	}

	if c.templateTemplateIDAnnotation.AccessToken != nil && *c.templateTemplateIDAnnotation.AccessToken == originalValue {
		c.templateTemplateIDAnnotation.SetAccessToken(newValue)
	}

	if c.templateTemplateIDAttachmentsSettings.AccessToken != nil && *c.templateTemplateIDAttachmentsSettings.AccessToken == originalValue {
		c.templateTemplateIDAttachmentsSettings.SetAccessToken(newValue)
	}

	if c.templateTemplateIDAttachmentsPlaceholders.AccessToken != nil && *c.templateTemplateIDAttachmentsPlaceholders.AccessToken == originalValue {
		c.templateTemplateIDAttachmentsPlaceholders.SetAccessToken(newValue)
	}

	if c.templateTemplateID.AccessToken != nil && *c.templateTemplateID.AccessToken == originalValue {
		c.templateTemplateID.SetAccessToken(newValue)
	}

	if c.template.AccessToken != nil && *c.template.AccessToken == originalValue {
		c.template.SetAccessToken(newValue)
	}

	if c.templates.AccessToken != nil && *c.templates.AccessToken == originalValue {
		c.templates.SetAccessToken(newValue)
	}

	if c.webhookID.AccessToken != nil && *c.webhookID.AccessToken == originalValue {
		c.webhookID.SetAccessToken(newValue)
	}

	if c.webhook.AccessToken != nil && *c.webhook.AccessToken == originalValue {
		c.webhook.SetAccessToken(newValue)
	}

	if c.webhooks.AccessToken != nil && *c.webhooks.AccessToken == originalValue {
		c.webhooks.SetAccessToken(newValue)
	}
}

// GetTemplateID returns the configuration for the TemplateID service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTemplateID() *signplusconfig.Config {
	return &c.templateID
}

// GetSignedDocuments returns the configuration for the SignedDocuments service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetSignedDocuments() *signplusconfig.Config {
	return &c.signedDocuments
}

// GetCertificate returns the configuration for the Certificate service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetCertificate() *signplusconfig.Config {
	return &c.certificate
}

// GetDocumentID returns the configuration for the DocumentID service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetDocumentID() *signplusconfig.Config {
	return &c.documentID
}

// GetDocument returns the configuration for the Document service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetDocument() *signplusconfig.Config {
	return &c.document
}

// GetDocuments returns the configuration for the Documents service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetDocuments() *signplusconfig.Config {
	return &c.documents
}

// GetDynamicFields returns the configuration for the DynamicFields service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetDynamicFields() *signplusconfig.Config {
	return &c.dynamicFields
}

// GetSigningSteps returns the configuration for the SigningSteps service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetSigningSteps() *signplusconfig.Config {
	return &c.signingSteps
}

// GetSettings returns the configuration for the Settings service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetSettings() *signplusconfig.Config {
	return &c.settings
}

// GetPlaceholders returns the configuration for the Placeholders service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetPlaceholders() *signplusconfig.Config {
	return &c.placeholders
}

// GetFileID returns the configuration for the FileID service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetFileID() *signplusconfig.Config {
	return &c.fileID
}

// GetSend returns the configuration for the Send service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetSend() *signplusconfig.Config {
	return &c.send
}

// GetDuplicate returns the configuration for the Duplicate service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetDuplicate() *signplusconfig.Config {
	return &c.duplicate
}

// GetVoid returns the configuration for the Void service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetVoid() *signplusconfig.Config {
	return &c.void
}

// GetRename returns the configuration for the Rename service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetRename() *signplusconfig.Config {
	return &c.rename
}

// GetSetComment returns the configuration for the SetComment service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetSetComment() *signplusconfig.Config {
	return &c.setComment
}

// GetSetNotification returns the configuration for the SetNotification service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetSetNotification() *signplusconfig.Config {
	return &c.setNotification
}

// GetSetExpirationDate returns the configuration for the SetExpirationDate service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetSetExpirationDate() *signplusconfig.Config {
	return &c.setExpirationDate
}

// GetSetLegalityLevel returns the configuration for the SetLegalityLevel service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetSetLegalityLevel() *signplusconfig.Config {
	return &c.setLegalityLevel
}

// GetEnvelopeEnvelopeIDAnnotationsDocumentID returns the configuration for the EnvelopeEnvelopeIDAnnotationsDocumentID service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetEnvelopeEnvelopeIDAnnotationsDocumentID() *signplusconfig.Config {
	return &c.envelopeEnvelopeIDAnnotationsDocumentID
}

// GetAnnotations returns the configuration for the Annotations service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetAnnotations() *signplusconfig.Config {
	return &c.annotations
}

// GetAnnotationID returns the configuration for the AnnotationID service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetAnnotationID() *signplusconfig.Config {
	return &c.annotationID
}

// GetAnnotation returns the configuration for the Annotation service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetAnnotation() *signplusconfig.Config {
	return &c.annotation
}

// GetEnvelopeID returns the configuration for the EnvelopeID service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetEnvelopeID() *signplusconfig.Config {
	return &c.envelopeID
}

// GetEnvelope returns the configuration for the Envelope service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetEnvelope() *signplusconfig.Config {
	return &c.envelope
}

// GetEnvelopes returns the configuration for the Envelopes service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetEnvelopes() *signplusconfig.Config {
	return &c.envelopes
}

// GetTemplateTemplateIDDuplicate returns the configuration for the TemplateTemplateIDDuplicate service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTemplateTemplateIDDuplicate() *signplusconfig.Config {
	return &c.templateTemplateIDDuplicate
}

// GetTemplateTemplateIDDocumentDocumentID returns the configuration for the TemplateTemplateIDDocumentDocumentID service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTemplateTemplateIDDocumentDocumentID() *signplusconfig.Config {
	return &c.templateTemplateIDDocumentDocumentID
}

// GetTemplateTemplateIDDocument returns the configuration for the TemplateTemplateIDDocument service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTemplateTemplateIDDocument() *signplusconfig.Config {
	return &c.templateTemplateIDDocument
}

// GetTemplateTemplateIDDocuments returns the configuration for the TemplateTemplateIDDocuments service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTemplateTemplateIDDocuments() *signplusconfig.Config {
	return &c.templateTemplateIDDocuments
}

// GetTemplateTemplateIDSigningSteps returns the configuration for the TemplateTemplateIDSigningSteps service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTemplateTemplateIDSigningSteps() *signplusconfig.Config {
	return &c.templateTemplateIDSigningSteps
}

// GetTemplateTemplateIDRename returns the configuration for the TemplateTemplateIDRename service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTemplateTemplateIDRename() *signplusconfig.Config {
	return &c.templateTemplateIDRename
}

// GetTemplateTemplateIDSetComment returns the configuration for the TemplateTemplateIDSetComment service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTemplateTemplateIDSetComment() *signplusconfig.Config {
	return &c.templateTemplateIDSetComment
}

// GetTemplateTemplateIDSetNotification returns the configuration for the TemplateTemplateIDSetNotification service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTemplateTemplateIDSetNotification() *signplusconfig.Config {
	return &c.templateTemplateIDSetNotification
}

// GetTemplateTemplateIDAnnotationsDocumentID returns the configuration for the TemplateTemplateIDAnnotationsDocumentID service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTemplateTemplateIDAnnotationsDocumentID() *signplusconfig.Config {
	return &c.templateTemplateIDAnnotationsDocumentID
}

// GetTemplateTemplateIDAnnotations returns the configuration for the TemplateTemplateIDAnnotations service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTemplateTemplateIDAnnotations() *signplusconfig.Config {
	return &c.templateTemplateIDAnnotations
}

// GetTemplateTemplateIDAnnotationAnnotationID returns the configuration for the TemplateTemplateIDAnnotationAnnotationID service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTemplateTemplateIDAnnotationAnnotationID() *signplusconfig.Config {
	return &c.templateTemplateIDAnnotationAnnotationID
}

// GetTemplateTemplateIDAnnotation returns the configuration for the TemplateTemplateIDAnnotation service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTemplateTemplateIDAnnotation() *signplusconfig.Config {
	return &c.templateTemplateIDAnnotation
}

// GetTemplateTemplateIDAttachmentsSettings returns the configuration for the TemplateTemplateIDAttachmentsSettings service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTemplateTemplateIDAttachmentsSettings() *signplusconfig.Config {
	return &c.templateTemplateIDAttachmentsSettings
}

// GetTemplateTemplateIDAttachmentsPlaceholders returns the configuration for the TemplateTemplateIDAttachmentsPlaceholders service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTemplateTemplateIDAttachmentsPlaceholders() *signplusconfig.Config {
	return &c.templateTemplateIDAttachmentsPlaceholders
}

// GetTemplateTemplateID returns the configuration for the TemplateTemplateID service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTemplateTemplateID() *signplusconfig.Config {
	return &c.templateTemplateID
}

// GetTemplate returns the configuration for the Template service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTemplate() *signplusconfig.Config {
	return &c.template
}

// GetTemplates returns the configuration for the Templates service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetTemplates() *signplusconfig.Config {
	return &c.templates
}

// GetWebhookID returns the configuration for the WebhookID service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetWebhookID() *signplusconfig.Config {
	return &c.webhookID
}

// GetWebhook returns the configuration for the Webhook service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetWebhook() *signplusconfig.Config {
	return &c.webhook
}

// GetWebhooks returns the configuration for the Webhooks service.
// Returns a pointer to the service-specific config for use in API calls.
func (c *ConfigManager) GetWebhooks() *signplusconfig.Config {
	return &c.webhooks
}

// GetBaseURL returns the currently configured base URL.
// All services share the same base URL; this reads it from the first service's config.
func (c *ConfigManager) GetBaseURL() string {
	return c.templateID.BaseURL
}
