# DeploymentHistoryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commit** | Pointer to [**Commit**](Commit.md) |  | [optional] 
**Status** | Pointer to [**DeploymentHistoryStatusEnum**](DeploymentHistoryStatusEnum.md) |  | [optional] 

## Methods

### NewDeploymentHistoryAllOf

`func NewDeploymentHistoryAllOf() *DeploymentHistoryAllOf`

NewDeploymentHistoryAllOf instantiates a new DeploymentHistoryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentHistoryAllOfWithDefaults

`func NewDeploymentHistoryAllOfWithDefaults() *DeploymentHistoryAllOf`

NewDeploymentHistoryAllOfWithDefaults instantiates a new DeploymentHistoryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommit

`func (o *DeploymentHistoryAllOf) GetCommit() Commit`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *DeploymentHistoryAllOf) GetCommitOk() (*Commit, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *DeploymentHistoryAllOf) SetCommit(v Commit)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *DeploymentHistoryAllOf) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetStatus

`func (o *DeploymentHistoryAllOf) GetStatus() DeploymentHistoryStatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentHistoryAllOf) GetStatusOk() (*DeploymentHistoryStatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentHistoryAllOf) SetStatus(v DeploymentHistoryStatusEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeploymentHistoryAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


