# DeployAllRequestJobsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | id of the job to be updated. | [optional] 
**ImageTag** | Pointer to **string** | new tag for the job image. Use only if job is an image source. Can be empty only if the service has been already deployed (in this case the service version won&#39;t be changed) | [optional] 
**GitCommitId** | Pointer to **string** | Commit ID to deploy. Use only if job is a repository source. Can be empty only if the service has been already deployed (in this case the service version won&#39;t be changed) | [optional] 

## Methods

### NewDeployAllRequestJobsInner

`func NewDeployAllRequestJobsInner() *DeployAllRequestJobsInner`

NewDeployAllRequestJobsInner instantiates a new DeployAllRequestJobsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployAllRequestJobsInnerWithDefaults

`func NewDeployAllRequestJobsInnerWithDefaults() *DeployAllRequestJobsInner`

NewDeployAllRequestJobsInnerWithDefaults instantiates a new DeployAllRequestJobsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeployAllRequestJobsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeployAllRequestJobsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeployAllRequestJobsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeployAllRequestJobsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageTag

`func (o *DeployAllRequestJobsInner) GetImageTag() string`

GetImageTag returns the ImageTag field if non-nil, zero value otherwise.

### GetImageTagOk

`func (o *DeployAllRequestJobsInner) GetImageTagOk() (*string, bool)`

GetImageTagOk returns a tuple with the ImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTag

`func (o *DeployAllRequestJobsInner) SetImageTag(v string)`

SetImageTag sets ImageTag field to given value.

### HasImageTag

`func (o *DeployAllRequestJobsInner) HasImageTag() bool`

HasImageTag returns a boolean if a field has been set.

### GetGitCommitId

`func (o *DeployAllRequestJobsInner) GetGitCommitId() string`

GetGitCommitId returns the GitCommitId field if non-nil, zero value otherwise.

### GetGitCommitIdOk

`func (o *DeployAllRequestJobsInner) GetGitCommitIdOk() (*string, bool)`

GetGitCommitIdOk returns a tuple with the GitCommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCommitId

`func (o *DeployAllRequestJobsInner) SetGitCommitId(v string)`

SetGitCommitId sets GitCommitId field to given value.

### HasGitCommitId

`func (o *DeployAllRequestJobsInner) HasGitCommitId() bool`

HasGitCommitId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


