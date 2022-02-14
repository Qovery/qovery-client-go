# ApplicationGitRepositoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasAccess** | Pointer to **bool** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | repository name | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**RootPath** | Pointer to **string** |  | [optional] 
**DeployedCommitId** | Pointer to **string** | Git commit ID corresponding to the deployed version of the app | [optional] 
**DeployedCommitDate** | Pointer to **time.Time** | Git commit date corresponding to the deployed version of the app | [optional] [readonly] 
**DeployedCommitContributor** | Pointer to **string** | Git commit user corresponding to the deployed version of the app | [optional] 
**DeployedCommitTag** | Pointer to **string** |  | [optional] 

## Methods

### NewApplicationGitRepositoryResponse

`func NewApplicationGitRepositoryResponse() *ApplicationGitRepositoryResponse`

NewApplicationGitRepositoryResponse instantiates a new ApplicationGitRepositoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationGitRepositoryResponseWithDefaults

`func NewApplicationGitRepositoryResponseWithDefaults() *ApplicationGitRepositoryResponse`

NewApplicationGitRepositoryResponseWithDefaults instantiates a new ApplicationGitRepositoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasAccess

`func (o *ApplicationGitRepositoryResponse) GetHasAccess() bool`

GetHasAccess returns the HasAccess field if non-nil, zero value otherwise.

### GetHasAccessOk

`func (o *ApplicationGitRepositoryResponse) GetHasAccessOk() (*bool, bool)`

GetHasAccessOk returns a tuple with the HasAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccess

`func (o *ApplicationGitRepositoryResponse) SetHasAccess(v bool)`

SetHasAccess sets HasAccess field to given value.

### HasHasAccess

`func (o *ApplicationGitRepositoryResponse) HasHasAccess() bool`

HasHasAccess returns a boolean if a field has been set.

### GetProvider

`func (o *ApplicationGitRepositoryResponse) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ApplicationGitRepositoryResponse) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ApplicationGitRepositoryResponse) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ApplicationGitRepositoryResponse) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetOwner

`func (o *ApplicationGitRepositoryResponse) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ApplicationGitRepositoryResponse) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ApplicationGitRepositoryResponse) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ApplicationGitRepositoryResponse) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetUrl

`func (o *ApplicationGitRepositoryResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ApplicationGitRepositoryResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ApplicationGitRepositoryResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ApplicationGitRepositoryResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetName

`func (o *ApplicationGitRepositoryResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationGitRepositoryResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationGitRepositoryResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationGitRepositoryResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBranch

`func (o *ApplicationGitRepositoryResponse) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *ApplicationGitRepositoryResponse) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *ApplicationGitRepositoryResponse) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *ApplicationGitRepositoryResponse) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetRootPath

`func (o *ApplicationGitRepositoryResponse) GetRootPath() string`

GetRootPath returns the RootPath field if non-nil, zero value otherwise.

### GetRootPathOk

`func (o *ApplicationGitRepositoryResponse) GetRootPathOk() (*string, bool)`

GetRootPathOk returns a tuple with the RootPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPath

`func (o *ApplicationGitRepositoryResponse) SetRootPath(v string)`

SetRootPath sets RootPath field to given value.

### HasRootPath

`func (o *ApplicationGitRepositoryResponse) HasRootPath() bool`

HasRootPath returns a boolean if a field has been set.

### GetDeployedCommitId

`func (o *ApplicationGitRepositoryResponse) GetDeployedCommitId() string`

GetDeployedCommitId returns the DeployedCommitId field if non-nil, zero value otherwise.

### GetDeployedCommitIdOk

`func (o *ApplicationGitRepositoryResponse) GetDeployedCommitIdOk() (*string, bool)`

GetDeployedCommitIdOk returns a tuple with the DeployedCommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedCommitId

`func (o *ApplicationGitRepositoryResponse) SetDeployedCommitId(v string)`

SetDeployedCommitId sets DeployedCommitId field to given value.

### HasDeployedCommitId

`func (o *ApplicationGitRepositoryResponse) HasDeployedCommitId() bool`

HasDeployedCommitId returns a boolean if a field has been set.

### GetDeployedCommitDate

`func (o *ApplicationGitRepositoryResponse) GetDeployedCommitDate() time.Time`

GetDeployedCommitDate returns the DeployedCommitDate field if non-nil, zero value otherwise.

### GetDeployedCommitDateOk

`func (o *ApplicationGitRepositoryResponse) GetDeployedCommitDateOk() (*time.Time, bool)`

GetDeployedCommitDateOk returns a tuple with the DeployedCommitDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedCommitDate

`func (o *ApplicationGitRepositoryResponse) SetDeployedCommitDate(v time.Time)`

SetDeployedCommitDate sets DeployedCommitDate field to given value.

### HasDeployedCommitDate

`func (o *ApplicationGitRepositoryResponse) HasDeployedCommitDate() bool`

HasDeployedCommitDate returns a boolean if a field has been set.

### GetDeployedCommitContributor

`func (o *ApplicationGitRepositoryResponse) GetDeployedCommitContributor() string`

GetDeployedCommitContributor returns the DeployedCommitContributor field if non-nil, zero value otherwise.

### GetDeployedCommitContributorOk

`func (o *ApplicationGitRepositoryResponse) GetDeployedCommitContributorOk() (*string, bool)`

GetDeployedCommitContributorOk returns a tuple with the DeployedCommitContributor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedCommitContributor

`func (o *ApplicationGitRepositoryResponse) SetDeployedCommitContributor(v string)`

SetDeployedCommitContributor sets DeployedCommitContributor field to given value.

### HasDeployedCommitContributor

`func (o *ApplicationGitRepositoryResponse) HasDeployedCommitContributor() bool`

HasDeployedCommitContributor returns a boolean if a field has been set.

### GetDeployedCommitTag

`func (o *ApplicationGitRepositoryResponse) GetDeployedCommitTag() string`

GetDeployedCommitTag returns the DeployedCommitTag field if non-nil, zero value otherwise.

### GetDeployedCommitTagOk

`func (o *ApplicationGitRepositoryResponse) GetDeployedCommitTagOk() (*string, bool)`

GetDeployedCommitTagOk returns a tuple with the DeployedCommitTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedCommitTag

`func (o *ApplicationGitRepositoryResponse) SetDeployedCommitTag(v string)`

SetDeployedCommitTag sets DeployedCommitTag field to given value.

### HasDeployedCommitTag

`func (o *ApplicationGitRepositoryResponse) HasDeployedCommitTag() bool`

HasDeployedCommitTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


