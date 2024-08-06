# ContainerRegistryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to [**ContainerRegistryKindEnum**](ContainerRegistryKindEnum.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** | URL of the container registry | [optional] 
**Cluster** | Pointer to [**ContainerRegistryResponseAllOfCluster**](ContainerRegistryResponseAllOfCluster.md) |  | [optional] 
**AssociatedServicesCount** | Pointer to **int32** | The number of services using this container registry | [optional] 
**Config** | Pointer to [**ContainerRegistryResponseAllOfConfig**](ContainerRegistryResponseAllOfConfig.md) |  | [optional] 

## Methods

### NewContainerRegistryResponse

`func NewContainerRegistryResponse(id string, createdAt time.Time, ) *ContainerRegistryResponse`

NewContainerRegistryResponse instantiates a new ContainerRegistryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerRegistryResponseWithDefaults

`func NewContainerRegistryResponseWithDefaults() *ContainerRegistryResponse`

NewContainerRegistryResponseWithDefaults instantiates a new ContainerRegistryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContainerRegistryResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContainerRegistryResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContainerRegistryResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ContainerRegistryResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ContainerRegistryResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ContainerRegistryResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ContainerRegistryResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ContainerRegistryResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ContainerRegistryResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ContainerRegistryResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

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

`func (o *ContainerRegistryResponse) GetKind() ContainerRegistryKindEnum`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ContainerRegistryResponse) GetKindOk() (*ContainerRegistryKindEnum, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ContainerRegistryResponse) SetKind(v ContainerRegistryKindEnum)`

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

### GetCluster

`func (o *ContainerRegistryResponse) GetCluster() ContainerRegistryResponseAllOfCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *ContainerRegistryResponse) GetClusterOk() (*ContainerRegistryResponseAllOfCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *ContainerRegistryResponse) SetCluster(v ContainerRegistryResponseAllOfCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *ContainerRegistryResponse) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetAssociatedServicesCount

`func (o *ContainerRegistryResponse) GetAssociatedServicesCount() int32`

GetAssociatedServicesCount returns the AssociatedServicesCount field if non-nil, zero value otherwise.

### GetAssociatedServicesCountOk

`func (o *ContainerRegistryResponse) GetAssociatedServicesCountOk() (*int32, bool)`

GetAssociatedServicesCountOk returns a tuple with the AssociatedServicesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedServicesCount

`func (o *ContainerRegistryResponse) SetAssociatedServicesCount(v int32)`

SetAssociatedServicesCount sets AssociatedServicesCount field to given value.

### HasAssociatedServicesCount

`func (o *ContainerRegistryResponse) HasAssociatedServicesCount() bool`

HasAssociatedServicesCount returns a boolean if a field has been set.

### GetConfig

`func (o *ContainerRegistryResponse) GetConfig() ContainerRegistryResponseAllOfConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ContainerRegistryResponse) GetConfigOk() (*ContainerRegistryResponseAllOfConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ContainerRegistryResponse) SetConfig(v ContainerRegistryResponseAllOfConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ContainerRegistryResponse) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


