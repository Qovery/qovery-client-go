# DeploymentHistoryApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Commit** | Pointer to [**NullableCommit**](Commit.md) |  | [optional] 
**Status** | Pointer to [**StateEnum**](StateEnum.md) |  | [optional] 

## Methods

### NewDeploymentHistoryApplication

`func NewDeploymentHistoryApplication(id string, createdAt time.Time, ) *DeploymentHistoryApplication`

NewDeploymentHistoryApplication instantiates a new DeploymentHistoryApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentHistoryApplicationWithDefaults

`func NewDeploymentHistoryApplicationWithDefaults() *DeploymentHistoryApplication`

NewDeploymentHistoryApplicationWithDefaults instantiates a new DeploymentHistoryApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentHistoryApplication) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentHistoryApplication) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentHistoryApplication) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *DeploymentHistoryApplication) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeploymentHistoryApplication) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeploymentHistoryApplication) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DeploymentHistoryApplication) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeploymentHistoryApplication) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeploymentHistoryApplication) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DeploymentHistoryApplication) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *DeploymentHistoryApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentHistoryApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentHistoryApplication) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentHistoryApplication) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCommit

`func (o *DeploymentHistoryApplication) GetCommit() Commit`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *DeploymentHistoryApplication) GetCommitOk() (*Commit, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *DeploymentHistoryApplication) SetCommit(v Commit)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *DeploymentHistoryApplication) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### SetCommitNil

`func (o *DeploymentHistoryApplication) SetCommitNil(b bool)`

 SetCommitNil sets the value for Commit to be an explicit nil

### UnsetCommit
`func (o *DeploymentHistoryApplication) UnsetCommit()`

UnsetCommit ensures that no value is present for Commit, not even an explicit nil
### GetStatus

`func (o *DeploymentHistoryApplication) GetStatus() StateEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentHistoryApplication) GetStatusOk() (*StateEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentHistoryApplication) SetStatus(v StateEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeploymentHistoryApplication) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


