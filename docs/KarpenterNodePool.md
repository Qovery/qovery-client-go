# KarpenterNodePool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requirements** | [**[]KarpenterNodePoolRequirement**](KarpenterNodePoolRequirement.md) |  | 
**StableOverride** | Pointer to [**KarpenterStableNodePoolOverride**](KarpenterStableNodePoolOverride.md) |  | [optional] 
**DefaultOverride** | Pointer to [**KarpenterDefaultNodePoolOverride**](KarpenterDefaultNodePoolOverride.md) |  | [optional] 

## Methods

### NewKarpenterNodePool

`func NewKarpenterNodePool(requirements []KarpenterNodePoolRequirement, ) *KarpenterNodePool`

NewKarpenterNodePool instantiates a new KarpenterNodePool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKarpenterNodePoolWithDefaults

`func NewKarpenterNodePoolWithDefaults() *KarpenterNodePool`

NewKarpenterNodePoolWithDefaults instantiates a new KarpenterNodePool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequirements

`func (o *KarpenterNodePool) GetRequirements() []KarpenterNodePoolRequirement`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *KarpenterNodePool) GetRequirementsOk() (*[]KarpenterNodePoolRequirement, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *KarpenterNodePool) SetRequirements(v []KarpenterNodePoolRequirement)`

SetRequirements sets Requirements field to given value.


### GetStableOverride

`func (o *KarpenterNodePool) GetStableOverride() KarpenterStableNodePoolOverride`

GetStableOverride returns the StableOverride field if non-nil, zero value otherwise.

### GetStableOverrideOk

`func (o *KarpenterNodePool) GetStableOverrideOk() (*KarpenterStableNodePoolOverride, bool)`

GetStableOverrideOk returns a tuple with the StableOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStableOverride

`func (o *KarpenterNodePool) SetStableOverride(v KarpenterStableNodePoolOverride)`

SetStableOverride sets StableOverride field to given value.

### HasStableOverride

`func (o *KarpenterNodePool) HasStableOverride() bool`

HasStableOverride returns a boolean if a field has been set.

### GetDefaultOverride

`func (o *KarpenterNodePool) GetDefaultOverride() KarpenterDefaultNodePoolOverride`

GetDefaultOverride returns the DefaultOverride field if non-nil, zero value otherwise.

### GetDefaultOverrideOk

`func (o *KarpenterNodePool) GetDefaultOverrideOk() (*KarpenterDefaultNodePoolOverride, bool)`

GetDefaultOverrideOk returns a tuple with the DefaultOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultOverride

`func (o *KarpenterNodePool) SetDefaultOverride(v KarpenterDefaultNodePoolOverride)`

SetDefaultOverride sets DefaultOverride field to given value.

### HasDefaultOverride

`func (o *KarpenterNodePool) HasDefaultOverride() bool`

HasDefaultOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


