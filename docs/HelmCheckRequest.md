# HelmCheckRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitRepository** | [**HelmGitRepositoryRequest**](HelmGitRepositoryRequest.md) |  | 

## Methods

### NewHelmCheckRequest

`func NewHelmCheckRequest(gitRepository HelmGitRepositoryRequest, ) *HelmCheckRequest`

NewHelmCheckRequest instantiates a new HelmCheckRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmCheckRequestWithDefaults

`func NewHelmCheckRequestWithDefaults() *HelmCheckRequest`

NewHelmCheckRequestWithDefaults instantiates a new HelmCheckRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitRepository

`func (o *HelmCheckRequest) GetGitRepository() HelmGitRepositoryRequest`

GetGitRepository returns the GitRepository field if non-nil, zero value otherwise.

### GetGitRepositoryOk

`func (o *HelmCheckRequest) GetGitRepositoryOk() (*HelmGitRepositoryRequest, bool)`

GetGitRepositoryOk returns a tuple with the GitRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepository

`func (o *HelmCheckRequest) SetGitRepository(v HelmGitRepositoryRequest)`

SetGitRepository sets GitRepository field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


