# ContainerStorageStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | [**StorageTypeEnum**](StorageTypeEnum.md) |  | 
**Size** | **int32** | unit is GB | 
**MountPoint** | **string** |  | 

## Methods

### NewContainerStorageStorage

`func NewContainerStorageStorage(type_ StorageTypeEnum, size int32, mountPoint string, ) *ContainerStorageStorage`

NewContainerStorageStorage instantiates a new ContainerStorageStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerStorageStorageWithDefaults

`func NewContainerStorageStorageWithDefaults() *ContainerStorageStorage`

NewContainerStorageStorageWithDefaults instantiates a new ContainerStorageStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContainerStorageStorage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContainerStorageStorage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContainerStorageStorage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ContainerStorageStorage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ContainerStorageStorage) GetType() StorageTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ContainerStorageStorage) GetTypeOk() (*StorageTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ContainerStorageStorage) SetType(v StorageTypeEnum)`

SetType sets Type field to given value.


### GetSize

`func (o *ContainerStorageStorage) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ContainerStorageStorage) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ContainerStorageStorage) SetSize(v int32)`

SetSize sets Size field to given value.


### GetMountPoint

`func (o *ContainerStorageStorage) GetMountPoint() string`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *ContainerStorageStorage) GetMountPointOk() (*string, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *ContainerStorageStorage) SetMountPoint(v string)`

SetMountPoint sets MountPoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


