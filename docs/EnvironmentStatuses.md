# EnvironmentStatuses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to [**EnvironmentStatus**](EnvironmentStatus.md) |  | [optional] 
**Applications** | Pointer to [**[]Status**](Status.md) |  | [optional] 
**Containers** | Pointer to [**[]Status**](Status.md) |  | [optional] 
**Jobs** | Pointer to [**[]Status**](Status.md) |  | [optional] 
**Databases** | Pointer to [**[]Status**](Status.md) |  | [optional] 
**Helms** | Pointer to [**[]Status**](Status.md) |  | [optional] 

## Methods

### NewEnvironmentStatuses

`func NewEnvironmentStatuses() *EnvironmentStatuses`

NewEnvironmentStatuses instantiates a new EnvironmentStatuses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentStatusesWithDefaults

`func NewEnvironmentStatusesWithDefaults() *EnvironmentStatuses`

NewEnvironmentStatusesWithDefaults instantiates a new EnvironmentStatuses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *EnvironmentStatuses) GetEnvironment() EnvironmentStatus`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *EnvironmentStatuses) GetEnvironmentOk() (*EnvironmentStatus, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *EnvironmentStatuses) SetEnvironment(v EnvironmentStatus)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *EnvironmentStatuses) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetApplications

`func (o *EnvironmentStatuses) GetApplications() []Status`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *EnvironmentStatuses) GetApplicationsOk() (*[]Status, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *EnvironmentStatuses) SetApplications(v []Status)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *EnvironmentStatuses) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetContainers

`func (o *EnvironmentStatuses) GetContainers() []Status`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *EnvironmentStatuses) GetContainersOk() (*[]Status, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *EnvironmentStatuses) SetContainers(v []Status)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *EnvironmentStatuses) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetJobs

`func (o *EnvironmentStatuses) GetJobs() []Status`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *EnvironmentStatuses) GetJobsOk() (*[]Status, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *EnvironmentStatuses) SetJobs(v []Status)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *EnvironmentStatuses) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetDatabases

`func (o *EnvironmentStatuses) GetDatabases() []Status`

GetDatabases returns the Databases field if non-nil, zero value otherwise.

### GetDatabasesOk

`func (o *EnvironmentStatuses) GetDatabasesOk() (*[]Status, bool)`

GetDatabasesOk returns a tuple with the Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabases

`func (o *EnvironmentStatuses) SetDatabases(v []Status)`

SetDatabases sets Databases field to given value.

### HasDatabases

`func (o *EnvironmentStatuses) HasDatabases() bool`

HasDatabases returns a boolean if a field has been set.

### GetHelms

`func (o *EnvironmentStatuses) GetHelms() []Status`

GetHelms returns the Helms field if non-nil, zero value otherwise.

### GetHelmsOk

`func (o *EnvironmentStatuses) GetHelmsOk() (*[]Status, bool)`

GetHelmsOk returns a tuple with the Helms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelms

`func (o *EnvironmentStatuses) SetHelms(v []Status)`

SetHelms sets Helms field to given value.

### HasHelms

`func (o *EnvironmentStatuses) HasHelms() bool`

HasHelms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


