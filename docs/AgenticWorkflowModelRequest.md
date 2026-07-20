# AgenticWorkflowModelRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**AgenticWorkflowModelType**](AgenticWorkflowModelType.md) |  | 
**ApiKey** | Pointer to **string** | Write-only. Provider API key; accepted on create/edit but never returned in responses. | [optional] [default to ""]
**Settings** | Pointer to **string** |  | [optional] [default to ""]

## Methods

### NewAgenticWorkflowModelRequest

`func NewAgenticWorkflowModelRequest(type_ AgenticWorkflowModelType, ) *AgenticWorkflowModelRequest`

NewAgenticWorkflowModelRequest instantiates a new AgenticWorkflowModelRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgenticWorkflowModelRequestWithDefaults

`func NewAgenticWorkflowModelRequestWithDefaults() *AgenticWorkflowModelRequest`

NewAgenticWorkflowModelRequestWithDefaults instantiates a new AgenticWorkflowModelRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AgenticWorkflowModelRequest) GetType() AgenticWorkflowModelType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AgenticWorkflowModelRequest) GetTypeOk() (*AgenticWorkflowModelType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AgenticWorkflowModelRequest) SetType(v AgenticWorkflowModelType)`

SetType sets Type field to given value.


### GetApiKey

`func (o *AgenticWorkflowModelRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *AgenticWorkflowModelRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *AgenticWorkflowModelRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *AgenticWorkflowModelRequest) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetSettings

`func (o *AgenticWorkflowModelRequest) GetSettings() string`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *AgenticWorkflowModelRequest) GetSettingsOk() (*string, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *AgenticWorkflowModelRequest) SetSettings(v string)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *AgenticWorkflowModelRequest) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


