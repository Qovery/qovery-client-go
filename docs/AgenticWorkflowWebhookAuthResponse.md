# AgenticWorkflowWebhookAuthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | [**AgenticWorkflowWebhookAuthMode**](AgenticWorkflowWebhookAuthMode.md) |  | 
**HeaderName** | **NullableString** | The header a caller must send the secret in, lowercased. Set for &#x60;SECRET_HEADER&#x60; only, null for every other mode. The secret itself is never returned. | 

## Methods

### NewAgenticWorkflowWebhookAuthResponse

`func NewAgenticWorkflowWebhookAuthResponse(mode AgenticWorkflowWebhookAuthMode, headerName NullableString, ) *AgenticWorkflowWebhookAuthResponse`

NewAgenticWorkflowWebhookAuthResponse instantiates a new AgenticWorkflowWebhookAuthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgenticWorkflowWebhookAuthResponseWithDefaults

`func NewAgenticWorkflowWebhookAuthResponseWithDefaults() *AgenticWorkflowWebhookAuthResponse`

NewAgenticWorkflowWebhookAuthResponseWithDefaults instantiates a new AgenticWorkflowWebhookAuthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *AgenticWorkflowWebhookAuthResponse) GetMode() AgenticWorkflowWebhookAuthMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *AgenticWorkflowWebhookAuthResponse) GetModeOk() (*AgenticWorkflowWebhookAuthMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *AgenticWorkflowWebhookAuthResponse) SetMode(v AgenticWorkflowWebhookAuthMode)`

SetMode sets Mode field to given value.


### GetHeaderName

`func (o *AgenticWorkflowWebhookAuthResponse) GetHeaderName() string`

GetHeaderName returns the HeaderName field if non-nil, zero value otherwise.

### GetHeaderNameOk

`func (o *AgenticWorkflowWebhookAuthResponse) GetHeaderNameOk() (*string, bool)`

GetHeaderNameOk returns a tuple with the HeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderName

`func (o *AgenticWorkflowWebhookAuthResponse) SetHeaderName(v string)`

SetHeaderName sets HeaderName field to given value.


### SetHeaderNameNil

`func (o *AgenticWorkflowWebhookAuthResponse) SetHeaderNameNil(b bool)`

 SetHeaderNameNil sets the value for HeaderName to be an explicit nil

### UnsetHeaderName
`func (o *AgenticWorkflowWebhookAuthResponse) UnsetHeaderName()`

UnsetHeaderName ensures that no value is present for HeaderName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


