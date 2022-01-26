# DeploymentHistoryEnvironmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Applications** | Pointer to [**Array**](array.md) |  | [optional] 
**Databases** | Pointer to [**Array**](array.md) |  | [optional] 
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewDeploymentHistoryEnvironmentResponse

`func NewDeploymentHistoryEnvironmentResponse(id string, createdAt time.Time, ) *DeploymentHistoryEnvironmentResponse`

NewDeploymentHistoryEnvironmentResponse instantiates a new DeploymentHistoryEnvironmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentHistoryEnvironmentResponseWithDefaults

`func NewDeploymentHistoryEnvironmentResponseWithDefaults() *DeploymentHistoryEnvironmentResponse`

NewDeploymentHistoryEnvironmentResponseWithDefaults instantiates a new DeploymentHistoryEnvironmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DeploymentHistoryEnvironmentResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentHistoryEnvironmentResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentHistoryEnvironmentResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeploymentHistoryEnvironmentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApplications

`func (o *DeploymentHistoryEnvironmentResponse) GetApplications() Array`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *DeploymentHistoryEnvironmentResponse) GetApplicationsOk() (*Array, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *DeploymentHistoryEnvironmentResponse) SetApplications(v Array)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *DeploymentHistoryEnvironmentResponse) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetDatabases

`func (o *DeploymentHistoryEnvironmentResponse) GetDatabases() Array`

GetDatabases returns the Databases field if non-nil, zero value otherwise.

### GetDatabasesOk

`func (o *DeploymentHistoryEnvironmentResponse) GetDatabasesOk() (*Array, bool)`

GetDatabasesOk returns a tuple with the Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabases

`func (o *DeploymentHistoryEnvironmentResponse) SetDatabases(v Array)`

SetDatabases sets Databases field to given value.

### HasDatabases

`func (o *DeploymentHistoryEnvironmentResponse) HasDatabases() bool`

HasDatabases returns a boolean if a field has been set.

### GetId

`func (o *DeploymentHistoryEnvironmentResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentHistoryEnvironmentResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentHistoryEnvironmentResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *DeploymentHistoryEnvironmentResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeploymentHistoryEnvironmentResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeploymentHistoryEnvironmentResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DeploymentHistoryEnvironmentResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeploymentHistoryEnvironmentResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeploymentHistoryEnvironmentResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DeploymentHistoryEnvironmentResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


