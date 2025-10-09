# KarpenterNodePoolLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | [default to false]
**MaxCpuInVcpu** | **int32** | CPU limit that will be applied for the node pool (in vCPU unit: 1 vCPU &#x3D; 1000 millicores) | 
**MaxMemoryInGibibytes** | **int32** | Memory limit that will be applied for the node pool (in Gibibytes unit: 1Gi &#x3D; 1024 mebibytes) | 
**MaxGpu** | **int32** | GPU limit that will be applied for the node pool | [default to 0]

## Methods

### NewKarpenterNodePoolLimits

`func NewKarpenterNodePoolLimits(enabled bool, maxCpuInVcpu int32, maxMemoryInGibibytes int32, maxGpu int32, ) *KarpenterNodePoolLimits`

NewKarpenterNodePoolLimits instantiates a new KarpenterNodePoolLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKarpenterNodePoolLimitsWithDefaults

`func NewKarpenterNodePoolLimitsWithDefaults() *KarpenterNodePoolLimits`

NewKarpenterNodePoolLimitsWithDefaults instantiates a new KarpenterNodePoolLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *KarpenterNodePoolLimits) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KarpenterNodePoolLimits) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KarpenterNodePoolLimits) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMaxCpuInVcpu

`func (o *KarpenterNodePoolLimits) GetMaxCpuInVcpu() int32`

GetMaxCpuInVcpu returns the MaxCpuInVcpu field if non-nil, zero value otherwise.

### GetMaxCpuInVcpuOk

`func (o *KarpenterNodePoolLimits) GetMaxCpuInVcpuOk() (*int32, bool)`

GetMaxCpuInVcpuOk returns a tuple with the MaxCpuInVcpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpuInVcpu

`func (o *KarpenterNodePoolLimits) SetMaxCpuInVcpu(v int32)`

SetMaxCpuInVcpu sets MaxCpuInVcpu field to given value.


### GetMaxMemoryInGibibytes

`func (o *KarpenterNodePoolLimits) GetMaxMemoryInGibibytes() int32`

GetMaxMemoryInGibibytes returns the MaxMemoryInGibibytes field if non-nil, zero value otherwise.

### GetMaxMemoryInGibibytesOk

`func (o *KarpenterNodePoolLimits) GetMaxMemoryInGibibytesOk() (*int32, bool)`

GetMaxMemoryInGibibytesOk returns a tuple with the MaxMemoryInGibibytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemoryInGibibytes

`func (o *KarpenterNodePoolLimits) SetMaxMemoryInGibibytes(v int32)`

SetMaxMemoryInGibibytes sets MaxMemoryInGibibytes field to given value.


### GetMaxGpu

`func (o *KarpenterNodePoolLimits) GetMaxGpu() int32`

GetMaxGpu returns the MaxGpu field if non-nil, zero value otherwise.

### GetMaxGpuOk

`func (o *KarpenterNodePoolLimits) GetMaxGpuOk() (*int32, bool)`

GetMaxGpuOk returns a tuple with the MaxGpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGpu

`func (o *KarpenterNodePoolLimits) SetMaxGpu(v int32)`

SetMaxGpu sets MaxGpu field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


