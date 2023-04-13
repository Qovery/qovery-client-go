# GetEnvironmentStatuses200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to [**EnvironmentStatus**](EnvironmentStatus.md) |  | [optional] 
**Applications** | Pointer to [**[]Status**](Status.md) |  | [optional] 
**Containers** | Pointer to [**[]Status**](Status.md) |  | [optional] 
**Jobs** | Pointer to [**[]Status**](Status.md) |  | [optional] 
**Databases** | Pointer to [**[]Status**](Status.md) |  | [optional] 

## Methods

### NewGetEnvironmentStatuses200Response

`func NewGetEnvironmentStatuses200Response() *GetEnvironmentStatuses200Response`

NewGetEnvironmentStatuses200Response instantiates a new GetEnvironmentStatuses200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEnvironmentStatuses200ResponseWithDefaults

`func NewGetEnvironmentStatuses200ResponseWithDefaults() *GetEnvironmentStatuses200Response`

NewGetEnvironmentStatuses200ResponseWithDefaults instantiates a new GetEnvironmentStatuses200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *GetEnvironmentStatuses200Response) GetEnvironment() EnvironmentStatus`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *GetEnvironmentStatuses200Response) GetEnvironmentOk() (*EnvironmentStatus, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *GetEnvironmentStatuses200Response) SetEnvironment(v EnvironmentStatus)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *GetEnvironmentStatuses200Response) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetApplications

`func (o *GetEnvironmentStatuses200Response) GetApplications() []Status`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *GetEnvironmentStatuses200Response) GetApplicationsOk() (*[]Status, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *GetEnvironmentStatuses200Response) SetApplications(v []Status)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *GetEnvironmentStatuses200Response) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetContainers

`func (o *GetEnvironmentStatuses200Response) GetContainers() []Status`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *GetEnvironmentStatuses200Response) GetContainersOk() (*[]Status, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *GetEnvironmentStatuses200Response) SetContainers(v []Status)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *GetEnvironmentStatuses200Response) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetJobs

`func (o *GetEnvironmentStatuses200Response) GetJobs() []Status`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *GetEnvironmentStatuses200Response) GetJobsOk() (*[]Status, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *GetEnvironmentStatuses200Response) SetJobs(v []Status)`

SetJobs sets Jobs field to given value.

### HasJobs

`func (o *GetEnvironmentStatuses200Response) HasJobs() bool`

HasJobs returns a boolean if a field has been set.

### GetDatabases

`func (o *GetEnvironmentStatuses200Response) GetDatabases() []Status`

GetDatabases returns the Databases field if non-nil, zero value otherwise.

### GetDatabasesOk

`func (o *GetEnvironmentStatuses200Response) GetDatabasesOk() (*[]Status, bool)`

GetDatabasesOk returns a tuple with the Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabases

`func (o *GetEnvironmentStatuses200Response) SetDatabases(v []Status)`

SetDatabases sets Databases field to given value.

### HasDatabases

`func (o *GetEnvironmentStatuses200Response) HasDatabases() bool`

HasDatabases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


