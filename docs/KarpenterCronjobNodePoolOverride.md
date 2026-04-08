# KarpenterCronjobNodePoolOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Consolidation** | Pointer to [**KarpenterNodePoolConsolidation**](KarpenterNodePoolConsolidation.md) |  | [optional] 
**Limits** | Pointer to [**KarpenterNodePoolLimits**](KarpenterNodePoolLimits.md) |  | [optional] 
**ConsolidateAfter** | Pointer to **string** | Time to wait before consolidating empty or underutilized nodes (e.g., 1m, 10m, 1h). Maximum: 24h | [optional] 

## Methods

### NewKarpenterCronjobNodePoolOverride

`func NewKarpenterCronjobNodePoolOverride() *KarpenterCronjobNodePoolOverride`

NewKarpenterCronjobNodePoolOverride instantiates a new KarpenterCronjobNodePoolOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKarpenterCronjobNodePoolOverrideWithDefaults

`func NewKarpenterCronjobNodePoolOverrideWithDefaults() *KarpenterCronjobNodePoolOverride`

NewKarpenterCronjobNodePoolOverrideWithDefaults instantiates a new KarpenterCronjobNodePoolOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsolidation

`func (o *KarpenterCronjobNodePoolOverride) GetConsolidation() KarpenterNodePoolConsolidation`

GetConsolidation returns the Consolidation field if non-nil, zero value otherwise.

### GetConsolidationOk

`func (o *KarpenterCronjobNodePoolOverride) GetConsolidationOk() (*KarpenterNodePoolConsolidation, bool)`

GetConsolidationOk returns a tuple with the Consolidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolidation

`func (o *KarpenterCronjobNodePoolOverride) SetConsolidation(v KarpenterNodePoolConsolidation)`

SetConsolidation sets Consolidation field to given value.

### HasConsolidation

`func (o *KarpenterCronjobNodePoolOverride) HasConsolidation() bool`

HasConsolidation returns a boolean if a field has been set.

### GetLimits

`func (o *KarpenterCronjobNodePoolOverride) GetLimits() KarpenterNodePoolLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *KarpenterCronjobNodePoolOverride) GetLimitsOk() (*KarpenterNodePoolLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *KarpenterCronjobNodePoolOverride) SetLimits(v KarpenterNodePoolLimits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *KarpenterCronjobNodePoolOverride) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetConsolidateAfter

`func (o *KarpenterCronjobNodePoolOverride) GetConsolidateAfter() string`

GetConsolidateAfter returns the ConsolidateAfter field if non-nil, zero value otherwise.

### GetConsolidateAfterOk

`func (o *KarpenterCronjobNodePoolOverride) GetConsolidateAfterOk() (*string, bool)`

GetConsolidateAfterOk returns a tuple with the ConsolidateAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolidateAfter

`func (o *KarpenterCronjobNodePoolOverride) SetConsolidateAfter(v string)`

SetConsolidateAfter sets ConsolidateAfter field to given value.

### HasConsolidateAfter

`func (o *KarpenterCronjobNodePoolOverride) HasConsolidateAfter() bool`

HasConsolidateAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


