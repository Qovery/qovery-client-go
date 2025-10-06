# KarpenterGpuNodePoolOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Consolidation** | Pointer to [**KarpenterNodePoolConsolidation**](KarpenterNodePoolConsolidation.md) |  | [optional] 
**Limits** | Pointer to [**KarpenterNodePoolLimits**](KarpenterNodePoolLimits.md) |  | [optional] 
**Requirements** | Pointer to [**[]KarpenterNodePoolRequirement**](KarpenterNodePoolRequirement.md) |  | [optional] 
**DiskSizeInGib** | Pointer to **int32** |  | [optional] [default to 100]

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


