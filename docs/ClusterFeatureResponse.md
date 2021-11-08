# ClusterFeatureResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**CostPerMonthInCents** | Pointer to **NullableInt32** |  | [optional] 
**CostPerMonth** | Pointer to **NullableFloat32** |  | [optional] 
**CurrencyCode** | Pointer to **NullableString** |  | [optional] 
**ValueType** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **NullableString** |  | [optional] 
**IsValueUpdatable** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewClusterFeatureResponse

`func NewClusterFeatureResponse() *ClusterFeatureResponse`

NewClusterFeatureResponse instantiates a new ClusterFeatureResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterFeatureResponseWithDefaults

`func NewClusterFeatureResponseWithDefaults() *ClusterFeatureResponse`

NewClusterFeatureResponseWithDefaults instantiates a new ClusterFeatureResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterFeatureResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterFeatureResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterFeatureResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterFeatureResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *ClusterFeatureResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ClusterFeatureResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ClusterFeatureResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ClusterFeatureResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *ClusterFeatureResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterFeatureResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterFeatureResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterFeatureResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ClusterFeatureResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ClusterFeatureResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCostPerMonthInCents

`func (o *ClusterFeatureResponse) GetCostPerMonthInCents() int32`

GetCostPerMonthInCents returns the CostPerMonthInCents field if non-nil, zero value otherwise.

### GetCostPerMonthInCentsOk

`func (o *ClusterFeatureResponse) GetCostPerMonthInCentsOk() (*int32, bool)`

GetCostPerMonthInCentsOk returns a tuple with the CostPerMonthInCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerMonthInCents

`func (o *ClusterFeatureResponse) SetCostPerMonthInCents(v int32)`

SetCostPerMonthInCents sets CostPerMonthInCents field to given value.

### HasCostPerMonthInCents

`func (o *ClusterFeatureResponse) HasCostPerMonthInCents() bool`

HasCostPerMonthInCents returns a boolean if a field has been set.

### SetCostPerMonthInCentsNil

`func (o *ClusterFeatureResponse) SetCostPerMonthInCentsNil(b bool)`

 SetCostPerMonthInCentsNil sets the value for CostPerMonthInCents to be an explicit nil

### UnsetCostPerMonthInCents
`func (o *ClusterFeatureResponse) UnsetCostPerMonthInCents()`

UnsetCostPerMonthInCents ensures that no value is present for CostPerMonthInCents, not even an explicit nil
### GetCostPerMonth

`func (o *ClusterFeatureResponse) GetCostPerMonth() float32`

GetCostPerMonth returns the CostPerMonth field if non-nil, zero value otherwise.

### GetCostPerMonthOk

`func (o *ClusterFeatureResponse) GetCostPerMonthOk() (*float32, bool)`

GetCostPerMonthOk returns a tuple with the CostPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerMonth

`func (o *ClusterFeatureResponse) SetCostPerMonth(v float32)`

SetCostPerMonth sets CostPerMonth field to given value.

### HasCostPerMonth

`func (o *ClusterFeatureResponse) HasCostPerMonth() bool`

HasCostPerMonth returns a boolean if a field has been set.

### SetCostPerMonthNil

`func (o *ClusterFeatureResponse) SetCostPerMonthNil(b bool)`

 SetCostPerMonthNil sets the value for CostPerMonth to be an explicit nil

### UnsetCostPerMonth
`func (o *ClusterFeatureResponse) UnsetCostPerMonth()`

UnsetCostPerMonth ensures that no value is present for CostPerMonth, not even an explicit nil
### GetCurrencyCode

`func (o *ClusterFeatureResponse) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *ClusterFeatureResponse) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *ClusterFeatureResponse) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *ClusterFeatureResponse) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *ClusterFeatureResponse) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *ClusterFeatureResponse) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetValueType

`func (o *ClusterFeatureResponse) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *ClusterFeatureResponse) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *ClusterFeatureResponse) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *ClusterFeatureResponse) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### GetValue

`func (o *ClusterFeatureResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ClusterFeatureResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ClusterFeatureResponse) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ClusterFeatureResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ClusterFeatureResponse) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ClusterFeatureResponse) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetIsValueUpdatable

`func (o *ClusterFeatureResponse) GetIsValueUpdatable() bool`

GetIsValueUpdatable returns the IsValueUpdatable field if non-nil, zero value otherwise.

### GetIsValueUpdatableOk

`func (o *ClusterFeatureResponse) GetIsValueUpdatableOk() (*bool, bool)`

GetIsValueUpdatableOk returns a tuple with the IsValueUpdatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValueUpdatable

`func (o *ClusterFeatureResponse) SetIsValueUpdatable(v bool)`

SetIsValueUpdatable sets IsValueUpdatable field to given value.

### HasIsValueUpdatable

`func (o *ClusterFeatureResponse) HasIsValueUpdatable() bool`

HasIsValueUpdatable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


