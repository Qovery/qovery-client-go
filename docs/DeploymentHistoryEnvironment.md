# DeploymentHistoryEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Status** | Pointer to [**GlobalDeploymentStatus**](GlobalDeploymentStatus.md) |  | [optional] 
**Applications** | Pointer to [**[]DeploymentHistoryApplication**](DeploymentHistoryApplication.md) |  | [optional] 
**Databases** | Pointer to [**[]DeploymentHistoryDatabase**](DeploymentHistoryDatabase.md) |  | [optional] 

## Methods

### NewDeploymentHistoryEnvironment

`func NewDeploymentHistoryEnvironment(id string, createdAt time.Time, ) *DeploymentHistoryEnvironment`

NewDeploymentHistoryEnvironment instantiates a new DeploymentHistoryEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentHistoryEnvironmentWithDefaults

`func NewDeploymentHistoryEnvironmentWithDefaults() *DeploymentHistoryEnvironment`

NewDeploymentHistoryEnvironmentWithDefaults instantiates a new DeploymentHistoryEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentHistoryEnvironment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentHistoryEnvironment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentHistoryEnvironment) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *DeploymentHistoryEnvironment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeploymentHistoryEnvironment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeploymentHistoryEnvironment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DeploymentHistoryEnvironment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeploymentHistoryEnvironment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeploymentHistoryEnvironment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DeploymentHistoryEnvironment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *DeploymentHistoryEnvironment) GetStatus() GlobalDeploymentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentHistoryEnvironment) GetStatusOk() (*GlobalDeploymentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentHistoryEnvironment) SetStatus(v GlobalDeploymentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeploymentHistoryEnvironment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApplications

`func (o *DeploymentHistoryEnvironment) GetApplications() []DeploymentHistoryApplication`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *DeploymentHistoryEnvironment) GetApplicationsOk() (*[]DeploymentHistoryApplication, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *DeploymentHistoryEnvironment) SetApplications(v []DeploymentHistoryApplication)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *DeploymentHistoryEnvironment) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetDatabases

`func (o *DeploymentHistoryEnvironment) GetDatabases() []DeploymentHistoryDatabase`

GetDatabases returns the Databases field if non-nil, zero value otherwise.

### GetDatabasesOk

`func (o *DeploymentHistoryEnvironment) GetDatabasesOk() (*[]DeploymentHistoryDatabase, bool)`

GetDatabasesOk returns a tuple with the Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabases

`func (o *DeploymentHistoryEnvironment) SetDatabases(v []DeploymentHistoryDatabase)`

SetDatabases sets Databases field to given value.

### HasDatabases

`func (o *DeploymentHistoryEnvironment) HasDatabases() bool`

HasDatabases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


