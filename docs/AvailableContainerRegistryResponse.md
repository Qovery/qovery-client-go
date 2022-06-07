# AvailableContainerRegistryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to [**ContainerRegistryKind**](ContainerRegistryKind.md) |  | [optional] [default to CONTAINERREGISTRYKIND_ECR]
**RequiredConfig** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewAvailableContainerRegistryResponse

`func NewAvailableContainerRegistryResponse() *AvailableContainerRegistryResponse`

NewAvailableContainerRegistryResponse instantiates a new AvailableContainerRegistryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailableContainerRegistryResponseWithDefaults

`func NewAvailableContainerRegistryResponseWithDefaults() *AvailableContainerRegistryResponse`

NewAvailableContainerRegistryResponseWithDefaults instantiates a new AvailableContainerRegistryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *AvailableContainerRegistryResponse) GetKind() ContainerRegistryKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AvailableContainerRegistryResponse) GetKindOk() (*ContainerRegistryKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AvailableContainerRegistryResponse) SetKind(v ContainerRegistryKind)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AvailableContainerRegistryResponse) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetRequiredConfig

`func (o *AvailableContainerRegistryResponse) GetRequiredConfig() map[string]interface{}`

GetRequiredConfig returns the RequiredConfig field if non-nil, zero value otherwise.

### GetRequiredConfigOk

`func (o *AvailableContainerRegistryResponse) GetRequiredConfigOk() (*map[string]interface{}, bool)`

GetRequiredConfigOk returns a tuple with the RequiredConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredConfig

`func (o *AvailableContainerRegistryResponse) SetRequiredConfig(v map[string]interface{})`

SetRequiredConfig sets RequiredConfig field to given value.

### HasRequiredConfig

`func (o *AvailableContainerRegistryResponse) HasRequiredConfig() bool`

HasRequiredConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


