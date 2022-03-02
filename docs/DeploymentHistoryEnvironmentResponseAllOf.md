# DeploymentHistoryEnvironmentResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Applications** | Pointer to [**[]DeploymentHistoryApplicationResponse**](DeploymentHistoryApplicationResponse.md) |  | [optional] 
**Databases** | Pointer to [**[]DeploymentHistoryDatabaseResponse**](DeploymentHistoryDatabaseResponse.md) |  | [optional] 

## Methods

### NewDeploymentHistoryEnvironmentResponseAllOf

`func NewDeploymentHistoryEnvironmentResponseAllOf() *DeploymentHistoryEnvironmentResponseAllOf`

NewDeploymentHistoryEnvironmentResponseAllOf instantiates a new DeploymentHistoryEnvironmentResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentHistoryEnvironmentResponseAllOfWithDefaults

`func NewDeploymentHistoryEnvironmentResponseAllOfWithDefaults() *DeploymentHistoryEnvironmentResponseAllOf`

NewDeploymentHistoryEnvironmentResponseAllOfWithDefaults instantiates a new DeploymentHistoryEnvironmentResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DeploymentHistoryEnvironmentResponseAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentHistoryEnvironmentResponseAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentHistoryEnvironmentResponseAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeploymentHistoryEnvironmentResponseAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApplications

`func (o *DeploymentHistoryEnvironmentResponseAllOf) GetApplications() []DeploymentHistoryApplicationResponse`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *DeploymentHistoryEnvironmentResponseAllOf) GetApplicationsOk() (*[]DeploymentHistoryApplicationResponse, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *DeploymentHistoryEnvironmentResponseAllOf) SetApplications(v []DeploymentHistoryApplicationResponse)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *DeploymentHistoryEnvironmentResponseAllOf) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetDatabases

`func (o *DeploymentHistoryEnvironmentResponseAllOf) GetDatabases() []DeploymentHistoryDatabaseResponse`

GetDatabases returns the Databases field if non-nil, zero value otherwise.

### GetDatabasesOk

`func (o *DeploymentHistoryEnvironmentResponseAllOf) GetDatabasesOk() (*[]DeploymentHistoryDatabaseResponse, bool)`

GetDatabasesOk returns a tuple with the Databases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabases

`func (o *DeploymentHistoryEnvironmentResponseAllOf) SetDatabases(v []DeploymentHistoryDatabaseResponse)`

SetDatabases sets Databases field to given value.

### HasDatabases

`func (o *DeploymentHistoryEnvironmentResponseAllOf) HasDatabases() bool`

HasDatabases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


