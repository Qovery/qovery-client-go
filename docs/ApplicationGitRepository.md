# ApplicationGitRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasAccess** | Pointer to **bool** |  | [optional] 
**Provider** | Pointer to [**GitProviderEnum**](GitProviderEnum.md) |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | repository name | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**RootPath** | Pointer to **string** |  | [optional] 
**DeployedCommitId** | Pointer to **string** | Git commit ID corresponding to the deployed version of the app | [optional] 
**DeployedCommitDate** | Pointer to **time.Time** | Git commit date corresponding to the deployed version of the app | [optional] [readonly] 
**DeployedCommitContributor** | Pointer to **string** | Git commit user corresponding to the deployed version of the app | [optional] 
**DeployedCommitTag** | Pointer to **string** |  | [optional] 
**GitTokenId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewApplicationGitRepository

`func NewApplicationGitRepository() *ApplicationGitRepository`

NewApplicationGitRepository instantiates a new ApplicationGitRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationGitRepositoryWithDefaults

`func NewApplicationGitRepositoryWithDefaults() *ApplicationGitRepository`

NewApplicationGitRepositoryWithDefaults instantiates a new ApplicationGitRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasAccess

`func (o *ApplicationGitRepository) GetHasAccess() bool`

GetHasAccess returns the HasAccess field if non-nil, zero value otherwise.

### GetHasAccessOk

`func (o *ApplicationGitRepository) GetHasAccessOk() (*bool, bool)`

GetHasAccessOk returns a tuple with the HasAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccess

`func (o *ApplicationGitRepository) SetHasAccess(v bool)`

SetHasAccess sets HasAccess field to given value.

### HasHasAccess

`func (o *ApplicationGitRepository) HasHasAccess() bool`

HasHasAccess returns a boolean if a field has been set.

### GetProvider

`func (o *ApplicationGitRepository) GetProvider() GitProviderEnum`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ApplicationGitRepository) GetProviderOk() (*GitProviderEnum, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ApplicationGitRepository) SetProvider(v GitProviderEnum)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ApplicationGitRepository) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetOwner

`func (o *ApplicationGitRepository) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ApplicationGitRepository) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ApplicationGitRepository) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ApplicationGitRepository) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetUrl

`func (o *ApplicationGitRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ApplicationGitRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ApplicationGitRepository) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ApplicationGitRepository) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetName

`func (o *ApplicationGitRepository) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationGitRepository) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationGitRepository) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationGitRepository) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBranch

`func (o *ApplicationGitRepository) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *ApplicationGitRepository) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *ApplicationGitRepository) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *ApplicationGitRepository) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetRootPath

`func (o *ApplicationGitRepository) GetRootPath() string`

GetRootPath returns the RootPath field if non-nil, zero value otherwise.

### GetRootPathOk

`func (o *ApplicationGitRepository) GetRootPathOk() (*string, bool)`

GetRootPathOk returns a tuple with the RootPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPath

`func (o *ApplicationGitRepository) SetRootPath(v string)`

SetRootPath sets RootPath field to given value.

### HasRootPath

`func (o *ApplicationGitRepository) HasRootPath() bool`

HasRootPath returns a boolean if a field has been set.

### GetDeployedCommitId

`func (o *ApplicationGitRepository) GetDeployedCommitId() string`

GetDeployedCommitId returns the DeployedCommitId field if non-nil, zero value otherwise.

### GetDeployedCommitIdOk

`func (o *ApplicationGitRepository) GetDeployedCommitIdOk() (*string, bool)`

GetDeployedCommitIdOk returns a tuple with the DeployedCommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedCommitId

`func (o *ApplicationGitRepository) SetDeployedCommitId(v string)`

SetDeployedCommitId sets DeployedCommitId field to given value.

### HasDeployedCommitId

`func (o *ApplicationGitRepository) HasDeployedCommitId() bool`

HasDeployedCommitId returns a boolean if a field has been set.

### GetDeployedCommitDate

`func (o *ApplicationGitRepository) GetDeployedCommitDate() time.Time`

GetDeployedCommitDate returns the DeployedCommitDate field if non-nil, zero value otherwise.

### GetDeployedCommitDateOk

`func (o *ApplicationGitRepository) GetDeployedCommitDateOk() (*time.Time, bool)`

GetDeployedCommitDateOk returns a tuple with the DeployedCommitDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedCommitDate

`func (o *ApplicationGitRepository) SetDeployedCommitDate(v time.Time)`

SetDeployedCommitDate sets DeployedCommitDate field to given value.

### HasDeployedCommitDate

`func (o *ApplicationGitRepository) HasDeployedCommitDate() bool`

HasDeployedCommitDate returns a boolean if a field has been set.

### GetDeployedCommitContributor

`func (o *ApplicationGitRepository) GetDeployedCommitContributor() string`

GetDeployedCommitContributor returns the DeployedCommitContributor field if non-nil, zero value otherwise.

### GetDeployedCommitContributorOk

`func (o *ApplicationGitRepository) GetDeployedCommitContributorOk() (*string, bool)`

GetDeployedCommitContributorOk returns a tuple with the DeployedCommitContributor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedCommitContributor

`func (o *ApplicationGitRepository) SetDeployedCommitContributor(v string)`

SetDeployedCommitContributor sets DeployedCommitContributor field to given value.

### HasDeployedCommitContributor

`func (o *ApplicationGitRepository) HasDeployedCommitContributor() bool`

HasDeployedCommitContributor returns a boolean if a field has been set.

### GetDeployedCommitTag

`func (o *ApplicationGitRepository) GetDeployedCommitTag() string`

GetDeployedCommitTag returns the DeployedCommitTag field if non-nil, zero value otherwise.

### GetDeployedCommitTagOk

`func (o *ApplicationGitRepository) GetDeployedCommitTagOk() (*string, bool)`

GetDeployedCommitTagOk returns a tuple with the DeployedCommitTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedCommitTag

`func (o *ApplicationGitRepository) SetDeployedCommitTag(v string)`

SetDeployedCommitTag sets DeployedCommitTag field to given value.

### HasDeployedCommitTag

`func (o *ApplicationGitRepository) HasDeployedCommitTag() bool`

HasDeployedCommitTag returns a boolean if a field has been set.

### GetGitTokenId

`func (o *ApplicationGitRepository) GetGitTokenId() string`

GetGitTokenId returns the GitTokenId field if non-nil, zero value otherwise.

### GetGitTokenIdOk

`func (o *ApplicationGitRepository) GetGitTokenIdOk() (*string, bool)`

GetGitTokenIdOk returns a tuple with the GitTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitTokenId

`func (o *ApplicationGitRepository) SetGitTokenId(v string)`

SetGitTokenId sets GitTokenId field to given value.

### HasGitTokenId

`func (o *ApplicationGitRepository) HasGitTokenId() bool`

HasGitTokenId returns a boolean if a field has been set.

### SetGitTokenIdNil

`func (o *ApplicationGitRepository) SetGitTokenIdNil(b bool)`

 SetGitTokenIdNil sets the value for GitTokenId to be an explicit nil

### UnsetGitTokenId
`func (o *ApplicationGitRepository) UnsetGitTokenId()`

UnsetGitTokenId ensures that no value is present for GitTokenId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


