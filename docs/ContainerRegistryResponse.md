# ContainerRegistryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to [**ContainerRegistryKind**](ContainerRegistryKind.md) |  | [optional] [default to CONTAINERREGISTRYKIND_ECR]
**Description** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** | URL of the container registry | [optional] 
**Config** | Pointer to **string** | authentification configuration | [optional] 

## Methods

### NewContainerRegistryResponse

`func NewContainerRegistryResponse() *ContainerRegistryResponse`

NewContainerRegistryResponse instantiates a new ContainerRegistryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerRegistryResponseWithDefaults

`func NewContainerRegistryResponseWithDefaults() *ContainerRegistryResponse`

NewContainerRegistryResponseWithDefaults instantiates a new ContainerRegistryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ContainerRegistryResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerRegistryResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerRegistryResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContainerRegistryResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKind

`func (o *ContainerRegistryResponse) GetKind() ContainerRegistryKind`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ContainerRegistryResponse) GetKindOk() (*ContainerRegistryKind, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ContainerRegistryResponse) SetKind(v ContainerRegistryKind)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ContainerRegistryResponse) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetDescription

`func (o *ContainerRegistryResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContainerRegistryResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContainerRegistryResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContainerRegistryResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *ContainerRegistryResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ContainerRegistryResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ContainerRegistryResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ContainerRegistryResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetConfig

`func (o *ContainerRegistryResponse) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ContainerRegistryResponse) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ContainerRegistryResponse) SetConfig(v string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ContainerRegistryResponse) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


