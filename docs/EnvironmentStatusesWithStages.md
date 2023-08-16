# EnvironmentStatusesWithStages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to [**EnvironmentStatus**](EnvironmentStatus.md) |  | [optional] 
**Stages** | Pointer to [**DeploymentStageWithServiceStatusesList**](DeploymentStageWithServiceStatusesList.md) |  | [optional] 

## Methods

### NewEnvironmentStatusesWithStages

`func NewEnvironmentStatusesWithStages() *EnvironmentStatusesWithStages`

NewEnvironmentStatusesWithStages instantiates a new EnvironmentStatusesWithStages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentStatusesWithStagesWithDefaults

`func NewEnvironmentStatusesWithStagesWithDefaults() *EnvironmentStatusesWithStages`

NewEnvironmentStatusesWithStagesWithDefaults instantiates a new EnvironmentStatusesWithStages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *EnvironmentStatusesWithStages) GetEnvironment() EnvironmentStatus`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *EnvironmentStatusesWithStages) GetEnvironmentOk() (*EnvironmentStatus, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *EnvironmentStatusesWithStages) SetEnvironment(v EnvironmentStatus)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *EnvironmentStatusesWithStages) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetStages

`func (o *EnvironmentStatusesWithStages) GetStages() DeploymentStageWithServiceStatusesList`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *EnvironmentStatusesWithStages) GetStagesOk() (*DeploymentStageWithServiceStatusesList, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *EnvironmentStatusesWithStages) SetStages(v DeploymentStageWithServiceStatusesList)`

SetStages sets Stages field to given value.

### HasStages

`func (o *EnvironmentStatusesWithStages) HasStages() bool`

HasStages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


