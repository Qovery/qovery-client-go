# DeploymentStageWithServicesStatuses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to [**[]Status**](Status.md) |  | [optional] 
**Containers** | Pointer to [**[]Status**](Status.md) |  | [optional] 
**Jobs** | Pointer to [**[]Status**](Status.md) |  | [optional] 
**Databases** | Pointer to [**[]Status**](Status.md) |  | [optional] 
**Stage** | Pointer to [**Stage**](Stage.md) |  | [optional] 

## Methods

### NewDeploymentStageWithServicesStatuses

`func NewDeploymentStageWithServicesStatuses() *DeploymentStageWithServicesStatuses`

NewDeploymentStageWithServicesStatuses instantiates a new DeploymentStageWithServicesStatuses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentStageWithServicesStatusesWithDefaults

`func NewDeploymentStageWithServicesStatusesWithDefaults() *DeploymentStageWithServicesStatuses`

NewDeploymentStageWithServicesStatusesWithDefaults instantiates a new DeploymentStageWithServicesStatuses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *DeploymentStageWithServicesStatuses) GetApplications() []Status`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *DeploymentStageWithServicesStatuses) GetApplicationsOk() (*[]Status, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *DeploymentStageWithServicesStatuses) SetApplications(v []Status)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *DeploymentStageWithServicesStatuses) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetContainers

`func (o *DeploymentStageWithServicesStatuses) GetContainers() []Status`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *DeploymentStageWithServicesStatuses) GetContainersOk() (*[]Status, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *DeploymentStageWithServicesStatuses) SetContainers(v []Status)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *DeploymentStageWithServicesStatuses) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetJobs

`func (o *DeploymentStageWithServicesStatuses) GetJobs() []Status`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *DeploymentStageWithServicesStatuses) GetJobsOk() (*[]Status, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *DeploymentStageWithServicesStatuses) SetJobs(v []Status)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *DeploymentStageWithServicesStatuses) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetDatabases

`func (o *DeploymentStageWithServicesStatuses) GetDatabases() []Status`

GetDatabases returns the Databases field if non-nil, zero value otherwise.

### GetDatabasesOk

`func (o *DeploymentStageWithServicesStatuses) GetDatabasesOk() (*[]Status, bool)`

GetDatabasesOk returns a tuple with the Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabases

`func (o *DeploymentStageWithServicesStatuses) SetDatabases(v []Status)`

SetDatabases sets Databases field to given value.

### HasDatabases

`func (o *DeploymentStageWithServicesStatuses) HasDatabases() bool`

HasDatabases returns a boolean if a field has been set.

### GetStage

`func (o *DeploymentStageWithServicesStatuses) GetStage() Stage`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *DeploymentStageWithServicesStatuses) GetStageOk() (*Stage, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *DeploymentStageWithServicesStatuses) SetStage(v Stage)`

SetStage sets Stage field to given value.

### HasStage

`func (o *DeploymentStageWithServicesStatuses) HasStage() bool`

HasStage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


