# ClusterEksAnywhereGitRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | EKS Anywhere git repository URL | 
**Branch** | Pointer to **string** | Name of the branch to use. This is optional. If not specified, the default branch of the repository is used.  | [optional] 
**CommitId** | Pointer to **string** | Optional git commit SHA to pin EKS Anywhere configuration on a specific revision. If omitted, the latest commit from the selected branch is used.  | [optional] 
**GitTokenId** | **string** | Qovery git token id used to access the repository | 
**Provider** | Pointer to [**GitProviderEnum**](GitProviderEnum.md) |  | [optional] 

## Methods

### NewClusterEksAnywhereGitRepository

`func NewClusterEksAnywhereGitRepository(url string, gitTokenId string, ) *ClusterEksAnywhereGitRepository`

NewClusterEksAnywhereGitRepository instantiates a new ClusterEksAnywhereGitRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterEksAnywhereGitRepositoryWithDefaults

`func NewClusterEksAnywhereGitRepositoryWithDefaults() *ClusterEksAnywhereGitRepository`

NewClusterEksAnywhereGitRepositoryWithDefaults instantiates a new ClusterEksAnywhereGitRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ClusterEksAnywhereGitRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ClusterEksAnywhereGitRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ClusterEksAnywhereGitRepository) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetBranch

`func (o *ClusterEksAnywhereGitRepository) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *ClusterEksAnywhereGitRepository) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *ClusterEksAnywhereGitRepository) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *ClusterEksAnywhereGitRepository) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetCommitId

`func (o *ClusterEksAnywhereGitRepository) GetCommitId() string`

GetCommitId returns the CommitId field if non-nil, zero value otherwise.

### GetCommitIdOk

`func (o *ClusterEksAnywhereGitRepository) GetCommitIdOk() (*string, bool)`

GetCommitIdOk returns a tuple with the CommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitId

`func (o *ClusterEksAnywhereGitRepository) SetCommitId(v string)`

SetCommitId sets CommitId field to given value.

### HasCommitId

`func (o *ClusterEksAnywhereGitRepository) HasCommitId() bool`

HasCommitId returns a boolean if a field has been set.

### GetGitTokenId

`func (o *ClusterEksAnywhereGitRepository) GetGitTokenId() string`

GetGitTokenId returns the GitTokenId field if non-nil, zero value otherwise.

### GetGitTokenIdOk

`func (o *ClusterEksAnywhereGitRepository) GetGitTokenIdOk() (*string, bool)`

GetGitTokenIdOk returns a tuple with the GitTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitTokenId

`func (o *ClusterEksAnywhereGitRepository) SetGitTokenId(v string)`

SetGitTokenId sets GitTokenId field to given value.


### GetProvider

`func (o *ClusterEksAnywhereGitRepository) GetProvider() GitProviderEnum`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ClusterEksAnywhereGitRepository) GetProviderOk() (*GitProviderEnum, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ClusterEksAnywhereGitRepository) SetProvider(v GitProviderEnum)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ClusterEksAnywhereGitRepository) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


