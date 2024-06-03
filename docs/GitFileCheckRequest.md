# GitFileCheckRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitRepository** | [**HelmGitRepositoryRequest**](HelmGitRepositoryRequest.md) |  | 
**Files** | **[]string** |  | 

## Methods

### NewGitFileCheckRequest

`func NewGitFileCheckRequest(gitRepository HelmGitRepositoryRequest, files []string, ) *GitFileCheckRequest`

NewGitFileCheckRequest instantiates a new GitFileCheckRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitFileCheckRequestWithDefaults

`func NewGitFileCheckRequestWithDefaults() *GitFileCheckRequest`

NewGitFileCheckRequestWithDefaults instantiates a new GitFileCheckRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitRepository

`func (o *GitFileCheckRequest) GetGitRepository() HelmGitRepositoryRequest`

GetGitRepository returns the GitRepository field if non-nil, zero value otherwise.

### GetGitRepositoryOk

`func (o *GitFileCheckRequest) GetGitRepositoryOk() (*HelmGitRepositoryRequest, bool)`

GetGitRepositoryOk returns a tuple with the GitRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepository

`func (o *GitFileCheckRequest) SetGitRepository(v HelmGitRepositoryRequest)`

SetGitRepository sets GitRepository field to given value.


### GetFiles

`func (o *GitFileCheckRequest) GetFiles() []string`

GetFiles returns the Files field if non-nil, zero value otherwise.

### GetFilesOk

`func (o *GitFileCheckRequest) GetFilesOk() (*[]string, bool)`

GetFilesOk returns a tuple with the Files field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiles

`func (o *GitFileCheckRequest) SetFiles(v []string)`

SetFiles sets Files field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


