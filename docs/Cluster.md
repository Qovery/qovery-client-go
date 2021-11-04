# Cluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | name is case-insensitive | 
**Description** | Pointer to **string** |  | [optional] 
**CloudProvider** | **string** |  | 
**Region** | **string** |  | 
**AutoUpdate** | Pointer to **bool** |  | [optional] 
**Cpu** | Pointer to **float32** | unit is millicores (m). 1000m &#x3D; 1 cpu | [optional] [default to 250]
**Memory** | Pointer to **float32** | unit is MB. 1024 MB &#x3D; 1GB | [optional] [default to 256]
**MinRunningNodes** | Pointer to **int32** |  | [optional] [default to 1]
**MaxRunningNodes** | Pointer to **int32** |  | [optional] [default to 1]

## Methods

### NewCluster

`func NewCluster(name string, cloudProvider string, region string, ) *Cluster`

NewCluster instantiates a new Cluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterWithDefaults

`func NewClusterWithDefaults() *Cluster`

NewClusterWithDefaults instantiates a new Cluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Cluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Cluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Cluster) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Cluster) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Cluster) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Cluster) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Cluster) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCloudProvider

`func (o *Cluster) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *Cluster) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *Cluster) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetRegion

`func (o *Cluster) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Cluster) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Cluster) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetAutoUpdate

`func (o *Cluster) GetAutoUpdate() bool`

GetAutoUpdate returns the AutoUpdate field if non-nil, zero value otherwise.

### GetAutoUpdateOk

`func (o *Cluster) GetAutoUpdateOk() (*bool, bool)`

GetAutoUpdateOk returns a tuple with the AutoUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpdate

`func (o *Cluster) SetAutoUpdate(v bool)`

SetAutoUpdate sets AutoUpdate field to given value.

### HasAutoUpdate

`func (o *Cluster) HasAutoUpdate() bool`

HasAutoUpdate returns a boolean if a field has been set.

### GetCpu

`func (o *Cluster) GetCpu() float32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *Cluster) GetCpuOk() (*float32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *Cluster) SetCpu(v float32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *Cluster) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *Cluster) GetMemory() float32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *Cluster) GetMemoryOk() (*float32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *Cluster) SetMemory(v float32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *Cluster) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetMinRunningNodes

`func (o *Cluster) GetMinRunningNodes() int32`

GetMinRunningNodes returns the MinRunningNodes field if non-nil, zero value otherwise.

### GetMinRunningNodesOk

`func (o *Cluster) GetMinRunningNodesOk() (*int32, bool)`

GetMinRunningNodesOk returns a tuple with the MinRunningNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRunningNodes

`func (o *Cluster) SetMinRunningNodes(v int32)`

SetMinRunningNodes sets MinRunningNodes field to given value.

### HasMinRunningNodes

`func (o *Cluster) HasMinRunningNodes() bool`

HasMinRunningNodes returns a boolean if a field has been set.

### GetMaxRunningNodes

`func (o *Cluster) GetMaxRunningNodes() int32`

GetMaxRunningNodes returns the MaxRunningNodes field if non-nil, zero value otherwise.

### GetMaxRunningNodesOk

`func (o *Cluster) GetMaxRunningNodesOk() (*int32, bool)`

GetMaxRunningNodesOk returns a tuple with the MaxRunningNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRunningNodes

`func (o *Cluster) SetMaxRunningNodes(v int32)`

SetMaxRunningNodes sets MaxRunningNodes field to given value.

### HasMaxRunningNodes

`func (o *Cluster) HasMaxRunningNodes() bool`

HasMaxRunningNodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


