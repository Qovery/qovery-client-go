# EnvironmentContainersStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Container** | **string** |  | 
**Disks** | Pointer to [**[]StorageDisk**](StorageDisk.md) |  | [optional] 

## Methods

### NewEnvironmentContainersStorage

`func NewEnvironmentContainersStorage(container string, ) *EnvironmentContainersStorage`

NewEnvironmentContainersStorage instantiates a new EnvironmentContainersStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentContainersStorageWithDefaults

`func NewEnvironmentContainersStorageWithDefaults() *EnvironmentContainersStorage`

NewEnvironmentContainersStorageWithDefaults instantiates a new EnvironmentContainersStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainer

`func (o *EnvironmentContainersStorage) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *EnvironmentContainersStorage) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *EnvironmentContainersStorage) SetContainer(v string)`

SetContainer sets Container field to given value.


### GetDisks

`func (o *EnvironmentContainersStorage) GetDisks() []StorageDisk`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *EnvironmentContainersStorage) GetDisksOk() (*[]StorageDisk, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *EnvironmentContainersStorage) SetDisks(v []StorageDisk)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *EnvironmentContainersStorage) HasDisks() bool`

HasDisks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


