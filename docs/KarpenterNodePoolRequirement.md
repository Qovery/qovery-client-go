# KarpenterNodePoolRequirement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | [**KarpenterNodePoolRequirementKey**](KarpenterNodePoolRequirementKey.md) |  | 
**Operator** | [**KarpenterNodePoolRequirementOperator**](KarpenterNodePoolRequirementOperator.md) |  | 
**Values** | **[]string** |  | 

## Methods

### NewKarpenterNodePoolRequirement

`func NewKarpenterNodePoolRequirement(key KarpenterNodePoolRequirementKey, operator KarpenterNodePoolRequirementOperator, values []string, ) *KarpenterNodePoolRequirement`

NewKarpenterNodePoolRequirement instantiates a new KarpenterNodePoolRequirement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKarpenterNodePoolRequirementWithDefaults

`func NewKarpenterNodePoolRequirementWithDefaults() *KarpenterNodePoolRequirement`

NewKarpenterNodePoolRequirementWithDefaults instantiates a new KarpenterNodePoolRequirement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *KarpenterNodePoolRequirement) GetKey() KarpenterNodePoolRequirementKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *KarpenterNodePoolRequirement) GetKeyOk() (*KarpenterNodePoolRequirementKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *KarpenterNodePoolRequirement) SetKey(v KarpenterNodePoolRequirementKey)`

SetKey sets Key field to given value.


### GetOperator

`func (o *KarpenterNodePoolRequirement) GetOperator() KarpenterNodePoolRequirementOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *KarpenterNodePoolRequirement) GetOperatorOk() (*KarpenterNodePoolRequirementOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *KarpenterNodePoolRequirement) SetOperator(v KarpenterNodePoolRequirementOperator)`

SetOperator sets Operator field to given value.


### GetValues

`func (o *KarpenterNodePoolRequirement) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *KarpenterNodePoolRequirement) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *KarpenterNodePoolRequirement) SetValues(v []string)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


