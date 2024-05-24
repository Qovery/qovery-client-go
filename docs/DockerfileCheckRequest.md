# DockerfileCheckRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitRepository** | [**GitRepository**](GitRepository.md) |  | 
**DockerfilePath** | **string** | path of the dockerfile with root_path as base path | 

## Methods

### NewDockerfileCheckRequest

`func NewDockerfileCheckRequest(gitRepository GitRepository, dockerfilePath string, ) *DockerfileCheckRequest`

NewDockerfileCheckRequest instantiates a new DockerfileCheckRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerfileCheckRequestWithDefaults

`func NewDockerfileCheckRequestWithDefaults() *DockerfileCheckRequest`

NewDockerfileCheckRequestWithDefaults instantiates a new DockerfileCheckRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitRepository

`func (o *DockerfileCheckRequest) GetGitRepository() GitRepository`

GetGitRepository returns the GitRepository field if non-nil, zero value otherwise.

### GetGitRepositoryOk

`func (o *DockerfileCheckRequest) GetGitRepositoryOk() (*GitRepository, bool)`

GetGitRepositoryOk returns a tuple with the GitRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepository

`func (o *DockerfileCheckRequest) SetGitRepository(v GitRepository)`

SetGitRepository sets GitRepository field to given value.


### GetDockerfilePath

`func (o *DockerfileCheckRequest) GetDockerfilePath() string`

GetDockerfilePath returns the DockerfilePath field if non-nil, zero value otherwise.

### GetDockerfilePathOk

`func (o *DockerfileCheckRequest) GetDockerfilePathOk() (*string, bool)`

GetDockerfilePathOk returns a tuple with the DockerfilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfilePath

`func (o *DockerfileCheckRequest) SetDockerfilePath(v string)`

SetDockerfilePath sets DockerfilePath field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


