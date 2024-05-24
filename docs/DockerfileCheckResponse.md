# DockerfileCheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DockerfilePath** | **string** |  | 
**Arg** | Pointer to **[]string** | All ARG variable declared in the Dockerfile | [optional] 
**Repositories** | Pointer to **[]string** | All image repositories we found declared in the Dockerfile | [optional] 

## Methods

### NewDockerfileCheckResponse

`func NewDockerfileCheckResponse(dockerfilePath string, ) *DockerfileCheckResponse`

NewDockerfileCheckResponse instantiates a new DockerfileCheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDockerfileCheckResponseWithDefaults

`func NewDockerfileCheckResponseWithDefaults() *DockerfileCheckResponse`

NewDockerfileCheckResponseWithDefaults instantiates a new DockerfileCheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDockerfilePath

`func (o *DockerfileCheckResponse) GetDockerfilePath() string`

GetDockerfilePath returns the DockerfilePath field if non-nil, zero value otherwise.

### GetDockerfilePathOk

`func (o *DockerfileCheckResponse) GetDockerfilePathOk() (*string, bool)`

GetDockerfilePathOk returns a tuple with the DockerfilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfilePath

`func (o *DockerfileCheckResponse) SetDockerfilePath(v string)`

SetDockerfilePath sets DockerfilePath field to given value.


### GetArg

`func (o *DockerfileCheckResponse) GetArg() []string`

GetArg returns the Arg field if non-nil, zero value otherwise.

### GetArgOk

`func (o *DockerfileCheckResponse) GetArgOk() (*[]string, bool)`

GetArgOk returns a tuple with the Arg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArg

`func (o *DockerfileCheckResponse) SetArg(v []string)`

SetArg sets Arg field to given value.

### HasArg

`func (o *DockerfileCheckResponse) HasArg() bool`

HasArg returns a boolean if a field has been set.

### GetRepositories

`func (o *DockerfileCheckResponse) GetRepositories() []string`

GetRepositories returns the Repositories field if non-nil, zero value otherwise.

### GetRepositoriesOk

`func (o *DockerfileCheckResponse) GetRepositoriesOk() (*[]string, bool)`

GetRepositoriesOk returns a tuple with the Repositories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepositories

`func (o *DockerfileCheckResponse) SetRepositories(v []string)`

SetRepositories sets Repositories field to given value.

### HasRepositories

`func (o *DockerfileCheckResponse) HasRepositories() bool`

HasRepositories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


