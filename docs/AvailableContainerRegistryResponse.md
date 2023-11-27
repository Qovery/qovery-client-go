# AvailableContainerRegistryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | [**ContainerRegistryKindEnum**](ContainerRegistryKindEnum.md) |  | 
**RequiredConfig** | **map[string]interface{}** |  | 
**IsMandatory** | **bool** |  | 

## Methods

### NewAvailableContainerRegistryResponse

`func NewAvailableContainerRegistryResponse(kind ContainerRegistryKindEnum, requiredConfig map[string]interface{}, isMandatory bool, ) *AvailableContainerRegistryResponse`

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

`func (o *AvailableContainerRegistryResponse) GetKind() ContainerRegistryKindEnum`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AvailableContainerRegistryResponse) GetKindOk() (*ContainerRegistryKindEnum, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AvailableContainerRegistryResponse) SetKind(v ContainerRegistryKindEnum)`

SetKind sets Kind field to given value.


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


### GetIsMandatory

`func (o *AvailableContainerRegistryResponse) GetIsMandatory() bool`

GetIsMandatory returns the IsMandatory field if non-nil, zero value otherwise.

### GetIsMandatoryOk

`func (o *AvailableContainerRegistryResponse) GetIsMandatoryOk() (*bool, bool)`

GetIsMandatoryOk returns a tuple with the IsMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMandatory

`func (o *AvailableContainerRegistryResponse) SetIsMandatory(v bool)`

SetIsMandatory sets IsMandatory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


