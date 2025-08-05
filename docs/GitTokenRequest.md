# GitTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Type** | [**GitProviderEnum**](GitProviderEnum.md) |  | 
**Token** | **string** | The token from your git provider side | 
**Workspace** | Pointer to **string** | Mandatory only for BITBUCKET git provider, to allow us to fetch repositories at creation/edition of a service | [optional] 
**GitApiUrl** | Pointer to **string** | custom git api url for the given git provider/type. I.e: Self-hosted version of Gitlab | [optional] 

## Methods

### NewGitTokenRequest

`func NewGitTokenRequest(name string, type_ GitProviderEnum, token string, ) *GitTokenRequest`

NewGitTokenRequest instantiates a new GitTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitTokenRequestWithDefaults

`func NewGitTokenRequestWithDefaults() *GitTokenRequest`

NewGitTokenRequestWithDefaults instantiates a new GitTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GitTokenRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GitTokenRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GitTokenRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GitTokenRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GitTokenRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GitTokenRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GitTokenRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *GitTokenRequest) GetType() GitProviderEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GitTokenRequest) GetTypeOk() (*GitProviderEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GitTokenRequest) SetType(v GitProviderEnum)`

SetType sets Type field to given value.


### GetToken

`func (o *GitTokenRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GitTokenRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GitTokenRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetWorkspace

`func (o *GitTokenRequest) GetWorkspace() string`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *GitTokenRequest) GetWorkspaceOk() (*string, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *GitTokenRequest) SetWorkspace(v string)`

SetWorkspace sets Workspace field to given value.

### HasWorkspace

`func (o *GitTokenRequest) HasWorkspace() bool`

HasWorkspace returns a boolean if a field has been set.

### GetGitApiUrl

`func (o *GitTokenRequest) GetGitApiUrl() string`

GetGitApiUrl returns the GitApiUrl field if non-nil, zero value otherwise.

### GetGitApiUrlOk

`func (o *GitTokenRequest) GetGitApiUrlOk() (*string, bool)`

GetGitApiUrlOk returns a tuple with the GitApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitApiUrl

`func (o *GitTokenRequest) SetGitApiUrl(v string)`

SetGitApiUrl sets GitApiUrl field to given value.

### HasGitApiUrl

`func (o *GitTokenRequest) HasGitApiUrl() bool`

HasGitApiUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


