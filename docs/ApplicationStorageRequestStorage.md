# ApplicationStorageRequestStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**StorageTypeEnum**](StorageTypeEnum.md) |  | 
**Size** | **int32** | unit is GB | 
**MountPoint** | **string** |  | 

## Methods

### NewApplicationStorageRequestStorage

`func NewApplicationStorageRequestStorage(type_ StorageTypeEnum, size int32, mountPoint string, ) *ApplicationStorageRequestStorage`

NewApplicationStorageRequestStorage instantiates a new ApplicationStorageRequestStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationStorageRequestStorageWithDefaults

`func NewApplicationStorageRequestStorageWithDefaults() *ApplicationStorageRequestStorage`

NewApplicationStorageRequestStorageWithDefaults instantiates a new ApplicationStorageRequestStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationStorageRequestStorage) GetType() StorageTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationStorageRequestStorage) GetTypeOk() (*StorageTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationStorageRequestStorage) SetType(v StorageTypeEnum)`

SetType sets Type field to given value.


### GetSize

`func (o *ApplicationStorageRequestStorage) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ApplicationStorageRequestStorage) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ApplicationStorageRequestStorage) SetSize(v int32)`

SetSize sets Size field to given value.


### GetMountPoint

`func (o *ApplicationStorageRequestStorage) GetMountPoint() string`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *ApplicationStorageRequestStorage) GetMountPointOk() (*string, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *ApplicationStorageRequestStorage) SetMountPoint(v string)`

SetMountPoint sets MountPoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


