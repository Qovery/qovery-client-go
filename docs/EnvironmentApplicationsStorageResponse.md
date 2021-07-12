# EnvironmentApplicationsStorageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | **string** |  | 
**Disks** | Pointer to [**[]StorageDiskResponse**](StorageDiskResponse.md) |  | [optional] 

## Methods

### NewEnvironmentApplicationsStorageResponse

`func NewEnvironmentApplicationsStorageResponse(application string, ) *EnvironmentApplicationsStorageResponse`

NewEnvironmentApplicationsStorageResponse instantiates a new EnvironmentApplicationsStorageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentApplicationsStorageResponseWithDefaults

`func NewEnvironmentApplicationsStorageResponseWithDefaults() *EnvironmentApplicationsStorageResponse`

NewEnvironmentApplicationsStorageResponseWithDefaults instantiates a new EnvironmentApplicationsStorageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *EnvironmentApplicationsStorageResponse) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *EnvironmentApplicationsStorageResponse) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *EnvironmentApplicationsStorageResponse) SetApplication(v string)`

SetApplication sets Application field to given value.


### GetDisks

`func (o *EnvironmentApplicationsStorageResponse) GetDisks() []StorageDiskResponse`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *EnvironmentApplicationsStorageResponse) GetDisksOk() (*[]StorageDiskResponse, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *EnvironmentApplicationsStorageResponse) SetDisks(v []StorageDiskResponse)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *EnvironmentApplicationsStorageResponse) HasDisks() bool`

HasDisks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


