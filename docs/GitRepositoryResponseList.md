# GitRepositoryResponseList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]GitRepository**](GitRepository.md) |  | [optional] 

## Methods

### NewGitRepositoryResponseList

`func NewGitRepositoryResponseList() *GitRepositoryResponseList`

NewGitRepositoryResponseList instantiates a new GitRepositoryResponseList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitRepositoryResponseListWithDefaults

`func NewGitRepositoryResponseListWithDefaults() *GitRepositoryResponseList`

NewGitRepositoryResponseListWithDefaults instantiates a new GitRepositoryResponseList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *GitRepositoryResponseList) GetResults() []GitRepository`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *GitRepositoryResponseList) GetResultsOk() (*[]GitRepository, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *GitRepositoryResponseList) SetResults(v []GitRepository)`

SetResults sets Results field to given value.

### HasResults

`func (o *GitRepositoryResponseList) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


