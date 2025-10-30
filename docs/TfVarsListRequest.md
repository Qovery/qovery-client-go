# TfVarsListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitRepository** | [**GitRepositoryRequest**](GitRepositoryRequest.md) |  | 
**Mode** | [**TfVarsDiscoveryMode**](TfVarsDiscoveryMode.md) |  | 

## Methods

### NewTfVarsListRequest

`func NewTfVarsListRequest(gitRepository GitRepositoryRequest, mode TfVarsDiscoveryMode, ) *TfVarsListRequest`

NewTfVarsListRequest instantiates a new TfVarsListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTfVarsListRequestWithDefaults

`func NewTfVarsListRequestWithDefaults() *TfVarsListRequest`

NewTfVarsListRequestWithDefaults instantiates a new TfVarsListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitRepository

`func (o *TfVarsListRequest) GetGitRepository() GitRepositoryRequest`

GetGitRepository returns the GitRepository field if non-nil, zero value otherwise.

### GetGitRepositoryOk

`func (o *TfVarsListRequest) GetGitRepositoryOk() (*GitRepositoryRequest, bool)`

GetGitRepositoryOk returns a tuple with the GitRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepository

`func (o *TfVarsListRequest) SetGitRepository(v GitRepositoryRequest)`

SetGitRepository sets GitRepository field to given value.


### GetMode

`func (o *TfVarsListRequest) GetMode() TfVarsDiscoveryMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *TfVarsListRequest) GetModeOk() (*TfVarsDiscoveryMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *TfVarsListRequest) SetMode(v TfVarsDiscoveryMode)`

SetMode sets Mode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


