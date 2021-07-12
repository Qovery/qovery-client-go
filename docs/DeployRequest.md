# DeployRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitCommitId** | **string** | Commit ID to deploy | 

## Methods

### NewDeployRequest

`func NewDeployRequest(gitCommitId string, ) *DeployRequest`

NewDeployRequest instantiates a new DeployRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployRequestWithDefaults

`func NewDeployRequestWithDefaults() *DeployRequest`

NewDeployRequestWithDefaults instantiates a new DeployRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitCommitId

`func (o *DeployRequest) GetGitCommitId() string`

GetGitCommitId returns the GitCommitId field if non-nil, zero value otherwise.

### GetGitCommitIdOk

`func (o *DeployRequest) GetGitCommitIdOk() (*string, bool)`

GetGitCommitIdOk returns a tuple with the GitCommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCommitId

`func (o *DeployRequest) SetGitCommitId(v string)`

SetGitCommitId sets GitCommitId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


