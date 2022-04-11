# EnvironmentApplicationsStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | **string** |  | 
**Disks** | Pointer to [**[]StorageDisk**](StorageDisk.md) |  | [optional] 

## Methods

### NewEnvironmentApplicationsStorage

`func NewEnvironmentApplicationsStorage(application string, ) *EnvironmentApplicationsStorage`

NewEnvironmentApplicationsStorage instantiates a new EnvironmentApplicationsStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentApplicationsStorageWithDefaults

`func NewEnvironmentApplicationsStorageWithDefaults() *EnvironmentApplicationsStorage`

NewEnvironmentApplicationsStorageWithDefaults instantiates a new EnvironmentApplicationsStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *EnvironmentApplicationsStorage) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *EnvironmentApplicationsStorage) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *EnvironmentApplicationsStorage) SetApplication(v string)`

SetApplication sets Application field to given value.


### GetDisks

`func (o *EnvironmentApplicationsStorage) GetDisks() []StorageDisk`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *EnvironmentApplicationsStorage) GetDisksOk() (*[]StorageDisk, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *EnvironmentApplicationsStorage) SetDisks(v []StorageDisk)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *EnvironmentApplicationsStorage) HasDisks() bool`

HasDisks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


