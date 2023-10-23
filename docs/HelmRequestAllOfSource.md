# HelmRequestAllOfSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Git** | Pointer to [**NullableHelmRequestAllOfSourceGit**](HelmRequestAllOfSourceGit.md) |  | [optional] 
**Repository** | Pointer to [**NullableHelmRequestAllOfSourceRepository**](HelmRequestAllOfSourceRepository.md) |  | [optional] 

## Methods

### NewHelmRequestAllOfSource

`func NewHelmRequestAllOfSource() *HelmRequestAllOfSource`

NewHelmRequestAllOfSource instantiates a new HelmRequestAllOfSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmRequestAllOfSourceWithDefaults

`func NewHelmRequestAllOfSourceWithDefaults() *HelmRequestAllOfSource`

NewHelmRequestAllOfSourceWithDefaults instantiates a new HelmRequestAllOfSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGit

`func (o *HelmRequestAllOfSource) GetGit() HelmRequestAllOfSourceGit`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *HelmRequestAllOfSource) GetGitOk() (*HelmRequestAllOfSourceGit, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *HelmRequestAllOfSource) SetGit(v HelmRequestAllOfSourceGit)`

SetGit sets Git field to given value.

### HasGit

`func (o *HelmRequestAllOfSource) HasGit() bool`

HasGit returns a boolean if a field has been set.

### SetGitNil

`func (o *HelmRequestAllOfSource) SetGitNil(b bool)`

 SetGitNil sets the value for Git to be an explicit nil

### UnsetGit
`func (o *HelmRequestAllOfSource) UnsetGit()`

UnsetGit ensures that no value is present for Git, not even an explicit nil
### GetRepository

`func (o *HelmRequestAllOfSource) GetRepository() HelmRequestAllOfSourceRepository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *HelmRequestAllOfSource) GetRepositoryOk() (*HelmRequestAllOfSourceRepository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *HelmRequestAllOfSource) SetRepository(v HelmRequestAllOfSourceRepository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *HelmRequestAllOfSource) HasRepository() bool`

HasRepository returns a boolean if a field has been set.

### SetRepositoryNil

`func (o *HelmRequestAllOfSource) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *HelmRequestAllOfSource) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


