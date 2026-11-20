# ContainerGitRepositorySourceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | Git repository URL | 
**Branch** | Pointer to **string** | Name of the branch to use. This is optional If not specified, then the branch used is the &#x60;main&#x60; or &#x60;master&#x60; one  | [optional] 
**RootPath** | Pointer to **string** | indicates the root path in the repository. | [optional] [default to "/"]
**GitTokenId** | Pointer to **NullableString** | The git token id on Qovery side | [optional] 
**Provider** | [**GitProviderEnum**](GitProviderEnum.md) |  | 
**DockerfilePath** | **string** | The path of the associated Dockerfile | 
**DockerTargetBuildStage** | Pointer to **NullableString** | The target build stage in the Dockerfile to build. If null, the last stage of the Dockerfile is built. | [optional] 

## Methods

### NewContainerGitRepositorySourceRequest

`func NewContainerGitRepositorySourceRequest(url string, provider GitProviderEnum, dockerfilePath string, ) *ContainerGitRepositorySourceRequest`

NewContainerGitRepositorySourceRequest instantiates a new ContainerGitRepositorySourceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerGitRepositorySourceRequestWithDefaults

`func NewContainerGitRepositorySourceRequestWithDefaults() *ContainerGitRepositorySourceRequest`

NewContainerGitRepositorySourceRequestWithDefaults instantiates a new ContainerGitRepositorySourceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ContainerGitRepositorySourceRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ContainerGitRepositorySourceRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ContainerGitRepositorySourceRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetBranch

`func (o *ContainerGitRepositorySourceRequest) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *ContainerGitRepositorySourceRequest) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *ContainerGitRepositorySourceRequest) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *ContainerGitRepositorySourceRequest) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetRootPath

`func (o *ContainerGitRepositorySourceRequest) GetRootPath() string`

GetRootPath returns the RootPath field if non-nil, zero value otherwise.

### GetRootPathOk

`func (o *ContainerGitRepositorySourceRequest) GetRootPathOk() (*string, bool)`

GetRootPathOk returns a tuple with the RootPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPath

`func (o *ContainerGitRepositorySourceRequest) SetRootPath(v string)`

SetRootPath sets RootPath field to given value.

### HasRootPath

`func (o *ContainerGitRepositorySourceRequest) HasRootPath() bool`

HasRootPath returns a boolean if a field has been set.

### GetGitTokenId

`func (o *ContainerGitRepositorySourceRequest) GetGitTokenId() string`

GetGitTokenId returns the GitTokenId field if non-nil, zero value otherwise.

### GetGitTokenIdOk

`func (o *ContainerGitRepositorySourceRequest) GetGitTokenIdOk() (*string, bool)`

GetGitTokenIdOk returns a tuple with the GitTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitTokenId

`func (o *ContainerGitRepositorySourceRequest) SetGitTokenId(v string)`

SetGitTokenId sets GitTokenId field to given value.

### HasGitTokenId

`func (o *ContainerGitRepositorySourceRequest) HasGitTokenId() bool`

HasGitTokenId returns a boolean if a field has been set.

### SetGitTokenIdNil

`func (o *ContainerGitRepositorySourceRequest) SetGitTokenIdNil(b bool)`

 SetGitTokenIdNil sets the value for GitTokenId to be an explicit nil

### UnsetGitTokenId
`func (o *ContainerGitRepositorySourceRequest) UnsetGitTokenId()`

UnsetGitTokenId ensures that no value is present for GitTokenId, not even an explicit nil
### GetProvider

`func (o *ContainerGitRepositorySourceRequest) GetProvider() GitProviderEnum`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ContainerGitRepositorySourceRequest) GetProviderOk() (*GitProviderEnum, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ContainerGitRepositorySourceRequest) SetProvider(v GitProviderEnum)`

SetProvider sets Provider field to given value.


### GetDockerfilePath

`func (o *ContainerGitRepositorySourceRequest) GetDockerfilePath() string`

GetDockerfilePath returns the DockerfilePath field if non-nil, zero value otherwise.

### GetDockerfilePathOk

`func (o *ContainerGitRepositorySourceRequest) GetDockerfilePathOk() (*string, bool)`

GetDockerfilePathOk returns a tuple with the DockerfilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfilePath

`func (o *ContainerGitRepositorySourceRequest) SetDockerfilePath(v string)`

SetDockerfilePath sets DockerfilePath field to given value.


### GetDockerTargetBuildStage

`func (o *ContainerGitRepositorySourceRequest) GetDockerTargetBuildStage() string`

GetDockerTargetBuildStage returns the DockerTargetBuildStage field if non-nil, zero value otherwise.

### GetDockerTargetBuildStageOk

`func (o *ContainerGitRepositorySourceRequest) GetDockerTargetBuildStageOk() (*string, bool)`

GetDockerTargetBuildStageOk returns a tuple with the DockerTargetBuildStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerTargetBuildStage

`func (o *ContainerGitRepositorySourceRequest) SetDockerTargetBuildStage(v string)`

SetDockerTargetBuildStage sets DockerTargetBuildStage field to given value.

### HasDockerTargetBuildStage

`func (o *ContainerGitRepositorySourceRequest) HasDockerTargetBuildStage() bool`

HasDockerTargetBuildStage returns a boolean if a field has been set.

### SetDockerTargetBuildStageNil

`func (o *ContainerGitRepositorySourceRequest) SetDockerTargetBuildStageNil(b bool)`

 SetDockerTargetBuildStageNil sets the value for DockerTargetBuildStage to be an explicit nil

### UnsetDockerTargetBuildStage
`func (o *ContainerGitRepositorySourceRequest) UnsetDockerTargetBuildStage()`

UnsetDockerTargetBuildStage ensures that no value is present for DockerTargetBuildStage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


