# ApplicationEditRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Storage** | Pointer to [**[]ApplicationStorageResponseStorage**](ApplicationStorageResponseStorage.md) |  | [optional] 
**Ports** | Pointer to [**[]ApplicationPortResponsePorts**](ApplicationPortResponsePorts.md) |  | [optional] 
**Name** | Pointer to **string** | name is case insensitive | [optional] 
**Description** | Pointer to **string** | give a description to this application | [optional] 
**GitRepository** | Pointer to [**ApplicationGitRepositoryRequest**](ApplicationGitRepositoryRequest.md) |  | [optional] 
**BuildMode** | Pointer to **string** | &#x60;DOCKER&#x60; requires &#x60;dockerfile_path&#x60; &#x60;BUILDPACKS&#x60; does not require any &#x60;dockerfile_path&#x60;  | [optional] 
**DockerfilePath** | Pointer to **string** | The path of the associated Dockerfile | [optional] 
**BuildpackLanguage** | Pointer to [**NullableBuildPackLanguageEnum**](BuildPackLanguageEnum.md) |  | [optional] 
**Cpu** | Pointer to **int32** | unit is millicores (m). 1000m &#x3D; 1 cpu | [optional] [default to 250]
**Memory** | Pointer to **int32** | unit is MB. 1024 MB &#x3D; 1GB | [optional] [default to 256]
**MinRunningInstances** | Pointer to **int32** | Minimum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: 0 means that there is no application running.  | [optional] [default to 1]
**MaxRunningInstances** | Pointer to **int32** | Maximum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: -1 means that there is no limit.  | [optional] [default to 1]
**Healthcheck** | Pointer to [**Healthcheck**](Healthcheck.md) |  | [optional] 
**AutoPreview** | Pointer to **bool** | Specify if the environment preview option is activated or not for this application. If activated, a preview environment will be automatically cloned at each pull request.  | [optional] [default to true]
**StickySession** | Pointer to **bool** | Specify if the sticky session option (also called persistant session) is activated or not for this application. If activated, user will be redirected by the load balancer to the same instance each time he access to the application.  | [optional] [default to false]

## Methods

### NewApplicationEditRequest

`func NewApplicationEditRequest() *ApplicationEditRequest`

NewApplicationEditRequest instantiates a new ApplicationEditRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationEditRequestWithDefaults

`func NewApplicationEditRequestWithDefaults() *ApplicationEditRequest`

NewApplicationEditRequestWithDefaults instantiates a new ApplicationEditRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorage

`func (o *ApplicationEditRequest) GetStorage() []ApplicationStorageResponseStorage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ApplicationEditRequest) GetStorageOk() (*[]ApplicationStorageResponseStorage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ApplicationEditRequest) SetStorage(v []ApplicationStorageResponseStorage)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *ApplicationEditRequest) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetPorts

`func (o *ApplicationEditRequest) GetPorts() []ApplicationPortResponsePorts`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ApplicationEditRequest) GetPortsOk() (*[]ApplicationPortResponsePorts, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ApplicationEditRequest) SetPorts(v []ApplicationPortResponsePorts)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ApplicationEditRequest) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetName

`func (o *ApplicationEditRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationEditRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationEditRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationEditRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationEditRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationEditRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationEditRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationEditRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGitRepository

`func (o *ApplicationEditRequest) GetGitRepository() ApplicationGitRepositoryRequest`

GetGitRepository returns the GitRepository field if non-nil, zero value otherwise.

### GetGitRepositoryOk

`func (o *ApplicationEditRequest) GetGitRepositoryOk() (*ApplicationGitRepositoryRequest, bool)`

GetGitRepositoryOk returns a tuple with the GitRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepository

`func (o *ApplicationEditRequest) SetGitRepository(v ApplicationGitRepositoryRequest)`

SetGitRepository sets GitRepository field to given value.

### HasGitRepository

`func (o *ApplicationEditRequest) HasGitRepository() bool`

HasGitRepository returns a boolean if a field has been set.

### GetBuildMode

`func (o *ApplicationEditRequest) GetBuildMode() string`

GetBuildMode returns the BuildMode field if non-nil, zero value otherwise.

### GetBuildModeOk

`func (o *ApplicationEditRequest) GetBuildModeOk() (*string, bool)`

GetBuildModeOk returns a tuple with the BuildMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildMode

`func (o *ApplicationEditRequest) SetBuildMode(v string)`

SetBuildMode sets BuildMode field to given value.

### HasBuildMode

`func (o *ApplicationEditRequest) HasBuildMode() bool`

HasBuildMode returns a boolean if a field has been set.

### GetDockerfilePath

`func (o *ApplicationEditRequest) GetDockerfilePath() string`

GetDockerfilePath returns the DockerfilePath field if non-nil, zero value otherwise.

### GetDockerfilePathOk

`func (o *ApplicationEditRequest) GetDockerfilePathOk() (*string, bool)`

GetDockerfilePathOk returns a tuple with the DockerfilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfilePath

`func (o *ApplicationEditRequest) SetDockerfilePath(v string)`

SetDockerfilePath sets DockerfilePath field to given value.

### HasDockerfilePath

`func (o *ApplicationEditRequest) HasDockerfilePath() bool`

HasDockerfilePath returns a boolean if a field has been set.

### GetBuildpackLanguage

`func (o *ApplicationEditRequest) GetBuildpackLanguage() BuildPackLanguageEnum`

GetBuildpackLanguage returns the BuildpackLanguage field if non-nil, zero value otherwise.

### GetBuildpackLanguageOk

`func (o *ApplicationEditRequest) GetBuildpackLanguageOk() (*BuildPackLanguageEnum, bool)`

GetBuildpackLanguageOk returns a tuple with the BuildpackLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildpackLanguage

`func (o *ApplicationEditRequest) SetBuildpackLanguage(v BuildPackLanguageEnum)`

SetBuildpackLanguage sets BuildpackLanguage field to given value.

### HasBuildpackLanguage

`func (o *ApplicationEditRequest) HasBuildpackLanguage() bool`

HasBuildpackLanguage returns a boolean if a field has been set.

### SetBuildpackLanguageNil

`func (o *ApplicationEditRequest) SetBuildpackLanguageNil(b bool)`

 SetBuildpackLanguageNil sets the value for BuildpackLanguage to be an explicit nil

### UnsetBuildpackLanguage
`func (o *ApplicationEditRequest) UnsetBuildpackLanguage()`

UnsetBuildpackLanguage ensures that no value is present for BuildpackLanguage, not even an explicit nil
### GetCpu

`func (o *ApplicationEditRequest) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ApplicationEditRequest) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ApplicationEditRequest) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ApplicationEditRequest) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *ApplicationEditRequest) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ApplicationEditRequest) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ApplicationEditRequest) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ApplicationEditRequest) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetMinRunningInstances

`func (o *ApplicationEditRequest) GetMinRunningInstances() int32`

GetMinRunningInstances returns the MinRunningInstances field if non-nil, zero value otherwise.

### GetMinRunningInstancesOk

`func (o *ApplicationEditRequest) GetMinRunningInstancesOk() (*int32, bool)`

GetMinRunningInstancesOk returns a tuple with the MinRunningInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRunningInstances

`func (o *ApplicationEditRequest) SetMinRunningInstances(v int32)`

SetMinRunningInstances sets MinRunningInstances field to given value.

### HasMinRunningInstances

`func (o *ApplicationEditRequest) HasMinRunningInstances() bool`

HasMinRunningInstances returns a boolean if a field has been set.

### GetMaxRunningInstances

`func (o *ApplicationEditRequest) GetMaxRunningInstances() int32`

GetMaxRunningInstances returns the MaxRunningInstances field if non-nil, zero value otherwise.

### GetMaxRunningInstancesOk

`func (o *ApplicationEditRequest) GetMaxRunningInstancesOk() (*int32, bool)`

GetMaxRunningInstancesOk returns a tuple with the MaxRunningInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRunningInstances

`func (o *ApplicationEditRequest) SetMaxRunningInstances(v int32)`

SetMaxRunningInstances sets MaxRunningInstances field to given value.

### HasMaxRunningInstances

`func (o *ApplicationEditRequest) HasMaxRunningInstances() bool`

HasMaxRunningInstances returns a boolean if a field has been set.

### GetHealthcheck

`func (o *ApplicationEditRequest) GetHealthcheck() Healthcheck`

GetHealthcheck returns the Healthcheck field if non-nil, zero value otherwise.

### GetHealthcheckOk

`func (o *ApplicationEditRequest) GetHealthcheckOk() (*Healthcheck, bool)`

GetHealthcheckOk returns a tuple with the Healthcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthcheck

`func (o *ApplicationEditRequest) SetHealthcheck(v Healthcheck)`

SetHealthcheck sets Healthcheck field to given value.

### HasHealthcheck

`func (o *ApplicationEditRequest) HasHealthcheck() bool`

HasHealthcheck returns a boolean if a field has been set.

### GetAutoPreview

`func (o *ApplicationEditRequest) GetAutoPreview() bool`

GetAutoPreview returns the AutoPreview field if non-nil, zero value otherwise.

### GetAutoPreviewOk

`func (o *ApplicationEditRequest) GetAutoPreviewOk() (*bool, bool)`

GetAutoPreviewOk returns a tuple with the AutoPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPreview

`func (o *ApplicationEditRequest) SetAutoPreview(v bool)`

SetAutoPreview sets AutoPreview field to given value.

### HasAutoPreview

`func (o *ApplicationEditRequest) HasAutoPreview() bool`

HasAutoPreview returns a boolean if a field has been set.

### GetStickySession

`func (o *ApplicationEditRequest) GetStickySession() bool`

GetStickySession returns the StickySession field if non-nil, zero value otherwise.

### GetStickySessionOk

`func (o *ApplicationEditRequest) GetStickySessionOk() (*bool, bool)`

GetStickySessionOk returns a tuple with the StickySession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickySession

`func (o *ApplicationEditRequest) SetStickySession(v bool)`

SetStickySession sets StickySession field to given value.

### HasStickySession

`func (o *ApplicationEditRequest) HasStickySession() bool`

HasStickySession returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


