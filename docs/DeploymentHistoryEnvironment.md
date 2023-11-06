# DeploymentHistoryEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Status** | Pointer to [**StateEnum**](StateEnum.md) |  | [optional] 
**Origin** | Pointer to [**OrganizationEventOrigin**](OrganizationEventOrigin.md) |  | [optional] 
**TriggeredBy** | Pointer to **string** |  | [optional] 
**Applications** | Pointer to [**[]DeploymentHistoryApplication**](DeploymentHistoryApplication.md) |  | [optional] 
**Containers** | Pointer to [**[]DeploymentHistoryContainer**](DeploymentHistoryContainer.md) |  | [optional] 
**Databases** | Pointer to [**[]DeploymentHistoryDatabase**](DeploymentHistoryDatabase.md) |  | [optional] 
**Jobs** | Pointer to [**[]DeploymentHistoryJobResponse**](DeploymentHistoryJobResponse.md) |  | [optional] 

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

`func (o *DeploymentHistoryEnvironment) GetStatus() StateEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentHistoryEnvironment) GetStatusOk() (*StateEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentHistoryEnvironment) SetStatus(v StateEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeploymentHistoryEnvironment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetOrigin

`func (o *DeploymentHistoryEnvironment) GetOrigin() OrganizationEventOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *DeploymentHistoryEnvironment) GetOriginOk() (*OrganizationEventOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *DeploymentHistoryEnvironment) SetOrigin(v OrganizationEventOrigin)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *DeploymentHistoryEnvironment) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetTriggeredBy

`func (o *DeploymentHistoryEnvironment) GetTriggeredBy() string`

GetTriggeredBy returns the TriggeredBy field if non-nil, zero value otherwise.

### GetTriggeredByOk

`func (o *DeploymentHistoryEnvironment) GetTriggeredByOk() (*string, bool)`

GetTriggeredByOk returns a tuple with the TriggeredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredBy

`func (o *DeploymentHistoryEnvironment) SetTriggeredBy(v string)`

SetTriggeredBy sets TriggeredBy field to given value.

### HasTriggeredBy

`func (o *DeploymentHistoryEnvironment) HasTriggeredBy() bool`

HasTriggeredBy returns a boolean if a field has been set.

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

### GetContainers

`func (o *DeploymentHistoryEnvironment) GetContainers() []DeploymentHistoryContainer`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *DeploymentHistoryEnvironment) GetContainersOk() (*[]DeploymentHistoryContainer, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *DeploymentHistoryEnvironment) SetContainers(v []DeploymentHistoryContainer)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *DeploymentHistoryEnvironment) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

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

### GetJobs

`func (o *DeploymentHistoryEnvironment) GetJobs() []DeploymentHistoryJobResponse`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *DeploymentHistoryEnvironment) GetJobsOk() (*[]DeploymentHistoryJobResponse, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *DeploymentHistoryEnvironment) SetJobs(v []DeploymentHistoryJobResponse)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *DeploymentHistoryEnvironment) HasJobs() bool`

HasJobs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


