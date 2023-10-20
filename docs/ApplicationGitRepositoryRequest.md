# ApplicationGitRepositoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | application git repository URL | 
**Branch** | Pointer to **string** | Name of the branch to use. This is optional If not specified, then the branch used is the &#x60;main&#x60; or &#x60;master&#x60; one  | [optional] 
**RootPath** | Pointer to **string** | indicates the root path of the application. | [optional] [default to "/"]
**GitTokenId** | Pointer to **NullableString** | The git token id on Qovery side | [optional] 

## Methods

### NewApplicationGitRepositoryRequest

`func NewApplicationGitRepositoryRequest(url string, ) *ApplicationGitRepositoryRequest`

NewApplicationGitRepositoryRequest instantiates a new ApplicationGitRepositoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationGitRepositoryRequestWithDefaults

`func NewApplicationGitRepositoryRequestWithDefaults() *ApplicationGitRepositoryRequest`

NewApplicationGitRepositoryRequestWithDefaults instantiates a new ApplicationGitRepositoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ApplicationGitRepositoryRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ApplicationGitRepositoryRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ApplicationGitRepositoryRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetBranch

`func (o *ApplicationGitRepositoryRequest) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *ApplicationGitRepositoryRequest) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *ApplicationGitRepositoryRequest) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *ApplicationGitRepositoryRequest) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetRootPath

`func (o *ApplicationGitRepositoryRequest) GetRootPath() string`

GetRootPath returns the RootPath field if non-nil, zero value otherwise.

### GetRootPathOk

`func (o *ApplicationGitRepositoryRequest) GetRootPathOk() (*string, bool)`

GetRootPathOk returns a tuple with the RootPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPath

`func (o *ApplicationGitRepositoryRequest) SetRootPath(v string)`

SetRootPath sets RootPath field to given value.

### HasRootPath

`func (o *ApplicationGitRepositoryRequest) HasRootPath() bool`

HasRootPath returns a boolean if a field has been set.

### GetGitTokenId

`func (o *ApplicationGitRepositoryRequest) GetGitTokenId() string`

GetGitTokenId returns the GitTokenId field if non-nil, zero value otherwise.

### GetGitTokenIdOk

`func (o *ApplicationGitRepositoryRequest) GetGitTokenIdOk() (*string, bool)`

GetGitTokenIdOk returns a tuple with the GitTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitTokenId

`func (o *ApplicationGitRepositoryRequest) SetGitTokenId(v string)`

SetGitTokenId sets GitTokenId field to given value.

### HasGitTokenId

`func (o *ApplicationGitRepositoryRequest) HasGitTokenId() bool`

HasGitTokenId returns a boolean if a field has been set.

### SetGitTokenIdNil

`func (o *ApplicationGitRepositoryRequest) SetGitTokenIdNil(b bool)`

 SetGitTokenIdNil sets the value for GitTokenId to be an explicit nil

### UnsetGitTokenId
`func (o *ApplicationGitRepositoryRequest) UnsetGitTokenId()`

UnsetGitTokenId ensures that no value is present for GitTokenId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


