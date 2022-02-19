# ApplicationResponse

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
**DockerfilePath** | Pointer to **string** | The path of the associated Dockerfile. Only if you are using build_mode &#x3D; DOCKER | [optional] 
**BuildpackLanguage** | Pointer to **string** | Development language of the application | [optional] 
**Cpu** | Pointer to **int32** | unit is millicores (m). 1000m &#x3D; 1 cpu | [optional] [default to 250]
**Memory** | Pointer to **int32** | unit is MB. 1024 MB &#x3D; 1GB | [optional] [default to 256]
**MinRunningInstances** | Pointer to **int32** | Minimum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: 0 means that there is no application running.  | [optional] [default to 1]
**MaxRunningInstances** | Pointer to **int32** | Maximum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: -1 means that there is no limit.  | [optional] [default to 1]
**Healthcheck** | Pointer to [**Healthcheck**](Healthcheck.md) |  | [optional] 
**AutoPreview** | Pointer to **bool** | Specify if the environment preview option is activated or not for this application. If activated, a preview environment will be automatically cloned at each pull request.  | [optional] [default to true]
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Storage** | Pointer to [**[]ApplicationStorageResponseStorage**](ApplicationStorageResponseStorage.md) |  | [optional] 
**Ports** | Pointer to [**[]ApplicationPortResponsePorts**](ApplicationPortResponsePorts.md) |  | [optional] 

## Methods

### NewApplicationResponse

`func NewApplicationResponse(id string, createdAt time.Time, ) *ApplicationResponse`

NewApplicationResponse instantiates a new ApplicationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationResponseWithDefaults

`func NewApplicationResponseWithDefaults() *ApplicationResponse`

NewApplicationResponseWithDefaults instantiates a new ApplicationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *ApplicationResponse) GetEnvironment() ReferenceObject`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ApplicationResponse) GetEnvironmentOk() (*ReferenceObject, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ApplicationResponse) SetEnvironment(v ReferenceObject)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ApplicationResponse) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetGitRepository

`func (o *ApplicationResponse) GetGitRepository() ApplicationGitRepositoryResponse`

GetGitRepository returns the GitRepository field if non-nil, zero value otherwise.

### GetGitRepositoryOk

`func (o *ApplicationResponse) GetGitRepositoryOk() (*ApplicationGitRepositoryResponse, bool)`

GetGitRepositoryOk returns a tuple with the GitRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepository

`func (o *ApplicationResponse) SetGitRepository(v ApplicationGitRepositoryResponse)`

SetGitRepository sets GitRepository field to given value.

### HasGitRepository

`func (o *ApplicationResponse) HasGitRepository() bool`

HasGitRepository returns a boolean if a field has been set.

### GetMaximumCpu

`func (o *ApplicationResponse) GetMaximumCpu() int32`

GetMaximumCpu returns the MaximumCpu field if non-nil, zero value otherwise.

### GetMaximumCpuOk

`func (o *ApplicationResponse) GetMaximumCpuOk() (*int32, bool)`

GetMaximumCpuOk returns a tuple with the MaximumCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumCpu

`func (o *ApplicationResponse) SetMaximumCpu(v int32)`

SetMaximumCpu sets MaximumCpu field to given value.

### HasMaximumCpu

`func (o *ApplicationResponse) HasMaximumCpu() bool`

HasMaximumCpu returns a boolean if a field has been set.

### GetMaximumMemory

`func (o *ApplicationResponse) GetMaximumMemory() int32`

GetMaximumMemory returns the MaximumMemory field if non-nil, zero value otherwise.

### GetMaximumMemoryOk

`func (o *ApplicationResponse) GetMaximumMemoryOk() (*int32, bool)`

GetMaximumMemoryOk returns a tuple with the MaximumMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumMemory

`func (o *ApplicationResponse) SetMaximumMemory(v int32)`

SetMaximumMemory sets MaximumMemory field to given value.

### HasMaximumMemory

`func (o *ApplicationResponse) HasMaximumMemory() bool`

HasMaximumMemory returns a boolean if a field has been set.

### GetName

`func (o *ApplicationResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ApplicationResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ApplicationResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetBuildMode

`func (o *ApplicationResponse) GetBuildMode() string`

GetBuildMode returns the BuildMode field if non-nil, zero value otherwise.

### GetBuildModeOk

`func (o *ApplicationResponse) GetBuildModeOk() (*string, bool)`

GetBuildModeOk returns a tuple with the BuildMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildMode

`func (o *ApplicationResponse) SetBuildMode(v string)`

SetBuildMode sets BuildMode field to given value.

### HasBuildMode

`func (o *ApplicationResponse) HasBuildMode() bool`

HasBuildMode returns a boolean if a field has been set.

### GetDockerfilePath

`func (o *ApplicationResponse) GetDockerfilePath() string`

GetDockerfilePath returns the DockerfilePath field if non-nil, zero value otherwise.

### GetDockerfilePathOk

`func (o *ApplicationResponse) GetDockerfilePathOk() (*string, bool)`

GetDockerfilePathOk returns a tuple with the DockerfilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfilePath

`func (o *ApplicationResponse) SetDockerfilePath(v string)`

SetDockerfilePath sets DockerfilePath field to given value.

### HasDockerfilePath

`func (o *ApplicationResponse) HasDockerfilePath() bool`

HasDockerfilePath returns a boolean if a field has been set.

### GetBuildpackLanguage

`func (o *ApplicationResponse) GetBuildpackLanguage() string`

GetBuildpackLanguage returns the BuildpackLanguage field if non-nil, zero value otherwise.

### GetBuildpackLanguageOk

`func (o *ApplicationResponse) GetBuildpackLanguageOk() (*string, bool)`

GetBuildpackLanguageOk returns a tuple with the BuildpackLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildpackLanguage

`func (o *ApplicationResponse) SetBuildpackLanguage(v string)`

SetBuildpackLanguage sets BuildpackLanguage field to given value.

### HasBuildpackLanguage

`func (o *ApplicationResponse) HasBuildpackLanguage() bool`

HasBuildpackLanguage returns a boolean if a field has been set.

### GetCpu

`func (o *ApplicationResponse) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ApplicationResponse) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ApplicationResponse) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ApplicationResponse) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *ApplicationResponse) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ApplicationResponse) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ApplicationResponse) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ApplicationResponse) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetMinRunningInstances

`func (o *ApplicationResponse) GetMinRunningInstances() int32`

GetMinRunningInstances returns the MinRunningInstances field if non-nil, zero value otherwise.

### GetMinRunningInstancesOk

`func (o *ApplicationResponse) GetMinRunningInstancesOk() (*int32, bool)`

GetMinRunningInstancesOk returns a tuple with the MinRunningInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRunningInstances

`func (o *ApplicationResponse) SetMinRunningInstances(v int32)`

SetMinRunningInstances sets MinRunningInstances field to given value.

### HasMinRunningInstances

`func (o *ApplicationResponse) HasMinRunningInstances() bool`

HasMinRunningInstances returns a boolean if a field has been set.

### GetMaxRunningInstances

`func (o *ApplicationResponse) GetMaxRunningInstances() int32`

GetMaxRunningInstances returns the MaxRunningInstances field if non-nil, zero value otherwise.

### GetMaxRunningInstancesOk

`func (o *ApplicationResponse) GetMaxRunningInstancesOk() (*int32, bool)`

GetMaxRunningInstancesOk returns a tuple with the MaxRunningInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRunningInstances

`func (o *ApplicationResponse) SetMaxRunningInstances(v int32)`

SetMaxRunningInstances sets MaxRunningInstances field to given value.

### HasMaxRunningInstances

`func (o *ApplicationResponse) HasMaxRunningInstances() bool`

HasMaxRunningInstances returns a boolean if a field has been set.

### GetHealthcheck

`func (o *ApplicationResponse) GetHealthcheck() Healthcheck`

GetHealthcheck returns the Healthcheck field if non-nil, zero value otherwise.

### GetHealthcheckOk

`func (o *ApplicationResponse) GetHealthcheckOk() (*Healthcheck, bool)`

GetHealthcheckOk returns a tuple with the Healthcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthcheck

`func (o *ApplicationResponse) SetHealthcheck(v Healthcheck)`

SetHealthcheck sets Healthcheck field to given value.

### HasHealthcheck

`func (o *ApplicationResponse) HasHealthcheck() bool`

HasHealthcheck returns a boolean if a field has been set.

### GetAutoPreview

`func (o *ApplicationResponse) GetAutoPreview() bool`

GetAutoPreview returns the AutoPreview field if non-nil, zero value otherwise.

### GetAutoPreviewOk

`func (o *ApplicationResponse) GetAutoPreviewOk() (*bool, bool)`

GetAutoPreviewOk returns a tuple with the AutoPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPreview

`func (o *ApplicationResponse) SetAutoPreview(v bool)`

SetAutoPreview sets AutoPreview field to given value.

### HasAutoPreview

`func (o *ApplicationResponse) HasAutoPreview() bool`

HasAutoPreview returns a boolean if a field has been set.

### GetId

`func (o *ApplicationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ApplicationResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApplicationResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApplicationResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ApplicationResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ApplicationResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ApplicationResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ApplicationResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetStorage

`func (o *ApplicationResponse) GetStorage() []ApplicationStorageResponseStorage`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ApplicationResponse) GetStorageOk() (*[]ApplicationStorageResponseStorage, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ApplicationResponse) SetStorage(v []ApplicationStorageResponseStorage)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *ApplicationResponse) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetPorts

`func (o *ApplicationResponse) GetPorts() []ApplicationPortResponsePorts`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ApplicationResponse) GetPortsOk() (*[]ApplicationPortResponsePorts, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ApplicationResponse) SetPorts(v []ApplicationPortResponsePorts)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ApplicationResponse) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


