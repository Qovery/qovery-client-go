# JobDeployRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageTag** | Pointer to **string** | Image tag to deploy.   Cannot be set if &#x60;git_commit_id&#x60; is defined  | [optional] 
**GitCommitId** | Pointer to **string** | Commit to deploy Cannot be set if &#x60;image_tag&#x60; is defined  | [optional] 

## Methods

### NewJobDeployRequest

`func NewJobDeployRequest() *JobDeployRequest`

NewJobDeployRequest instantiates a new JobDeployRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobDeployRequestWithDefaults

`func NewJobDeployRequestWithDefaults() *JobDeployRequest`

NewJobDeployRequestWithDefaults instantiates a new JobDeployRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageTag

`func (o *JobDeployRequest) GetImageTag() string`

GetImageTag returns the ImageTag field if non-nil, zero value otherwise.

### GetImageTagOk

`func (o *JobDeployRequest) GetImageTagOk() (*string, bool)`

GetImageTagOk returns a tuple with the ImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTag

`func (o *JobDeployRequest) SetImageTag(v string)`

SetImageTag sets ImageTag field to given value.

### HasImageTag

`func (o *JobDeployRequest) HasImageTag() bool`

HasImageTag returns a boolean if a field has been set.

### GetGitCommitId

`func (o *JobDeployRequest) GetGitCommitId() string`

GetGitCommitId returns the GitCommitId field if non-nil, zero value otherwise.

### GetGitCommitIdOk

`func (o *JobDeployRequest) GetGitCommitIdOk() (*string, bool)`

GetGitCommitIdOk returns a tuple with the GitCommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCommitId

`func (o *JobDeployRequest) SetGitCommitId(v string)`

SetGitCommitId sets GitCommitId field to given value.

### HasGitCommitId

`func (o *JobDeployRequest) HasGitCommitId() bool`

HasGitCommitId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


