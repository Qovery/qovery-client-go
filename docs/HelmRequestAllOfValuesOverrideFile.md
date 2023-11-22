# HelmRequestAllOfValuesOverrideFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Raw** | Pointer to [**NullableHelmRequestAllOfValuesOverrideFileRaw**](HelmRequestAllOfValuesOverrideFileRaw.md) |  | [optional] 
**GitRepository** | Pointer to [**HelmValuesGitRepositoryRequest**](HelmValuesGitRepositoryRequest.md) |  | [optional] 

## Methods

### NewHelmRequestAllOfValuesOverrideFile

`func NewHelmRequestAllOfValuesOverrideFile() *HelmRequestAllOfValuesOverrideFile`

NewHelmRequestAllOfValuesOverrideFile instantiates a new HelmRequestAllOfValuesOverrideFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmRequestAllOfValuesOverrideFileWithDefaults

`func NewHelmRequestAllOfValuesOverrideFileWithDefaults() *HelmRequestAllOfValuesOverrideFile`

NewHelmRequestAllOfValuesOverrideFileWithDefaults instantiates a new HelmRequestAllOfValuesOverrideFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRaw

`func (o *HelmRequestAllOfValuesOverrideFile) GetRaw() HelmRequestAllOfValuesOverrideFileRaw`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *HelmRequestAllOfValuesOverrideFile) GetRawOk() (*HelmRequestAllOfValuesOverrideFileRaw, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *HelmRequestAllOfValuesOverrideFile) SetRaw(v HelmRequestAllOfValuesOverrideFileRaw)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *HelmRequestAllOfValuesOverrideFile) HasRaw() bool`

HasRaw returns a boolean if a field has been set.

### SetRawNil

`func (o *HelmRequestAllOfValuesOverrideFile) SetRawNil(b bool)`

 SetRawNil sets the value for Raw to be an explicit nil

### UnsetRaw
`func (o *HelmRequestAllOfValuesOverrideFile) UnsetRaw()`

UnsetRaw ensures that no value is present for Raw, not even an explicit nil
### GetGitRepository

`func (o *HelmRequestAllOfValuesOverrideFile) GetGitRepository() HelmValuesGitRepositoryRequest`

GetGitRepository returns the GitRepository field if non-nil, zero value otherwise.

### GetGitRepositoryOk

`func (o *HelmRequestAllOfValuesOverrideFile) GetGitRepositoryOk() (*HelmValuesGitRepositoryRequest, bool)`

GetGitRepositoryOk returns a tuple with the GitRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepository

`func (o *HelmRequestAllOfValuesOverrideFile) SetGitRepository(v HelmValuesGitRepositoryRequest)`

SetGitRepository sets GitRepository field to given value.

### HasGitRepository

`func (o *HelmRequestAllOfValuesOverrideFile) HasGitRepository() bool`

HasGitRepository returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


