# AvailableHelmRepositoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to [**HelmRepositoryKindEnum**](HelmRepositoryKindEnum.md) |  | [optional] 
**RequiredConfig** | Pointer to **map[string]interface{}** |  | [optional] 
**IsMandatory** | Pointer to **bool** |  | [optional] 

## Methods

### NewAvailableHelmRepositoryResponse

`func NewAvailableHelmRepositoryResponse() *AvailableHelmRepositoryResponse`

NewAvailableHelmRepositoryResponse instantiates a new AvailableHelmRepositoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailableHelmRepositoryResponseWithDefaults

`func NewAvailableHelmRepositoryResponseWithDefaults() *AvailableHelmRepositoryResponse`

NewAvailableHelmRepositoryResponseWithDefaults instantiates a new AvailableHelmRepositoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *AvailableHelmRepositoryResponse) GetKind() HelmRepositoryKindEnum`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AvailableHelmRepositoryResponse) GetKindOk() (*HelmRepositoryKindEnum, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AvailableHelmRepositoryResponse) SetKind(v HelmRepositoryKindEnum)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AvailableHelmRepositoryResponse) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetRequiredConfig

`func (o *AvailableHelmRepositoryResponse) GetRequiredConfig() map[string]interface{}`

GetRequiredConfig returns the RequiredConfig field if non-nil, zero value otherwise.

### GetRequiredConfigOk

`func (o *AvailableHelmRepositoryResponse) GetRequiredConfigOk() (*map[string]interface{}, bool)`

GetRequiredConfigOk returns a tuple with the RequiredConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredConfig

`func (o *AvailableHelmRepositoryResponse) SetRequiredConfig(v map[string]interface{})`

SetRequiredConfig sets RequiredConfig field to given value.

### HasRequiredConfig

`func (o *AvailableHelmRepositoryResponse) HasRequiredConfig() bool`

HasRequiredConfig returns a boolean if a field has been set.

### GetIsMandatory

`func (o *AvailableHelmRepositoryResponse) GetIsMandatory() bool`

GetIsMandatory returns the IsMandatory field if non-nil, zero value otherwise.

### GetIsMandatoryOk

`func (o *AvailableHelmRepositoryResponse) GetIsMandatoryOk() (*bool, bool)`

GetIsMandatoryOk returns a tuple with the IsMandatory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMandatory

`func (o *AvailableHelmRepositoryResponse) SetIsMandatory(v bool)`

SetIsMandatory sets IsMandatory field to given value.

### HasIsMandatory

`func (o *AvailableHelmRepositoryResponse) HasIsMandatory() bool`

HasIsMandatory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


