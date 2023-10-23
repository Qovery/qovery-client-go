# DeployHelmRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | version of the chart to deploy. Cannot be set if &#x60;git_commit_id&#x60; is defined  | [optional] 
**GitCommitId** | Pointer to **string** | Commit to deploy Cannot be set if &#x60;version&#x60; is defined  | [optional] 

## Methods

### NewDeployHelmRequest

`func NewDeployHelmRequest() *DeployHelmRequest`

NewDeployHelmRequest instantiates a new DeployHelmRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployHelmRequestWithDefaults

`func NewDeployHelmRequestWithDefaults() *DeployHelmRequest`

NewDeployHelmRequestWithDefaults instantiates a new DeployHelmRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *DeployHelmRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeployHelmRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeployHelmRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DeployHelmRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetGitCommitId

`func (o *DeployHelmRequest) GetGitCommitId() string`

GetGitCommitId returns the GitCommitId field if non-nil, zero value otherwise.

### GetGitCommitIdOk

`func (o *DeployHelmRequest) GetGitCommitIdOk() (*string, bool)`

GetGitCommitIdOk returns a tuple with the GitCommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCommitId

`func (o *DeployHelmRequest) SetGitCommitId(v string)`

SetGitCommitId sets GitCommitId field to given value.

### HasGitCommitId

`func (o *DeployHelmRequest) HasGitCommitId() bool`

HasGitCommitId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


