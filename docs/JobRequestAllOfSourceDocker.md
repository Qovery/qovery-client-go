# JobRequestAllOfSourceDocker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DockerfilePath** | Pointer to **NullableString** | The path of the associated Dockerfile. Only if you are using build_mode &#x3D; DOCKER | [optional] 
**GitRepository** | Pointer to [**ApplicationGitRepositoryRequest**](ApplicationGitRepositoryRequest.md) |  | [optional] 
**DockerfileRaw** | Pointer to **NullableString** | The content of your dockerfile if it is not stored inside your git repository | [optional] 

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

### SetDockerfilePathNil

`func (o *JobRequestAllOfSourceDocker) SetDockerfilePathNil(b bool)`

 SetDockerfilePathNil sets the value for DockerfilePath to be an explicit nil

### UnsetDockerfilePath
`func (o *JobRequestAllOfSourceDocker) UnsetDockerfilePath()`

UnsetDockerfilePath ensures that no value is present for DockerfilePath, not even an explicit nil
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

### SetDockerfileRawNil

`func (o *JobRequestAllOfSourceDocker) SetDockerfileRawNil(b bool)`

 SetDockerfileRawNil sets the value for DockerfileRaw to be an explicit nil

### UnsetDockerfileRaw
`func (o *JobRequestAllOfSourceDocker) UnsetDockerfileRaw()`

UnsetDockerfileRaw ensures that no value is present for DockerfileRaw, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


