# DeploymentHistoryApplicationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Commit** | Pointer to [**Commit**](Commit.md) |  | [optional] 
**Status** | Pointer to [**GlobalDeploymentStatus**](GlobalDeploymentStatus.md) |  | [optional] 

## Methods

### NewDeploymentHistoryApplicationAllOf

`func NewDeploymentHistoryApplicationAllOf() *DeploymentHistoryApplicationAllOf`

NewDeploymentHistoryApplicationAllOf instantiates a new DeploymentHistoryApplicationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentHistoryApplicationAllOfWithDefaults

`func NewDeploymentHistoryApplicationAllOfWithDefaults() *DeploymentHistoryApplicationAllOf`

NewDeploymentHistoryApplicationAllOfWithDefaults instantiates a new DeploymentHistoryApplicationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DeploymentHistoryApplicationAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentHistoryApplicationAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentHistoryApplicationAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentHistoryApplicationAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCommit

`func (o *DeploymentHistoryApplicationAllOf) GetCommit() Commit`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *DeploymentHistoryApplicationAllOf) GetCommitOk() (*Commit, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *DeploymentHistoryApplicationAllOf) SetCommit(v Commit)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *DeploymentHistoryApplicationAllOf) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### GetStatus

`func (o *DeploymentHistoryApplicationAllOf) GetStatus() GlobalDeploymentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentHistoryApplicationAllOf) GetStatusOk() (*GlobalDeploymentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentHistoryApplicationAllOf) SetStatus(v GlobalDeploymentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeploymentHistoryApplicationAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


