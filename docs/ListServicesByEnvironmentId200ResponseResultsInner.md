# ListServicesByEnvironmentId200ResponseResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Storage** | Pointer to **int32** | unit is GB | [optional] [default to 10]
**Environment** | [**ReferenceObject**](ReferenceObject.md) |  | 
**GitRepository** | Pointer to [**ApplicationGitRepository**](ApplicationGitRepository.md) |  | [optional] 
**MaximumCpu** | **int32** | Maximum cpu that can be allocated to the database based on organization cluster configuration. unit is millicores (m). 1000m &#x3D; 1 cpu | 
**MaximumMemory** | **int32** | Maximum memory that can be allocated to the database based on organization cluster configuration. unit is MB. 1024 MB &#x3D; 1GB | 
**Name** | **string** | name is case insensitive | 
**Description** | Pointer to **string** |  | [optional] 
**BuildMode** | Pointer to [**BuildModeEnum**](BuildModeEnum.md) |  | [optional] [default to BUILDMODEENUM_DOCKER]
**DockerfilePath** | Pointer to **NullableString** | The path of the associated Dockerfile. Only if you are using build_mode &#x3D; DOCKER | [optional] 
**Cpu** | **int32** | unit is millicores (m). 1000m &#x3D; 1 cpu This field will be ignored for managed DB (instance type will be used instead).  | [default to 250]
**Memory** | **int32** | unit is MB. 1024 MB &#x3D; 1GB This field will be ignored for managed DB (instance type will be used instead). Default value is linked to the database type: - MANAGED: &#x60;100&#x60; - CONTAINER   - POSTGRES: &#x60;100&#x60;   - REDIS: &#x60;100&#x60;   - MYSQL: &#x60;512&#x60;   - MONGODB: &#x60;256&#x60;  | 
**MinRunningInstances** | **int32** | Minimum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: 0 means that there is no container running.  | [default to 1]
**MaxRunningInstances** | **int32** | Maximum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: -1 means that there is no limit.  | [default to 1]
**Healthchecks** | [**Healthcheck**](Healthcheck.md) |  | 
**AutoPreview** | **bool** | Indicates if the &#39;environment preview option&#39; is enabled.   If enabled, a preview environment will be automatically cloned when &#x60;/preview&#x60; endpoint is called.   If not specified, it takes the value of the &#x60;auto_preview&#x60; property from the associated environment.  | 
**Ports** | Pointer to [**[]HelmResponseAllOfPorts**](HelmResponseAllOfPorts.md) |  | [optional] 
**Arguments** | **[]string** | The extra arguments to pass to helm | 
**Entrypoint** | Pointer to **string** | optional entrypoint when launching container | [optional] 
**AutoDeploy** | **bool** | Specify if the service will be automatically updated after receiving a new image tag or a new commit according to the source type.  | 
**AnnotationsGroups** | Pointer to [**[]OrganizationAnnotationsGroupResponse**](OrganizationAnnotationsGroupResponse.md) |  | [optional] 
**LabelsGroups** | Pointer to [**[]OrganizationLabelsGroupResponse**](OrganizationLabelsGroupResponse.md) |  | [optional] 
**IconUri** | **string** | Icon URI representing the helm service. | 
**ServiceType** | [**ServiceTypeEnum**](ServiceTypeEnum.md) |  | 
**ImageName** | **string** | The image name pattern differs according to chosen container registry provider: * &#x60;ECR&#x60;: &#x60;repository&#x60; * &#x60;SCALEWAY_CR&#x60;: &#x60;namespace/image&#x60; * &#x60;DOCKER_HUB&#x60;: &#x60;image&#x60; or &#x60;repository/image&#x60; * &#x60;PUBLIC_ECR&#x60;: &#x60;registry_alias/repository&#x60;  | 
**Tag** | **string** | tag of the image container | 
**RegistryId** | Pointer to **string** | tag of the image container | [optional] 
**Registry** | [**ContainerRegistryProviderDetailsResponse**](ContainerRegistryProviderDetailsResponse.md) |  | 
**Type** | [**DatabaseTypeEnum**](DatabaseTypeEnum.md) |  | 
**Version** | **string** |  | 
**Mode** | [**DatabaseModeEnum**](DatabaseModeEnum.md) |  | 
**Accessibility** | Pointer to [**DatabaseAccessibilityEnum**](DatabaseAccessibilityEnum.md) |  | [optional] [default to DATABASEACCESSIBILITYENUM_PRIVATE]
**InstanceType** | Pointer to **string** | Database instance type to be used for this database. The list of values can be retrieved via the endpoint /{CloudProvider}/managedDatabase/instanceType/{region}/{dbType}. This field is null for container DB. | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**DiskEncrypted** | Pointer to **bool** | indicates if the database disk is encrypted or not | [optional] 
**TimeoutSec** | Pointer to **int32** | Maximum number of seconds allowed for helm to run before killing it and mark it as failed  | [optional] [default to 600]
**Source** | [**HelmResponseAllOfSource**](HelmResponseAllOfSource.md) |  | 
**AllowClusterWideResources** | **bool** | If we should allow the chart to deploy object outside his specified namespace. Setting this flag to true, requires special rights  | [default to false]
**ValuesOverride** | [**HelmResponseAllOfValuesOverride**](HelmResponseAllOfValuesOverride.md) |  | 

## Methods

### NewListServicesByEnvironmentId200ResponseResultsInner

`func NewListServicesByEnvironmentId200ResponseResultsInner(id string, createdAt time.Time, environment ReferenceObject, maximumCpu int32, maximumMemory int32, name string, cpu int32, memory int32, minRunningInstances int32, maxRunningInstances int32, healthchecks Healthcheck, autoPreview bool, arguments []string, autoDeploy bool, iconUri string, serviceType ServiceTypeEnum, imageName string, tag string, registry ContainerRegistryProviderDetailsResponse, type_ DatabaseTypeEnum, version string, mode DatabaseModeEnum, source HelmResponseAllOfSource, allowClusterWideResources bool, valuesOverride HelmResponseAllOfValuesOverride, ) *ListServicesByEnvironmentId200ResponseResultsInner`

NewListServicesByEnvironmentId200ResponseResultsInner instantiates a new ListServicesByEnvironmentId200ResponseResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServicesByEnvironmentId200ResponseResultsInnerWithDefaults

`func NewListServicesByEnvironmentId200ResponseResultsInnerWithDefaults() *ListServicesByEnvironmentId200ResponseResultsInner`

NewListServicesByEnvironmentId200ResponseResultsInnerWithDefaults instantiates a new ListServicesByEnvironmentId200ResponseResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetStorage

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetStorage() int32`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetStorageOk() (*int32, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetStorage(v int32)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetEnvironment

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetEnvironment() ReferenceObject`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetEnvironmentOk() (*ReferenceObject, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetEnvironment(v ReferenceObject)`

SetEnvironment sets Environment field to given value.


### GetGitRepository

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetGitRepository() ApplicationGitRepository`

GetGitRepository returns the GitRepository field if non-nil, zero value otherwise.

### GetGitRepositoryOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetGitRepositoryOk() (*ApplicationGitRepository, bool)`

GetGitRepositoryOk returns a tuple with the GitRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRepository

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetGitRepository(v ApplicationGitRepository)`

SetGitRepository sets GitRepository field to given value.

### HasGitRepository

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) HasGitRepository() bool`

HasGitRepository returns a boolean if a field has been set.

### GetMaximumCpu

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetMaximumCpu() int32`

GetMaximumCpu returns the MaximumCpu field if non-nil, zero value otherwise.

### GetMaximumCpuOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetMaximumCpuOk() (*int32, bool)`

GetMaximumCpuOk returns a tuple with the MaximumCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumCpu

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetMaximumCpu(v int32)`

SetMaximumCpu sets MaximumCpu field to given value.


### GetMaximumMemory

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetMaximumMemory() int32`

GetMaximumMemory returns the MaximumMemory field if non-nil, zero value otherwise.

### GetMaximumMemoryOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetMaximumMemoryOk() (*int32, bool)`

GetMaximumMemoryOk returns a tuple with the MaximumMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumMemory

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetMaximumMemory(v int32)`

SetMaximumMemory sets MaximumMemory field to given value.


### GetName

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetBuildMode

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetBuildMode() BuildModeEnum`

GetBuildMode returns the BuildMode field if non-nil, zero value otherwise.

### GetBuildModeOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetBuildModeOk() (*BuildModeEnum, bool)`

GetBuildModeOk returns a tuple with the BuildMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildMode

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetBuildMode(v BuildModeEnum)`

SetBuildMode sets BuildMode field to given value.

### HasBuildMode

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) HasBuildMode() bool`

HasBuildMode returns a boolean if a field has been set.

### GetDockerfilePath

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetDockerfilePath() string`

GetDockerfilePath returns the DockerfilePath field if non-nil, zero value otherwise.

### GetDockerfilePathOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetDockerfilePathOk() (*string, bool)`

GetDockerfilePathOk returns a tuple with the DockerfilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerfilePath

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetDockerfilePath(v string)`

SetDockerfilePath sets DockerfilePath field to given value.

### HasDockerfilePath

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) HasDockerfilePath() bool`

HasDockerfilePath returns a boolean if a field has been set.

### SetDockerfilePathNil

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetDockerfilePathNil(b bool)`

 SetDockerfilePathNil sets the value for DockerfilePath to be an explicit nil

### UnsetDockerfilePath
`func (o *ListServicesByEnvironmentId200ResponseResultsInner) UnsetDockerfilePath()`

UnsetDockerfilePath ensures that no value is present for DockerfilePath, not even an explicit nil
### GetCpu

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetCpu(v int32)`

SetCpu sets Cpu field to given value.


### GetMemory

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetMemory(v int32)`

SetMemory sets Memory field to given value.


### GetMinRunningInstances

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetMinRunningInstances() int32`

GetMinRunningInstances returns the MinRunningInstances field if non-nil, zero value otherwise.

### GetMinRunningInstancesOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetMinRunningInstancesOk() (*int32, bool)`

GetMinRunningInstancesOk returns a tuple with the MinRunningInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRunningInstances

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetMinRunningInstances(v int32)`

SetMinRunningInstances sets MinRunningInstances field to given value.


### GetMaxRunningInstances

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetMaxRunningInstances() int32`

GetMaxRunningInstances returns the MaxRunningInstances field if non-nil, zero value otherwise.

### GetMaxRunningInstancesOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetMaxRunningInstancesOk() (*int32, bool)`

GetMaxRunningInstancesOk returns a tuple with the MaxRunningInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRunningInstances

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetMaxRunningInstances(v int32)`

SetMaxRunningInstances sets MaxRunningInstances field to given value.


### GetHealthchecks

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetHealthchecks() Healthcheck`

GetHealthchecks returns the Healthchecks field if non-nil, zero value otherwise.

### GetHealthchecksOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetHealthchecksOk() (*Healthcheck, bool)`

GetHealthchecksOk returns a tuple with the Healthchecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthchecks

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetHealthchecks(v Healthcheck)`

SetHealthchecks sets Healthchecks field to given value.


### GetAutoPreview

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetAutoPreview() bool`

GetAutoPreview returns the AutoPreview field if non-nil, zero value otherwise.

### GetAutoPreviewOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetAutoPreviewOk() (*bool, bool)`

GetAutoPreviewOk returns a tuple with the AutoPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPreview

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetAutoPreview(v bool)`

SetAutoPreview sets AutoPreview field to given value.


### GetPorts

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetPorts() []HelmResponseAllOfPorts`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetPortsOk() (*[]HelmResponseAllOfPorts, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetPorts(v []HelmResponseAllOfPorts)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetArguments

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetArguments() []string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetArgumentsOk() (*[]string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetArguments(v []string)`

SetArguments sets Arguments field to given value.


### GetEntrypoint

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetEntrypoint() string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetEntrypointOk() (*string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetEntrypoint(v string)`

SetEntrypoint sets Entrypoint field to given value.

### HasEntrypoint

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) HasEntrypoint() bool`

HasEntrypoint returns a boolean if a field has been set.

### GetAutoDeploy

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetAutoDeploy() bool`

GetAutoDeploy returns the AutoDeploy field if non-nil, zero value otherwise.

### GetAutoDeployOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetAutoDeployOk() (*bool, bool)`

GetAutoDeployOk returns a tuple with the AutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeploy

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetAutoDeploy(v bool)`

SetAutoDeploy sets AutoDeploy field to given value.


### GetAnnotationsGroups

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetAnnotationsGroups() []OrganizationAnnotationsGroupResponse`

GetAnnotationsGroups returns the AnnotationsGroups field if non-nil, zero value otherwise.

### GetAnnotationsGroupsOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetAnnotationsGroupsOk() (*[]OrganizationAnnotationsGroupResponse, bool)`

GetAnnotationsGroupsOk returns a tuple with the AnnotationsGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationsGroups

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetAnnotationsGroups(v []OrganizationAnnotationsGroupResponse)`

SetAnnotationsGroups sets AnnotationsGroups field to given value.

### HasAnnotationsGroups

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) HasAnnotationsGroups() bool`

HasAnnotationsGroups returns a boolean if a field has been set.

### GetLabelsGroups

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetLabelsGroups() []OrganizationLabelsGroupResponse`

GetLabelsGroups returns the LabelsGroups field if non-nil, zero value otherwise.

### GetLabelsGroupsOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetLabelsGroupsOk() (*[]OrganizationLabelsGroupResponse, bool)`

GetLabelsGroupsOk returns a tuple with the LabelsGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelsGroups

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetLabelsGroups(v []OrganizationLabelsGroupResponse)`

SetLabelsGroups sets LabelsGroups field to given value.

### HasLabelsGroups

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) HasLabelsGroups() bool`

HasLabelsGroups returns a boolean if a field has been set.

### GetIconUri

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetIconUri() string`

GetIconUri returns the IconUri field if non-nil, zero value otherwise.

### GetIconUriOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetIconUriOk() (*string, bool)`

GetIconUriOk returns a tuple with the IconUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUri

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetIconUri(v string)`

SetIconUri sets IconUri field to given value.


### GetServiceType

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetServiceType() ServiceTypeEnum`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetServiceTypeOk() (*ServiceTypeEnum, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetServiceType(v ServiceTypeEnum)`

SetServiceType sets ServiceType field to given value.


### GetImageName

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetImageName(v string)`

SetImageName sets ImageName field to given value.


### GetTag

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetTag(v string)`

SetTag sets Tag field to given value.


### GetRegistryId

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.

### HasRegistryId

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) HasRegistryId() bool`

HasRegistryId returns a boolean if a field has been set.

### GetRegistry

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetRegistry() ContainerRegistryProviderDetailsResponse`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetRegistryOk() (*ContainerRegistryProviderDetailsResponse, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetRegistry(v ContainerRegistryProviderDetailsResponse)`

SetRegistry sets Registry field to given value.


### GetType

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetType() DatabaseTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetTypeOk() (*DatabaseTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetType(v DatabaseTypeEnum)`

SetType sets Type field to given value.


### GetVersion

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetMode

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetMode() DatabaseModeEnum`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetModeOk() (*DatabaseModeEnum, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetMode(v DatabaseModeEnum)`

SetMode sets Mode field to given value.


### GetAccessibility

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetAccessibility() DatabaseAccessibilityEnum`

GetAccessibility returns the Accessibility field if non-nil, zero value otherwise.

### GetAccessibilityOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetAccessibilityOk() (*DatabaseAccessibilityEnum, bool)`

GetAccessibilityOk returns a tuple with the Accessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibility

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetAccessibility(v DatabaseAccessibilityEnum)`

SetAccessibility sets Accessibility field to given value.

### HasAccessibility

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) HasAccessibility() bool`

HasAccessibility returns a boolean if a field has been set.

### GetInstanceType

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetHost

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetDiskEncrypted

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetDiskEncrypted() bool`

GetDiskEncrypted returns the DiskEncrypted field if non-nil, zero value otherwise.

### GetDiskEncryptedOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetDiskEncryptedOk() (*bool, bool)`

GetDiskEncryptedOk returns a tuple with the DiskEncrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEncrypted

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetDiskEncrypted(v bool)`

SetDiskEncrypted sets DiskEncrypted field to given value.

### HasDiskEncrypted

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) HasDiskEncrypted() bool`

HasDiskEncrypted returns a boolean if a field has been set.

### GetTimeoutSec

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetTimeoutSec() int32`

GetTimeoutSec returns the TimeoutSec field if non-nil, zero value otherwise.

### GetTimeoutSecOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetTimeoutSecOk() (*int32, bool)`

GetTimeoutSecOk returns a tuple with the TimeoutSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSec

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetTimeoutSec(v int32)`

SetTimeoutSec sets TimeoutSec field to given value.

### HasTimeoutSec

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) HasTimeoutSec() bool`

HasTimeoutSec returns a boolean if a field has been set.

### GetSource

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetSource() HelmResponseAllOfSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetSourceOk() (*HelmResponseAllOfSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetSource(v HelmResponseAllOfSource)`

SetSource sets Source field to given value.


### GetAllowClusterWideResources

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetAllowClusterWideResources() bool`

GetAllowClusterWideResources returns the AllowClusterWideResources field if non-nil, zero value otherwise.

### GetAllowClusterWideResourcesOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetAllowClusterWideResourcesOk() (*bool, bool)`

GetAllowClusterWideResourcesOk returns a tuple with the AllowClusterWideResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowClusterWideResources

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetAllowClusterWideResources(v bool)`

SetAllowClusterWideResources sets AllowClusterWideResources field to given value.


### GetValuesOverride

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetValuesOverride() HelmResponseAllOfValuesOverride`

GetValuesOverride returns the ValuesOverride field if non-nil, zero value otherwise.

### GetValuesOverrideOk

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) GetValuesOverrideOk() (*HelmResponseAllOfValuesOverride, bool)`

GetValuesOverrideOk returns a tuple with the ValuesOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesOverride

`func (o *ListServicesByEnvironmentId200ResponseResultsInner) SetValuesOverride(v HelmResponseAllOfValuesOverride)`

SetValuesOverride sets ValuesOverride field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


