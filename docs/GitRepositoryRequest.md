# GitRepositoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | Git repository URL | 
**Branch** | Pointer to **NullableString** | Name of the branch to use (optional) | [optional] 
**RootPath** | Pointer to **string** | Root path within the repository | [optional] [default to "/"]
**GitTokenId** | Pointer to **NullableString** | The git token id on Qovery side | [optional] 
**Provider** | [**GitProvider**](GitProvider.md) |  | 

## Methods

### NewGitRepositoryRequest

`func NewGitRepositoryRequest(url string, provider GitProvider, ) *GitRepositoryRequest`

NewGitRepositoryRequest instantiates a new GitRepositoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitRepositoryRequestWithDefaults

`func NewGitRepositoryRequestWithDefaults() *GitRepositoryRequest`

NewGitRepositoryRequestWithDefaults instantiates a new GitRepositoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *GitRepositoryRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GitRepositoryRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GitRepositoryRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetBranch

`func (o *GitRepositoryRequest) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *GitRepositoryRequest) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *GitRepositoryRequest) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *GitRepositoryRequest) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### SetBranchNil

`func (o *GitRepositoryRequest) SetBranchNil(b bool)`

 SetBranchNil sets the value for Branch to be an explicit nil

### UnsetBranch
`func (o *GitRepositoryRequest) UnsetBranch()`

UnsetBranch ensures that no value is present for Branch, not even an explicit nil
### GetRootPath

`func (o *GitRepositoryRequest) GetRootPath() string`

GetRootPath returns the RootPath field if non-nil, zero value otherwise.

### GetRootPathOk

`func (o *GitRepositoryRequest) GetRootPathOk() (*string, bool)`

GetRootPathOk returns a tuple with the RootPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPath

`func (o *GitRepositoryRequest) SetRootPath(v string)`

SetRootPath sets RootPath field to given value.

### HasRootPath

`func (o *GitRepositoryRequest) HasRootPath() bool`

HasRootPath returns a boolean if a field has been set.

### GetGitTokenId

`func (o *GitRepositoryRequest) GetGitTokenId() string`

GetGitTokenId returns the GitTokenId field if non-nil, zero value otherwise.

### GetGitTokenIdOk

`func (o *GitRepositoryRequest) GetGitTokenIdOk() (*string, bool)`

GetGitTokenIdOk returns a tuple with the GitTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitTokenId

`func (o *GitRepositoryRequest) SetGitTokenId(v string)`

SetGitTokenId sets GitTokenId field to given value.

### HasGitTokenId

`func (o *GitRepositoryRequest) HasGitTokenId() bool`

HasGitTokenId returns a boolean if a field has been set.

### SetGitTokenIdNil

`func (o *GitRepositoryRequest) SetGitTokenIdNil(b bool)`

 SetGitTokenIdNil sets the value for GitTokenId to be an explicit nil

### UnsetGitTokenId
`func (o *GitRepositoryRequest) UnsetGitTokenId()`

UnsetGitTokenId ensures that no value is present for GitTokenId, not even an explicit nil
### GetProvider

`func (o *GitRepositoryRequest) GetProvider() GitProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *GitRepositoryRequest) GetProviderOk() (*GitProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *GitRepositoryRequest) SetProvider(v GitProvider)`

SetProvider sets Provider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


