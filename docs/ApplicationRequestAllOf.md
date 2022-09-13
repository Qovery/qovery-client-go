# ApplicationRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name is case insensitive | 
**Description** | Pointer to **NullableString** | give a description to this application | [optional] 
**GitRepository** | [**ApplicationGitRepositoryRequest**](ApplicationGitRepositoryRequest.md) |  | 
**BuildMode** | Pointer to [**BuildModeEnum**](BuildModeEnum.md) |  | [optional] [default to BUILDMODEENUM_BUILDPACKS]
**DockerfilePath** | Pointer to **NullableString** | The path of the associated Dockerfile. Only if you are using build_mode &#x3D; DOCKER | [optional] 
**BuildpackLanguage** | Pointer to [**NullableBuildPackLanguageEnum**](BuildPackLanguageEnum.md) |  | [optional] 
**Cpu** | Pointer to **int32** | unit is millicores (m). 1000m &#x3D; 1 cpu | [optional] [default to 500]
**Memory** | Pointer to **int32** | unit is MB. 1024 MB &#x3D; 1GB | [optional] [default to 512]
**MinRunningInstances** | Pointer to **int32** | Minimum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: 0 means that there is no application running.  | [optional] [default to 1]
**MaxRunningInstances** | Pointer to **int32** | Maximum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: -1 means that there is no limit.  | [optional] [default to 1]
**Healthcheck** | Pointer to [**Healthcheck**](Healthcheck.md) |  | [optional] 
**AutoPreview** | Pointer to **bool** | Specify if the environment preview option is activated or not for this application. If activated, a preview environment will be automatically cloned at each pull request.  | [optional] [default to true]
**Ports** | Pointer to [**[]ServicePort**](ServicePort.md) |  | [optional] 

## Methods

### NewApplicationRequestAllOf

`func NewApplicationRequestAllOf(name string, gitRepository ApplicationGitRepositoryRequest, ) *ApplicationRequestAllOf`

NewApplicationRequestAllOf instantiates a new ApplicationRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRequestAllOfWithDefaults

`func NewApplicationRequestAllOfWithDefaults() *ApplicationRequestAllOf`

NewApplicationRequestAllOfWithDefaults instantiates a new ApplicationRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApplicationRequestAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationRequestAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationRequestAllOf) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ApplicationRequestAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationRequestAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationRequestAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationRequestAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ApplicationRequestAllOf) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ApplicationRequestAllOf) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetGitRepository

`func (o *ApplicationRequestAllOf) GetGitRepository() ApplicationGitRepositoryRequest`

GetGitRepository returns the GitRepository field if non-nil, zero value otherwise.

### GetGitRepositoryOk

`func (o *ApplicationRequestAllOf) GetGitRepositoryOk() (*ApplicationGitRepositoryRequest, bool)`

GetGitRepositoryOk returns a tuple with the GitRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepository

`func (o *ApplicationRequestAllOf) SetGitRepository(v ApplicationGitRepositoryRequest)`

SetGitRepository sets GitRepository field to given value.


### GetBuildMode

`func (o *ApplicationRequestAllOf) GetBuildMode() BuildModeEnum`

GetBuildMode returns the BuildMode field if non-nil, zero value otherwise.

### GetBuildModeOk

`func (o *ApplicationRequestAllOf) GetBuildModeOk() (*BuildModeEnum, bool)`

GetBuildModeOk returns a tuple with the BuildMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildMode

`func (o *ApplicationRequestAllOf) SetBuildMode(v BuildModeEnum)`

SetBuildMode sets BuildMode field to given value.

### HasBuildMode

`func (o *ApplicationRequestAllOf) HasBuildMode() bool`

HasBuildMode returns a boolean if a field has been set.

### GetDockerfilePath

`func (o *ApplicationRequestAllOf) GetDockerfilePath() string`

GetDockerfilePath returns the DockerfilePath field if non-nil, zero value otherwise.

### GetDockerfilePathOk

`func (o *ApplicationRequestAllOf) GetDockerfilePathOk() (*string, bool)`

GetDockerfilePathOk returns a tuple with the DockerfilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfilePath

`func (o *ApplicationRequestAllOf) SetDockerfilePath(v string)`

SetDockerfilePath sets DockerfilePath field to given value.

### HasDockerfilePath

`func (o *ApplicationRequestAllOf) HasDockerfilePath() bool`

HasDockerfilePath returns a boolean if a field has been set.

### SetDockerfilePathNil

`func (o *ApplicationRequestAllOf) SetDockerfilePathNil(b bool)`

 SetDockerfilePathNil sets the value for DockerfilePath to be an explicit nil

### UnsetDockerfilePath
`func (o *ApplicationRequestAllOf) UnsetDockerfilePath()`

UnsetDockerfilePath ensures that no value is present for DockerfilePath, not even an explicit nil
### GetBuildpackLanguage

`func (o *ApplicationRequestAllOf) GetBuildpackLanguage() BuildPackLanguageEnum`

GetBuildpackLanguage returns the BuildpackLanguage field if non-nil, zero value otherwise.

### GetBuildpackLanguageOk

`func (o *ApplicationRequestAllOf) GetBuildpackLanguageOk() (*BuildPackLanguageEnum, bool)`

GetBuildpackLanguageOk returns a tuple with the BuildpackLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildpackLanguage

`func (o *ApplicationRequestAllOf) SetBuildpackLanguage(v BuildPackLanguageEnum)`

SetBuildpackLanguage sets BuildpackLanguage field to given value.

### HasBuildpackLanguage

`func (o *ApplicationRequestAllOf) HasBuildpackLanguage() bool`

HasBuildpackLanguage returns a boolean if a field has been set.

### SetBuildpackLanguageNil

`func (o *ApplicationRequestAllOf) SetBuildpackLanguageNil(b bool)`

 SetBuildpackLanguageNil sets the value for BuildpackLanguage to be an explicit nil

### UnsetBuildpackLanguage
`func (o *ApplicationRequestAllOf) UnsetBuildpackLanguage()`

UnsetBuildpackLanguage ensures that no value is present for BuildpackLanguage, not even an explicit nil
### GetCpu

`func (o *ApplicationRequestAllOf) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ApplicationRequestAllOf) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ApplicationRequestAllOf) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ApplicationRequestAllOf) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *ApplicationRequestAllOf) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ApplicationRequestAllOf) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ApplicationRequestAllOf) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ApplicationRequestAllOf) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetMinRunningInstances

`func (o *ApplicationRequestAllOf) GetMinRunningInstances() int32`

GetMinRunningInstances returns the MinRunningInstances field if non-nil, zero value otherwise.

### GetMinRunningInstancesOk

`func (o *ApplicationRequestAllOf) GetMinRunningInstancesOk() (*int32, bool)`

GetMinRunningInstancesOk returns a tuple with the MinRunningInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRunningInstances

`func (o *ApplicationRequestAllOf) SetMinRunningInstances(v int32)`

SetMinRunningInstances sets MinRunningInstances field to given value.

### HasMinRunningInstances

`func (o *ApplicationRequestAllOf) HasMinRunningInstances() bool`

HasMinRunningInstances returns a boolean if a field has been set.

### GetMaxRunningInstances

`func (o *ApplicationRequestAllOf) GetMaxRunningInstances() int32`

GetMaxRunningInstances returns the MaxRunningInstances field if non-nil, zero value otherwise.

### GetMaxRunningInstancesOk

`func (o *ApplicationRequestAllOf) GetMaxRunningInstancesOk() (*int32, bool)`

GetMaxRunningInstancesOk returns a tuple with the MaxRunningInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRunningInstances

`func (o *ApplicationRequestAllOf) SetMaxRunningInstances(v int32)`

SetMaxRunningInstances sets MaxRunningInstances field to given value.

### HasMaxRunningInstances

`func (o *ApplicationRequestAllOf) HasMaxRunningInstances() bool`

HasMaxRunningInstances returns a boolean if a field has been set.

### GetHealthcheck

`func (o *ApplicationRequestAllOf) GetHealthcheck() Healthcheck`

GetHealthcheck returns the Healthcheck field if non-nil, zero value otherwise.

### GetHealthcheckOk

`func (o *ApplicationRequestAllOf) GetHealthcheckOk() (*Healthcheck, bool)`

GetHealthcheckOk returns a tuple with the Healthcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthcheck

`func (o *ApplicationRequestAllOf) SetHealthcheck(v Healthcheck)`

SetHealthcheck sets Healthcheck field to given value.

### HasHealthcheck

`func (o *ApplicationRequestAllOf) HasHealthcheck() bool`

HasHealthcheck returns a boolean if a field has been set.

### GetAutoPreview

`func (o *ApplicationRequestAllOf) GetAutoPreview() bool`

GetAutoPreview returns the AutoPreview field if non-nil, zero value otherwise.

### GetAutoPreviewOk

`func (o *ApplicationRequestAllOf) GetAutoPreviewOk() (*bool, bool)`

GetAutoPreviewOk returns a tuple with the AutoPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPreview

`func (o *ApplicationRequestAllOf) SetAutoPreview(v bool)`

SetAutoPreview sets AutoPreview field to given value.

### HasAutoPreview

`func (o *ApplicationRequestAllOf) HasAutoPreview() bool`

HasAutoPreview returns a boolean if a field has been set.

### GetPorts

`func (o *ApplicationRequestAllOf) GetPorts() []ServicePort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ApplicationRequestAllOf) GetPortsOk() (*[]ServicePort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ApplicationRequestAllOf) SetPorts(v []ServicePort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ApplicationRequestAllOf) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


