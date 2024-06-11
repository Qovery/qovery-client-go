# HelmResponseAllOfSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Git** | [**HelmSourceGitResponse**](HelmSourceGitResponse.md) |  | 
**Repository** | [**HelmSourceRepositoryResponse**](HelmSourceRepositoryResponse.md) |  | 

## Methods

### NewHelmResponseAllOfSource

`func NewHelmResponseAllOfSource(git HelmSourceGitResponse, repository HelmSourceRepositoryResponse, ) *HelmResponseAllOfSource`

NewHelmResponseAllOfSource instantiates a new HelmResponseAllOfSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelmResponseAllOfSourceWithDefaults

`func NewHelmResponseAllOfSourceWithDefaults() *HelmResponseAllOfSource`

NewHelmResponseAllOfSourceWithDefaults instantiates a new HelmResponseAllOfSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGit

`func (o *HelmResponseAllOfSource) GetGit() HelmSourceGitResponse`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *HelmResponseAllOfSource) GetGitOk() (*HelmSourceGitResponse, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *HelmResponseAllOfSource) SetGit(v HelmSourceGitResponse)`

SetGit sets Git field to given value.


### GetRepository

`func (o *HelmResponseAllOfSource) GetRepository() HelmSourceRepositoryResponse`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *HelmResponseAllOfSource) GetRepositoryOk() (*HelmSourceRepositoryResponse, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *HelmResponseAllOfSource) SetRepository(v HelmSourceRepositoryResponse)`

SetRepository sets Repository field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


