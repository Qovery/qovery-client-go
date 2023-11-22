# HelmRequestAllOfSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitRepository** | Pointer to [**HelmGitRepositoryRequest**](HelmGitRepositoryRequest.md) |  | [optional] 
**HelmRepository** | Pointer to [**HelmRequestAllOfSourceOneOf1HelmRepository**](HelmRequestAllOfSourceOneOf1HelmRepository.md) |  | [optional] 

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

### GetGitRepository

`func (o *HelmRequestAllOfSource) GetGitRepository() HelmGitRepositoryRequest`

GetGitRepository returns the GitRepository field if non-nil, zero value otherwise.

### GetGitRepositoryOk

`func (o *HelmRequestAllOfSource) GetGitRepositoryOk() (*HelmGitRepositoryRequest, bool)`

GetGitRepositoryOk returns a tuple with the GitRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepository

`func (o *HelmRequestAllOfSource) SetGitRepository(v HelmGitRepositoryRequest)`

SetGitRepository sets GitRepository field to given value.

### HasGitRepository

`func (o *HelmRequestAllOfSource) HasGitRepository() bool`

HasGitRepository returns a boolean if a field has been set.

### GetHelmRepository

`func (o *HelmRequestAllOfSource) GetHelmRepository() HelmRequestAllOfSourceOneOf1HelmRepository`

GetHelmRepository returns the HelmRepository field if non-nil, zero value otherwise.

### GetHelmRepositoryOk

`func (o *HelmRequestAllOfSource) GetHelmRepositoryOk() (*HelmRequestAllOfSourceOneOf1HelmRepository, bool)`

GetHelmRepositoryOk returns a tuple with the HelmRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelmRepository

`func (o *HelmRequestAllOfSource) SetHelmRepository(v HelmRequestAllOfSourceOneOf1HelmRepository)`

SetHelmRepository sets HelmRepository field to given value.

### HasHelmRepository

`func (o *HelmRequestAllOfSource) HasHelmRepository() bool`

HasHelmRepository returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


