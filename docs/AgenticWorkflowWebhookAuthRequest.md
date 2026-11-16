# AgenticWorkflowWebhookAuthRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | [**AgenticWorkflowWebhookAuthMode**](AgenticWorkflowWebhookAuthMode.md) |  | 
**Secret** | Pointer to **string** | Write-only, never returned. Ignored for &#x60;NONE&#x60;.  For &#x60;SECRET_HEADER&#x60;, the secret the caller will send. It must be 16 to 1024 printable ASCII characters, without leading or trailing spaces. Required when entering the mode. Once in the mode, omit it (or send a blank value) to keep the current secret; send a new one to replace it, which is how the secret is rotated. Only its SHA-256 hash is stored.  For the provider modes, the provider&#39;s webhook signing secret. Required when entering the mode; omitted or blank keeps the stored one while the mode does not change.  An invalid or missing secret is rejected with 400 on &#x60;webhook_auth.secret&#x60;.  | [optional] 
**HeaderName** | Pointer to **string** | &#x60;SECRET_HEADER&#x60; only: the HTTP header that carries the secret. Must be a valid header name (RFC 9110 token) of at most 128 characters, and not a header that proxies rewrite or consume (&#x60;authorization&#x60;, &#x60;cookie&#x60;, &#x60;host&#x60;, &#x60;content-type&#x60;, &#x60;x-forwarded-*&#x60;...). It is trimmed and stored lowercased. Omitted or blank keeps the current name, or uses &#x60;x-qovery-webhook-secret&#x60; when entering the mode. An invalid name is rejected with 400 on &#x60;webhook_auth.header_name&#x60;. | [optional] 

## Methods

### NewAgenticWorkflowWebhookAuthRequest

`func NewAgenticWorkflowWebhookAuthRequest(mode AgenticWorkflowWebhookAuthMode, ) *AgenticWorkflowWebhookAuthRequest`

NewAgenticWorkflowWebhookAuthRequest instantiates a new AgenticWorkflowWebhookAuthRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgenticWorkflowWebhookAuthRequestWithDefaults

`func NewAgenticWorkflowWebhookAuthRequestWithDefaults() *AgenticWorkflowWebhookAuthRequest`

NewAgenticWorkflowWebhookAuthRequestWithDefaults instantiates a new AgenticWorkflowWebhookAuthRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *AgenticWorkflowWebhookAuthRequest) GetMode() AgenticWorkflowWebhookAuthMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *AgenticWorkflowWebhookAuthRequest) GetModeOk() (*AgenticWorkflowWebhookAuthMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *AgenticWorkflowWebhookAuthRequest) SetMode(v AgenticWorkflowWebhookAuthMode)`

SetMode sets Mode field to given value.


### GetSecret

`func (o *AgenticWorkflowWebhookAuthRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *AgenticWorkflowWebhookAuthRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *AgenticWorkflowWebhookAuthRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *AgenticWorkflowWebhookAuthRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetHeaderName

`func (o *AgenticWorkflowWebhookAuthRequest) GetHeaderName() string`

GetHeaderName returns the HeaderName field if non-nil, zero value otherwise.

### GetHeaderNameOk

`func (o *AgenticWorkflowWebhookAuthRequest) GetHeaderNameOk() (*string, bool)`

GetHeaderNameOk returns a tuple with the HeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderName

`func (o *AgenticWorkflowWebhookAuthRequest) SetHeaderName(v string)`

SetHeaderName sets HeaderName field to given value.

### HasHeaderName

`func (o *AgenticWorkflowWebhookAuthRequest) HasHeaderName() bool`

HasHeaderName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


