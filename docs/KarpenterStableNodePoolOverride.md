# KarpenterStableNodePoolOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Consolidation** | Pointer to [**KarpenterNodePoolConsolidation**](KarpenterNodePoolConsolidation.md) |  | [optional] 
**Limits** | Pointer to [**KarpenterNodePoolLimits**](KarpenterNodePoolLimits.md) |  | [optional] 
**ConsolidateAfter** | Pointer to **string** | Time to wait before consolidating empty or underutilized nodes (e.g., 1m, 10m, 1h). Maximum: 24h | [optional] 

## Methods

### NewKarpenterStableNodePoolOverride

`func NewKarpenterStableNodePoolOverride() *KarpenterStableNodePoolOverride`

NewKarpenterStableNodePoolOverride instantiates a new KarpenterStableNodePoolOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKarpenterStableNodePoolOverrideWithDefaults

`func NewKarpenterStableNodePoolOverrideWithDefaults() *KarpenterStableNodePoolOverride`

NewKarpenterStableNodePoolOverrideWithDefaults instantiates a new KarpenterStableNodePoolOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsolidation

`func (o *KarpenterStableNodePoolOverride) GetConsolidation() KarpenterNodePoolConsolidation`

GetConsolidation returns the Consolidation field if non-nil, zero value otherwise.

### GetConsolidationOk

`func (o *KarpenterStableNodePoolOverride) GetConsolidationOk() (*KarpenterNodePoolConsolidation, bool)`

GetConsolidationOk returns a tuple with the Consolidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolidation

`func (o *KarpenterStableNodePoolOverride) SetConsolidation(v KarpenterNodePoolConsolidation)`

SetConsolidation sets Consolidation field to given value.

### HasConsolidation

`func (o *KarpenterStableNodePoolOverride) HasConsolidation() bool`

HasConsolidation returns a boolean if a field has been set.

### GetLimits

`func (o *KarpenterStableNodePoolOverride) GetLimits() KarpenterNodePoolLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *KarpenterStableNodePoolOverride) GetLimitsOk() (*KarpenterNodePoolLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *KarpenterStableNodePoolOverride) SetLimits(v KarpenterNodePoolLimits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *KarpenterStableNodePoolOverride) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetConsolidateAfter

`func (o *KarpenterStableNodePoolOverride) GetConsolidateAfter() string`

GetConsolidateAfter returns the ConsolidateAfter field if non-nil, zero value otherwise.

### GetConsolidateAfterOk

`func (o *KarpenterStableNodePoolOverride) GetConsolidateAfterOk() (*string, bool)`

GetConsolidateAfterOk returns a tuple with the ConsolidateAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolidateAfter

`func (o *KarpenterStableNodePoolOverride) SetConsolidateAfter(v string)`

SetConsolidateAfter sets ConsolidateAfter field to given value.

### HasConsolidateAfter

`func (o *KarpenterStableNodePoolOverride) HasConsolidateAfter() bool`

HasConsolidateAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


