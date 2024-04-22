# Application

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Storage** | Pointer to [**[]ServiceStorageStorageInner**](ServiceStorageStorageInner.md) |  | [optional] 
**Environment** | [**ReferenceObject**](ReferenceObject.md) |  | 
**GitRepository** | Pointer to [**ApplicationGitRepository**](ApplicationGitRepository.md) |  | [optional] 
**MaximumCpu** | Pointer to **int32** | Maximum cpu that can be allocated to the application based on organization cluster configuration. unit is millicores (m). 1000m &#x3D; 1 cpu | [optional] 
**MaximumMemory** | Pointer to **int32** | Maximum memory that can be allocated to the application based on organization cluster configuration. unit is MB. 1024 MB &#x3D; 1GB | [optional] 
**Name** | **string** | name is case insensitive | 
**Description** | Pointer to **string** | give a description to this application | [optional] 
**BuildMode** | Pointer to [**BuildModeEnum**](BuildModeEnum.md) |  | [optional] [default to BUILDMODEENUM_BUILDPACKS]
**DockerfilePath** | Pointer to **NullableString** | The path of the associated Dockerfile. Only if you are using build_mode &#x3D; DOCKER | [optional] 
**BuildpackLanguage** | Pointer to [**NullableBuildPackLanguageEnum**](BuildPackLanguageEnum.md) |  | [optional] 
**Cpu** | Pointer to **int32** | unit is millicores (m). 1000m &#x3D; 1 cpu | [optional] 
**Memory** | Pointer to **int32** | unit is MB. 1024 MB &#x3D; 1GB | [optional] 
**MinRunningInstances** | Pointer to **int32** | Minimum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: 0 means that there is no application running.  | [optional] [default to 1]
**MaxRunningInstances** | Pointer to **int32** | Maximum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: -1 means that there is no limit.  | [optional] [default to 1]
**Healthchecks** | [**Healthcheck**](Healthcheck.md) |  | 
**AutoPreview** | Pointer to **bool** | Specify if the environment preview option is activated or not for this application.   If activated, a preview environment will be automatically cloned at each pull request.   If not specified, it takes the value of the &#x60;auto_preview&#x60; property from the associated environment.  | [optional] [default to true]
**Ports** | Pointer to [**[]ServicePort**](ServicePort.md) |  | [optional] 
**Arguments** | Pointer to **[]string** |  | [optional] 
**Entrypoint** | Pointer to **string** | optional entrypoint when launching container | [optional] 
**AutoDeploy** | Pointer to **bool** | Specify if the application will be automatically updated after receiving a new commit. | [optional] 
**AnnotationsGroups** | Pointer to [**OrganizationAnnotationsGroupResponseList**](OrganizationAnnotationsGroupResponseList.md) |  | [optional] 

## Methods

### NewApplication

`func NewApplication(id string, createdAt time.Time, environment ReferenceObject, name string, healthchecks Healthcheck, ) *Application`

NewApplication instantiates a new Application object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationWithDefaults

`func NewApplicationWithDefaults() *Application`

NewApplicationWithDefaults instantiates a new Application object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Application) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Application) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Application) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Application) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Application) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Application) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Application) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Application) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Application) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Application) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetStorage

`func (o *Application) GetStorage() []ServiceStorageStorageInner`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *Application) GetStorageOk() (*[]ServiceStorageStorageInner, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *Application) SetStorage(v []ServiceStorageStorageInner)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *Application) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetEnvironment

`func (o *Application) GetEnvironment() ReferenceObject`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *Application) GetEnvironmentOk() (*ReferenceObject, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *Application) SetEnvironment(v ReferenceObject)`

SetEnvironment sets Environment field to given value.


### GetGitRepository

`func (o *Application) GetGitRepository() ApplicationGitRepository`

GetGitRepository returns the GitRepository field if non-nil, zero value otherwise.

### GetGitRepositoryOk

`func (o *Application) GetGitRepositoryOk() (*ApplicationGitRepository, bool)`

GetGitRepositoryOk returns a tuple with the GitRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepository

`func (o *Application) SetGitRepository(v ApplicationGitRepository)`

SetGitRepository sets GitRepository field to given value.

### HasGitRepository

`func (o *Application) HasGitRepository() bool`

HasGitRepository returns a boolean if a field has been set.

### GetMaximumCpu

`func (o *Application) GetMaximumCpu() int32`

GetMaximumCpu returns the MaximumCpu field if non-nil, zero value otherwise.

### GetMaximumCpuOk

`func (o *Application) GetMaximumCpuOk() (*int32, bool)`

GetMaximumCpuOk returns a tuple with the MaximumCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumCpu

`func (o *Application) SetMaximumCpu(v int32)`

SetMaximumCpu sets MaximumCpu field to given value.

### HasMaximumCpu

`func (o *Application) HasMaximumCpu() bool`

HasMaximumCpu returns a boolean if a field has been set.

### GetMaximumMemory

`func (o *Application) GetMaximumMemory() int32`

GetMaximumMemory returns the MaximumMemory field if non-nil, zero value otherwise.

### GetMaximumMemoryOk

`func (o *Application) GetMaximumMemoryOk() (*int32, bool)`

GetMaximumMemoryOk returns a tuple with the MaximumMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumMemory

`func (o *Application) SetMaximumMemory(v int32)`

SetMaximumMemory sets MaximumMemory field to given value.

### HasMaximumMemory

`func (o *Application) HasMaximumMemory() bool`

HasMaximumMemory returns a boolean if a field has been set.

### GetName

`func (o *Application) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Application) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Application) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Application) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Application) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Application) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Application) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBuildMode

`func (o *Application) GetBuildMode() BuildModeEnum`

GetBuildMode returns the BuildMode field if non-nil, zero value otherwise.

### GetBuildModeOk

`func (o *Application) GetBuildModeOk() (*BuildModeEnum, bool)`

GetBuildModeOk returns a tuple with the BuildMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildMode

`func (o *Application) SetBuildMode(v BuildModeEnum)`

SetBuildMode sets BuildMode field to given value.

### HasBuildMode

`func (o *Application) HasBuildMode() bool`

HasBuildMode returns a boolean if a field has been set.

### GetDockerfilePath

`func (o *Application) GetDockerfilePath() string`

GetDockerfilePath returns the DockerfilePath field if non-nil, zero value otherwise.

### GetDockerfilePathOk

`func (o *Application) GetDockerfilePathOk() (*string, bool)`

GetDockerfilePathOk returns a tuple with the DockerfilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfilePath

`func (o *Application) SetDockerfilePath(v string)`

SetDockerfilePath sets DockerfilePath field to given value.

### HasDockerfilePath

`func (o *Application) HasDockerfilePath() bool`

HasDockerfilePath returns a boolean if a field has been set.

### SetDockerfilePathNil

`func (o *Application) SetDockerfilePathNil(b bool)`

 SetDockerfilePathNil sets the value for DockerfilePath to be an explicit nil

### UnsetDockerfilePath
`func (o *Application) UnsetDockerfilePath()`

UnsetDockerfilePath ensures that no value is present for DockerfilePath, not even an explicit nil
### GetBuildpackLanguage

`func (o *Application) GetBuildpackLanguage() BuildPackLanguageEnum`

GetBuildpackLanguage returns the BuildpackLanguage field if non-nil, zero value otherwise.

### GetBuildpackLanguageOk

`func (o *Application) GetBuildpackLanguageOk() (*BuildPackLanguageEnum, bool)`

GetBuildpackLanguageOk returns a tuple with the BuildpackLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildpackLanguage

`func (o *Application) SetBuildpackLanguage(v BuildPackLanguageEnum)`

SetBuildpackLanguage sets BuildpackLanguage field to given value.

### HasBuildpackLanguage

`func (o *Application) HasBuildpackLanguage() bool`

HasBuildpackLanguage returns a boolean if a field has been set.

### SetBuildpackLanguageNil

`func (o *Application) SetBuildpackLanguageNil(b bool)`

 SetBuildpackLanguageNil sets the value for BuildpackLanguage to be an explicit nil

### UnsetBuildpackLanguage
`func (o *Application) UnsetBuildpackLanguage()`

UnsetBuildpackLanguage ensures that no value is present for BuildpackLanguage, not even an explicit nil
### GetCpu

`func (o *Application) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *Application) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *Application) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *Application) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *Application) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *Application) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *Application) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *Application) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetMinRunningInstances

`func (o *Application) GetMinRunningInstances() int32`

GetMinRunningInstances returns the MinRunningInstances field if non-nil, zero value otherwise.

### GetMinRunningInstancesOk

`func (o *Application) GetMinRunningInstancesOk() (*int32, bool)`

GetMinRunningInstancesOk returns a tuple with the MinRunningInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRunningInstances

`func (o *Application) SetMinRunningInstances(v int32)`

SetMinRunningInstances sets MinRunningInstances field to given value.

### HasMinRunningInstances

`func (o *Application) HasMinRunningInstances() bool`

HasMinRunningInstances returns a boolean if a field has been set.

### GetMaxRunningInstances

`func (o *Application) GetMaxRunningInstances() int32`

GetMaxRunningInstances returns the MaxRunningInstances field if non-nil, zero value otherwise.

### GetMaxRunningInstancesOk

`func (o *Application) GetMaxRunningInstancesOk() (*int32, bool)`

GetMaxRunningInstancesOk returns a tuple with the MaxRunningInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRunningInstances

`func (o *Application) SetMaxRunningInstances(v int32)`

SetMaxRunningInstances sets MaxRunningInstances field to given value.

### HasMaxRunningInstances

`func (o *Application) HasMaxRunningInstances() bool`

HasMaxRunningInstances returns a boolean if a field has been set.

### GetHealthchecks

`func (o *Application) GetHealthchecks() Healthcheck`

GetHealthchecks returns the Healthchecks field if non-nil, zero value otherwise.

### GetHealthchecksOk

`func (o *Application) GetHealthchecksOk() (*Healthcheck, bool)`

GetHealthchecksOk returns a tuple with the Healthchecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthchecks

`func (o *Application) SetHealthchecks(v Healthcheck)`

SetHealthchecks sets Healthchecks field to given value.


### GetAutoPreview

`func (o *Application) GetAutoPreview() bool`

GetAutoPreview returns the AutoPreview field if non-nil, zero value otherwise.

### GetAutoPreviewOk

`func (o *Application) GetAutoPreviewOk() (*bool, bool)`

GetAutoPreviewOk returns a tuple with the AutoPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPreview

`func (o *Application) SetAutoPreview(v bool)`

SetAutoPreview sets AutoPreview field to given value.

### HasAutoPreview

`func (o *Application) HasAutoPreview() bool`

HasAutoPreview returns a boolean if a field has been set.

### GetPorts

`func (o *Application) GetPorts() []ServicePort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *Application) GetPortsOk() (*[]ServicePort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *Application) SetPorts(v []ServicePort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *Application) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetArguments

`func (o *Application) GetArguments() []string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *Application) GetArgumentsOk() (*[]string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *Application) SetArguments(v []string)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *Application) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetEntrypoint

`func (o *Application) GetEntrypoint() string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *Application) GetEntrypointOk() (*string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *Application) SetEntrypoint(v string)`

SetEntrypoint sets Entrypoint field to given value.

### HasEntrypoint

`func (o *Application) HasEntrypoint() bool`

HasEntrypoint returns a boolean if a field has been set.

### GetAutoDeploy

`func (o *Application) GetAutoDeploy() bool`

GetAutoDeploy returns the AutoDeploy field if non-nil, zero value otherwise.

### GetAutoDeployOk

`func (o *Application) GetAutoDeployOk() (*bool, bool)`

GetAutoDeployOk returns a tuple with the AutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeploy

`func (o *Application) SetAutoDeploy(v bool)`

SetAutoDeploy sets AutoDeploy field to given value.

### HasAutoDeploy

`func (o *Application) HasAutoDeploy() bool`

HasAutoDeploy returns a boolean if a field has been set.

### GetAnnotationsGroups

`func (o *Application) GetAnnotationsGroups() OrganizationAnnotationsGroupResponseList`

GetAnnotationsGroups returns the AnnotationsGroups field if non-nil, zero value otherwise.

### GetAnnotationsGroupsOk

`func (o *Application) GetAnnotationsGroupsOk() (*OrganizationAnnotationsGroupResponseList, bool)`

GetAnnotationsGroupsOk returns a tuple with the AnnotationsGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationsGroups

`func (o *Application) SetAnnotationsGroups(v OrganizationAnnotationsGroupResponseList)`

SetAnnotationsGroups sets AnnotationsGroups field to given value.

### HasAnnotationsGroups

`func (o *Application) HasAnnotationsGroups() bool`

HasAnnotationsGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


