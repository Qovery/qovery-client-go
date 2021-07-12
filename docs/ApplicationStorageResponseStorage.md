# ApplicationStorageResponseStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**Size** | **float32** | unit is GB | 
**MountPoint** | **string** |  | 

## Methods

### NewApplicationStorageResponseStorage

`func NewApplicationStorageResponseStorage(type_ string, size float32, mountPoint string, ) *ApplicationStorageResponseStorage`

NewApplicationStorageResponseStorage instantiates a new ApplicationStorageResponseStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationStorageResponseStorageWithDefaults

`func NewApplicationStorageResponseStorageWithDefaults() *ApplicationStorageResponseStorage`

NewApplicationStorageResponseStorageWithDefaults instantiates a new ApplicationStorageResponseStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationStorageResponseStorage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationStorageResponseStorage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationStorageResponseStorage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationStorageResponseStorage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ApplicationStorageResponseStorage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationStorageResponseStorage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationStorageResponseStorage) SetType(v string)`

SetType sets Type field to given value.


### GetSize

`func (o *ApplicationStorageResponseStorage) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ApplicationStorageResponseStorage) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ApplicationStorageResponseStorage) SetSize(v float32)`

SetSize sets Size field to given value.


### GetMountPoint

`func (o *ApplicationStorageResponseStorage) GetMountPoint() string`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *ApplicationStorageResponseStorage) GetMountPointOk() (*string, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *ApplicationStorageResponseStorage) SetMountPoint(v string)`

SetMountPoint sets MountPoint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


