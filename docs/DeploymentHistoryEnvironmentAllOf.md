# DeploymentHistoryEnvironmentAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**StateEnum**](StateEnum.md) |  | [optional] 
**Applications** | Pointer to [**[]DeploymentHistoryApplication**](DeploymentHistoryApplication.md) |  | [optional] 
**Containers** | Pointer to [**[]DeploymentHistoryContainer**](DeploymentHistoryContainer.md) |  | [optional] 
**Databases** | Pointer to [**[]DeploymentHistoryDatabase**](DeploymentHistoryDatabase.md) |  | [optional] 

## Methods

### NewDeploymentHistoryEnvironmentAllOf

`func NewDeploymentHistoryEnvironmentAllOf() *DeploymentHistoryEnvironmentAllOf`

NewDeploymentHistoryEnvironmentAllOf instantiates a new DeploymentHistoryEnvironmentAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentHistoryEnvironmentAllOfWithDefaults

`func NewDeploymentHistoryEnvironmentAllOfWithDefaults() *DeploymentHistoryEnvironmentAllOf`

NewDeploymentHistoryEnvironmentAllOfWithDefaults instantiates a new DeploymentHistoryEnvironmentAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DeploymentHistoryEnvironmentAllOf) GetStatus() StateEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentHistoryEnvironmentAllOf) GetStatusOk() (*StateEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentHistoryEnvironmentAllOf) SetStatus(v StateEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeploymentHistoryEnvironmentAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApplications

`func (o *DeploymentHistoryEnvironmentAllOf) GetApplications() []DeploymentHistoryApplication`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *DeploymentHistoryEnvironmentAllOf) GetApplicationsOk() (*[]DeploymentHistoryApplication, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *DeploymentHistoryEnvironmentAllOf) SetApplications(v []DeploymentHistoryApplication)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *DeploymentHistoryEnvironmentAllOf) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetContainers

`func (o *DeploymentHistoryEnvironmentAllOf) GetContainers() []DeploymentHistoryContainer`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *DeploymentHistoryEnvironmentAllOf) GetContainersOk() (*[]DeploymentHistoryContainer, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *DeploymentHistoryEnvironmentAllOf) SetContainers(v []DeploymentHistoryContainer)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *DeploymentHistoryEnvironmentAllOf) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetDatabases

`func (o *DeploymentHistoryEnvironmentAllOf) GetDatabases() []DeploymentHistoryDatabase`

GetDatabases returns the Databases field if non-nil, zero value otherwise.

### GetDatabasesOk

`func (o *DeploymentHistoryEnvironmentAllOf) GetDatabasesOk() (*[]DeploymentHistoryDatabase, bool)`

GetDatabasesOk returns a tuple with the Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabases

`func (o *DeploymentHistoryEnvironmentAllOf) SetDatabases(v []DeploymentHistoryDatabase)`

SetDatabases sets Databases field to given value.

### HasDatabases

`func (o *DeploymentHistoryEnvironmentAllOf) HasDatabases() bool`

HasDatabases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


