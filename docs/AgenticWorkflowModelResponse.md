# AgenticWorkflowModelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**AgenticWorkflowModelType**](AgenticWorkflowModelType.md) |  | 
**Settings** | **string** |  | 
**LlmProviderId** | Pointer to **NullableString** | The LLM provider the workflow takes its credential from, or null when it carries its own &#x60;api_key&#x60;. Unlike &#x60;api_key&#x60; this is returned: it names a credential rather than carrying one. | [optional] 

## Methods

### NewAgenticWorkflowModelResponse

`func NewAgenticWorkflowModelResponse(type_ AgenticWorkflowModelType, settings string, ) *AgenticWorkflowModelResponse`

NewAgenticWorkflowModelResponse instantiates a new AgenticWorkflowModelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgenticWorkflowModelResponseWithDefaults

`func NewAgenticWorkflowModelResponseWithDefaults() *AgenticWorkflowModelResponse`

NewAgenticWorkflowModelResponseWithDefaults instantiates a new AgenticWorkflowModelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AgenticWorkflowModelResponse) GetType() AgenticWorkflowModelType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgenticWorkflowModelResponse) GetTypeOk() (*AgenticWorkflowModelType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgenticWorkflowModelResponse) SetType(v AgenticWorkflowModelType)`

SetType sets Type field to given value.


### GetSettings

`func (o *AgenticWorkflowModelResponse) GetSettings() string`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *AgenticWorkflowModelResponse) GetSettingsOk() (*string, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *AgenticWorkflowModelResponse) SetSettings(v string)`

SetSettings sets Settings field to given value.


### GetLlmProviderId

`func (o *AgenticWorkflowModelResponse) GetLlmProviderId() string`

GetLlmProviderId returns the LlmProviderId field if non-nil, zero value otherwise.

### GetLlmProviderIdOk

`func (o *AgenticWorkflowModelResponse) GetLlmProviderIdOk() (*string, bool)`

GetLlmProviderIdOk returns a tuple with the LlmProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLlmProviderId

`func (o *AgenticWorkflowModelResponse) SetLlmProviderId(v string)`

SetLlmProviderId sets LlmProviderId field to given value.

### HasLlmProviderId

`func (o *AgenticWorkflowModelResponse) HasLlmProviderId() bool`

HasLlmProviderId returns a boolean if a field has been set.

### SetLlmProviderIdNil

`func (o *AgenticWorkflowModelResponse) SetLlmProviderIdNil(b bool)`

 SetLlmProviderIdNil sets the value for LlmProviderId to be an explicit nil

### UnsetLlmProviderId
`func (o *AgenticWorkflowModelResponse) UnsetLlmProviderId()`

UnsetLlmProviderId ensures that no value is present for LlmProviderId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


