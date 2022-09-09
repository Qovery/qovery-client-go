# ContainerResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | [**ReferenceObject**](ReferenceObject.md) |  | 
**Registry** | [**ReferenceObject**](ReferenceObject.md) |  | 
**MaximumCpu** | **int32** | Maximum cpu that can be allocated to the container based on organization cluster configuration. unit is millicores (m). 1000m &#x3D; 1 cpu | 
**MaximumMemory** | **int32** | Maximum memory that can be allocated to the container based on organization cluster configuration. unit is MB. 1024 MB &#x3D; 1GB | 
**Name** | **string** | name is case insensitive | 
**ImageName** | **string** | name of the image container | 
**Tag** | **string** | tag of the image container | 
**Arguments** | Pointer to **[]string** |  | [optional] 
**Entrypoint** | Pointer to **string** | optional entrypoint when launching container | [optional] 
**Cpu** | **int32** | unit is millicores (m). 1000m &#x3D; 1 cpu | 
**Memory** | **int32** | unit is MB. 1024 MB &#x3D; 1GB | 
**MinRunningInstances** | **int32** | Minimum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: 0 means that there is no container running.  | [default to 1]
**MaxRunningInstances** | **int32** | Maximum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: -1 means that there is no limit.  | [default to 1]
**AutoPreview** | **bool** | Specify if the environment preview option is activated or not for this container. If activated, a preview environment will be automatically cloned at each pull request.  | 
**Ports** | Pointer to [**ServicePortResponseList**](ServicePortResponseList.md) |  | [optional] 

## Methods

### NewContainerResponseAllOf

`func NewContainerResponseAllOf(environment ReferenceObject, registry ReferenceObject, maximumCpu int32, maximumMemory int32, name string, imageName string, tag string, cpu int32, memory int32, minRunningInstances int32, maxRunningInstances int32, autoPreview bool, ) *ContainerResponseAllOf`

NewContainerResponseAllOf instantiates a new ContainerResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerResponseAllOfWithDefaults

`func NewContainerResponseAllOfWithDefaults() *ContainerResponseAllOf`

NewContainerResponseAllOfWithDefaults instantiates a new ContainerResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *ContainerResponseAllOf) GetEnvironment() ReferenceObject`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ContainerResponseAllOf) GetEnvironmentOk() (*ReferenceObject, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ContainerResponseAllOf) SetEnvironment(v ReferenceObject)`

SetEnvironment sets Environment field to given value.


### GetRegistry

`func (o *ContainerResponseAllOf) GetRegistry() ReferenceObject`

GetRegistry returns the Registry field if non-nil, zero value otherwise.

### GetRegistryOk

`func (o *ContainerResponseAllOf) GetRegistryOk() (*ReferenceObject, bool)`

GetRegistryOk returns a tuple with the Registry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistry

`func (o *ContainerResponseAllOf) SetRegistry(v ReferenceObject)`

SetRegistry sets Registry field to given value.


### GetMaximumCpu

`func (o *ContainerResponseAllOf) GetMaximumCpu() int32`

GetMaximumCpu returns the MaximumCpu field if non-nil, zero value otherwise.

### GetMaximumCpuOk

`func (o *ContainerResponseAllOf) GetMaximumCpuOk() (*int32, bool)`

GetMaximumCpuOk returns a tuple with the MaximumCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumCpu

`func (o *ContainerResponseAllOf) SetMaximumCpu(v int32)`

SetMaximumCpu sets MaximumCpu field to given value.


### GetMaximumMemory

`func (o *ContainerResponseAllOf) GetMaximumMemory() int32`

GetMaximumMemory returns the MaximumMemory field if non-nil, zero value otherwise.

### GetMaximumMemoryOk

`func (o *ContainerResponseAllOf) GetMaximumMemoryOk() (*int32, bool)`

GetMaximumMemoryOk returns a tuple with the MaximumMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumMemory

`func (o *ContainerResponseAllOf) SetMaximumMemory(v int32)`

SetMaximumMemory sets MaximumMemory field to given value.


### GetName

`func (o *ContainerResponseAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContainerResponseAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContainerResponseAllOf) SetName(v string)`

SetName sets Name field to given value.


### GetImageName

`func (o *ContainerResponseAllOf) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *ContainerResponseAllOf) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *ContainerResponseAllOf) SetImageName(v string)`

SetImageName sets ImageName field to given value.


### GetTag

`func (o *ContainerResponseAllOf) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ContainerResponseAllOf) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ContainerResponseAllOf) SetTag(v string)`

SetTag sets Tag field to given value.


### GetArguments

`func (o *ContainerResponseAllOf) GetArguments() []string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *ContainerResponseAllOf) GetArgumentsOk() (*[]string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *ContainerResponseAllOf) SetArguments(v []string)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *ContainerResponseAllOf) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### GetEntrypoint

`func (o *ContainerResponseAllOf) GetEntrypoint() string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *ContainerResponseAllOf) GetEntrypointOk() (*string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *ContainerResponseAllOf) SetEntrypoint(v string)`

SetEntrypoint sets Entrypoint field to given value.

### HasEntrypoint

`func (o *ContainerResponseAllOf) HasEntrypoint() bool`

HasEntrypoint returns a boolean if a field has been set.

### GetCpu

`func (o *ContainerResponseAllOf) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ContainerResponseAllOf) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ContainerResponseAllOf) SetCpu(v int32)`

SetCpu sets Cpu field to given value.


### GetMemory

`func (o *ContainerResponseAllOf) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ContainerResponseAllOf) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ContainerResponseAllOf) SetMemory(v int32)`

SetMemory sets Memory field to given value.


### GetMinRunningInstances

`func (o *ContainerResponseAllOf) GetMinRunningInstances() int32`

GetMinRunningInstances returns the MinRunningInstances field if non-nil, zero value otherwise.

### GetMinRunningInstancesOk

`func (o *ContainerResponseAllOf) GetMinRunningInstancesOk() (*int32, bool)`

GetMinRunningInstancesOk returns a tuple with the MinRunningInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRunningInstances

`func (o *ContainerResponseAllOf) SetMinRunningInstances(v int32)`

SetMinRunningInstances sets MinRunningInstances field to given value.


### GetMaxRunningInstances

`func (o *ContainerResponseAllOf) GetMaxRunningInstances() int32`

GetMaxRunningInstances returns the MaxRunningInstances field if non-nil, zero value otherwise.

### GetMaxRunningInstancesOk

`func (o *ContainerResponseAllOf) GetMaxRunningInstancesOk() (*int32, bool)`

GetMaxRunningInstancesOk returns a tuple with the MaxRunningInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRunningInstances

`func (o *ContainerResponseAllOf) SetMaxRunningInstances(v int32)`

SetMaxRunningInstances sets MaxRunningInstances field to given value.


### GetAutoPreview

`func (o *ContainerResponseAllOf) GetAutoPreview() bool`

GetAutoPreview returns the AutoPreview field if non-nil, zero value otherwise.

### GetAutoPreviewOk

`func (o *ContainerResponseAllOf) GetAutoPreviewOk() (*bool, bool)`

GetAutoPreviewOk returns a tuple with the AutoPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPreview

`func (o *ContainerResponseAllOf) SetAutoPreview(v bool)`

SetAutoPreview sets AutoPreview field to given value.


### GetPorts

`func (o *ContainerResponseAllOf) GetPorts() ServicePortResponseList`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ContainerResponseAllOf) GetPortsOk() (*ServicePortResponseList, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ContainerResponseAllOf) SetPorts(v ServicePortResponseList)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ContainerResponseAllOf) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


