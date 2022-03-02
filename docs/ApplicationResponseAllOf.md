# ApplicationResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to [**ReferenceObject**](ReferenceObject.md) |  | [optional] 
**GitRepository** | Pointer to [**ApplicationGitRepositoryResponse**](ApplicationGitRepositoryResponse.md) |  | [optional] 
**MaximumCpu** | Pointer to **int32** | Maximum cpu that can be allocated to the application based on organization cluster configuration. unit is millicores (m). 1000m &#x3D; 1 cpu | [optional] [default to 250]
**MaximumMemory** | Pointer to **int32** | Maximum memory that can be allocated to the application based on organization cluster configuration. unit is MB. 1024 MB &#x3D; 1GB | [optional] [default to 256]
**Name** | Pointer to **string** | name is case insensitive | [optional] 
**Description** | Pointer to **NullableString** | give a description to this application | [optional] 
**BuildMode** | Pointer to **string** | &#x60;DOCKER&#x60; requires &#x60;dockerfile_path&#x60; &#x60;BUILDPACKS&#x60; does not require any &#x60;dockerfile_path&#x60;  | [optional] [default to "BUILDPACKS"]
**DockerfilePath** | Pointer to **NullableString** | The path of the associated Dockerfile. Only if you are using build_mode &#x3D; DOCKER | [optional] 
**BuildpackLanguage** | Pointer to [**NullableBuildPackLanguageEnum**](BuildPackLanguageEnum.md) |  | [optional] 
**Cpu** | Pointer to **int32** | unit is millicores (m). 1000m &#x3D; 1 cpu | [optional] [default to 250]
**Memory** | Pointer to **int32** | unit is MB. 1024 MB &#x3D; 1GB | [optional] [default to 256]
**MinRunningInstances** | Pointer to **int32** | Minimum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: 0 means that there is no application running.  | [optional] [default to 1]
**MaxRunningInstances** | Pointer to **int32** | Maximum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: -1 means that there is no limit.  | [optional] [default to 1]
**Healthcheck** | Pointer to [**Healthcheck**](Healthcheck.md) |  | [optional] 
**AutoPreview** | Pointer to **bool** | Specify if the environment preview option is activated or not for this application. If activated, a preview environment will be automatically cloned at each pull request.  | [optional] [default to true]

## Methods

### NewApplicationResponseAllOf

`func NewApplicationResponseAllOf() *ApplicationResponseAllOf`

NewApplicationResponseAllOf instantiates a new ApplicationResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationResponseAllOfWithDefaults

`func NewApplicationResponseAllOfWithDefaults() *ApplicationResponseAllOf`

NewApplicationResponseAllOfWithDefaults instantiates a new ApplicationResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *ApplicationResponseAllOf) GetEnvironment() ReferenceObject`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ApplicationResponseAllOf) GetEnvironmentOk() (*ReferenceObject, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ApplicationResponseAllOf) SetEnvironment(v ReferenceObject)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ApplicationResponseAllOf) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetGitRepository

`func (o *ApplicationResponseAllOf) GetGitRepository() ApplicationGitRepositoryResponse`

GetGitRepository returns the GitRepository field if non-nil, zero value otherwise.

### GetGitRepositoryOk

`func (o *ApplicationResponseAllOf) GetGitRepositoryOk() (*ApplicationGitRepositoryResponse, bool)`

GetGitRepositoryOk returns a tuple with the GitRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepository

`func (o *ApplicationResponseAllOf) SetGitRepository(v ApplicationGitRepositoryResponse)`

SetGitRepository sets GitRepository field to given value.

### HasGitRepository

`func (o *ApplicationResponseAllOf) HasGitRepository() bool`

HasGitRepository returns a boolean if a field has been set.

### GetMaximumCpu

`func (o *ApplicationResponseAllOf) GetMaximumCpu() int32`

GetMaximumCpu returns the MaximumCpu field if non-nil, zero value otherwise.

### GetMaximumCpuOk

`func (o *ApplicationResponseAllOf) GetMaximumCpuOk() (*int32, bool)`

GetMaximumCpuOk returns a tuple with the MaximumCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumCpu

`func (o *ApplicationResponseAllOf) SetMaximumCpu(v int32)`

SetMaximumCpu sets MaximumCpu field to given value.

### HasMaximumCpu

`func (o *ApplicationResponseAllOf) HasMaximumCpu() bool`

HasMaximumCpu returns a boolean if a field has been set.

### GetMaximumMemory

`func (o *ApplicationResponseAllOf) GetMaximumMemory() int32`

GetMaximumMemory returns the MaximumMemory field if non-nil, zero value otherwise.

### GetMaximumMemoryOk

`func (o *ApplicationResponseAllOf) GetMaximumMemoryOk() (*int32, bool)`

GetMaximumMemoryOk returns a tuple with the MaximumMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumMemory

`func (o *ApplicationResponseAllOf) SetMaximumMemory(v int32)`

SetMaximumMemory sets MaximumMemory field to given value.

### HasMaximumMemory

`func (o *ApplicationResponseAllOf) HasMaximumMemory() bool`

HasMaximumMemory returns a boolean if a field has been set.

### GetName

`func (o *ApplicationResponseAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationResponseAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationResponseAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationResponseAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationResponseAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationResponseAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationResponseAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationResponseAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ApplicationResponseAllOf) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ApplicationResponseAllOf) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBuildMode

`func (o *ApplicationResponseAllOf) GetBuildMode() string`

GetBuildMode returns the BuildMode field if non-nil, zero value otherwise.

### GetBuildModeOk

`func (o *ApplicationResponseAllOf) GetBuildModeOk() (*string, bool)`

GetBuildModeOk returns a tuple with the BuildMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildMode

`func (o *ApplicationResponseAllOf) SetBuildMode(v string)`

SetBuildMode sets BuildMode field to given value.

### HasBuildMode

`func (o *ApplicationResponseAllOf) HasBuildMode() bool`

HasBuildMode returns a boolean if a field has been set.

### GetDockerfilePath

`func (o *ApplicationResponseAllOf) GetDockerfilePath() string`

GetDockerfilePath returns the DockerfilePath field if non-nil, zero value otherwise.

### GetDockerfilePathOk

`func (o *ApplicationResponseAllOf) GetDockerfilePathOk() (*string, bool)`

GetDockerfilePathOk returns a tuple with the DockerfilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfilePath

`func (o *ApplicationResponseAllOf) SetDockerfilePath(v string)`

SetDockerfilePath sets DockerfilePath field to given value.

### HasDockerfilePath

`func (o *ApplicationResponseAllOf) HasDockerfilePath() bool`

HasDockerfilePath returns a boolean if a field has been set.

### SetDockerfilePathNil

`func (o *ApplicationResponseAllOf) SetDockerfilePathNil(b bool)`

 SetDockerfilePathNil sets the value for DockerfilePath to be an explicit nil

### UnsetDockerfilePath
`func (o *ApplicationResponseAllOf) UnsetDockerfilePath()`

UnsetDockerfilePath ensures that no value is present for DockerfilePath, not even an explicit nil
### GetBuildpackLanguage

`func (o *ApplicationResponseAllOf) GetBuildpackLanguage() BuildPackLanguageEnum`

GetBuildpackLanguage returns the BuildpackLanguage field if non-nil, zero value otherwise.

### GetBuildpackLanguageOk

`func (o *ApplicationResponseAllOf) GetBuildpackLanguageOk() (*BuildPackLanguageEnum, bool)`

GetBuildpackLanguageOk returns a tuple with the BuildpackLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildpackLanguage

`func (o *ApplicationResponseAllOf) SetBuildpackLanguage(v BuildPackLanguageEnum)`

SetBuildpackLanguage sets BuildpackLanguage field to given value.

### HasBuildpackLanguage

`func (o *ApplicationResponseAllOf) HasBuildpackLanguage() bool`

HasBuildpackLanguage returns a boolean if a field has been set.

### SetBuildpackLanguageNil

`func (o *ApplicationResponseAllOf) SetBuildpackLanguageNil(b bool)`

 SetBuildpackLanguageNil sets the value for BuildpackLanguage to be an explicit nil

### UnsetBuildpackLanguage
`func (o *ApplicationResponseAllOf) UnsetBuildpackLanguage()`

UnsetBuildpackLanguage ensures that no value is present for BuildpackLanguage, not even an explicit nil
### GetCpu

`func (o *ApplicationResponseAllOf) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ApplicationResponseAllOf) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ApplicationResponseAllOf) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ApplicationResponseAllOf) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *ApplicationResponseAllOf) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ApplicationResponseAllOf) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ApplicationResponseAllOf) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ApplicationResponseAllOf) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetMinRunningInstances

`func (o *ApplicationResponseAllOf) GetMinRunningInstances() int32`

GetMinRunningInstances returns the MinRunningInstances field if non-nil, zero value otherwise.

### GetMinRunningInstancesOk

`func (o *ApplicationResponseAllOf) GetMinRunningInstancesOk() (*int32, bool)`

GetMinRunningInstancesOk returns a tuple with the MinRunningInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRunningInstances

`func (o *ApplicationResponseAllOf) SetMinRunningInstances(v int32)`

SetMinRunningInstances sets MinRunningInstances field to given value.

### HasMinRunningInstances

`func (o *ApplicationResponseAllOf) HasMinRunningInstances() bool`

HasMinRunningInstances returns a boolean if a field has been set.

### GetMaxRunningInstances

`func (o *ApplicationResponseAllOf) GetMaxRunningInstances() int32`

GetMaxRunningInstances returns the MaxRunningInstances field if non-nil, zero value otherwise.

### GetMaxRunningInstancesOk

`func (o *ApplicationResponseAllOf) GetMaxRunningInstancesOk() (*int32, bool)`

GetMaxRunningInstancesOk returns a tuple with the MaxRunningInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRunningInstances

`func (o *ApplicationResponseAllOf) SetMaxRunningInstances(v int32)`

SetMaxRunningInstances sets MaxRunningInstances field to given value.

### HasMaxRunningInstances

`func (o *ApplicationResponseAllOf) HasMaxRunningInstances() bool`

HasMaxRunningInstances returns a boolean if a field has been set.

### GetHealthcheck

`func (o *ApplicationResponseAllOf) GetHealthcheck() Healthcheck`

GetHealthcheck returns the Healthcheck field if non-nil, zero value otherwise.

### GetHealthcheckOk

`func (o *ApplicationResponseAllOf) GetHealthcheckOk() (*Healthcheck, bool)`

GetHealthcheckOk returns a tuple with the Healthcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthcheck

`func (o *ApplicationResponseAllOf) SetHealthcheck(v Healthcheck)`

SetHealthcheck sets Healthcheck field to given value.

### HasHealthcheck

`func (o *ApplicationResponseAllOf) HasHealthcheck() bool`

HasHealthcheck returns a boolean if a field has been set.

### GetAutoPreview

`func (o *ApplicationResponseAllOf) GetAutoPreview() bool`

GetAutoPreview returns the AutoPreview field if non-nil, zero value otherwise.

### GetAutoPreviewOk

`func (o *ApplicationResponseAllOf) GetAutoPreviewOk() (*bool, bool)`

GetAutoPreviewOk returns a tuple with the AutoPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPreview

`func (o *ApplicationResponseAllOf) SetAutoPreview(v bool)`

SetAutoPreview sets AutoPreview field to given value.

### HasAutoPreview

`func (o *ApplicationResponseAllOf) HasAutoPreview() bool`

HasAutoPreview returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


