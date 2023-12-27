# HelmResponseAllOfValuesOverrideFileGit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitRepository** | [**ApplicationGitRepository**](ApplicationGitRepository.md) |  | 
**Paths** | **[]string** | List of path inside your git repository to locate values file. Must start by a / | 

## Methods

### NewHelmResponseAllOfValuesOverrideFileGit

`func NewHelmResponseAllOfValuesOverrideFileGit(gitRepository ApplicationGitRepository, paths []string, ) *HelmResponseAllOfValuesOverrideFileGit`

NewHelmResponseAllOfValuesOverrideFileGit instantiates a new HelmResponseAllOfValuesOverrideFileGit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmResponseAllOfValuesOverrideFileGitWithDefaults

`func NewHelmResponseAllOfValuesOverrideFileGitWithDefaults() *HelmResponseAllOfValuesOverrideFileGit`

NewHelmResponseAllOfValuesOverrideFileGitWithDefaults instantiates a new HelmResponseAllOfValuesOverrideFileGit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitRepository

`func (o *HelmResponseAllOfValuesOverrideFileGit) GetGitRepository() ApplicationGitRepository`

GetGitRepository returns the GitRepository field if non-nil, zero value otherwise.

### GetGitRepositoryOk

`func (o *HelmResponseAllOfValuesOverrideFileGit) GetGitRepositoryOk() (*ApplicationGitRepository, bool)`

GetGitRepositoryOk returns a tuple with the GitRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepository

`func (o *HelmResponseAllOfValuesOverrideFileGit) SetGitRepository(v ApplicationGitRepository)`

SetGitRepository sets GitRepository field to given value.


### GetPaths

`func (o *HelmResponseAllOfValuesOverrideFileGit) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *HelmResponseAllOfValuesOverrideFileGit) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *HelmResponseAllOfValuesOverrideFileGit) SetPaths(v []string)`

SetPaths sets Paths field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


