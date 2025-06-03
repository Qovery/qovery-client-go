# JobRequestAllOfSourceDocker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitRepository** | Pointer to [**ApplicationGitRepositoryRequest**](ApplicationGitRepositoryRequest.md) |  | [optional] 
**DockerfilePath** | Pointer to **string** | The path of the associated Dockerfile. Only if you are using build_mode &#x3D; DOCKER | [optional] 
**DockerfileRaw** | Pointer to **string** | The content of your dockerfile if it is not stored inside your git repository | [optional] 
**DockerTargetBuildStage** | Pointer to **string** | The target build stage in the Dockerfile to build | [optional] 

## Methods

### NewJobRequestAllOfSourceDocker

`func NewJobRequestAllOfSourceDocker() *JobRequestAllOfSourceDocker`

NewJobRequestAllOfSourceDocker instantiates a new JobRequestAllOfSourceDocker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobRequestAllOfSourceDockerWithDefaults

`func NewJobRequestAllOfSourceDockerWithDefaults() *JobRequestAllOfSourceDocker`

NewJobRequestAllOfSourceDockerWithDefaults instantiates a new JobRequestAllOfSourceDocker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitRepository

`func (o *JobRequestAllOfSourceDocker) GetGitRepository() ApplicationGitRepositoryRequest`

GetGitRepository returns the GitRepository field if non-nil, zero value otherwise.

### GetGitRepositoryOk

`func (o *JobRequestAllOfSourceDocker) GetGitRepositoryOk() (*ApplicationGitRepositoryRequest, bool)`

GetGitRepositoryOk returns a tuple with the GitRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepository

`func (o *JobRequestAllOfSourceDocker) SetGitRepository(v ApplicationGitRepositoryRequest)`

SetGitRepository sets GitRepository field to given value.

### HasGitRepository

`func (o *JobRequestAllOfSourceDocker) HasGitRepository() bool`

HasGitRepository returns a boolean if a field has been set.

### GetDockerfilePath

`func (o *JobRequestAllOfSourceDocker) GetDockerfilePath() string`

GetDockerfilePath returns the DockerfilePath field if non-nil, zero value otherwise.

### GetDockerfilePathOk

`func (o *JobRequestAllOfSourceDocker) GetDockerfilePathOk() (*string, bool)`

GetDockerfilePathOk returns a tuple with the DockerfilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfilePath

`func (o *JobRequestAllOfSourceDocker) SetDockerfilePath(v string)`

SetDockerfilePath sets DockerfilePath field to given value.

### HasDockerfilePath

`func (o *JobRequestAllOfSourceDocker) HasDockerfilePath() bool`

HasDockerfilePath returns a boolean if a field has been set.

### GetDockerfileRaw

`func (o *JobRequestAllOfSourceDocker) GetDockerfileRaw() string`

GetDockerfileRaw returns the DockerfileRaw field if non-nil, zero value otherwise.

### GetDockerfileRawOk

`func (o *JobRequestAllOfSourceDocker) GetDockerfileRawOk() (*string, bool)`

GetDockerfileRawOk returns a tuple with the DockerfileRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfileRaw

`func (o *JobRequestAllOfSourceDocker) SetDockerfileRaw(v string)`

SetDockerfileRaw sets DockerfileRaw field to given value.

### HasDockerfileRaw

`func (o *JobRequestAllOfSourceDocker) HasDockerfileRaw() bool`

HasDockerfileRaw returns a boolean if a field has been set.

### GetDockerTargetBuildStage

`func (o *JobRequestAllOfSourceDocker) GetDockerTargetBuildStage() string`

GetDockerTargetBuildStage returns the DockerTargetBuildStage field if non-nil, zero value otherwise.

### GetDockerTargetBuildStageOk

`func (o *JobRequestAllOfSourceDocker) GetDockerTargetBuildStageOk() (*string, bool)`

GetDockerTargetBuildStageOk returns a tuple with the DockerTargetBuildStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerTargetBuildStage

`func (o *JobRequestAllOfSourceDocker) SetDockerTargetBuildStage(v string)`

SetDockerTargetBuildStage sets DockerTargetBuildStage field to given value.

### HasDockerTargetBuildStage

`func (o *JobRequestAllOfSourceDocker) HasDockerTargetBuildStage() bool`

HasDockerTargetBuildStage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


