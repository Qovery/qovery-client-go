# ApplicationStorageStorageInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | [**StorageTypeEnum**](StorageTypeEnum.md) |  | 
**Size** | **int32** | unit is GB | 
**MountPoint** | **string** |  | 

## Methods

### NewApplicationStorageStorageInner

`func NewApplicationStorageStorageInner(id string, type_ StorageTypeEnum, size int32, mountPoint string, ) *ApplicationStorageStorageInner`

NewApplicationStorageStorageInner instantiates a new ApplicationStorageStorageInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationStorageStorageInnerWithDefaults

`func NewApplicationStorageStorageInnerWithDefaults() *ApplicationStorageStorageInner`

NewApplicationStorageStorageInnerWithDefaults instantiates a new ApplicationStorageStorageInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationStorageStorageInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationStorageStorageInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationStorageStorageInner) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ApplicationStorageStorageInner) GetType() StorageTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationStorageStorageInner) GetTypeOk() (*StorageTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationStorageStorageInner) SetType(v StorageTypeEnum)`

SetType sets Type field to given value.


### GetSize

`func (o *ApplicationStorageStorageInner) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ApplicationStorageStorageInner) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ApplicationStorageStorageInner) SetSize(v int32)`

SetSize sets Size field to given value.


### GetMountPoint

`func (o *ApplicationStorageStorageInner) GetMountPoint() string`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *ApplicationStorageStorageInner) GetMountPointOk() (*string, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *ApplicationStorageStorageInner) SetMountPoint(v string)`

SetMountPoint sets MountPoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


