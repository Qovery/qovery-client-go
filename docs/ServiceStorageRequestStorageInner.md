# ServiceStorageRequestStorageInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**StorageTypeEnum**](StorageTypeEnum.md) |  | 
**Size** | **int32** | unit is GB | 
**MountPoint** | **string** |  | 

## Methods

### NewServiceStorageRequestStorageInner

`func NewServiceStorageRequestStorageInner(type_ StorageTypeEnum, size int32, mountPoint string, ) *ServiceStorageRequestStorageInner`

NewServiceStorageRequestStorageInner instantiates a new ServiceStorageRequestStorageInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceStorageRequestStorageInnerWithDefaults

`func NewServiceStorageRequestStorageInnerWithDefaults() *ServiceStorageRequestStorageInner`

NewServiceStorageRequestStorageInnerWithDefaults instantiates a new ServiceStorageRequestStorageInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ServiceStorageRequestStorageInner) GetType() StorageTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServiceStorageRequestStorageInner) GetTypeOk() (*StorageTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServiceStorageRequestStorageInner) SetType(v StorageTypeEnum)`

SetType sets Type field to given value.


### GetSize

`func (o *ServiceStorageRequestStorageInner) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ServiceStorageRequestStorageInner) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ServiceStorageRequestStorageInner) SetSize(v int32)`

SetSize sets Size field to given value.


### GetMountPoint

`func (o *ServiceStorageRequestStorageInner) GetMountPoint() string`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *ServiceStorageRequestStorageInner) GetMountPointOk() (*string, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *ServiceStorageRequestStorageInner) SetMountPoint(v string)`

SetMountPoint sets MountPoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


