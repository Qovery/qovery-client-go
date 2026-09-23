# LlmProviderModelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Anthropic model id, or Bedrock inference profile id | 
**DisplayName** | **string** |  | 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewLlmProviderModelResponse

`func NewLlmProviderModelResponse(id string, displayName string, ) *LlmProviderModelResponse`

NewLlmProviderModelResponse instantiates a new LlmProviderModelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLlmProviderModelResponseWithDefaults

`func NewLlmProviderModelResponseWithDefaults() *LlmProviderModelResponse`

NewLlmProviderModelResponseWithDefaults instantiates a new LlmProviderModelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LlmProviderModelResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LlmProviderModelResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LlmProviderModelResponse) SetId(v string)`

SetId sets Id field to given value.


### GetDisplayName

`func (o *LlmProviderModelResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *LlmProviderModelResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *LlmProviderModelResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetCreatedAt

`func (o *LlmProviderModelResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LlmProviderModelResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LlmProviderModelResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LlmProviderModelResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *LlmProviderModelResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *LlmProviderModelResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


