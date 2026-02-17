# KarpenterDefaultNodePoolOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limits** | Pointer to [**KarpenterNodePoolLimits**](KarpenterNodePoolLimits.md) |  | [optional] 
**ConsolidateAfter** | Pointer to **string** | Time to wait before consolidating empty or underutilized nodes (e.g., 1m, 10m, 1h). Maximum: 24h | [optional] 

## Methods

### NewKarpenterDefaultNodePoolOverride

`func NewKarpenterDefaultNodePoolOverride() *KarpenterDefaultNodePoolOverride`

NewKarpenterDefaultNodePoolOverride instantiates a new KarpenterDefaultNodePoolOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKarpenterDefaultNodePoolOverrideWithDefaults

`func NewKarpenterDefaultNodePoolOverrideWithDefaults() *KarpenterDefaultNodePoolOverride`

NewKarpenterDefaultNodePoolOverrideWithDefaults instantiates a new KarpenterDefaultNodePoolOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimits

`func (o *KarpenterDefaultNodePoolOverride) GetLimits() KarpenterNodePoolLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *KarpenterDefaultNodePoolOverride) GetLimitsOk() (*KarpenterNodePoolLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *KarpenterDefaultNodePoolOverride) SetLimits(v KarpenterNodePoolLimits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *KarpenterDefaultNodePoolOverride) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetConsolidateAfter

`func (o *KarpenterDefaultNodePoolOverride) GetConsolidateAfter() string`

GetConsolidateAfter returns the ConsolidateAfter field if non-nil, zero value otherwise.

### GetConsolidateAfterOk

`func (o *KarpenterDefaultNodePoolOverride) GetConsolidateAfterOk() (*string, bool)`

GetConsolidateAfterOk returns a tuple with the ConsolidateAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsolidateAfter

`func (o *KarpenterDefaultNodePoolOverride) SetConsolidateAfter(v string)`

SetConsolidateAfter sets ConsolidateAfter field to given value.

### HasConsolidateAfter

`func (o *KarpenterDefaultNodePoolOverride) HasConsolidateAfter() bool`

HasConsolidateAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


