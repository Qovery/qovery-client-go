# ContainerGitRepositorySourceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasAccess** | Pointer to **bool** |  | [optional] 
**Provider** | [**GitProviderEnum**](GitProviderEnum.md) |  | 
**Owner** | **string** |  | 
**Url** | **string** |  | 
**Name** | **string** | repository name | 
**Branch** | Pointer to **string** |  | [optional] 
**RootPath** | Pointer to **string** |  | [optional] 
**DeployedCommitId** | Pointer to **string** | Git commit ID corresponding to the deployed version of the app | [optional] 
**DeployedCommitDate** | Pointer to **time.Time** | Git commit date corresponding to the deployed version of the app | [optional] [readonly] 
**DeployedCommitContributor** | Pointer to **string** | Git commit user corresponding to the deployed version of the app | [optional] 
**DeployedCommitTag** | Pointer to **string** |  | [optional] 
**GitTokenId** | Pointer to **NullableString** |  | [optional] 
**GitTokenName** | Pointer to **NullableString** |  | [optional] 
**DockerfilePath** | **string** | The path of the associated Dockerfile | 
**DockerTargetBuildStage** | Pointer to **NullableString** | The target build stage in the Dockerfile to build. If null, the last stage of the Dockerfile is built. | [optional] 

## Methods

### NewContainerGitRepositorySourceResponse

`func NewContainerGitRepositorySourceResponse(provider GitProviderEnum, owner string, url string, name string, dockerfilePath string, ) *ContainerGitRepositorySourceResponse`

NewContainerGitRepositorySourceResponse instantiates a new ContainerGitRepositorySourceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerGitRepositorySourceResponseWithDefaults

`func NewContainerGitRepositorySourceResponseWithDefaults() *ContainerGitRepositorySourceResponse`

NewContainerGitRepositorySourceResponseWithDefaults instantiates a new ContainerGitRepositorySourceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasAccess

`func (o *ContainerGitRepositorySourceResponse) GetHasAccess() bool`

GetHasAccess returns the HasAccess field if non-nil, zero value otherwise.

### GetHasAccessOk

`func (o *ContainerGitRepositorySourceResponse) GetHasAccessOk() (*bool, bool)`

GetHasAccessOk returns a tuple with the HasAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccess

`func (o *ContainerGitRepositorySourceResponse) SetHasAccess(v bool)`

SetHasAccess sets HasAccess field to given value.

### HasHasAccess

`func (o *ContainerGitRepositorySourceResponse) HasHasAccess() bool`

HasHasAccess returns a boolean if a field has been set.

### GetProvider

`func (o *ContainerGitRepositorySourceResponse) GetProvider() GitProviderEnum`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ContainerGitRepositorySourceResponse) GetProviderOk() (*GitProviderEnum, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ContainerGitRepositorySourceResponse) SetProvider(v GitProviderEnum)`

SetProvider sets Provider field to given value.


### GetOwner

`func (o *ContainerGitRepositorySourceResponse) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ContainerGitRepositorySourceResponse) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ContainerGitRepositorySourceResponse) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetUrl

`func (o *ContainerGitRepositorySourceResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ContainerGitRepositorySourceResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ContainerGitRepositorySourceResponse) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetName

`func (o *ContainerGitRepositorySourceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerGitRepositorySourceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerGitRepositorySourceResponse) SetName(v string)`

SetName sets Name field to given value.


### GetBranch

`func (o *ContainerGitRepositorySourceResponse) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *ContainerGitRepositorySourceResponse) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *ContainerGitRepositorySourceResponse) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *ContainerGitRepositorySourceResponse) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetRootPath

`func (o *ContainerGitRepositorySourceResponse) GetRootPath() string`

GetRootPath returns the RootPath field if non-nil, zero value otherwise.

### GetRootPathOk

`func (o *ContainerGitRepositorySourceResponse) GetRootPathOk() (*string, bool)`

GetRootPathOk returns a tuple with the RootPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPath

`func (o *ContainerGitRepositorySourceResponse) SetRootPath(v string)`

SetRootPath sets RootPath field to given value.

### HasRootPath

`func (o *ContainerGitRepositorySourceResponse) HasRootPath() bool`

HasRootPath returns a boolean if a field has been set.

### GetDeployedCommitId

`func (o *ContainerGitRepositorySourceResponse) GetDeployedCommitId() string`

GetDeployedCommitId returns the DeployedCommitId field if non-nil, zero value otherwise.

### GetDeployedCommitIdOk

`func (o *ContainerGitRepositorySourceResponse) GetDeployedCommitIdOk() (*string, bool)`

GetDeployedCommitIdOk returns a tuple with the DeployedCommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedCommitId

`func (o *ContainerGitRepositorySourceResponse) SetDeployedCommitId(v string)`

SetDeployedCommitId sets DeployedCommitId field to given value.

### HasDeployedCommitId

`func (o *ContainerGitRepositorySourceResponse) HasDeployedCommitId() bool`

HasDeployedCommitId returns a boolean if a field has been set.

### GetDeployedCommitDate

`func (o *ContainerGitRepositorySourceResponse) GetDeployedCommitDate() time.Time`

GetDeployedCommitDate returns the DeployedCommitDate field if non-nil, zero value otherwise.

### GetDeployedCommitDateOk

`func (o *ContainerGitRepositorySourceResponse) GetDeployedCommitDateOk() (*time.Time, bool)`

GetDeployedCommitDateOk returns a tuple with the DeployedCommitDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedCommitDate

`func (o *ContainerGitRepositorySourceResponse) SetDeployedCommitDate(v time.Time)`

SetDeployedCommitDate sets DeployedCommitDate field to given value.

### HasDeployedCommitDate

`func (o *ContainerGitRepositorySourceResponse) HasDeployedCommitDate() bool`

HasDeployedCommitDate returns a boolean if a field has been set.

### GetDeployedCommitContributor

`func (o *ContainerGitRepositorySourceResponse) GetDeployedCommitContributor() string`

GetDeployedCommitContributor returns the DeployedCommitContributor field if non-nil, zero value otherwise.

### GetDeployedCommitContributorOk

`func (o *ContainerGitRepositorySourceResponse) GetDeployedCommitContributorOk() (*string, bool)`

GetDeployedCommitContributorOk returns a tuple with the DeployedCommitContributor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedCommitContributor

`func (o *ContainerGitRepositorySourceResponse) SetDeployedCommitContributor(v string)`

SetDeployedCommitContributor sets DeployedCommitContributor field to given value.

### HasDeployedCommitContributor

`func (o *ContainerGitRepositorySourceResponse) HasDeployedCommitContributor() bool`

HasDeployedCommitContributor returns a boolean if a field has been set.

### GetDeployedCommitTag

`func (o *ContainerGitRepositorySourceResponse) GetDeployedCommitTag() string`

GetDeployedCommitTag returns the DeployedCommitTag field if non-nil, zero value otherwise.

### GetDeployedCommitTagOk

`func (o *ContainerGitRepositorySourceResponse) GetDeployedCommitTagOk() (*string, bool)`

GetDeployedCommitTagOk returns a tuple with the DeployedCommitTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedCommitTag

`func (o *ContainerGitRepositorySourceResponse) SetDeployedCommitTag(v string)`

SetDeployedCommitTag sets DeployedCommitTag field to given value.

### HasDeployedCommitTag

`func (o *ContainerGitRepositorySourceResponse) HasDeployedCommitTag() bool`

HasDeployedCommitTag returns a boolean if a field has been set.

### GetGitTokenId

`func (o *ContainerGitRepositorySourceResponse) GetGitTokenId() string`

GetGitTokenId returns the GitTokenId field if non-nil, zero value otherwise.

### GetGitTokenIdOk

`func (o *ContainerGitRepositorySourceResponse) GetGitTokenIdOk() (*string, bool)`

GetGitTokenIdOk returns a tuple with the GitTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitTokenId

`func (o *ContainerGitRepositorySourceResponse) SetGitTokenId(v string)`

SetGitTokenId sets GitTokenId field to given value.

### HasGitTokenId

`func (o *ContainerGitRepositorySourceResponse) HasGitTokenId() bool`

HasGitTokenId returns a boolean if a field has been set.

### SetGitTokenIdNil

`func (o *ContainerGitRepositorySourceResponse) SetGitTokenIdNil(b bool)`

 SetGitTokenIdNil sets the value for GitTokenId to be an explicit nil

### UnsetGitTokenId
`func (o *ContainerGitRepositorySourceResponse) UnsetGitTokenId()`

UnsetGitTokenId ensures that no value is present for GitTokenId, not even an explicit nil
### GetGitTokenName

`func (o *ContainerGitRepositorySourceResponse) GetGitTokenName() string`

GetGitTokenName returns the GitTokenName field if non-nil, zero value otherwise.

### GetGitTokenNameOk

`func (o *ContainerGitRepositorySourceResponse) GetGitTokenNameOk() (*string, bool)`

GetGitTokenNameOk returns a tuple with the GitTokenName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitTokenName

`func (o *ContainerGitRepositorySourceResponse) SetGitTokenName(v string)`

SetGitTokenName sets GitTokenName field to given value.

### HasGitTokenName

`func (o *ContainerGitRepositorySourceResponse) HasGitTokenName() bool`

HasGitTokenName returns a boolean if a field has been set.

### SetGitTokenNameNil

`func (o *ContainerGitRepositorySourceResponse) SetGitTokenNameNil(b bool)`

 SetGitTokenNameNil sets the value for GitTokenName to be an explicit nil

### UnsetGitTokenName
`func (o *ContainerGitRepositorySourceResponse) UnsetGitTokenName()`

UnsetGitTokenName ensures that no value is present for GitTokenName, not even an explicit nil
### GetDockerfilePath

`func (o *ContainerGitRepositorySourceResponse) GetDockerfilePath() string`

GetDockerfilePath returns the DockerfilePath field if non-nil, zero value otherwise.

### GetDockerfilePathOk

`func (o *ContainerGitRepositorySourceResponse) GetDockerfilePathOk() (*string, bool)`

GetDockerfilePathOk returns a tuple with the DockerfilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfilePath

`func (o *ContainerGitRepositorySourceResponse) SetDockerfilePath(v string)`

SetDockerfilePath sets DockerfilePath field to given value.


### GetDockerTargetBuildStage

`func (o *ContainerGitRepositorySourceResponse) GetDockerTargetBuildStage() string`

GetDockerTargetBuildStage returns the DockerTargetBuildStage field if non-nil, zero value otherwise.

### GetDockerTargetBuildStageOk

`func (o *ContainerGitRepositorySourceResponse) GetDockerTargetBuildStageOk() (*string, bool)`

GetDockerTargetBuildStageOk returns a tuple with the DockerTargetBuildStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerTargetBuildStage

`func (o *ContainerGitRepositorySourceResponse) SetDockerTargetBuildStage(v string)`

SetDockerTargetBuildStage sets DockerTargetBuildStage field to given value.

### HasDockerTargetBuildStage

`func (o *ContainerGitRepositorySourceResponse) HasDockerTargetBuildStage() bool`

HasDockerTargetBuildStage returns a boolean if a field has been set.

### SetDockerTargetBuildStageNil

`func (o *ContainerGitRepositorySourceResponse) SetDockerTargetBuildStageNil(b bool)`

 SetDockerTargetBuildStageNil sets the value for DockerTargetBuildStage to be an explicit nil

### UnsetDockerTargetBuildStage
`func (o *ContainerGitRepositorySourceResponse) UnsetDockerTargetBuildStage()`

UnsetDockerTargetBuildStage ensures that no value is present for DockerTargetBuildStage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


