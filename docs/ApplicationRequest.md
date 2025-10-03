# ApplicationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Storage** | Pointer to [**[]ServiceStorageRequestStorageInner**](ServiceStorageRequestStorageInner.md) |  | [optional] 
**Ports** | Pointer to [**[]ServicePortRequestPortsInner**](ServicePortRequestPortsInner.md) |  | [optional] 
**Name** | **string** | name is case insensitive | 
**Description** | Pointer to **NullableString** | give a description to this application | [optional] 
**GitRepository** | [**ApplicationGitRepositoryRequest**](ApplicationGitRepositoryRequest.md) |  | 
**BuildMode** | Pointer to [**BuildModeEnum**](BuildModeEnum.md) |  | [optional] [default to BUILDMODEENUM_DOCKER]
**DockerfilePath** | Pointer to **NullableString** | The path of the associated Dockerfile. Only if you are using build_mode &#x3D; DOCKER | [optional] 
**Cpu** | Pointer to **int32** | unit is millicores (m). 1000m &#x3D; 1 cpu | [optional] [default to 500]
**Memory** | Pointer to **int32** | unit is MB. 1024 MB &#x3D; 1GB | [optional] [default to 512]
**Gpu** | Pointer to **int32** |  | [optional] [default to 0]
**MinRunningInstances** | Pointer to **int32** | Minimum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: 0 means that there is no application running.  | [optional] [default to 1]
**MaxRunningInstances** | Pointer to **int32** | Maximum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: -1 means that there is no limit.  | [optional] [default to 1]
**Healthchecks** | [**Healthcheck**](Healthcheck.md) |  | 
**AutoPreview** | Pointer to **bool** | Specify if the environment preview option is activated or not for this application.   If activated, a preview environment will be automatically cloned at each pull request.   If not specified, it takes the value of the &#x60;auto_preview&#x60; property from the associated environment.  | [optional] [default to true]
**Arguments** | Pointer to **[]string** |  | [optional] 
**Entrypoint** | Pointer to **string** | optional entrypoint when launching container | [optional] 
**AutoDeploy** | Pointer to **NullableBool** | Specify if the application will be automatically updated after receiving a new commit. | [optional] 
**AnnotationsGroups** | Pointer to [**[]ServiceAnnotationRequest**](ServiceAnnotationRequest.md) |  | [optional] 
**LabelsGroups** | Pointer to [**[]ServiceLabelRequest**](ServiceLabelRequest.md) |  | [optional] 
**IconUri** | Pointer to **string** | Icon URI representing the application. | [optional] 
**DockerTargetBuildStage** | Pointer to **NullableString** | The target build stage in the Dockerfile to build | [optional] 

## Methods

### NewApplicationRequest

`func NewApplicationRequest(name string, gitRepository ApplicationGitRepositoryRequest, healthchecks Healthcheck, ) *ApplicationRequest`

NewApplicationRequest instantiates a new ApplicationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRequestWithDefaults

`func NewApplicationRequestWithDefaults() *ApplicationRequest`

NewApplicationRequestWithDefaults instantiates a new ApplicationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorage

`func (o *ApplicationRequest) GetStorage() []ServiceStorageRequestStorageInner`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ApplicationRequest) GetStorageOk() (*[]ServiceStorageRequestStorageInner, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ApplicationRequest) SetStorage(v []ServiceStorageRequestStorageInner)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *ApplicationRequest) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetPorts

`func (o *ApplicationRequest) GetPorts() []ServicePortRequestPortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ApplicationRequest) GetPortsOk() (*[]ServicePortRequestPortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ApplicationRequest) SetPorts(v []ServicePortRequestPortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ApplicationRequest) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetName

`func (o *ApplicationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ApplicationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ApplicationRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ApplicationRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetGitRepository

`func (o *ApplicationRequest) GetGitRepository() ApplicationGitRepositoryRequest`

GetGitRepository returns the GitRepository field if non-nil, zero value otherwise.

### GetGitRepositoryOk

`func (o *ApplicationRequest) GetGitRepositoryOk() (*ApplicationGitRepositoryRequest, bool)`

GetGitRepositoryOk returns a tuple with the GitRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepository

`func (o *ApplicationRequest) SetGitRepository(v ApplicationGitRepositoryRequest)`

SetGitRepository sets GitRepository field to given value.


### GetBuildMode

`func (o *ApplicationRequest) GetBuildMode() BuildModeEnum`

GetBuildMode returns the BuildMode field if non-nil, zero value otherwise.

### GetBuildModeOk

`func (o *ApplicationRequest) GetBuildModeOk() (*BuildModeEnum, bool)`

GetBuildModeOk returns a tuple with the BuildMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildMode

`func (o *ApplicationRequest) SetBuildMode(v BuildModeEnum)`

SetBuildMode sets BuildMode field to given value.

### HasBuildMode

`func (o *ApplicationRequest) HasBuildMode() bool`

HasBuildMode returns a boolean if a field has been set.

### GetDockerfilePath

`func (o *ApplicationRequest) GetDockerfilePath() string`

GetDockerfilePath returns the DockerfilePath field if non-nil, zero value otherwise.

### GetDockerfilePathOk

`func (o *ApplicationRequest) GetDockerfilePathOk() (*string, bool)`

GetDockerfilePathOk returns a tuple with the DockerfilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfilePath

`func (o *ApplicationRequest) SetDockerfilePath(v string)`

SetDockerfilePath sets DockerfilePath field to given value.

### HasDockerfilePath

`func (o *ApplicationRequest) HasDockerfilePath() bool`

HasDockerfilePath returns a boolean if a field has been set.

### SetDockerfilePathNil

`func (o *ApplicationRequest) SetDockerfilePathNil(b bool)`

 SetDockerfilePathNil sets the value for DockerfilePath to be an explicit nil

### UnsetDockerfilePath
`func (o *ApplicationRequest) UnsetDockerfilePath()`

UnsetDockerfilePath ensures that no value is present for DockerfilePath, not even an explicit nil
### GetCpu

`func (o *ApplicationRequest) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ApplicationRequest) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ApplicationRequest) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ApplicationRequest) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *ApplicationRequest) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ApplicationRequest) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ApplicationRequest) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ApplicationRequest) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetGpu

`func (o *ApplicationRequest) GetGpu() int32`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *ApplicationRequest) GetGpuOk() (*int32, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *ApplicationRequest) SetGpu(v int32)`

SetGpu sets Gpu field to given value.

### HasGpu

`func (o *ApplicationRequest) HasGpu() bool`

HasGpu returns a boolean if a field has been set.

### GetMinRunningInstances

`func (o *ApplicationRequest) GetMinRunningInstances() int32`

GetMinRunningInstances returns the MinRunningInstances field if non-nil, zero value otherwise.

### GetMinRunningInstancesOk

`func (o *ApplicationRequest) GetMinRunningInstancesOk() (*int32, bool)`

GetMinRunningInstancesOk returns a tuple with the MinRunningInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRunningInstances

`func (o *ApplicationRequest) SetMinRunningInstances(v int32)`

SetMinRunningInstances sets MinRunningInstances field to given value.

### HasMinRunningInstances

`func (o *ApplicationRequest) HasMinRunningInstances() bool`

HasMinRunningInstances returns a boolean if a field has been set.

### GetMaxRunningInstances

`func (o *ApplicationRequest) GetMaxRunningInstances() int32`

GetMaxRunningInstances returns the MaxRunningInstances field if non-nil, zero value otherwise.

### GetMaxRunningInstancesOk

`func (o *ApplicationRequest) GetMaxRunningInstancesOk() (*int32, bool)`

GetMaxRunningInstancesOk returns a tuple with the MaxRunningInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRunningInstances

`func (o *ApplicationRequest) SetMaxRunningInstances(v int32)`

SetMaxRunningInstances sets MaxRunningInstances field to given value.

### HasMaxRunningInstances

`func (o *ApplicationRequest) HasMaxRunningInstances() bool`

HasMaxRunningInstances returns a boolean if a field has been set.

### GetHealthchecks

`func (o *ApplicationRequest) GetHealthchecks() Healthcheck`

GetHealthchecks returns the Healthchecks field if non-nil, zero value otherwise.

### GetHealthchecksOk

`func (o *ApplicationRequest) GetHealthchecksOk() (*Healthcheck, bool)`

GetHealthchecksOk returns a tuple with the Healthchecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthchecks

`func (o *ApplicationRequest) SetHealthchecks(v Healthcheck)`

SetHealthchecks sets Healthchecks field to given value.


### GetAutoPreview

`func (o *ApplicationRequest) GetAutoPreview() bool`

GetAutoPreview returns the AutoPreview field if non-nil, zero value otherwise.

### GetAutoPreviewOk

`func (o *ApplicationRequest) GetAutoPreviewOk() (*bool, bool)`

GetAutoPreviewOk returns a tuple with the AutoPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPreview

`func (o *ApplicationRequest) SetAutoPreview(v bool)`

SetAutoPreview sets AutoPreview field to given value.

### HasAutoPreview

`func (o *ApplicationRequest) HasAutoPreview() bool`

HasAutoPreview returns a boolean if a field has been set.

### GetArguments

`func (o *ApplicationRequest) GetArguments() []string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *ApplicationRequest) GetArgumentsOk() (*[]string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *ApplicationRequest) SetArguments(v []string)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *ApplicationRequest) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetEntrypoint

`func (o *ApplicationRequest) GetEntrypoint() string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *ApplicationRequest) GetEntrypointOk() (*string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *ApplicationRequest) SetEntrypoint(v string)`

SetEntrypoint sets Entrypoint field to given value.

### HasEntrypoint

`func (o *ApplicationRequest) HasEntrypoint() bool`

HasEntrypoint returns a boolean if a field has been set.

### GetAutoDeploy

`func (o *ApplicationRequest) GetAutoDeploy() bool`

GetAutoDeploy returns the AutoDeploy field if non-nil, zero value otherwise.

### GetAutoDeployOk

`func (o *ApplicationRequest) GetAutoDeployOk() (*bool, bool)`

GetAutoDeployOk returns a tuple with the AutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeploy

`func (o *ApplicationRequest) SetAutoDeploy(v bool)`

SetAutoDeploy sets AutoDeploy field to given value.

### HasAutoDeploy

`func (o *ApplicationRequest) HasAutoDeploy() bool`

HasAutoDeploy returns a boolean if a field has been set.

### SetAutoDeployNil

`func (o *ApplicationRequest) SetAutoDeployNil(b bool)`

 SetAutoDeployNil sets the value for AutoDeploy to be an explicit nil

### UnsetAutoDeploy
`func (o *ApplicationRequest) UnsetAutoDeploy()`

UnsetAutoDeploy ensures that no value is present for AutoDeploy, not even an explicit nil
### GetAnnotationsGroups

`func (o *ApplicationRequest) GetAnnotationsGroups() []ServiceAnnotationRequest`

GetAnnotationsGroups returns the AnnotationsGroups field if non-nil, zero value otherwise.

### GetAnnotationsGroupsOk

`func (o *ApplicationRequest) GetAnnotationsGroupsOk() (*[]ServiceAnnotationRequest, bool)`

GetAnnotationsGroupsOk returns a tuple with the AnnotationsGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationsGroups

`func (o *ApplicationRequest) SetAnnotationsGroups(v []ServiceAnnotationRequest)`

SetAnnotationsGroups sets AnnotationsGroups field to given value.

### HasAnnotationsGroups

`func (o *ApplicationRequest) HasAnnotationsGroups() bool`

HasAnnotationsGroups returns a boolean if a field has been set.

### GetLabelsGroups

`func (o *ApplicationRequest) GetLabelsGroups() []ServiceLabelRequest`

GetLabelsGroups returns the LabelsGroups field if non-nil, zero value otherwise.

### GetLabelsGroupsOk

`func (o *ApplicationRequest) GetLabelsGroupsOk() (*[]ServiceLabelRequest, bool)`

GetLabelsGroupsOk returns a tuple with the LabelsGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelsGroups

`func (o *ApplicationRequest) SetLabelsGroups(v []ServiceLabelRequest)`

SetLabelsGroups sets LabelsGroups field to given value.

### HasLabelsGroups

`func (o *ApplicationRequest) HasLabelsGroups() bool`

HasLabelsGroups returns a boolean if a field has been set.

### GetIconUri

`func (o *ApplicationRequest) GetIconUri() string`

GetIconUri returns the IconUri field if non-nil, zero value otherwise.

### GetIconUriOk

`func (o *ApplicationRequest) GetIconUriOk() (*string, bool)`

GetIconUriOk returns a tuple with the IconUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUri

`func (o *ApplicationRequest) SetIconUri(v string)`

SetIconUri sets IconUri field to given value.

### HasIconUri

`func (o *ApplicationRequest) HasIconUri() bool`

HasIconUri returns a boolean if a field has been set.

### GetDockerTargetBuildStage

`func (o *ApplicationRequest) GetDockerTargetBuildStage() string`

GetDockerTargetBuildStage returns the DockerTargetBuildStage field if non-nil, zero value otherwise.

### GetDockerTargetBuildStageOk

`func (o *ApplicationRequest) GetDockerTargetBuildStageOk() (*string, bool)`

GetDockerTargetBuildStageOk returns a tuple with the DockerTargetBuildStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerTargetBuildStage

`func (o *ApplicationRequest) SetDockerTargetBuildStage(v string)`

SetDockerTargetBuildStage sets DockerTargetBuildStage field to given value.

### HasDockerTargetBuildStage

`func (o *ApplicationRequest) HasDockerTargetBuildStage() bool`

HasDockerTargetBuildStage returns a boolean if a field has been set.

### SetDockerTargetBuildStageNil

`func (o *ApplicationRequest) SetDockerTargetBuildStageNil(b bool)`

 SetDockerTargetBuildStageNil sets the value for DockerTargetBuildStage to be an explicit nil

### UnsetDockerTargetBuildStage
`func (o *ApplicationRequest) UnsetDockerTargetBuildStage()`

UnsetDockerTargetBuildStage ensures that no value is present for DockerTargetBuildStage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


