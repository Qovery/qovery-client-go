# ContainerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Storage** | Pointer to [**[]ServiceStorageStorageInner**](ServiceStorageStorageInner.md) |  | [optional] 
**ImageName** | **string** | The image name pattern differs according to chosen container registry provider: * &#x60;ECR&#x60;: &#x60;repository&#x60; * &#x60;SCALEWAY_CR&#x60;: &#x60;namespace/image&#x60; * &#x60;DOCKER_HUB&#x60;: &#x60;image&#x60; or &#x60;repository/image&#x60; * &#x60;PUBLIC_ECR&#x60;: &#x60;registry_alias/repository&#x60;  | 
**Tag** | **string** | tag of the image container | 
**RegistryId** | Pointer to **string** | tag of the image container | [optional] 
**Registry** | [**ContainerRegistryProviderDetailsResponse**](ContainerRegistryProviderDetailsResponse.md) |  | 
**Environment** | [**ReferenceObject**](ReferenceObject.md) |  | 
**MaximumCpu** | **int32** | Maximum cpu that can be allocated to the container based on organization cluster configuration. unit is millicores (m). 1000m &#x3D; 1 cpu | 
**MaximumMemory** | **int32** | Maximum memory that can be allocated to the container based on organization cluster configuration. unit is MB. 1024 MB &#x3D; 1GB | 
**MaximumGpu** | **int32** | Maximum memory that can be allocated to the container based on organization cluster configuration. unit is MB. 1024 MB &#x3D; 1GB | [default to 0]
**Name** | **string** | name is case insensitive | 
**Description** | Pointer to **string** | give a description to this container | [optional] 
**Arguments** | Pointer to **[]string** |  | [optional] 
**Entrypoint** | Pointer to **string** | optional entrypoint when launching container | [optional] 
**Cpu** | **int32** | unit is millicores (m). 1000m &#x3D; 1 cpu | 
**Memory** | **int32** | unit is MB. 1024 MB &#x3D; 1GB | 
**Gpu** | **int32** |  | [default to 0]
**MinRunningInstances** | **int32** | Minimum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: 0 means that there is no container running.  | [default to 1]
**MaxRunningInstances** | **int32** | Maximum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: -1 means that there is no limit.  | [default to 1]
**Healthchecks** | [**Healthcheck**](Healthcheck.md) |  | 
**AutoPreview** | **bool** | Indicates if the &#39;environment preview option&#39; is enabled for this container.   If enabled, a preview environment will be automatically cloned when &#x60;/preview&#x60; endpoint is called.   If not specified, it takes the value of the &#x60;auto_preview&#x60; property from the associated environment.  | 
**Ports** | Pointer to [**[]ServicePort**](ServicePort.md) |  | [optional] 
**AutoDeploy** | Pointer to **bool** | Specify if the container will be automatically updated after receiving a new image tag.  The new image tag shall be communicated via the \&quot;Auto Deploy container\&quot; endpoint https://api-doc.qovery.com/#tag/Containers/operation/autoDeployContainerEnvironments  | [optional] 
**AnnotationsGroups** | Pointer to [**[]OrganizationAnnotationsGroupResponse**](OrganizationAnnotationsGroupResponse.md) |  | [optional] 
**LabelsGroups** | Pointer to [**[]OrganizationLabelsGroupResponse**](OrganizationLabelsGroupResponse.md) |  | [optional] 
**IconUri** | **string** | Icon URI representing the container. | 
**ServiceType** | [**ServiceTypeEnum**](ServiceTypeEnum.md) |  | 
**Autoscaling** | Pointer to [**AutoscalingPolicyResponse**](AutoscalingPolicyResponse.md) |  | [optional] 

## Methods

### NewContainerResponse

`func NewContainerResponse(id string, createdAt time.Time, imageName string, tag string, registry ContainerRegistryProviderDetailsResponse, environment ReferenceObject, maximumCpu int32, maximumMemory int32, maximumGpu int32, name string, cpu int32, memory int32, gpu int32, minRunningInstances int32, maxRunningInstances int32, healthchecks Healthcheck, autoPreview bool, iconUri string, serviceType ServiceTypeEnum, ) *ContainerResponse`

NewContainerResponse instantiates a new ContainerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerResponseWithDefaults

`func NewContainerResponseWithDefaults() *ContainerResponse`

NewContainerResponseWithDefaults instantiates a new ContainerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ContainerResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContainerResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContainerResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ContainerResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ContainerResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ContainerResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ContainerResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ContainerResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ContainerResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ContainerResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetStorage

`func (o *ContainerResponse) GetStorage() []ServiceStorageStorageInner`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ContainerResponse) GetStorageOk() (*[]ServiceStorageStorageInner, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ContainerResponse) SetStorage(v []ServiceStorageStorageInner)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *ContainerResponse) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetImageName

`func (o *ContainerResponse) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *ContainerResponse) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *ContainerResponse) SetImageName(v string)`

SetImageName sets ImageName field to given value.


### GetTag

`func (o *ContainerResponse) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ContainerResponse) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ContainerResponse) SetTag(v string)`

SetTag sets Tag field to given value.


### GetRegistryId

`func (o *ContainerResponse) GetRegistryId() string`

GetRegistryId returns the RegistryId field if non-nil, zero value otherwise.

### GetRegistryIdOk

`func (o *ContainerResponse) GetRegistryIdOk() (*string, bool)`

GetRegistryIdOk returns a tuple with the RegistryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistryId

`func (o *ContainerResponse) SetRegistryId(v string)`

SetRegistryId sets RegistryId field to given value.

### HasRegistryId

`func (o *ContainerResponse) HasRegistryId() bool`

HasRegistryId returns a boolean if a field has been set.

### GetRegistry

`func (o *ContainerResponse) GetRegistry() ContainerRegistryProviderDetailsResponse`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *ContainerResponse) GetRegistryOk() (*ContainerRegistryProviderDetailsResponse, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *ContainerResponse) SetRegistry(v ContainerRegistryProviderDetailsResponse)`

SetRegistry sets Registry field to given value.


### GetEnvironment

`func (o *ContainerResponse) GetEnvironment() ReferenceObject`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ContainerResponse) GetEnvironmentOk() (*ReferenceObject, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ContainerResponse) SetEnvironment(v ReferenceObject)`

SetEnvironment sets Environment field to given value.


### GetMaximumCpu

`func (o *ContainerResponse) GetMaximumCpu() int32`

GetMaximumCpu returns the MaximumCpu field if non-nil, zero value otherwise.

### GetMaximumCpuOk

`func (o *ContainerResponse) GetMaximumCpuOk() (*int32, bool)`

GetMaximumCpuOk returns a tuple with the MaximumCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumCpu

`func (o *ContainerResponse) SetMaximumCpu(v int32)`

SetMaximumCpu sets MaximumCpu field to given value.


### GetMaximumMemory

`func (o *ContainerResponse) GetMaximumMemory() int32`

GetMaximumMemory returns the MaximumMemory field if non-nil, zero value otherwise.

### GetMaximumMemoryOk

`func (o *ContainerResponse) GetMaximumMemoryOk() (*int32, bool)`

GetMaximumMemoryOk returns a tuple with the MaximumMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumMemory

`func (o *ContainerResponse) SetMaximumMemory(v int32)`

SetMaximumMemory sets MaximumMemory field to given value.


### GetMaximumGpu

`func (o *ContainerResponse) GetMaximumGpu() int32`

GetMaximumGpu returns the MaximumGpu field if non-nil, zero value otherwise.

### GetMaximumGpuOk

`func (o *ContainerResponse) GetMaximumGpuOk() (*int32, bool)`

GetMaximumGpuOk returns a tuple with the MaximumGpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumGpu

`func (o *ContainerResponse) SetMaximumGpu(v int32)`

SetMaximumGpu sets MaximumGpu field to given value.


### GetName

`func (o *ContainerResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ContainerResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContainerResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContainerResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContainerResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetArguments

`func (o *ContainerResponse) GetArguments() []string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *ContainerResponse) GetArgumentsOk() (*[]string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *ContainerResponse) SetArguments(v []string)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *ContainerResponse) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetEntrypoint

`func (o *ContainerResponse) GetEntrypoint() string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *ContainerResponse) GetEntrypointOk() (*string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *ContainerResponse) SetEntrypoint(v string)`

SetEntrypoint sets Entrypoint field to given value.

### HasEntrypoint

`func (o *ContainerResponse) HasEntrypoint() bool`

HasEntrypoint returns a boolean if a field has been set.

### GetCpu

`func (o *ContainerResponse) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ContainerResponse) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ContainerResponse) SetCpu(v int32)`

SetCpu sets Cpu field to given value.


### GetMemory

`func (o *ContainerResponse) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ContainerResponse) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ContainerResponse) SetMemory(v int32)`

SetMemory sets Memory field to given value.


### GetGpu

`func (o *ContainerResponse) GetGpu() int32`

GetGpu returns the Gpu field if non-nil, zero value otherwise.

### GetGpuOk

`func (o *ContainerResponse) GetGpuOk() (*int32, bool)`

GetGpuOk returns a tuple with the Gpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpu

`func (o *ContainerResponse) SetGpu(v int32)`

SetGpu sets Gpu field to given value.


### GetMinRunningInstances

`func (o *ContainerResponse) GetMinRunningInstances() int32`

GetMinRunningInstances returns the MinRunningInstances field if non-nil, zero value otherwise.

### GetMinRunningInstancesOk

`func (o *ContainerResponse) GetMinRunningInstancesOk() (*int32, bool)`

GetMinRunningInstancesOk returns a tuple with the MinRunningInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRunningInstances

`func (o *ContainerResponse) SetMinRunningInstances(v int32)`

SetMinRunningInstances sets MinRunningInstances field to given value.


### GetMaxRunningInstances

`func (o *ContainerResponse) GetMaxRunningInstances() int32`

GetMaxRunningInstances returns the MaxRunningInstances field if non-nil, zero value otherwise.

### GetMaxRunningInstancesOk

`func (o *ContainerResponse) GetMaxRunningInstancesOk() (*int32, bool)`

GetMaxRunningInstancesOk returns a tuple with the MaxRunningInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRunningInstances

`func (o *ContainerResponse) SetMaxRunningInstances(v int32)`

SetMaxRunningInstances sets MaxRunningInstances field to given value.


### GetHealthchecks

`func (o *ContainerResponse) GetHealthchecks() Healthcheck`

GetHealthchecks returns the Healthchecks field if non-nil, zero value otherwise.

### GetHealthchecksOk

`func (o *ContainerResponse) GetHealthchecksOk() (*Healthcheck, bool)`

GetHealthchecksOk returns a tuple with the Healthchecks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthchecks

`func (o *ContainerResponse) SetHealthchecks(v Healthcheck)`

SetHealthchecks sets Healthchecks field to given value.


### GetAutoPreview

`func (o *ContainerResponse) GetAutoPreview() bool`

GetAutoPreview returns the AutoPreview field if non-nil, zero value otherwise.

### GetAutoPreviewOk

`func (o *ContainerResponse) GetAutoPreviewOk() (*bool, bool)`

GetAutoPreviewOk returns a tuple with the AutoPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPreview

`func (o *ContainerResponse) SetAutoPreview(v bool)`

SetAutoPreview sets AutoPreview field to given value.


### GetPorts

`func (o *ContainerResponse) GetPorts() []ServicePort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ContainerResponse) GetPortsOk() (*[]ServicePort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ContainerResponse) SetPorts(v []ServicePort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ContainerResponse) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetAutoDeploy

`func (o *ContainerResponse) GetAutoDeploy() bool`

GetAutoDeploy returns the AutoDeploy field if non-nil, zero value otherwise.

### GetAutoDeployOk

`func (o *ContainerResponse) GetAutoDeployOk() (*bool, bool)`

GetAutoDeployOk returns a tuple with the AutoDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDeploy

`func (o *ContainerResponse) SetAutoDeploy(v bool)`

SetAutoDeploy sets AutoDeploy field to given value.

### HasAutoDeploy

`func (o *ContainerResponse) HasAutoDeploy() bool`

HasAutoDeploy returns a boolean if a field has been set.

### GetAnnotationsGroups

`func (o *ContainerResponse) GetAnnotationsGroups() []OrganizationAnnotationsGroupResponse`

GetAnnotationsGroups returns the AnnotationsGroups field if non-nil, zero value otherwise.

### GetAnnotationsGroupsOk

`func (o *ContainerResponse) GetAnnotationsGroupsOk() (*[]OrganizationAnnotationsGroupResponse, bool)`

GetAnnotationsGroupsOk returns a tuple with the AnnotationsGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationsGroups

`func (o *ContainerResponse) SetAnnotationsGroups(v []OrganizationAnnotationsGroupResponse)`

SetAnnotationsGroups sets AnnotationsGroups field to given value.

### HasAnnotationsGroups

`func (o *ContainerResponse) HasAnnotationsGroups() bool`

HasAnnotationsGroups returns a boolean if a field has been set.

### GetLabelsGroups

`func (o *ContainerResponse) GetLabelsGroups() []OrganizationLabelsGroupResponse`

GetLabelsGroups returns the LabelsGroups field if non-nil, zero value otherwise.

### GetLabelsGroupsOk

`func (o *ContainerResponse) GetLabelsGroupsOk() (*[]OrganizationLabelsGroupResponse, bool)`

GetLabelsGroupsOk returns a tuple with the LabelsGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelsGroups

`func (o *ContainerResponse) SetLabelsGroups(v []OrganizationLabelsGroupResponse)`

SetLabelsGroups sets LabelsGroups field to given value.

### HasLabelsGroups

`func (o *ContainerResponse) HasLabelsGroups() bool`

HasLabelsGroups returns a boolean if a field has been set.

### GetIconUri

`func (o *ContainerResponse) GetIconUri() string`

GetIconUri returns the IconUri field if non-nil, zero value otherwise.

### GetIconUriOk

`func (o *ContainerResponse) GetIconUriOk() (*string, bool)`

GetIconUriOk returns a tuple with the IconUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUri

`func (o *ContainerResponse) SetIconUri(v string)`

SetIconUri sets IconUri field to given value.


### GetServiceType

`func (o *ContainerResponse) GetServiceType() ServiceTypeEnum`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *ContainerResponse) GetServiceTypeOk() (*ServiceTypeEnum, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *ContainerResponse) SetServiceType(v ServiceTypeEnum)`

SetServiceType sets ServiceType field to given value.


### GetAutoscaling

`func (o *ContainerResponse) GetAutoscaling() AutoscalingPolicyResponse`

GetAutoscaling returns the Autoscaling field if non-nil, zero value otherwise.

### GetAutoscalingOk

`func (o *ContainerResponse) GetAutoscalingOk() (*AutoscalingPolicyResponse, bool)`

GetAutoscalingOk returns a tuple with the Autoscaling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoscaling

`func (o *ContainerResponse) SetAutoscaling(v AutoscalingPolicyResponse)`

SetAutoscaling sets Autoscaling field to given value.

### HasAutoscaling

`func (o *ContainerResponse) HasAutoscaling() bool`

HasAutoscaling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


