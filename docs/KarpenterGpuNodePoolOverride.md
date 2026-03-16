# KarpenterGpuNodePoolOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Consolidation** | Pointer to [**KarpenterNodePoolConsolidation**](KarpenterNodePoolConsolidation.md) |  | [optional] 
**Limits** | Pointer to [**KarpenterNodePoolLimits**](KarpenterNodePoolLimits.md) |  | [optional] 
**Requirements** | Pointer to [**[]KarpenterNodePoolRequirement**](KarpenterNodePoolRequirement.md) |  | [optional] 
**DiskSizeInGib** | Pointer to **int32** |  | [optional] [default to 100]
**DiskIops** | Pointer to **int32** | Unit is operation/seconds. The disk IOPS to be used for the GPU node pool configuration | [optional] 
**DiskThroughput** | Pointer to **int32** | Unit is in MB/s. The disk throughput to be used for the GPU node pool configuration | [optional] 
**SpotEnabled** | Pointer to **bool** |  | [optional] [default to false]
**ConsolidateAfter** | Pointer to **string** | Time to wait before consolidating empty or underutilized nodes (e.g., 1m, 10m, 1h). Maximum: 24h | [optional] 

## Methods

### NewKarpenterGpuNodePoolOverride

`func NewKarpenterGpuNodePoolOverride() *KarpenterGpuNodePoolOverride`

NewKarpenterGpuNodePoolOverride instantiates a new KarpenterGpuNodePoolOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKarpenterGpuNodePoolOverrideWithDefaults

`func NewKarpenterGpuNodePoolOverrideWithDefaults() *KarpenterGpuNodePoolOverride`

NewKarpenterGpuNodePoolOverrideWithDefaults instantiates a new KarpenterGpuNodePoolOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsolidation

`func (o *KarpenterGpuNodePoolOverride) GetConsolidation() KarpenterNodePoolConsolidation`

GetConsolidation returns the Consolidation field if non-nil, zero value otherwise.

### GetConsolidationOk

`func (o *KarpenterGpuNodePoolOverride) GetConsolidationOk() (*KarpenterNodePoolConsolidation, bool)`

GetConsolidationOk returns a tuple with the Consolidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolidation

`func (o *KarpenterGpuNodePoolOverride) SetConsolidation(v KarpenterNodePoolConsolidation)`

SetConsolidation sets Consolidation field to given value.

### HasConsolidation

`func (o *KarpenterGpuNodePoolOverride) HasConsolidation() bool`

HasConsolidation returns a boolean if a field has been set.

### GetLimits

`func (o *KarpenterGpuNodePoolOverride) GetLimits() KarpenterNodePoolLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *KarpenterGpuNodePoolOverride) GetLimitsOk() (*KarpenterNodePoolLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *KarpenterGpuNodePoolOverride) SetLimits(v KarpenterNodePoolLimits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *KarpenterGpuNodePoolOverride) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetRequirements

`func (o *KarpenterGpuNodePoolOverride) GetRequirements() []KarpenterNodePoolRequirement`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *KarpenterGpuNodePoolOverride) GetRequirementsOk() (*[]KarpenterNodePoolRequirement, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *KarpenterGpuNodePoolOverride) SetRequirements(v []KarpenterNodePoolRequirement)`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *KarpenterGpuNodePoolOverride) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.

### GetDiskSizeInGib

`func (o *KarpenterGpuNodePoolOverride) GetDiskSizeInGib() int32`

GetDiskSizeInGib returns the DiskSizeInGib field if non-nil, zero value otherwise.

### GetDiskSizeInGibOk

`func (o *KarpenterGpuNodePoolOverride) GetDiskSizeInGibOk() (*int32, bool)`

GetDiskSizeInGibOk returns a tuple with the DiskSizeInGib field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeInGib

`func (o *KarpenterGpuNodePoolOverride) SetDiskSizeInGib(v int32)`

SetDiskSizeInGib sets DiskSizeInGib field to given value.

### HasDiskSizeInGib

`func (o *KarpenterGpuNodePoolOverride) HasDiskSizeInGib() bool`

HasDiskSizeInGib returns a boolean if a field has been set.

### GetDiskIops

`func (o *KarpenterGpuNodePoolOverride) GetDiskIops() int32`

GetDiskIops returns the DiskIops field if non-nil, zero value otherwise.

### GetDiskIopsOk

`func (o *KarpenterGpuNodePoolOverride) GetDiskIopsOk() (*int32, bool)`

GetDiskIopsOk returns a tuple with the DiskIops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskIops

`func (o *KarpenterGpuNodePoolOverride) SetDiskIops(v int32)`

SetDiskIops sets DiskIops field to given value.

### HasDiskIops

`func (o *KarpenterGpuNodePoolOverride) HasDiskIops() bool`

HasDiskIops returns a boolean if a field has been set.

### GetDiskThroughput

`func (o *KarpenterGpuNodePoolOverride) GetDiskThroughput() int32`

GetDiskThroughput returns the DiskThroughput field if non-nil, zero value otherwise.

### GetDiskThroughputOk

`func (o *KarpenterGpuNodePoolOverride) GetDiskThroughputOk() (*int32, bool)`

GetDiskThroughputOk returns a tuple with the DiskThroughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskThroughput

`func (o *KarpenterGpuNodePoolOverride) SetDiskThroughput(v int32)`

SetDiskThroughput sets DiskThroughput field to given value.

### HasDiskThroughput

`func (o *KarpenterGpuNodePoolOverride) HasDiskThroughput() bool`

HasDiskThroughput returns a boolean if a field has been set.

### GetSpotEnabled

`func (o *KarpenterGpuNodePoolOverride) GetSpotEnabled() bool`

GetSpotEnabled returns the SpotEnabled field if non-nil, zero value otherwise.

### GetSpotEnabledOk

`func (o *KarpenterGpuNodePoolOverride) GetSpotEnabledOk() (*bool, bool)`

GetSpotEnabledOk returns a tuple with the SpotEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotEnabled

`func (o *KarpenterGpuNodePoolOverride) SetSpotEnabled(v bool)`

SetSpotEnabled sets SpotEnabled field to given value.

### HasSpotEnabled

`func (o *KarpenterGpuNodePoolOverride) HasSpotEnabled() bool`

HasSpotEnabled returns a boolean if a field has been set.

### GetConsolidateAfter

`func (o *KarpenterGpuNodePoolOverride) GetConsolidateAfter() string`

GetConsolidateAfter returns the ConsolidateAfter field if non-nil, zero value otherwise.

### GetConsolidateAfterOk

`func (o *KarpenterGpuNodePoolOverride) GetConsolidateAfterOk() (*string, bool)`

GetConsolidateAfterOk returns a tuple with the ConsolidateAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolidateAfter

`func (o *KarpenterGpuNodePoolOverride) SetConsolidateAfter(v string)`

SetConsolidateAfter sets ConsolidateAfter field to given value.

### HasConsolidateAfter

`func (o *KarpenterGpuNodePoolOverride) HasConsolidateAfter() bool`

HasConsolidateAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


