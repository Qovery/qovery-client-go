# BaseJobResponseAllOfSourceOneOf1Docker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DockerfilePath** | Pointer to **NullableString** | The path of the associated Dockerfile. Only if you are using build_mode &#x3D; DOCKER | [optional] 
**GitRepository** | Pointer to [**ApplicationGitRepository**](ApplicationGitRepository.md) |  | [optional] 

## Methods

### NewBaseJobResponseAllOfSourceOneOf1Docker

`func NewBaseJobResponseAllOfSourceOneOf1Docker() *BaseJobResponseAllOfSourceOneOf1Docker`

NewBaseJobResponseAllOfSourceOneOf1Docker instantiates a new BaseJobResponseAllOfSourceOneOf1Docker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseJobResponseAllOfSourceOneOf1DockerWithDefaults

`func NewBaseJobResponseAllOfSourceOneOf1DockerWithDefaults() *BaseJobResponseAllOfSourceOneOf1Docker`

NewBaseJobResponseAllOfSourceOneOf1DockerWithDefaults instantiates a new BaseJobResponseAllOfSourceOneOf1Docker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDockerfilePath

`func (o *BaseJobResponseAllOfSourceOneOf1Docker) GetDockerfilePath() string`

GetDockerfilePath returns the DockerfilePath field if non-nil, zero value otherwise.

### GetDockerfilePathOk

`func (o *BaseJobResponseAllOfSourceOneOf1Docker) GetDockerfilePathOk() (*string, bool)`

GetDockerfilePathOk returns a tuple with the DockerfilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfilePath

`func (o *BaseJobResponseAllOfSourceOneOf1Docker) SetDockerfilePath(v string)`

SetDockerfilePath sets DockerfilePath field to given value.

### HasDockerfilePath

`func (o *BaseJobResponseAllOfSourceOneOf1Docker) HasDockerfilePath() bool`

HasDockerfilePath returns a boolean if a field has been set.

### SetDockerfilePathNil

`func (o *BaseJobResponseAllOfSourceOneOf1Docker) SetDockerfilePathNil(b bool)`

 SetDockerfilePathNil sets the value for DockerfilePath to be an explicit nil

### UnsetDockerfilePath
`func (o *BaseJobResponseAllOfSourceOneOf1Docker) UnsetDockerfilePath()`

UnsetDockerfilePath ensures that no value is present for DockerfilePath, not even an explicit nil
### GetGitRepository

`func (o *BaseJobResponseAllOfSourceOneOf1Docker) GetGitRepository() ApplicationGitRepository`

GetGitRepository returns the GitRepository field if non-nil, zero value otherwise.

### GetGitRepositoryOk

`func (o *BaseJobResponseAllOfSourceOneOf1Docker) GetGitRepositoryOk() (*ApplicationGitRepository, bool)`

GetGitRepositoryOk returns a tuple with the GitRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepository

`func (o *BaseJobResponseAllOfSourceOneOf1Docker) SetGitRepository(v ApplicationGitRepository)`

SetGitRepository sets GitRepository field to given value.

### HasGitRepository

`func (o *BaseJobResponseAllOfSourceOneOf1Docker) HasGitRepository() bool`

HasGitRepository returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


