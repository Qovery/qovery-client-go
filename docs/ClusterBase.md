# ClusterBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name is case-insensitive | 
**Description** | Pointer to **string** |  | [optional] 
**CloudProvider** | [**CloudProviderEnum**](CloudProviderEnum.md) |  | 
**Region** | **string** |  | 
**AutoUpdate** | Pointer to **bool** |  | [optional] 
**Cpu** | Pointer to **int32** | unit is millicores (m). 1000m &#x3D; 1 cpu | [optional] [default to 250]
**Memory** | Pointer to **int32** | unit is MB. 1024 MB &#x3D; 1GB | [optional] [default to 256]
**MinRunningNodes** | Pointer to **int32** |  | [optional] [default to 1]
**MaxRunningNodes** | Pointer to **int32** |  | [optional] [default to 1]
**InstanceType** | Pointer to **string** | the instance type to be used for this cluster. The list of values can be retrieved via the endpoint /{CloudProvider}/instanceType | [optional] 

## Methods

### NewClusterBase

`func NewClusterBase(name string, cloudProvider CloudProviderEnum, region string, ) *ClusterBase`

NewClusterBase instantiates a new ClusterBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterBaseWithDefaults

`func NewClusterBaseWithDefaults() *ClusterBase`

NewClusterBaseWithDefaults instantiates a new ClusterBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClusterBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterBase) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ClusterBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCloudProvider

`func (o *ClusterBase) GetCloudProvider() CloudProviderEnum`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *ClusterBase) GetCloudProviderOk() (*CloudProviderEnum, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *ClusterBase) SetCloudProvider(v CloudProviderEnum)`

SetCloudProvider sets CloudProvider field to given value.


### GetRegion

`func (o *ClusterBase) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ClusterBase) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ClusterBase) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetAutoUpdate

`func (o *ClusterBase) GetAutoUpdate() bool`

GetAutoUpdate returns the AutoUpdate field if non-nil, zero value otherwise.

### GetAutoUpdateOk

`func (o *ClusterBase) GetAutoUpdateOk() (*bool, bool)`

GetAutoUpdateOk returns a tuple with the AutoUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpdate

`func (o *ClusterBase) SetAutoUpdate(v bool)`

SetAutoUpdate sets AutoUpdate field to given value.

### HasAutoUpdate

`func (o *ClusterBase) HasAutoUpdate() bool`

HasAutoUpdate returns a boolean if a field has been set.

### GetCpu

`func (o *ClusterBase) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ClusterBase) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ClusterBase) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ClusterBase) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *ClusterBase) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ClusterBase) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ClusterBase) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ClusterBase) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetMinRunningNodes

`func (o *ClusterBase) GetMinRunningNodes() int32`

GetMinRunningNodes returns the MinRunningNodes field if non-nil, zero value otherwise.

### GetMinRunningNodesOk

`func (o *ClusterBase) GetMinRunningNodesOk() (*int32, bool)`

GetMinRunningNodesOk returns a tuple with the MinRunningNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRunningNodes

`func (o *ClusterBase) SetMinRunningNodes(v int32)`

SetMinRunningNodes sets MinRunningNodes field to given value.

### HasMinRunningNodes

`func (o *ClusterBase) HasMinRunningNodes() bool`

HasMinRunningNodes returns a boolean if a field has been set.

### GetMaxRunningNodes

`func (o *ClusterBase) GetMaxRunningNodes() int32`

GetMaxRunningNodes returns the MaxRunningNodes field if non-nil, zero value otherwise.

### GetMaxRunningNodesOk

`func (o *ClusterBase) GetMaxRunningNodesOk() (*int32, bool)`

GetMaxRunningNodesOk returns a tuple with the MaxRunningNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRunningNodes

`func (o *ClusterBase) SetMaxRunningNodes(v int32)`

SetMaxRunningNodes sets MaxRunningNodes field to given value.

### HasMaxRunningNodes

`func (o *ClusterBase) HasMaxRunningNodes() bool`

HasMaxRunningNodes returns a boolean if a field has been set.

### GetInstanceType

`func (o *ClusterBase) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *ClusterBase) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *ClusterBase) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *ClusterBase) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


