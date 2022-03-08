# EnvironmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | **string** | name is case insensitive | 
**Project** | Pointer to [**ReferenceObject**](ReferenceObject.md) |  | [optional] 
**LastUpdatedBy** | Pointer to **string** | uuid of the user that made the last update | [optional] 
**CloudProvider** | [**EnvironmentResponseAllOfCloudProvider**](EnvironmentResponseAllOfCloudProvider.md) |  | 
**Mode** | [**EnvironmentModeEnum**](EnvironmentModeEnum.md) |  | 
**ClusterId** | **string** |  | 

## Methods

### NewEnvironmentResponse

`func NewEnvironmentResponse(id string, createdAt time.Time, name string, cloudProvider EnvironmentResponseAllOfCloudProvider, mode EnvironmentModeEnum, clusterId string, ) *EnvironmentResponse`

NewEnvironmentResponse instantiates a new EnvironmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentResponseWithDefaults

`func NewEnvironmentResponseWithDefaults() *EnvironmentResponse`

NewEnvironmentResponseWithDefaults instantiates a new EnvironmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *EnvironmentResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnvironmentResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnvironmentResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EnvironmentResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnvironmentResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnvironmentResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EnvironmentResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *EnvironmentResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentResponse) SetName(v string)`

SetName sets Name field to given value.


### GetProject

`func (o *EnvironmentResponse) GetProject() ReferenceObject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *EnvironmentResponse) GetProjectOk() (*ReferenceObject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *EnvironmentResponse) SetProject(v ReferenceObject)`

SetProject sets Project field to given value.

### HasProject

`func (o *EnvironmentResponse) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *EnvironmentResponse) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *EnvironmentResponse) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *EnvironmentResponse) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *EnvironmentResponse) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.

### GetCloudProvider

`func (o *EnvironmentResponse) GetCloudProvider() EnvironmentResponseAllOfCloudProvider`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *EnvironmentResponse) GetCloudProviderOk() (*EnvironmentResponseAllOfCloudProvider, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *EnvironmentResponse) SetCloudProvider(v EnvironmentResponseAllOfCloudProvider)`

SetCloudProvider sets CloudProvider field to given value.


### GetMode

`func (o *EnvironmentResponse) GetMode() EnvironmentModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *EnvironmentResponse) GetModeOk() (*EnvironmentModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *EnvironmentResponse) SetMode(v EnvironmentModeEnum)`

SetMode sets Mode field to given value.


### GetClusterId

`func (o *EnvironmentResponse) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *EnvironmentResponse) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *EnvironmentResponse) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


