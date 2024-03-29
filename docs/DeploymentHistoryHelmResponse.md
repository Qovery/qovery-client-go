# DeploymentHistoryHelmResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** | name of the helm | [optional] 
**Status** | Pointer to [**StateEnum**](StateEnum.md) |  | [optional] 
**Commit** | Pointer to [**NullableCommit**](Commit.md) |  | [optional] 
**Repository** | Pointer to [**NullableDeploymentHistoryHelmResponseAllOfRepository**](DeploymentHistoryHelmResponseAllOfRepository.md) |  | [optional] 

## Methods

### NewDeploymentHistoryHelmResponse

`func NewDeploymentHistoryHelmResponse(id string, createdAt time.Time, ) *DeploymentHistoryHelmResponse`

NewDeploymentHistoryHelmResponse instantiates a new DeploymentHistoryHelmResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentHistoryHelmResponseWithDefaults

`func NewDeploymentHistoryHelmResponseWithDefaults() *DeploymentHistoryHelmResponse`

NewDeploymentHistoryHelmResponseWithDefaults instantiates a new DeploymentHistoryHelmResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentHistoryHelmResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentHistoryHelmResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentHistoryHelmResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *DeploymentHistoryHelmResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeploymentHistoryHelmResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeploymentHistoryHelmResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *DeploymentHistoryHelmResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DeploymentHistoryHelmResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DeploymentHistoryHelmResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DeploymentHistoryHelmResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *DeploymentHistoryHelmResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeploymentHistoryHelmResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeploymentHistoryHelmResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeploymentHistoryHelmResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *DeploymentHistoryHelmResponse) GetStatus() StateEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentHistoryHelmResponse) GetStatusOk() (*StateEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentHistoryHelmResponse) SetStatus(v StateEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeploymentHistoryHelmResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCommit

`func (o *DeploymentHistoryHelmResponse) GetCommit() Commit`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *DeploymentHistoryHelmResponse) GetCommitOk() (*Commit, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *DeploymentHistoryHelmResponse) SetCommit(v Commit)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *DeploymentHistoryHelmResponse) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### SetCommitNil

`func (o *DeploymentHistoryHelmResponse) SetCommitNil(b bool)`

 SetCommitNil sets the value for Commit to be an explicit nil

### UnsetCommit
`func (o *DeploymentHistoryHelmResponse) UnsetCommit()`

UnsetCommit ensures that no value is present for Commit, not even an explicit nil
### GetRepository

`func (o *DeploymentHistoryHelmResponse) GetRepository() DeploymentHistoryHelmResponseAllOfRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *DeploymentHistoryHelmResponse) GetRepositoryOk() (*DeploymentHistoryHelmResponseAllOfRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *DeploymentHistoryHelmResponse) SetRepository(v DeploymentHistoryHelmResponseAllOfRepository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *DeploymentHistoryHelmResponse) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *DeploymentHistoryHelmResponse) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *DeploymentHistoryHelmResponse) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


