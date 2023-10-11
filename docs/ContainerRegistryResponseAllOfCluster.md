# ContainerRegistryResponseAllOfCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Id of the cluster of which the registry belongs to | 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the cluster of which the registry belongs to | [optional] 

## Methods

### NewContainerRegistryResponseAllOfCluster

`func NewContainerRegistryResponseAllOfCluster(id string, createdAt time.Time, ) *ContainerRegistryResponseAllOfCluster`

NewContainerRegistryResponseAllOfCluster instantiates a new ContainerRegistryResponseAllOfCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerRegistryResponseAllOfClusterWithDefaults

`func NewContainerRegistryResponseAllOfClusterWithDefaults() *ContainerRegistryResponseAllOfCluster`

NewContainerRegistryResponseAllOfClusterWithDefaults instantiates a new ContainerRegistryResponseAllOfCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContainerRegistryResponseAllOfCluster) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContainerRegistryResponseAllOfCluster) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContainerRegistryResponseAllOfCluster) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ContainerRegistryResponseAllOfCluster) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ContainerRegistryResponseAllOfCluster) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ContainerRegistryResponseAllOfCluster) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ContainerRegistryResponseAllOfCluster) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ContainerRegistryResponseAllOfCluster) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ContainerRegistryResponseAllOfCluster) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ContainerRegistryResponseAllOfCluster) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *ContainerRegistryResponseAllOfCluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerRegistryResponseAllOfCluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerRegistryResponseAllOfCluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContainerRegistryResponseAllOfCluster) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


