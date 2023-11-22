# HelmValuesGitRepositoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | application git repository URL | 
**Branch** | **string** | Name of the branch to use. This is optional If not specified, then the branch used is the &#x60;main&#x60; or &#x60;master&#x60; one  | 
**Paths** | **[]string** | List of path inside your git repository to locate values file. Must start by a / | 
**GitTokenId** | Pointer to **NullableString** | The git token id on Qovery side | [optional] 

## Methods

### NewHelmValuesGitRepositoryRequest

`func NewHelmValuesGitRepositoryRequest(url string, branch string, paths []string, ) *HelmValuesGitRepositoryRequest`

NewHelmValuesGitRepositoryRequest instantiates a new HelmValuesGitRepositoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmValuesGitRepositoryRequestWithDefaults

`func NewHelmValuesGitRepositoryRequestWithDefaults() *HelmValuesGitRepositoryRequest`

NewHelmValuesGitRepositoryRequestWithDefaults instantiates a new HelmValuesGitRepositoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *HelmValuesGitRepositoryRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *HelmValuesGitRepositoryRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *HelmValuesGitRepositoryRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetBranch

`func (o *HelmValuesGitRepositoryRequest) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *HelmValuesGitRepositoryRequest) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *HelmValuesGitRepositoryRequest) SetBranch(v string)`

SetBranch sets Branch field to given value.


### GetPaths

`func (o *HelmValuesGitRepositoryRequest) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *HelmValuesGitRepositoryRequest) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *HelmValuesGitRepositoryRequest) SetPaths(v []string)`

SetPaths sets Paths field to given value.


### GetGitTokenId

`func (o *HelmValuesGitRepositoryRequest) GetGitTokenId() string`

GetGitTokenId returns the GitTokenId field if non-nil, zero value otherwise.

### GetGitTokenIdOk

`func (o *HelmValuesGitRepositoryRequest) GetGitTokenIdOk() (*string, bool)`

GetGitTokenIdOk returns a tuple with the GitTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitTokenId

`func (o *HelmValuesGitRepositoryRequest) SetGitTokenId(v string)`

SetGitTokenId sets GitTokenId field to given value.

### HasGitTokenId

`func (o *HelmValuesGitRepositoryRequest) HasGitTokenId() bool`

HasGitTokenId returns a boolean if a field has been set.

### SetGitTokenIdNil

`func (o *HelmValuesGitRepositoryRequest) SetGitTokenIdNil(b bool)`

 SetGitTokenIdNil sets the value for GitTokenId to be an explicit nil

### UnsetGitTokenId
`func (o *HelmValuesGitRepositoryRequest) UnsetGitTokenId()`

UnsetGitTokenId ensures that no value is present for GitTokenId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


