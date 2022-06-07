# ApplicationStorageRequestStorageInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**StorageTypeEnum**](StorageTypeEnum.md) |  | 
**Size** | **int32** | unit is GB | 
**MountPoint** | **string** |  | 

## Methods

### NewApplicationStorageRequestStorageInner

`func NewApplicationStorageRequestStorageInner(type_ StorageTypeEnum, size int32, mountPoint string, ) *ApplicationStorageRequestStorageInner`

NewApplicationStorageRequestStorageInner instantiates a new ApplicationStorageRequestStorageInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationStorageRequestStorageInnerWithDefaults

`func NewApplicationStorageRequestStorageInnerWithDefaults() *ApplicationStorageRequestStorageInner`

NewApplicationStorageRequestStorageInnerWithDefaults instantiates a new ApplicationStorageRequestStorageInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationStorageRequestStorageInner) GetType() StorageTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationStorageRequestStorageInner) GetTypeOk() (*StorageTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationStorageRequestStorageInner) SetType(v StorageTypeEnum)`

SetType sets Type field to given value.


### GetSize

`func (o *ApplicationStorageRequestStorageInner) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ApplicationStorageRequestStorageInner) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ApplicationStorageRequestStorageInner) SetSize(v int32)`

SetSize sets Size field to given value.


### GetMountPoint

`func (o *ApplicationStorageRequestStorageInner) GetMountPoint() string`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *ApplicationStorageRequestStorageInner) GetMountPointOk() (*string, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *ApplicationStorageRequestStorageInner) SetMountPoint(v string)`

SetMountPoint sets MountPoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


