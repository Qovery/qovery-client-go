# DeployAllRequestContainersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | id of the container to be updated. | 
**ImageTag** | Pointer to **string** | new tag for the container. Can be empty only if the service has been already deployed (in this case the service version won&#39;t be changed)  For a container with a Git repository source: - an absent or empty &#x60;image_tag&#x60;, or one equal to the &#x60;tag&#x60; that &#x60;GET /container/{containerId}&#x60; returns, deploys the container at &#x60;git_commit_id&#x60;, or at its stored commit when &#x60;git_commit_id&#x60; is absent or empty; - a different &#x60;image_tag&#x60; returns 400.  | [optional] 
**GitCommitId** | Pointer to **string** | Commit to deploy, for a container with a Git repository source. When absent or empty, the container&#39;s stored commit is deployed. On a container with a Git repository source, a value that is not a 40-character SHA-1 commit id returns 400. Ignored for a container with an image source.  | [optional] 

## Methods

### NewDeployAllRequestContainersInner

`func NewDeployAllRequestContainersInner(id string, ) *DeployAllRequestContainersInner`

NewDeployAllRequestContainersInner instantiates a new DeployAllRequestContainersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeployAllRequestContainersInnerWithDefaults

`func NewDeployAllRequestContainersInnerWithDefaults() *DeployAllRequestContainersInner`

NewDeployAllRequestContainersInnerWithDefaults instantiates a new DeployAllRequestContainersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeployAllRequestContainersInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeployAllRequestContainersInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeployAllRequestContainersInner) SetId(v string)`

SetId sets Id field to given value.


### GetImageTag

`func (o *DeployAllRequestContainersInner) GetImageTag() string`

GetImageTag returns the ImageTag field if non-nil, zero value otherwise.

### GetImageTagOk

`func (o *DeployAllRequestContainersInner) GetImageTagOk() (*string, bool)`

GetImageTagOk returns a tuple with the ImageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTag

`func (o *DeployAllRequestContainersInner) SetImageTag(v string)`

SetImageTag sets ImageTag field to given value.

### HasImageTag

`func (o *DeployAllRequestContainersInner) HasImageTag() bool`

HasImageTag returns a boolean if a field has been set.

### GetGitCommitId

`func (o *DeployAllRequestContainersInner) GetGitCommitId() string`

GetGitCommitId returns the GitCommitId field if non-nil, zero value otherwise.

### GetGitCommitIdOk

`func (o *DeployAllRequestContainersInner) GetGitCommitIdOk() (*string, bool)`

GetGitCommitIdOk returns a tuple with the GitCommitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCommitId

`func (o *DeployAllRequestContainersInner) SetGitCommitId(v string)`

SetGitCommitId sets GitCommitId field to given value.

### HasGitCommitId

`func (o *DeployAllRequestContainersInner) HasGitCommitId() bool`

HasGitCommitId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


