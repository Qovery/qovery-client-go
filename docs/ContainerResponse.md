# ContainerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Storage** | Pointer to [**[]ApplicationStorageStorageInner**](ApplicationStorageStorageInner.md) |  | [optional] 
**Ports** | Pointer to [**[]ApplicationPortPortsInner**](ApplicationPortPortsInner.md) |  | [optional] 
**Environment** | Pointer to [**ReferenceObject**](ReferenceObject.md) |  | [optional] 
**MaximumCpu** | Pointer to **int32** | Maximum cpu that can be allocated to the container based on organization cluster configuration. unit is millicores (m). 1000m &#x3D; 1 cpu | [optional] [default to 250]
**MaximumMemory** | Pointer to **int32** | Maximum memory that can be allocated to the container based on organization cluster configuration. unit is MB. 1024 MB &#x3D; 1GB | [optional] [default to 256]
**Name** | Pointer to **string** | name is case insensitive | [optional] 
**Description** | Pointer to **NullableString** | give a description to this container | [optional] 
**RegistryId** | Pointer to **string** | id of the linked registry | [optional] 
**ImageName** | Pointer to **string** | name of the image container | [optional] 
**Arguments** | Pointer to **string** |  | [optional] 
**Cpu** | Pointer to **int32** | unit is millicores (m). 1000m &#x3D; 1 cpu | [optional] [default to 250]
**Memory** | Pointer to **int32** | unit is MB. 1024 MB &#x3D; 1GB | [optional] [default to 256]
**MinRunningInstances** | Pointer to **int32** | Minimum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: 0 means that there is no container running.  | [optional] [default to 1]
**MaxRunningInstances** | Pointer to **int32** | Maximum number of instances running. This resource auto-scale based on the CPU and Memory consumption. Note: -1 means that there is no limit.  | [optional] [default to 1]
**Healthcheck** | Pointer to [**Healthcheck**](Healthcheck.md) |  | [optional] 

## Methods

### NewContainerResponse

`func NewContainerResponse(id string, createdAt time.Time, ) *ContainerResponse`

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

`func (o *ContainerResponse) GetStorage() []ApplicationStorageStorageInner`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *ContainerResponse) GetStorageOk() (*[]ApplicationStorageStorageInner, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *ContainerResponse) SetStorage(v []ApplicationStorageStorageInner)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *ContainerResponse) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetPorts

`func (o *ContainerResponse) GetPorts() []ApplicationPortPortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ContainerResponse) GetPortsOk() (*[]ApplicationPortPortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ContainerResponse) SetPorts(v []ApplicationPortPortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ContainerResponse) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

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

### HasEnvironment

`func (o *ContainerResponse) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

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

### HasMaximumCpu

`func (o *ContainerResponse) HasMaximumCpu() bool`

HasMaximumCpu returns a boolean if a field has been set.

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

### HasMaximumMemory

`func (o *ContainerResponse) HasMaximumMemory() bool`

HasMaximumMemory returns a boolean if a field has been set.

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

### HasName

`func (o *ContainerResponse) HasName() bool`

HasName returns a boolean if a field has been set.

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

### SetDescriptionNil

`func (o *ContainerResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ContainerResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
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

### HasImageName

`func (o *ContainerResponse) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetArguments

`func (o *ContainerResponse) GetArguments() string`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *ContainerResponse) GetArgumentsOk() (*string, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *ContainerResponse) SetArguments(v string)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *ContainerResponse) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

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

### HasCpu

`func (o *ContainerResponse) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

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

### HasMemory

`func (o *ContainerResponse) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

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

### HasMinRunningInstances

`func (o *ContainerResponse) HasMinRunningInstances() bool`

HasMinRunningInstances returns a boolean if a field has been set.

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

### HasMaxRunningInstances

`func (o *ContainerResponse) HasMaxRunningInstances() bool`

HasMaxRunningInstances returns a boolean if a field has been set.

### GetHealthcheck

`func (o *ContainerResponse) GetHealthcheck() Healthcheck`

GetHealthcheck returns the Healthcheck field if non-nil, zero value otherwise.

### GetHealthcheckOk

`func (o *ContainerResponse) GetHealthcheckOk() (*Healthcheck, bool)`

GetHealthcheckOk returns a tuple with the Healthcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthcheck

`func (o *ContainerResponse) SetHealthcheck(v Healthcheck)`

SetHealthcheck sets Healthcheck field to given value.

### HasHealthcheck

`func (o *ContainerResponse) HasHealthcheck() bool`

HasHealthcheck returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


