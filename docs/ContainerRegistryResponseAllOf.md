# ContainerRegistryResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to [**ContainerRegistryKindEnum**](ContainerRegistryKindEnum.md) |  | [optional] [default to CONTAINERREGISTRYKINDENUM_ECR]
**Description** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** | URL of the container registry | [optional] 

## Methods

### NewContainerRegistryResponseAllOf

`func NewContainerRegistryResponseAllOf() *ContainerRegistryResponseAllOf`

NewContainerRegistryResponseAllOf instantiates a new ContainerRegistryResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerRegistryResponseAllOfWithDefaults

`func NewContainerRegistryResponseAllOfWithDefaults() *ContainerRegistryResponseAllOf`

NewContainerRegistryResponseAllOfWithDefaults instantiates a new ContainerRegistryResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContainerRegistryResponseAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerRegistryResponseAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerRegistryResponseAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContainerRegistryResponseAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKind

`func (o *ContainerRegistryResponseAllOf) GetKind() ContainerRegistryKindEnum`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ContainerRegistryResponseAllOf) GetKindOk() (*ContainerRegistryKindEnum, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ContainerRegistryResponseAllOf) SetKind(v ContainerRegistryKindEnum)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ContainerRegistryResponseAllOf) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetDescription

`func (o *ContainerRegistryResponseAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContainerRegistryResponseAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContainerRegistryResponseAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContainerRegistryResponseAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *ContainerRegistryResponseAllOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ContainerRegistryResponseAllOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ContainerRegistryResponseAllOf) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ContainerRegistryResponseAllOf) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


