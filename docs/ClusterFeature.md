# ClusterFeature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**CostPerMonthInCents** | Pointer to **NullableInt32** |  | [optional] 
**CostPerMonth** | Pointer to **NullableFloat32** |  | [optional] 
**CurrencyCode** | Pointer to **NullableString** |  | [optional] 
**IsCloudProviderPayingFeature** | Pointer to **bool** |  | [optional] 
**CloudProviderFeatureDocumentation** | Pointer to **NullableString** |  | [optional] 
**IsQoveryPayingFeature** | Pointer to **bool** |  | [optional] 
**QoveryFeatureDocumentation** | Pointer to **NullableString** |  | [optional] 
**ValueType** | Pointer to **string** |  | [optional] 
**Value** | Pointer to [**NullableClusterFeatureValue**](ClusterFeatureValue.md) |  | [optional] 
**IsValueUpdatable** | Pointer to **bool** |  | [optional] [default to false]
**AcceptedValues** | Pointer to [**[]ClusterFeatureAcceptedValuesInner**](ClusterFeatureAcceptedValuesInner.md) |  | [optional] 

## Methods

### NewClusterFeature

`func NewClusterFeature() *ClusterFeature`

NewClusterFeature instantiates a new ClusterFeature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterFeatureWithDefaults

`func NewClusterFeatureWithDefaults() *ClusterFeature`

NewClusterFeatureWithDefaults instantiates a new ClusterFeature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterFeature) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterFeature) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterFeature) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterFeature) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *ClusterFeature) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ClusterFeature) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ClusterFeature) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ClusterFeature) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *ClusterFeature) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterFeature) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterFeature) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterFeature) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ClusterFeature) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ClusterFeature) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCostPerMonthInCents

`func (o *ClusterFeature) GetCostPerMonthInCents() int32`

GetCostPerMonthInCents returns the CostPerMonthInCents field if non-nil, zero value otherwise.

### GetCostPerMonthInCentsOk

`func (o *ClusterFeature) GetCostPerMonthInCentsOk() (*int32, bool)`

GetCostPerMonthInCentsOk returns a tuple with the CostPerMonthInCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerMonthInCents

`func (o *ClusterFeature) SetCostPerMonthInCents(v int32)`

SetCostPerMonthInCents sets CostPerMonthInCents field to given value.

### HasCostPerMonthInCents

`func (o *ClusterFeature) HasCostPerMonthInCents() bool`

HasCostPerMonthInCents returns a boolean if a field has been set.

### SetCostPerMonthInCentsNil

`func (o *ClusterFeature) SetCostPerMonthInCentsNil(b bool)`

 SetCostPerMonthInCentsNil sets the value for CostPerMonthInCents to be an explicit nil

### UnsetCostPerMonthInCents
`func (o *ClusterFeature) UnsetCostPerMonthInCents()`

UnsetCostPerMonthInCents ensures that no value is present for CostPerMonthInCents, not even an explicit nil
### GetCostPerMonth

`func (o *ClusterFeature) GetCostPerMonth() float32`

GetCostPerMonth returns the CostPerMonth field if non-nil, zero value otherwise.

### GetCostPerMonthOk

`func (o *ClusterFeature) GetCostPerMonthOk() (*float32, bool)`

GetCostPerMonthOk returns a tuple with the CostPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerMonth

`func (o *ClusterFeature) SetCostPerMonth(v float32)`

SetCostPerMonth sets CostPerMonth field to given value.

### HasCostPerMonth

`func (o *ClusterFeature) HasCostPerMonth() bool`

HasCostPerMonth returns a boolean if a field has been set.

### SetCostPerMonthNil

`func (o *ClusterFeature) SetCostPerMonthNil(b bool)`

 SetCostPerMonthNil sets the value for CostPerMonth to be an explicit nil

### UnsetCostPerMonth
`func (o *ClusterFeature) UnsetCostPerMonth()`

UnsetCostPerMonth ensures that no value is present for CostPerMonth, not even an explicit nil
### GetCurrencyCode

`func (o *ClusterFeature) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *ClusterFeature) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *ClusterFeature) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *ClusterFeature) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### SetCurrencyCodeNil

`func (o *ClusterFeature) SetCurrencyCodeNil(b bool)`

 SetCurrencyCodeNil sets the value for CurrencyCode to be an explicit nil

### UnsetCurrencyCode
`func (o *ClusterFeature) UnsetCurrencyCode()`

UnsetCurrencyCode ensures that no value is present for CurrencyCode, not even an explicit nil
### GetIsCloudProviderPayingFeature

`func (o *ClusterFeature) GetIsCloudProviderPayingFeature() bool`

GetIsCloudProviderPayingFeature returns the IsCloudProviderPayingFeature field if non-nil, zero value otherwise.

### GetIsCloudProviderPayingFeatureOk

`func (o *ClusterFeature) GetIsCloudProviderPayingFeatureOk() (*bool, bool)`

GetIsCloudProviderPayingFeatureOk returns a tuple with the IsCloudProviderPayingFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudProviderPayingFeature

`func (o *ClusterFeature) SetIsCloudProviderPayingFeature(v bool)`

SetIsCloudProviderPayingFeature sets IsCloudProviderPayingFeature field to given value.

### HasIsCloudProviderPayingFeature

`func (o *ClusterFeature) HasIsCloudProviderPayingFeature() bool`

HasIsCloudProviderPayingFeature returns a boolean if a field has been set.

### GetCloudProviderFeatureDocumentation

`func (o *ClusterFeature) GetCloudProviderFeatureDocumentation() string`

GetCloudProviderFeatureDocumentation returns the CloudProviderFeatureDocumentation field if non-nil, zero value otherwise.

### GetCloudProviderFeatureDocumentationOk

`func (o *ClusterFeature) GetCloudProviderFeatureDocumentationOk() (*string, bool)`

GetCloudProviderFeatureDocumentationOk returns a tuple with the CloudProviderFeatureDocumentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderFeatureDocumentation

`func (o *ClusterFeature) SetCloudProviderFeatureDocumentation(v string)`

SetCloudProviderFeatureDocumentation sets CloudProviderFeatureDocumentation field to given value.

### HasCloudProviderFeatureDocumentation

`func (o *ClusterFeature) HasCloudProviderFeatureDocumentation() bool`

HasCloudProviderFeatureDocumentation returns a boolean if a field has been set.

### SetCloudProviderFeatureDocumentationNil

`func (o *ClusterFeature) SetCloudProviderFeatureDocumentationNil(b bool)`

 SetCloudProviderFeatureDocumentationNil sets the value for CloudProviderFeatureDocumentation to be an explicit nil

### UnsetCloudProviderFeatureDocumentation
`func (o *ClusterFeature) UnsetCloudProviderFeatureDocumentation()`

UnsetCloudProviderFeatureDocumentation ensures that no value is present for CloudProviderFeatureDocumentation, not even an explicit nil
### GetIsQoveryPayingFeature

`func (o *ClusterFeature) GetIsQoveryPayingFeature() bool`

GetIsQoveryPayingFeature returns the IsQoveryPayingFeature field if non-nil, zero value otherwise.

### GetIsQoveryPayingFeatureOk

`func (o *ClusterFeature) GetIsQoveryPayingFeatureOk() (*bool, bool)`

GetIsQoveryPayingFeatureOk returns a tuple with the IsQoveryPayingFeature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsQoveryPayingFeature

`func (o *ClusterFeature) SetIsQoveryPayingFeature(v bool)`

SetIsQoveryPayingFeature sets IsQoveryPayingFeature field to given value.

### HasIsQoveryPayingFeature

`func (o *ClusterFeature) HasIsQoveryPayingFeature() bool`

HasIsQoveryPayingFeature returns a boolean if a field has been set.

### GetQoveryFeatureDocumentation

`func (o *ClusterFeature) GetQoveryFeatureDocumentation() string`

GetQoveryFeatureDocumentation returns the QoveryFeatureDocumentation field if non-nil, zero value otherwise.

### GetQoveryFeatureDocumentationOk

`func (o *ClusterFeature) GetQoveryFeatureDocumentationOk() (*string, bool)`

GetQoveryFeatureDocumentationOk returns a tuple with the QoveryFeatureDocumentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQoveryFeatureDocumentation

`func (o *ClusterFeature) SetQoveryFeatureDocumentation(v string)`

SetQoveryFeatureDocumentation sets QoveryFeatureDocumentation field to given value.

### HasQoveryFeatureDocumentation

`func (o *ClusterFeature) HasQoveryFeatureDocumentation() bool`

HasQoveryFeatureDocumentation returns a boolean if a field has been set.

### SetQoveryFeatureDocumentationNil

`func (o *ClusterFeature) SetQoveryFeatureDocumentationNil(b bool)`

 SetQoveryFeatureDocumentationNil sets the value for QoveryFeatureDocumentation to be an explicit nil

### UnsetQoveryFeatureDocumentation
`func (o *ClusterFeature) UnsetQoveryFeatureDocumentation()`

UnsetQoveryFeatureDocumentation ensures that no value is present for QoveryFeatureDocumentation, not even an explicit nil
### GetValueType

`func (o *ClusterFeature) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *ClusterFeature) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *ClusterFeature) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *ClusterFeature) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### GetValue

`func (o *ClusterFeature) GetValue() ClusterFeatureValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ClusterFeature) GetValueOk() (*ClusterFeatureValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ClusterFeature) SetValue(v ClusterFeatureValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *ClusterFeature) HasValue() bool`

HasValue returns a boolean if a field has been set.

### SetValueNil

`func (o *ClusterFeature) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *ClusterFeature) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetIsValueUpdatable

`func (o *ClusterFeature) GetIsValueUpdatable() bool`

GetIsValueUpdatable returns the IsValueUpdatable field if non-nil, zero value otherwise.

### GetIsValueUpdatableOk

`func (o *ClusterFeature) GetIsValueUpdatableOk() (*bool, bool)`

GetIsValueUpdatableOk returns a tuple with the IsValueUpdatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsValueUpdatable

`func (o *ClusterFeature) SetIsValueUpdatable(v bool)`

SetIsValueUpdatable sets IsValueUpdatable field to given value.

### HasIsValueUpdatable

`func (o *ClusterFeature) HasIsValueUpdatable() bool`

HasIsValueUpdatable returns a boolean if a field has been set.

### GetAcceptedValues

`func (o *ClusterFeature) GetAcceptedValues() []ClusterFeatureAcceptedValuesInner`

GetAcceptedValues returns the AcceptedValues field if non-nil, zero value otherwise.

### GetAcceptedValuesOk

`func (o *ClusterFeature) GetAcceptedValuesOk() (*[]ClusterFeatureAcceptedValuesInner, bool)`

GetAcceptedValuesOk returns a tuple with the AcceptedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedValues

`func (o *ClusterFeature) SetAcceptedValues(v []ClusterFeatureAcceptedValuesInner)`

SetAcceptedValues sets AcceptedValues field to given value.

### HasAcceptedValues

`func (o *ClusterFeature) HasAcceptedValues() bool`

HasAcceptedValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


