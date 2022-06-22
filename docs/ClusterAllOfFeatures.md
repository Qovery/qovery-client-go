# ClusterAllOfFeatures

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**CostPerMonthInCents** | Pointer to **NullableInt32** |  | [optional] 
**CostPerMonth** | Pointer to **NullableFloat32** |  | [optional] 
**ValueType** | Pointer to **string** |  | [optional] 
**Value** | Pointer to [**NullableOneOfstringboolean**](oneOf&lt;string,boolean&gt;.md) |  | [optional] 
**IsValueUpdatable** | Pointer to **bool** |  | [optional] [default to false]
**AcceptedValues** | Pointer to [**[]OneOfstringboolean**](OneOfstringboolean.md) |  | [optional] 

## Methods

### NewClusterAllOfFeatures

`func NewClusterAllOfFeatures() *ClusterAllOfFeatures`

NewClusterAllOfFeatures instantiates a new ClusterAllOfFeatures object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterAllOfFeaturesWithDefaults

`func NewClusterAllOfFeaturesWithDefaults() *ClusterAllOfFeatures`

NewClusterAllOfFeaturesWithDefaults instantiates a new ClusterAllOfFeatures object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterAllOfFeatures) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterAllOfFeatures) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterAllOfFeatures) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterAllOfFeatures) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *ClusterAllOfFeatures) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ClusterAllOfFeatures) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ClusterAllOfFeatures) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ClusterAllOfFeatures) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *ClusterAllOfFeatures) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterAllOfFeatures) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterAllOfFeatures) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterAllOfFeatures) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ClusterAllOfFeatures) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ClusterAllOfFeatures) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCostPerMonthInCents

`func (o *ClusterAllOfFeatures) GetCostPerMonthInCents() int32`

GetCostPerMonthInCents returns the CostPerMonthInCents field if non-nil, zero value otherwise.

### GetCostPerMonthInCentsOk

`func (o *ClusterAllOfFeatures) GetCostPerMonthInCentsOk() (*int32, bool)`

GetCostPerMonthInCentsOk returns a tuple with the CostPerMonthInCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerMonthInCents

`func (o *ClusterAllOfFeatures) SetCostPerMonthInCents(v int32)`

SetCostPerMonthInCents sets CostPerMonthInCents field to given value.

### HasCostPerMonthInCents

`func (o *ClusterAllOfFeatures) HasCostPerMonthInCents() bool`

HasCostPerMonthInCents returns a boolean if a field has been set.

### SetCostPerMonthInCentsNil

`func (o *ClusterAllOfFeatures) SetCostPerMonthInCentsNil(b bool)`

 SetCostPerMonthInCentsNil sets the value for CostPerMonthInCents to be an explicit nil

### UnsetCostPerMonthInCents
`func (o *ClusterAllOfFeatures) UnsetCostPerMonthInCents()`

UnsetCostPerMonthInCents ensures that no value is present for CostPerMonthInCents, not even an explicit nil
### GetCostPerMonth

`func (o *ClusterAllOfFeatures) GetCostPerMonth() float32`

GetCostPerMonth returns the CostPerMonth field if non-nil, zero value otherwise.

### GetCostPerMonthOk

`func (o *ClusterAllOfFeatures) GetCostPerMonthOk() (*float32, bool)`

GetCostPerMonthOk returns a tuple with the CostPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerMonth

`func (o *ClusterAllOfFeatures) SetCostPerMonth(v float32)`

SetCostPerMonth sets CostPerMonth field to given value.

### HasCostPerMonth

`func (o *ClusterAllOfFeatures) HasCostPerMonth() bool`

HasCostPerMonth returns a boolean if a field has been set.

### SetCostPerMonthNil

`func (o *ClusterAllOfFeatures) SetCostPerMonthNil(b bool)`

 SetCostPerMonthNil sets the value for CostPerMonth to be an explicit nil

### UnsetCostPerMonth
`func (o *ClusterAllOfFeatures) UnsetCostPerMonth()`

UnsetCostPerMonth ensures that no value is present for CostPerMonth, not even an explicit nil
### GetValueType

`func (o *ClusterAllOfFeatures) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *ClusterAllOfFeatures) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *ClusterAllOfFeatures) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *ClusterAllOfFeatures) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### GetValue

`func (o *ClusterAllOfFeatures) GetValue() OneOfstringboolean`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ClusterAllOfFeatures) GetValueOk() (*OneOfstringboolean, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ClusterAllOfFeatures) SetValue(v OneOfstringboolean)`

SetValue sets Value field to given value.

### HasValue

`func (o *ClusterAllOfFeatures) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ClusterAllOfFeatures) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ClusterAllOfFeatures) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetIsValueUpdatable

`func (o *ClusterAllOfFeatures) GetIsValueUpdatable() bool`

GetIsValueUpdatable returns the IsValueUpdatable field if non-nil, zero value otherwise.

### GetIsValueUpdatableOk

`func (o *ClusterAllOfFeatures) GetIsValueUpdatableOk() (*bool, bool)`

GetIsValueUpdatableOk returns a tuple with the IsValueUpdatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValueUpdatable

`func (o *ClusterAllOfFeatures) SetIsValueUpdatable(v bool)`

SetIsValueUpdatable sets IsValueUpdatable field to given value.

### HasIsValueUpdatable

`func (o *ClusterAllOfFeatures) HasIsValueUpdatable() bool`

HasIsValueUpdatable returns a boolean if a field has been set.

### GetAcceptedValues

`func (o *ClusterAllOfFeatures) GetAcceptedValues() []OneOfstringboolean`

GetAcceptedValues returns the AcceptedValues field if non-nil, zero value otherwise.

### GetAcceptedValuesOk

`func (o *ClusterAllOfFeatures) GetAcceptedValuesOk() (*[]OneOfstringboolean, bool)`

GetAcceptedValuesOk returns a tuple with the AcceptedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedValues

`func (o *ClusterAllOfFeatures) SetAcceptedValues(v []OneOfstringboolean)`

SetAcceptedValues sets AcceptedValues field to given value.

### HasAcceptedValues

`func (o *ClusterAllOfFeatures) HasAcceptedValues() bool`

HasAcceptedValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


