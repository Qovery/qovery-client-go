# HelmRequestAllOfValuesOverrideFileGit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitRepository** | Pointer to [**ApplicationGitRepositoryRequest**](ApplicationGitRepositoryRequest.md) |  | [optional] 
**Paths** | Pointer to **[]string** | List of path inside your git repository to locate values file. Must start by a / | [optional] 

## Methods

### NewHelmRequestAllOfValuesOverrideFileGit

`func NewHelmRequestAllOfValuesOverrideFileGit() *HelmRequestAllOfValuesOverrideFileGit`

NewHelmRequestAllOfValuesOverrideFileGit instantiates a new HelmRequestAllOfValuesOverrideFileGit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmRequestAllOfValuesOverrideFileGitWithDefaults

`func NewHelmRequestAllOfValuesOverrideFileGitWithDefaults() *HelmRequestAllOfValuesOverrideFileGit`

NewHelmRequestAllOfValuesOverrideFileGitWithDefaults instantiates a new HelmRequestAllOfValuesOverrideFileGit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitRepository

`func (o *HelmRequestAllOfValuesOverrideFileGit) GetGitRepository() ApplicationGitRepositoryRequest`

GetGitRepository returns the GitRepository field if non-nil, zero value otherwise.

### GetGitRepositoryOk

`func (o *HelmRequestAllOfValuesOverrideFileGit) GetGitRepositoryOk() (*ApplicationGitRepositoryRequest, bool)`

GetGitRepositoryOk returns a tuple with the GitRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepository

`func (o *HelmRequestAllOfValuesOverrideFileGit) SetGitRepository(v ApplicationGitRepositoryRequest)`

SetGitRepository sets GitRepository field to given value.

### HasGitRepository

`func (o *HelmRequestAllOfValuesOverrideFileGit) HasGitRepository() bool`

HasGitRepository returns a boolean if a field has been set.

### GetPaths

`func (o *HelmRequestAllOfValuesOverrideFileGit) GetPaths() []string`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *HelmRequestAllOfValuesOverrideFileGit) GetPathsOk() (*[]string, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *HelmRequestAllOfValuesOverrideFileGit) SetPaths(v []string)`

SetPaths sets Paths field to given value.

### HasPaths

`func (o *HelmRequestAllOfValuesOverrideFileGit) HasPaths() bool`

HasPaths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


