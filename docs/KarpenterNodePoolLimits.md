# KarpenterNodePoolLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxCpuInVcpu** | **float32** | CPU limit that will be applied for the node pool (in vCPU unit: 1 vCPU &#x3D; 1000 millicores) | 
**MaxMemoryInGibibytes** | **float32** | Memory limit that will be applied for the node pool (in Gibibytes unit: 1Gi &#x3D; 1024 mebibytes) | 

## Methods

### NewKarpenterNodePoolLimits

`func NewKarpenterNodePoolLimits(maxCpuInVcpu float32, maxMemoryInGibibytes float32, ) *KarpenterNodePoolLimits`

NewKarpenterNodePoolLimits instantiates a new KarpenterNodePoolLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKarpenterNodePoolLimitsWithDefaults

`func NewKarpenterNodePoolLimitsWithDefaults() *KarpenterNodePoolLimits`

NewKarpenterNodePoolLimitsWithDefaults instantiates a new KarpenterNodePoolLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxCpuInVcpu

`func (o *KarpenterNodePoolLimits) GetMaxCpuInVcpu() float32`

GetMaxCpuInVcpu returns the MaxCpuInVcpu field if non-nil, zero value otherwise.

### GetMaxCpuInVcpuOk

`func (o *KarpenterNodePoolLimits) GetMaxCpuInVcpuOk() (*float32, bool)`

GetMaxCpuInVcpuOk returns a tuple with the MaxCpuInVcpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpuInVcpu

`func (o *KarpenterNodePoolLimits) SetMaxCpuInVcpu(v float32)`

SetMaxCpuInVcpu sets MaxCpuInVcpu field to given value.


### GetMaxMemoryInGibibytes

`func (o *KarpenterNodePoolLimits) GetMaxMemoryInGibibytes() float32`

GetMaxMemoryInGibibytes returns the MaxMemoryInGibibytes field if non-nil, zero value otherwise.

### GetMaxMemoryInGibibytesOk

`func (o *KarpenterNodePoolLimits) GetMaxMemoryInGibibytesOk() (*float32, bool)`

GetMaxMemoryInGibibytesOk returns a tuple with the MaxMemoryInGibibytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemoryInGibibytes

`func (o *KarpenterNodePoolLimits) SetMaxMemoryInGibibytes(v float32)`

SetMaxMemoryInGibibytes sets MaxMemoryInGibibytes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


