# AgenticWorkflowProjectRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Branch** | **string** |  | 
**GitTokenId** | Pointer to **NullableString** | Qovery git token id used to clone the repository. Omit it (or send null) for a public repository, which is cloned without credentials.  | [optional] 

## Methods

### NewAgenticWorkflowProjectRepository

`func NewAgenticWorkflowProjectRepository(url string, branch string, ) *AgenticWorkflowProjectRepository`

NewAgenticWorkflowProjectRepository instantiates a new AgenticWorkflowProjectRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgenticWorkflowProjectRepositoryWithDefaults

`func NewAgenticWorkflowProjectRepositoryWithDefaults() *AgenticWorkflowProjectRepository`

NewAgenticWorkflowProjectRepositoryWithDefaults instantiates a new AgenticWorkflowProjectRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *AgenticWorkflowProjectRepository) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AgenticWorkflowProjectRepository) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AgenticWorkflowProjectRepository) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetBranch

`func (o *AgenticWorkflowProjectRepository) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *AgenticWorkflowProjectRepository) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *AgenticWorkflowProjectRepository) SetBranch(v string)`

SetBranch sets Branch field to given value.


### GetGitTokenId

`func (o *AgenticWorkflowProjectRepository) GetGitTokenId() string`

GetGitTokenId returns the GitTokenId field if non-nil, zero value otherwise.

### GetGitTokenIdOk

`func (o *AgenticWorkflowProjectRepository) GetGitTokenIdOk() (*string, bool)`

GetGitTokenIdOk returns a tuple with the GitTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitTokenId

`func (o *AgenticWorkflowProjectRepository) SetGitTokenId(v string)`

SetGitTokenId sets GitTokenId field to given value.

### HasGitTokenId

`func (o *AgenticWorkflowProjectRepository) HasGitTokenId() bool`

HasGitTokenId returns a boolean if a field has been set.

### SetGitTokenIdNil

`func (o *AgenticWorkflowProjectRepository) SetGitTokenIdNil(b bool)`

 SetGitTokenIdNil sets the value for GitTokenId to be an explicit nil

### UnsetGitTokenId
`func (o *AgenticWorkflowProjectRepository) UnsetGitTokenId()`

UnsetGitTokenId ensures that no value is present for GitTokenId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


