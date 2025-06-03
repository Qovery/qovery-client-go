# JobSourceDockerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitRepository** | Pointer to [**ApplicationGitRepository**](ApplicationGitRepository.md) |  | [optional] 
**DockerfilePath** | Pointer to **string** | The path of the associated Dockerfile. Only if you are using build_mode &#x3D; DOCKER | [optional] 
**DockerfileRaw** | Pointer to **string** | The content of your dockerfile if it is not stored inside your git repository | [optional] 
**DockerTargetBuildStage** | Pointer to **string** | The target build stage in the Dockerfile to build | [optional] 

## Methods

### NewJobSourceDockerResponse

`func NewJobSourceDockerResponse() *JobSourceDockerResponse`

NewJobSourceDockerResponse instantiates a new JobSourceDockerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobSourceDockerResponseWithDefaults

`func NewJobSourceDockerResponseWithDefaults() *JobSourceDockerResponse`

NewJobSourceDockerResponseWithDefaults instantiates a new JobSourceDockerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitRepository

`func (o *JobSourceDockerResponse) GetGitRepository() ApplicationGitRepository`

GetGitRepository returns the GitRepository field if non-nil, zero value otherwise.

### GetGitRepositoryOk

`func (o *JobSourceDockerResponse) GetGitRepositoryOk() (*ApplicationGitRepository, bool)`

GetGitRepositoryOk returns a tuple with the GitRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepository

`func (o *JobSourceDockerResponse) SetGitRepository(v ApplicationGitRepository)`

SetGitRepository sets GitRepository field to given value.

### HasGitRepository

`func (o *JobSourceDockerResponse) HasGitRepository() bool`

HasGitRepository returns a boolean if a field has been set.

### GetDockerfilePath

`func (o *JobSourceDockerResponse) GetDockerfilePath() string`

GetDockerfilePath returns the DockerfilePath field if non-nil, zero value otherwise.

### GetDockerfilePathOk

`func (o *JobSourceDockerResponse) GetDockerfilePathOk() (*string, bool)`

GetDockerfilePathOk returns a tuple with the DockerfilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfilePath

`func (o *JobSourceDockerResponse) SetDockerfilePath(v string)`

SetDockerfilePath sets DockerfilePath field to given value.

### HasDockerfilePath

`func (o *JobSourceDockerResponse) HasDockerfilePath() bool`

HasDockerfilePath returns a boolean if a field has been set.

### GetDockerfileRaw

`func (o *JobSourceDockerResponse) GetDockerfileRaw() string`

GetDockerfileRaw returns the DockerfileRaw field if non-nil, zero value otherwise.

### GetDockerfileRawOk

`func (o *JobSourceDockerResponse) GetDockerfileRawOk() (*string, bool)`

GetDockerfileRawOk returns a tuple with the DockerfileRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfileRaw

`func (o *JobSourceDockerResponse) SetDockerfileRaw(v string)`

SetDockerfileRaw sets DockerfileRaw field to given value.

### HasDockerfileRaw

`func (o *JobSourceDockerResponse) HasDockerfileRaw() bool`

HasDockerfileRaw returns a boolean if a field has been set.

### GetDockerTargetBuildStage

`func (o *JobSourceDockerResponse) GetDockerTargetBuildStage() string`

GetDockerTargetBuildStage returns the DockerTargetBuildStage field if non-nil, zero value otherwise.

### GetDockerTargetBuildStageOk

`func (o *JobSourceDockerResponse) GetDockerTargetBuildStageOk() (*string, bool)`

GetDockerTargetBuildStageOk returns a tuple with the DockerTargetBuildStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerTargetBuildStage

`func (o *JobSourceDockerResponse) SetDockerTargetBuildStage(v string)`

SetDockerTargetBuildStage sets DockerTargetBuildStage field to given value.

### HasDockerTargetBuildStage

`func (o *JobSourceDockerResponse) HasDockerTargetBuildStage() bool`

HasDockerTargetBuildStage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


