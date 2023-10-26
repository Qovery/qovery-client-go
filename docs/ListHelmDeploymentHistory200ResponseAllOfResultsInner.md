# ListHelmDeploymentHistory200ResponseAllOfResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Name** | Pointer to **string** | name of the helm | [optional] 
**Status** | Pointer to [**StateEnum**](StateEnum.md) |  | [optional] 
**Commit** | Pointer to [**NullableCommit**](Commit.md) |  | [optional] 
**Repository** | Pointer to [**NullableListHelmDeploymentHistory200ResponseAllOfResultsInnerAllOfRepository**](ListHelmDeploymentHistory200ResponseAllOfResultsInnerAllOfRepository.md) |  | [optional] 

## Methods

### NewListHelmDeploymentHistory200ResponseAllOfResultsInner

`func NewListHelmDeploymentHistory200ResponseAllOfResultsInner(id string, createdAt time.Time, ) *ListHelmDeploymentHistory200ResponseAllOfResultsInner`

NewListHelmDeploymentHistory200ResponseAllOfResultsInner instantiates a new ListHelmDeploymentHistory200ResponseAllOfResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHelmDeploymentHistory200ResponseAllOfResultsInnerWithDefaults

`func NewListHelmDeploymentHistory200ResponseAllOfResultsInnerWithDefaults() *ListHelmDeploymentHistory200ResponseAllOfResultsInner`

NewListHelmDeploymentHistory200ResponseAllOfResultsInnerWithDefaults instantiates a new ListHelmDeploymentHistory200ResponseAllOfResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListHelmDeploymentHistory200ResponseAllOfResultsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListHelmDeploymentHistory200ResponseAllOfResultsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListHelmDeploymentHistory200ResponseAllOfResultsInner) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ListHelmDeploymentHistory200ResponseAllOfResultsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListHelmDeploymentHistory200ResponseAllOfResultsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListHelmDeploymentHistory200ResponseAllOfResultsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ListHelmDeploymentHistory200ResponseAllOfResultsInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListHelmDeploymentHistory200ResponseAllOfResultsInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListHelmDeploymentHistory200ResponseAllOfResultsInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ListHelmDeploymentHistory200ResponseAllOfResultsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *ListHelmDeploymentHistory200ResponseAllOfResultsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListHelmDeploymentHistory200ResponseAllOfResultsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListHelmDeploymentHistory200ResponseAllOfResultsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListHelmDeploymentHistory200ResponseAllOfResultsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ListHelmDeploymentHistory200ResponseAllOfResultsInner) GetStatus() StateEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListHelmDeploymentHistory200ResponseAllOfResultsInner) GetStatusOk() (*StateEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListHelmDeploymentHistory200ResponseAllOfResultsInner) SetStatus(v StateEnum)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListHelmDeploymentHistory200ResponseAllOfResultsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCommit

`func (o *ListHelmDeploymentHistory200ResponseAllOfResultsInner) GetCommit() Commit`

GetCommit returns the Commit field if non-nil, zero value otherwise.

### GetCommitOk

`func (o *ListHelmDeploymentHistory200ResponseAllOfResultsInner) GetCommitOk() (*Commit, bool)`

GetCommitOk returns a tuple with the Commit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommit

`func (o *ListHelmDeploymentHistory200ResponseAllOfResultsInner) SetCommit(v Commit)`

SetCommit sets Commit field to given value.

### HasCommit

`func (o *ListHelmDeploymentHistory200ResponseAllOfResultsInner) HasCommit() bool`

HasCommit returns a boolean if a field has been set.

### SetCommitNil

`func (o *ListHelmDeploymentHistory200ResponseAllOfResultsInner) SetCommitNil(b bool)`

 SetCommitNil sets the value for Commit to be an explicit nil

### UnsetCommit
`func (o *ListHelmDeploymentHistory200ResponseAllOfResultsInner) UnsetCommit()`

UnsetCommit ensures that no value is present for Commit, not even an explicit nil
### GetRepository

`func (o *ListHelmDeploymentHistory200ResponseAllOfResultsInner) GetRepository() ListHelmDeploymentHistory200ResponseAllOfResultsInnerAllOfRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ListHelmDeploymentHistory200ResponseAllOfResultsInner) GetRepositoryOk() (*ListHelmDeploymentHistory200ResponseAllOfResultsInnerAllOfRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ListHelmDeploymentHistory200ResponseAllOfResultsInner) SetRepository(v ListHelmDeploymentHistory200ResponseAllOfResultsInnerAllOfRepository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *ListHelmDeploymentHistory200ResponseAllOfResultsInner) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *ListHelmDeploymentHistory200ResponseAllOfResultsInner) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *ListHelmDeploymentHistory200ResponseAllOfResultsInner) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

